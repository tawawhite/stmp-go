package stmp

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"io"
	"net"
	"sync/atomic"
	"time"
)

// the options to make a call
//
// The options to keep response packet.
//  func ExampleCallOptions_keepPacket() {
//      p := stmp.NewPacket()
//      someClient.DoSomething(ctx, input, stmp.NewCallOptions().KeepPacket(p))
//      log.Printf("Kind=%d, Mid=%d, Status=%d, Payload=%s.", p.Kind, p.Mid, p.Status, string(p.Payload))
//      // Output: Kind=4, Mid=1, Status=0, Payload=.
//  }
type CallOptions struct {
	useNotify  bool
	keepPacket *Packet
}

// set default options
func (o *CallOptions) ApplyDefault() *CallOptions {
	no := o
	if no == nil {
		no = NewCallOptions()
	}
	return no
}

// call with notify, which means the conn will not wait for the response
func (o *CallOptions) Notify() *CallOptions {
	o.useNotify = true
	return o
}

// keep response packet, you can get the raw information from the packet
func (o *CallOptions) KeepPacket(p *Packet) *CallOptions {
	o.keepPacket = p
	return o
}

// create a default call options
func NewCallOptions() *CallOptions {
	return &CallOptions{}
}

// pick the first call options, if the input is empty, will build a default options
func PickCallOptions(opts ...*CallOptions) *CallOptions {
	if len(opts) == 0 {
		return NewCallOptions().ApplyDefault()
	}
	return opts[0]
}

// the default notify options, this is used for protoc-gen-stmp
var NotifyOptions = NewCallOptions().ApplyDefault().Notify()

// the connection configurations
type ConnOptions struct {
	logger           *zap.Logger
	handshakeTimeout time.Duration
	readTimeout      time.Duration
	writeTimeout     time.Duration
	router           *Router
	maxPacketSize    uint64
	writeQueueSize   int
	compressLevel    int
}

// set custom logger, default is zap.NewProduction()
func (o *ConnOptions) WithLogger(logger *zap.Logger) *ConnOptions {
	o.logger = logger
	return o
}

func (o *ConnOptions) WithCompress(level int) *ConnOptions {
	o.compressLevel = level
	return o
}

// set max write queue size, default is 8
func (o *ConnOptions) WithWriteQueueLimit(max int) *ConnOptions {
	o.writeQueueSize = max
	return o
}

// set the max packet size, default is 16mb
func (o *ConnOptions) WithPacketSizeLimit(max uint64) *ConnOptions {
	o.maxPacketSize = max
	return o
}

// set timeouts, the read timeout means ping interval
func (o *ConnOptions) WithTimeout(handshake, read, write time.Duration) *ConnOptions {
	o.handshakeTimeout = handshake
	o.readTimeout = read
	o.writeTimeout = write
	return o
}

// build final options, check default options
func (o *ConnOptions) ApplyDefault() *ConnOptions {
	no := o
	if no == nil {
		no = NewConnOptions()
	}
	if no.logger == nil {
		var err error
		no.logger, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	}
	return no
}

// create a new conn options
func NewConnOptions() *ConnOptions {
	return &ConnOptions{
		handshakeTimeout: time.Minute,
		// ping interval
		readTimeout:  time.Minute * 2,
		writeTimeout: time.Minute,
		// 16 mb
		maxPacketSize:  1 << 24,
		writeQueueSize: 8,
		compressLevel:  6,
	}
}

// send write signal to write bump
//
// if p is ClosePacket, r is nil, means receive close event from another peer, else means the close event
// is sent by self peer, it may occur when read error, or user call close method.
//
// else the r could be nil or not, if is not nil, the write err will send to it.
type writeEvent struct {
	p *Packet
	r chan StatusError
}

var writePingEvent = writeEvent{p: &Packet{Kind: MessageKindPing}}

