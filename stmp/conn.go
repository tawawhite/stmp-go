// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-26 18:16:34
package stmp

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type CallOptions struct {
	UseNotify bool
	Packet    *Packet
}

func (o *CallOptions) Notify() *CallOptions {
	o.UseNotify = true
	return o
}

func (o *CallOptions) KeepPacket(p *Packet) *CallOptions {
	o.Packet = p
	return o
}

func NewCallOptions() *CallOptions {
	return &CallOptions{}
}

var NotifyOptions = NewCallOptions().Notify()

type ConnOptions struct {
	Logger           *zap.Logger
	HandshakeTimeout time.Duration
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	Router           *Router
	MaxPacketSize    uint64
	WriteQueueSize   int
}

func (o *ConnOptions) WithLogger(logger *zap.Logger) *ConnOptions {
	o.Logger = logger
	return o
}

func (o *ConnOptions) WithRouter(r *Router) *ConnOptions {
	o.Router = r
	return o
}

func (o *ConnOptions) WithWriteQueueLimit(max int) *ConnOptions {
	o.WriteQueueSize = max
	return o
}

func (o *ConnOptions) WithPacketSizeLimit(max uint64) *ConnOptions {
	o.MaxPacketSize = max
	return o
}

func (o *ConnOptions) WithTimeout(handshake, read, write time.Duration) *ConnOptions {
	o.HandshakeTimeout = handshake
	o.ReadTimeout = read
	o.WriteTimeout = write
	return o
}

func (o *ConnOptions) ApplyDefault() *ConnOptions {
	no := o
	if no == nil {
		no = NewConnOptions()
	}
	if no.Router == nil {
		no.Router = NewRouter()
	}
	if no.Logger == nil {
		var err error
		no.Logger, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	}
	return no
}

func NewConnOptions() *ConnOptions {
	return &ConnOptions{
		HandshakeTimeout: time.Minute,
		// ping timeout
		ReadTimeout:  time.Minute * 2,
		WriteTimeout: time.Minute,
		Router:       nil,
		// 16 mb
		MaxPacketSize:  1 << 24,
		WriteQueueSize: 16,
	}
}

type writeEvent struct {
	p *Packet
	r chan error
}

// the struct will only keep the required fields for the connection to save space at server side
type Conn struct {
	net.Conn
	opts *ConnOptions
	mu   sync.Mutex
	mid  *uint32

	// Write signal, if closed, means the conn is closed
	// both the read and write channel will stop
	writeChan chan *writeEvent
	// lock for write requests
	// pending requests, waiting for response
	requests map[uint16]chan *Packet

	// the stmp major version
	Major byte
	// the stmp minor version
	Minor byte

	// content-type codec
	Media MediaCodec

	// client writeHandshakeResponse request header
	ClientHeader Header
	// server writeHandshakeResponse response header
	ServerHeader Header
}

func (c *Conn) Call(ctx context.Context, method string, payload []byte, opts *CallOptions) (out interface{}, err error) {
	action := ms.methods[method]
	p := &Packet{Fin: true, Kind: MessageKindRequest, Action: action, Payload: payload}
	we := &writeEvent{p: p}
	var r chan *Packet
	if opts.UseNotify {
		p.Kind = MessageKindNotify
	} else {
		we.r = make(chan error, 1)
		r = make(chan *Packet, 1)
		p.Mid = uint16(atomic.AddUint32(c.mid, 1))
		c.mu.Lock()
		c.requests[p.Mid] = r
		c.mu.Unlock()
	}
	c.writeChan <- we
	if r == nil {
		return
	}
	select {
	case err = <-we.r:
	case <-ctx.Done():
		err = NewStatusError(StatusCancelled, ctx.Err())
		return
	}
	if err != nil {
		c.mu.Lock()
		delete(c.requests, p.Mid)
		c.mu.Unlock()
		err = NewStatusError(StatusNetworkError, err)
		return
	}
	select {
	case p = <-r:
	case <-ctx.Done():
		err = NewStatusError(StatusCancelled, ctx.Err())
		return
	}
	if opts.Packet != nil {
		*opts.Packet = *p
	}
	if p.Status != StatusOk {
		err = NewStatusError(p.Status, string(p.Payload))
		return
	}
	if p.Payload == nil {
		return
	}
	out = ms.actions[action].output()
	err = c.Media.Unmarshal(p.Payload, out)
	return
}

