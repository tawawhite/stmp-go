// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-26 18:16:34
package stmp

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type writeEvent struct {
	p *Packet
	r chan error
}

type Conn struct {
	net.Conn
	// router to dispatch actions
	*router
	mu *sync.Mutex
	// the stmp major version
	Major byte
	// the stmp minor version
	Minor byte
	// client writeHandshakeResponse request header
	ClientHeader Header
	// server writeHandshakeResponse response header
	ServerHeader Header
	// client writeHandshakeResponse request message
	ClientMessage string
	// server writeHandshakeResponse response message
	ServerMessage string
	// network conn
	// content-type codec
	media MediaCodec
	// Write signal
	writeChan chan *writeEvent
	// close state
	closeChan chan struct{}
	// pending requests & responses
	requests map[uint16]chan *Packet

	mid *uint32

	onClosed CloseHandlerFunc
}

func (c *Conn) SetCloseHandler(h CloseHandlerFunc) {
	c.onClosed = h
}

// TODO check connection Status
func (c *Conn) call(ctx context.Context, action uint64, payload []byte, options *CallOptions) (out interface{}, err error) {
	p := &Packet{Fin: true, Kind: MessageKindRequest, Action: action, Payload: payload}
	we := &writeEvent{p: p}
	var r chan *Packet
	if options.Notify {
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
	if options.Response != nil {
		*options.Response = p.Payload
	}
	if p.Status != StatusOk {
		err = NewStatusError(p.Status, string(p.Payload))
		return
	}
	if p.Payload == nil {
		return
	}
	out = ms.actions[action].output()
	err = c.media.Unmarshal(p.Payload, out)
	return
}

func (c *Conn) Invoke(ctx context.Context, send *SendOptions, opts ...CallOption) (interface{}, error) {
	callOptions := NewCallOptions(opts...)
	var payload []byte
	if send.input != nil {
		var err error
		payload, err = c.media.Marshal(send.input)
		if err != nil {
			return nil, NewStatusError(StatusMarshalError, err.Error())
		}
	}
	return c.call(ctx, send.action, payload, callOptions)
}

func (c *Conn) Close(status Status, message string) error {
	// TODO
	return nil
}

func newConn(nc net.Conn) *Conn {
	return &Conn{
		Major:     1,
		Minor:     0,
		mu:        &sync.Mutex{},
		Conn:      nc,
		writeChan: make(chan *writeEvent),
		closeChan: make(chan struct{}),
		requests:  map[uint16]chan *Packet{},
	}
}

func (c *Conn) writeHandshakeResponse(status Status) error {
	// TODO
	return nil
}

func (c *Conn) websocketWriteHandshakeResponse(status Status) error {
	// TODO
	return nil
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

func (c *Conn) binaryReadChannel(r io.ReadCloser, timeout time.Duration) {
	p := new(Packet)
	buf := make([]byte, MaxStreamHeadSize, MaxStreamHeadSize)
	var err error
	for {
		c.Conn.SetReadDeadline(time.Now().Add(timeout))
		err = p.Read(r, buf)
		if err != nil {
			// TODO
			return
		}
		c.dispatchPacket(p)
	}
}

func (c *Conn) binaryWriteChannel(w EncodingWriter, timeout time.Duration) {
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
		c.Conn.SetWriteDeadline(time.Now().Add(timeout))
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

func (c *Conn) websocketBinaryReadChannel(wc *websocket.Conn, timeout time.Duration) {
	p := new(Packet)
	var err error
	var data []byte
	for {
		wc.SetReadDeadline(time.Now().Add(timeout))
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

func (c *Conn) websocketBinaryWriteChannel(wc *websocket.Conn, timeout time.Duration) {
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
		wc.SetWriteDeadline(time.Now().Add(timeout))
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

func (c *Conn) websocketTextReadChannel(wc *websocket.Conn, timeout time.Duration) {
	p := new(Packet)
	var err error
	var data []byte
	for {
		wc.SetReadDeadline(time.Now().Add(timeout))
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

func (c *Conn) websocketTextWriteChannel(wc *websocket.Conn, timeout time.Duration) {
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
		wc.SetWriteDeadline(time.Now().Add(timeout))
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
	c.dispatch(ctx, action, payload, c.media)
}

func (c *Conn) handleRequest(mid uint16, action uint64, payload []byte) {
	ctx := WithConn(context.Background(), c)
	status, payload := c.dispatch(ctx, action, payload, c.media)
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
		if c.media = GetMediaCodec(inputValue); c.media != nil {
			c.ServerHeader.Set(DetermineContentType, inputValue)
			break
		}
	}
	if c.media == nil {
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
	c.media = GetMediaCodec(c.ServerHeader.Get(DetermineContentType))
	if c.media == nil {
		err = errors.New("cannot find the codec for content-type: " + c.ServerHeader.Get(DetermineContentType) + ", please register it at first")
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