// the struct will only keep the required fields for the connection to save space at server side
type Conn struct {
	net.Conn
	State
	opts *ConnOptions
	mid  *uint32

	// the packets to send
	writeChan chan writeEvent

	// lock for write requests
	// pending requests, waiting for response
	requests map[uint16]chan *Packet

	// the stmp major version
	Major byte
	// the stmp minor version
	Minor byte

	// the content-type codec
	Media MediaCodec

	// client handshake request header
	ClientHeader Header
	// server handshake response header
	ServerHeader Header
}

// create a conn from net.Conn and options
func NewConn(nc net.Conn, opts *ConnOptions) *Conn {
	return &Conn{
		State:     NewState(),
		Conn:      nc,
		opts:      opts,
		mid:       new(uint32),
		writeChan: make(chan writeEvent, opts.writeQueueSize),
		requests:  make(map[uint16]chan *Packet),
	}
}

// invoke a method a marshaled payload
func (c *Conn) Call(ctx context.Context, method string, payload []byte, opts *CallOptions) (out interface{}, err error) {
	action, ok := ms.methods[method]
	if !ok {
		err = NewStatusError(StatusBadRequest, "method "+method+" is not registered")
		return
	}
	p := &Packet{Action: action, Payload: payload}
	var r chan *Packet
	if opts.useNotify {
		p.Kind = MessageKindNotify
	} else {
		p.Mid = uint16(atomic.AddUint32(c.mid, 1))
		r = make(chan *Packet, 1)
		c.Lock()
		c.requests[p.Mid] = r
		c.Unlock()
	}
	err = c.send(ctx, p, true)
	if r == nil {
		return
	}
	// send error
	if err != nil {
		c.Lock()
		delete(c.requests, p.Mid)
		c.Unlock()
		return
	}
	// wait response
	select {
	case p = <-r:
	case <-ctx.Done():
		err = NewStatusError(StatusRequestTimeout, "wait response error: "+ctx.Err().Error())
		c.Lock()
		delete(c.requests, p.Mid)
		c.Unlock()
		return
	}
	if opts.keepPacket != nil {
		*opts.keepPacket = *p
	}
	if p.Status != StatusOk {
		err = NewStatusError(p.Status, string(p.Payload))
		return
	}
	if p.Payload == nil {
		return
	}
	out = ms.actions[action].Output()
	err = c.Media.Unmarshal(p.Payload, out)
	if err != nil {
		out = nil
		err = DetectError(err, StatusUnknown)
	}
	return
}

// invoke a method with raw input, will marshal it with conn's media codec
func (c *Conn) Invoke(ctx context.Context, method string, in interface{}, opts *CallOptions) (interface{}, error) {
	var payload []byte
	if !isNil(in) {
		var err error
		payload, err = c.Media.Marshal(in)
		if err != nil {
			return nil, NewStatusError(StatusBadRequest, err)
		}
	}
	return c.Call(ctx, method, payload, opts)
}

func (c *Conn) terminate() {
	res := &Packet{Kind: MessageKindResponse, Status: StatusNetworkError, Payload: []byte("connection closed")}
	c.Lock()
	for _, p := range c.requests {
		p <- res
	}
	c.requests = nil
	c.Unlock()
	close(c.writeChan)
	c.Conn.Close()
	for {
		e, ok := <-c.writeChan
		if !ok {
			break
		}
		if e.r != nil {
			e.r <- NewStatusError(StatusNetworkError, "connection closed")
		}
	}
}

func (c *Conn) send(ctx context.Context, p *Packet, wait bool) (se StatusError) {
	defer func() {
		if recover() != nil {
			se = NewStatusError(StatusNetworkError, "connection closed already")
		}
	}()
	we := writeEvent{p: p}
	if wait {
		we.r = make(chan StatusError, 1)
	}
	select {
	case c.writeChan <- we:
	case <-ctx.Done():
		se = NewStatusError(StatusRequestTimeout, "pending timeout: "+ctx.Err().Error())
		return
	}
	if wait {
		select {
		case se = <-we.r:
			return
		case <-ctx.Done():
			se = NewStatusError(StatusRequestTimeout, "wait timeout: "+ctx.Err().Error())
		}
	}
	return
}