func (c *Conn) Invoke(ctx context.Context, method string, in interface{}, opts *CallOptions) (interface{}, error) {
	var payload []byte
	if in != nil {
		var err error
		payload, err = c.Media.Marshal(in)
		if err != nil {
			return nil, NewStatusError(StatusMarshalError, err.Error())
		}
	}
	return c.Call(ctx, method, payload, opts)
}

func (c *Conn) Close(status Status, message string) error {
	// TODO
	return nil
}

func NewConn(nc net.Conn, opts *ConnOptions) *Conn {
	return &Conn{
		Conn:      nc,
		opts:      opts,
		Major:     1,
		Minor:     0,
		writeChan: make(chan *writeEvent, opts.WriteQueueSize),
		requests:  map[uint16]chan *Packet{},
	}
}

func (c *Conn) checkPendingVolume() {
}

func (c *Conn) dispatchPacket(p *Packet) {
	switch p.Kind {
	case MessageKindPing:
		c.handlePing()
	case MessageKindPong:
		c.handlePong()
	case MessageKindRequest:
		go c.handleRequest(p.Mid, p.Action, p.Payload)
	case MessageKindNotify:
		go c.handleNotify(p.Action, p.Payload)
	case MessageKindResponse:
		c.handleResponse(p.Mid, p.Status, p.Payload)
	case MessageKindClose:
		c.handleClose(p.Status, string(p.Payload))
	}
}

func (c *Conn) binaryReadChannel(r io.ReadCloser) error {
	p := new(Packet)
	buf := make([]byte, MaxStreamHeadSize, MaxStreamHeadSize)
	var err error
	for {
		c.Conn.SetReadDeadline(time.Now().Add(c.opts.ReadTimeout))
		err = p.Read(r, buf)
		if err != nil {
			// TODO
			return err
		}
		c.dispatchPacket(p)
	}
}

func (c *Conn) binaryWriteChannel(w EncodingWriter) {
	var e *writeEvent
	var ok bool
	var err error
	buf := make([]byte, MaxStreamHeadSize, MaxStreamHeadSize)
	for {
		e, ok = <-c.writeChan
		if !ok {
			// TODO
			return
		}
		c.Conn.SetWriteDeadline(time.Now().Add(c.opts.WriteTimeout))
		err = e.p.Write(w, buf)
		if e.r != nil {
			e.r <- err
		}
		if err != nil {
			// TODO
			return
		}
	}
}

func (c *Conn) websocketBinaryReadChannel(wc *websocket.Conn) {
	p := new(Packet)
	var err error
	var data []byte
	for {
		wc.SetReadDeadline(time.Now().Add(c.opts.ReadTimeout))
		_, data, err = wc.ReadMessage()
		if err != nil {
			// TODO
			return
		}
		err = p.UnmarshalBinary(data)
		if err != nil {
			// TODO
			return
		}
		c.dispatchPacket(p)
	}
}

func (c *Conn) websocketBinaryWriteChannel(wc *websocket.Conn) {
	var e *writeEvent
	var ok bool
	var err error
	var data []byte
	buf := make([]byte, MaxBinaryHeadSize, MaxBinaryHeadSize)
	for {
		e, ok = <-c.writeChan
		if !ok {
			// TODO
			return
		}
		wc.SetWriteDeadline(time.Now().Add(c.opts.WriteTimeout))
		data = e.p.MarshalBinary(buf)
		err = wc.WriteMessage(websocket.BinaryMessage, data)
		if e.r != nil {
			e.r <- err
		}
		if err != nil {
			// TODO
			return
		}
	}
}

func (c *Conn) websocketTextReadChannel(wc *websocket.Conn) {
	p := new(Packet)
	var err error
	var data []byte
	for {
		wc.SetReadDeadline(time.Now().Add(c.opts.ReadTimeout))
		_, data, err = wc.ReadMessage()
		if err != nil {
			// TODO
			return
		}
		err = p.UnmarshalText(data)
		if err != nil {
			// TODO
			return
		}
		c.dispatchPacket(p)
	}
}

