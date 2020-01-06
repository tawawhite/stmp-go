package stmp

import (
	"context"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"io"
	"net"
	"sync"
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
}

// set custom logger, default is zap.NewProduction()
func (o *ConnOptions) WithLogger(logger *zap.Logger) *ConnOptions {
	o.logger = logger
	return o
}

// set custom router
func (o *ConnOptions) WithRouter(r *Router) *ConnOptions {
	o.router = r
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
	if no.router == nil {
		no.router = NewRouter()
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
		// ping timeout
		readTimeout:  time.Minute * 2,
		writeTimeout: time.Minute,
		router:       nil,
		// 16 mb
		maxPacketSize:  1 << 24,
		writeQueueSize: 8,
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
		Conn:      nc,
		opts:      opts,
		Major:     1,
		Minor:     0,
		mid:       new(uint32),
		writeChan: make(chan *writeEvent, opts.writeQueueSize),
		requests:  map[uint16]chan *Packet{},
	}
}

// Private: invoke a method arbitrary
func (c *Conn) Call(ctx context.Context, method string, payload []byte, opts *CallOptions) (out interface{}, err error) {
	action := ms.methods[method]
	p := &Packet{Fin: true, Kind: MessageKindRequest, Action: action, Payload: payload}
	we := &writeEvent{p: p}
	var r chan *Packet
	if opts.useNotify {
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
	out = ms.actions[action].output()
	err = c.Media.Unmarshal(p.Payload, out)
	return
}

// Private: invoke a method arbitrary, this method will marshal the input to payload bytes with conn's codec.
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

// error occurred, or received close message, or sent close message, interrupt the connection only
func (c *Conn) terminate() error {
	// TODO
	return nil
}

// close connection manually, this should be manipulated by Server or ClientConn
func (c *Conn) close(status Status, message string) error {
	// TODO
	return nil
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
	buf := make([]byte, maxStreamHeadSize, maxStreamHeadSize)
	var err error
	for {
		c.Conn.SetReadDeadline(time.Now().Add(c.opts.readTimeout))
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
	buf := make([]byte, maxStreamHeadSize, maxStreamHeadSize)
	for {
		e, ok = <-c.writeChan
		if !ok {
			// TODO
			return
		}
		c.Conn.SetWriteDeadline(time.Now().Add(c.opts.writeTimeout))
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
		wc.SetReadDeadline(time.Now().Add(c.opts.readTimeout))
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
	buf := make([]byte, maxBinaryHeadSize, maxBinaryHeadSize)
	for {
		e, ok = <-c.writeChan
		if !ok {
			// TODO
			return
		}
		wc.SetWriteDeadline(time.Now().Add(c.opts.writeTimeout))
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
		wc.SetReadDeadline(time.Now().Add(c.opts.readTimeout))
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
	buf := make([]byte, maxTextHeadSize, maxTextHeadSize)
	for {
		e, ok = <-c.writeChan
		if !ok {
			// TODO
			return
		}
		wc.SetWriteDeadline(time.Now().Add(c.opts.writeTimeout))
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
	c.opts.router.dispatch(ctx, action, payload, c.Media)
}

func (c *Conn) handleRequest(mid uint16, action uint64, payload []byte) {
	ctx := WithConn(context.Background(), c)
	status, payload := c.opts.router.dispatch(ctx, action, payload, c.Media)
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
		return NewStatusError(StatusUnsupportedContentType, "")
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
		mediaInput = mediaInput[inputLength:]
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
		err = NewStatusError(StatusUnsupportedContentType, "cannot find codec: "+c.ServerHeader.Get(DetermineContentType)+", please register it first")
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