// close the connection manually
func (c *Conn) Close(status Status, message string) (err error) {
	if status == StatusServerShutdown {
		return errors.New("reserved close status: " + MapStatus[StatusServerShutdown])
	}
	return c.send(context.Background(), NewClosePacket(status, message), true)
}

const maxStreamHeadSize = 23

func (c *Conn) read() {
	p := new(Packet)
	buf := make([]byte, maxStreamHeadSize, maxStreamHeadSize)
	var err error
	var r io.ReadCloser
	ec := GetEncodingCodec(c.ServerHeader.Get(DetermineEncoding))
	if ec == nil {
		r = plainEncoding{Conn: c.Conn}
	} else {
		r, err = ec.Reader(c)
		if err != nil {
			c.send(context.Background(), NewClosePacket(StatusProtocolError, "init encoding reader error: "+err.Error()), false)
			return
		}
	}
	for {
		c.Conn.SetReadDeadline(time.Now().Add(c.opts.readTimeout))
		se := p.Read(r, buf, c.opts.maxPacketSize)
		if se != nil {
			c.send(context.Background(), NewClosePacket(se.Code(), se.Message()), false)
			break
		}
		c.dispatchPacket(p)
	}
	r.Close()
}

// response value never be nil, if close normal, it will be status ok
func (c *Conn) write() StatusError {
	var err error
	var w EncodingWriter
	ec := GetEncodingCodec(c.ServerHeader.Get(DetermineEncoding))
	if ec == nil {
		w = plainEncoding{Conn: c.Conn}
	} else {
		w, err = ec.Writer(c, c.opts.compressLevel)
		if err != nil {
			c.Conn.Close()
			return NewStatusError(StatusProtocolError, "init encoding writer error: "+err.Error())
		}
	}
	var se StatusError
	var e writeEvent
	buf := make([]byte, maxStreamHeadSize, maxStreamHeadSize)
	ticker := time.NewTicker(c.opts.readTimeout)
	for {
		select {
		case <-ticker.C:
			e = writePingEvent
		case e = <-c.writeChan:
		}
		c.Conn.SetWriteDeadline(time.Now().Add(c.opts.writeTimeout))
		se = e.p.Write(w, buf)
		if e.r != nil {
			e.r <- se
		}
		// if write error occurs, will stop any write immediately, for client side cannot split packet
		// if write more packet.
		if se != nil {
			break
		}
		if e.p.Kind == MessageKindClose && e.r != nil {
			// use close manually, which means reader is reading
			for {
				ce := <-c.writeChan
				if ce.p.Kind == MessageKindClose && ce.r == nil {
					// receive passive close event, maybe sent by remote peer, or read error from local peer
					// the channel will done
					break
				} else if ce.r != nil {
					ce.r <- NewStatusError(StatusNetworkError, "connection is closing")
				}
			}
		} else if e.p.Kind == MessageKindClose {
			se = NewStatusError(e.p.Status, string(e.p.Payload))
			break
		}
	}
	w.Close()
	ticker.Stop()
	c.terminate()
	return se
}

func (c *Conn) readBinaryWebsocket(wc *websocket.Conn) {
	p := new(Packet)
	var err error
	var data []byte
	for {
		wc.SetReadDeadline(time.Now().Add(c.opts.readTimeout))
		if _, data, err = wc.ReadMessage(); err != nil {
			c.send(context.Background(), NewClosePacket(StatusNetworkError, "read packet error: "+err.Error()), false)
			break
		}
		if se := p.UnmarshalBinary(data); se != nil {
			c.send(context.Background(), NewClosePacket(se.Code(), se.Message()), false)
			break
		}
		c.dispatchPacket(p)
	}
}

const maxBinaryHeadSize = 13