func (c *Conn) websocketTextWriteChannel(wc *websocket.Conn) {
	var e *writeEvent
	var ok bool
	var err error
	var data []byte
	buf := make([]byte, MaxTextHeadSize, MaxTextHeadSize)
	for {
		e, ok = <-c.writeChan
		if !ok {
			// TODO
			return
		}
		wc.SetWriteDeadline(time.Now().Add(c.opts.WriteTimeout))
		data = e.p.MarshalText(buf)
		err = wc.WriteMessage(websocket.TextMessage, data)
		if e.r != nil {
			e.r <- err
		}
		if err != nil {
			// TODO
			return
		}
	}
}

func (c *Conn) handlePing() {
	c.writeChan <- &writeEvent{p: &Packet{Kind: MessageKindPong}}
}

func (c *Conn) handlePong() {
	// do nothing for deadline limited pong message receive rate
}

func (c *Conn) handleClose(status Status, message string) {
	// TODO close connection
}

func (c *Conn) handleNotify(action uint64, payload []byte) {
	ctx := WithConn(context.Background(), c)
	c.opts.Router.dispatch(ctx, action, payload, c.Media)
}

func (c *Conn) handleRequest(mid uint16, action uint64, payload []byte) {
	ctx := WithConn(context.Background(), c)
	status, payload := c.opts.Router.dispatch(ctx, action, payload, c.Media)
	we := &writeEvent{p: &Packet{
		Fin:     true,
		Kind:    MessageKindResponse,
		Mid:     mid,
		Action:  action,
		Status:  status,
		Payload: payload,
	}}
	c.writeChan <- we
}

func (c *Conn) handleResponse(mid uint16, status Status, payload []byte) {
	c.mu.Lock()
	q, ok := c.requests[mid]
	delete(c.requests, mid)
	c.mu.Unlock()
	if ok {
		q <- &Packet{Status: status, Payload: payload}
	}
}

func (c *Conn) negotiate() *StatusError {
	mediaInput := c.ClientHeader.Get(AcceptContentType)
	inputLength := 0
	var inputValue string
	for inputLength < len(mediaInput) {
		inputValue, inputLength = ReadNegotiate(mediaInput)
		if c.Media = GetMediaCodec(inputValue); c.Media != nil {
			c.ServerHeader.Set(DetermineContentType, inputValue)
			break
		}
	}
	if c.Media == nil {
		return NewStatusError(StatusUnsupportedContentType, "")
	}
	encodingInput := c.ClientHeader.Get(AcceptEncoding)
	inputLength = 0
	var encoding EncodingCodec
	for inputLength < len(encodingInput) {
		inputValue, inputLength = ReadNegotiate(encodingInput)
		if encoding = GetEncodingCodec(inputValue); encoding != nil {
			c.ServerHeader.Set(DetermineEncoding, inputValue)
			break
		}
	}
	packetFormat := c.ClientHeader.Get(DeterminePacketFormat)
	if packetFormat != "" {
		c.ServerHeader.Set(DeterminePacketFormat, packetFormat)
	}
	return nil
}

func (c *Conn) initEncoding(compressLevel int) (r io.ReadCloser, w EncodingWriter, err error) {
	c.Media = GetMediaCodec(c.ServerHeader.Get(DetermineContentType))
	if c.Media == nil {
		err = errors.New("cannot find content-type: " + c.ServerHeader.Get(DetermineContentType) + ", please register it first")
		return
	}
	ec := GetEncodingCodec(c.ServerHeader.Get(DetermineEncoding))
	if ec == nil {
		rw := plainEncoding{Conn: c.Conn}
		r = rw
		w = rw
		return
	} else {
		r, err = ec.Reader(c.Conn)
		if err != nil {
			return
		}
		w, err = ec.Writer(c.Conn, compressLevel)
		if err != nil {
			r.Close()
			r = nil
			return
		}
		return
	}
}
