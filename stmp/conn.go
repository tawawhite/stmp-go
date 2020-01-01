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
	p *packet
	r chan error
}

type Conn struct {
	net.Conn
	// router to dispatch actions
	*Router
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
	// write signal
	writeChan chan *writeEvent
	// close state
	closeChan chan struct{}
	// pending requests & responses
	requests map[uint16]chan *packet

	mid uint32

	onClosed CloseHandlerFunc
}

func (c *Conn) SetCloseHandler(h CloseHandlerFunc) {
	c.onClosed = h
}

func (c *Conn) Request(options SendContext) (out []byte, err error) {
	payload, err := options.Marshal(c.media)
	if err != nil {
		return nil, err
	}
	mid := uint16(atomic.AddUint32(&c.mid, 1))
	we := &writeEvent{
		p: &packet{
			fin:     true,
			kind:    MessageKindRequest,
			mid:     mid,
			action:  options.Action,
			payload: payload,
		},
		r: make(chan error, 1),
	}
	r := make(chan *packet, 1)
	c.mu.Lock()
	c.requests[mid] = r
	c.mu.Unlock()
	c.writeChan <- we
	err = <-we.r
	if err != nil {
		c.mu.Lock()
		delete(c.requests, mid)
		c.mu.Unlock()
		err = NewStatusError(StatusNetworkError, err)
		return
	}
	p := <-r
	if p.status != StatusOk {
		err = NewStatusError(p.status, string(p.payload))
	}
	out = p.payload
	if out == nil {
		options.Output = nil
		return
	}
	err = c.media.Unmarshal(out, options.Output)
	return
}

func (c *Conn) Notify(options SendContext) (err error) {
	options.Marshal(c.media)
	payload, err := options.Marshal(c.media)
	if err != nil {
		return err
	}
	we := &writeEvent{
		p: &packet{
			fin:     true,
			kind:    MessageKindNotify,
			action:  options.Action,
			payload: payload,
		},
		r: make(chan error, 1),
	}
	c.writeChan <- we
	return <-we.r
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
		requests:  map[uint16]chan *packet{},
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

func (c *Conn) dispatchPacket(p *packet) {
	switch p.kind {
	case MessageKindPing:
		c.handlePing()
	case MessageKindPong:
		c.handlePong()
	case MessageKindRequest:
		go c.handleRequest(p.mid, p.action, p.payload)
	case MessageKindNotify:
		go c.handleNotify(p.action, p.payload)
	case MessageKindResponse:
		c.handleResponse(p.mid, p.status, p.payload)
	case MessageKindClose:
		c.handleClose(p.status, string(p.payload))
	}
}

func (c *Conn) binaryReadChannel(r io.ReadCloser, timeout time.Duration) {
	p := new(packet)
	buf := make([]byte, maxStreamHeadSize, maxStreamHeadSize)
	var err error
	for {
		c.Conn.SetReadDeadline(time.Now().Add(timeout))
		err = p.read(r, buf)
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
	buf := make([]byte, maxStreamHeadSize, maxStreamHeadSize)
	for {
		e, ok = <-c.writeChan
		if !ok {
			// TODO
			return
		}
		c.Conn.SetWriteDeadline(time.Now().Add(timeout))
		err = e.p.write(w, buf)
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
	p := new(packet)
	var err error
	var data []byte
	for {
		wc.SetReadDeadline(time.Now().Add(timeout))
		_, data, err = wc.ReadMessage()
		if err != nil {
			// TODO
			return
		}
		err = p.unmarshalBinary(data)
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
	buf := make([]byte, maxBinaryHeadSize, maxBinaryHeadSize)
	for {
		e, ok = <-c.writeChan
		if !ok {
			// TODO
			return
		}
		wc.SetWriteDeadline(time.Now().Add(timeout))
		data = e.p.marshalBinary(buf)
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
	p := new(packet)
	var err error
	var data []byte
	for {
		wc.SetReadDeadline(time.Now().Add(timeout))
		_, data, err = wc.ReadMessage()
		if err != nil {
			// TODO
			return
		}
		err = p.unmarshalText(data)
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
	buf := make([]byte, maxTextHeadSize, maxTextHeadSize)
	for {
		e, ok = <-c.writeChan
		if !ok {
			// TODO
			return
		}
		wc.SetWriteDeadline(time.Now().Add(timeout))
		data = e.p.marshalText(buf)
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
	c.writeChan <- &writeEvent{p: &packet{kind: MessageKindPong}}
}

func (c *Conn) handlePong() {
	// do nothing for deadline limited pong message receive rate
}

func (c *Conn) handleClose(status Status, message string) {
	// TODO close connection
}

func (c *Conn) handleNotify(action uint64, payload []byte) {
	ctx := WithConn(context.Background(), c)
	c.Dispatch(ctx, action, payload, c.media)
}

func (c *Conn) handleRequest(mid uint16, action uint64, payload []byte) {
	ctx := WithConn(context.Background(), c)
	status, payload := c.Dispatch(ctx, action, payload, c.media)
	we := &writeEvent{p: &packet{
		fin:     true,
		kind:    MessageKindResponse,
		mid:     mid,
		action:  action,
		status:  status,
		payload: payload,
	}}
	c.writeChan <- we
}

func (c *Conn) handleResponse(mid uint16, status Status, payload []byte) {
	c.mu.Lock()
	q, ok := c.requests[mid]
	delete(c.requests, mid)
	c.mu.Unlock()
	if ok {
		q <- &packet{status: status, payload: payload}
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