func (c *Conn) writeBinaryWebsocket(wc *websocket.Conn) StatusError {
	var e writeEvent
	var err error
	var se StatusError
	var data []byte
	buf := make([]byte, maxBinaryHeadSize, maxBinaryHeadSize)
	ticker := time.NewTicker(c.opts.readTimeout)
	for {
		select {
		case <-ticker.C:
			e = writePingEvent
		case e = <-c.writeChan:
		}
		if e.p.Kind == MessageKindClose && e.r == nil {
			break
		}
		wc.SetWriteDeadline(time.Now().Add(c.opts.writeTimeout))
		data = e.p.MarshalBinary(buf)
		err = wc.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			se = NewStatusError(StatusNetworkError, "write packet error: "+err.Error())
		}
		if e.r != nil {
			e.r <- se
		}
		// if write error occurs, will stop any write immediately, for client side cannot split packet
		// if write more packet.
		if se != nil {
			break
		}
		if e.p.Kind == MessageKindClose && e.r != nil {
			// use close manually, which means reader is reading
			for {
				ce := <-c.writeChan
				if ce.p.Kind == MessageKindClose && ce.r == nil {
					// receive passive close event, maybe sent by remote peer, or read error from local peer
					// the channel will done
					break
				} else if ce.r != nil {
					ce.r <- NewStatusError(StatusNetworkError, "connection is closing")
				}
			}
		} else if e.p.Kind == MessageKindClose {
			se = NewStatusError(e.p.Status, string(e.p.Payload))
			break
		}
	}
	wc.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "OK"), time.Time{})
	ticker.Stop()
	c.terminate()
	return se
}

func (c *Conn) readTextWebsocket(wc *websocket.Conn) {
	p := new(Packet)
	var err error
	var data []byte
	for {
		wc.SetReadDeadline(time.Now().Add(c.opts.readTimeout))
		if _, data, err = wc.ReadMessage(); err != nil {
			c.send(context.Background(), NewClosePacket(StatusNetworkError, "read packet error: "+err.Error()), false)
			break
		}
		if se := p.UnmarshalText(data); se != nil {
			c.send(context.Background(), NewClosePacket(se.Code(), se.Message()), false)
			break
		}
		c.dispatchPacket(p)
	}
}

const maxTextHeadSize = 21

func (c *Conn) writeTextWebsocket(wc *websocket.Conn) StatusError {
	var e writeEvent
	var err error
	var se StatusError
	var data []byte
	buf := make([]byte, maxTextHeadSize, maxTextHeadSize)
	ticker := time.NewTicker(c.opts.readTimeout)
	for {
		select {
		case <-ticker.C:
			e = writePingEvent
		case e = <-c.writeChan:
		}
		if e.p.Kind == MessageKindClose && e.r == nil {
			break
		}
		wc.SetWriteDeadline(time.Now().Add(c.opts.writeTimeout))
		data = e.p.MarshalText(buf)
		err = wc.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			se = NewStatusError(StatusNetworkError, "write packet error: "+err.Error())
		}
		if e.r != nil {
			e.r <- se
		}
		// if write error occurs, will stop any write immediately, for client side cannot split packet
		// if write more packet.
		if se != nil {
			break
		}
		if e.p.Kind == MessageKindClose && e.r != nil {
			// use close manually, which means reader is reading
			for {
				ce := <-c.writeChan
				if ce.p.Kind == MessageKindClose && ce.r == nil {
					// receive passive close event, maybe sent by remote peer, or read error from local peer
					// the channel will done
					break
				} else if ce.r != nil {
					ce.r <- NewStatusError(StatusNetworkError, "connection is closing")
				}
			}
		} else if e.p.Kind == MessageKindClose {
			se = NewStatusError(e.p.Status, string(e.p.Payload))
			break
		}
	}
	wc.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "OK"), time.Time{})
	ticker.Stop()
	c.terminate()
	return se
}

func (c *Conn) dispatchPacket(p *Packet) {
	switch p.Kind {
	case MessageKindPing:
		c.handlePing()
	case MessageKindPong:
		c.handlePong()
	case MessageKindRequest, MessageKindNotify:
		go c.handleRequest(p)
	case MessageKindResponse:
		c.handleResponse(p)
	case MessageKindClose:
		c.handleClose(p)
	}
}

func (c *Conn) handlePing() {
	c.send(context.Background(), &Packet{Kind: MessageKindPong}, false)
}

func (c *Conn) handlePong() {
	// do nothing for deadline limited pong message receive rate
}

func (c *Conn) handleClose(p *Packet) {
	c.send(context.Background(), p, false)
}

// TODO with context
func (c *Conn) logRequest(p *Packet, status Status, res []byte, err error) {
	fields := []zap.Field{
		zap.String("addr", c.RemoteAddr().String()),
		zap.String("server", c.LocalAddr().String()),
		zap.String("Mid", hexFormatUint64(uint64(p.Mid))),
		zap.String("Status", hexFormatUint64(uint64(status))),
	}
	method := ms.actions[p.Action]
	if method == nil {
		fields = append(fields, zap.String("Action", hexFormatUint64(p.Action)))
	} else {
		fields = append(fields, zap.String("Method", method.Method), zap.String("Action", method.ActionHex))
	}
	if status != StatusOk {
		fields = append(fields, zap.String("Reason", status.Error()+": "+string(res)))
	}
	if err != nil {
		fields = append(fields, zap.String("Fail", err.Error()))
	}
	if status == StatusOk {
		c.opts.logger.Debug("packet", fields...)
	} else {
		c.opts.logger.Warn("packet error", fields...)
	}
}

func (c *Conn) handleRequest(p *Packet) {
	status, res := c.opts.router.dispatch(c, p)
	var err error
	if p.Kind == MessageKindRequest {
		err = c.send(context.Background(), &Packet{Kind: MessageKindResponse, Payload: res}, true)
	}
	c.logRequest(p, status, res, err)
}

func (c *Conn) handleResponse(p *Packet) {
	c.Lock()
	q, ok := c.requests[p.Mid]
	if ok {
		delete(c.requests, p.Mid)
		q <- p
	}
	c.Unlock()
}

func (c *Conn) negotiate() error {
	mediaInput := c.ClientHeader.Get(AcceptContentType)
	for {
		inputValue, inputLength := readNegotiate(mediaInput)
		if c.Media = GetMediaCodec(inputValue); c.Media != nil {
			c.ServerHeader.Set(DetermineContentType, inputValue)
			break
		}
		if inputLength == len(mediaInput) {
			break
		}
		mediaInput = mediaInput[inputLength:]
	}
	if c.Media == nil {
		return NewStatusError(StatusBadRequest, "no supported content-type in candidates: "+c.ClientHeader.Get(AcceptContentType))
	}
	encodingInput := c.ClientHeader.Get(AcceptEncoding)
	var encoding EncodingCodec
	for {
		inputValue, inputLength := readNegotiate(encodingInput)
		if encoding = GetEncodingCodec(inputValue); encoding != nil {
			c.ServerHeader.Set(DetermineEncoding, inputValue)
			break
		}
		if inputLength == len(encodingInput) {
			break
		}
		encodingInput = encodingInput[inputLength:]
	}
	packetFormat := c.ClientHeader.Get(DeterminePacketFormat)
	if packetFormat != "" {
		c.ServerHeader.Set(DeterminePacketFormat, packetFormat)
	}
	return nil
}

func (c *Conn) initEncoding() StatusError {
	contentType := c.ServerHeader.Get(DetermineContentType)
	c.Media = GetMediaCodec(contentType)
	if c.Media == nil {
		return NewStatusError(StatusBadRequest, "unsupported content-type: "+contentType)
	}
	if c.ServerHeader.Has(DetermineEncoding) {
		encoding := c.ServerHeader.Get(DetermineEncoding)
		if GetEncodingCodec(encoding) == nil {
			// this may never occur
			return NewStatusError(StatusBadRequest, "unsupported encoding: "+encoding)
		}
	}
	return nil
}
