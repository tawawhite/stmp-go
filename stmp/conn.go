// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-26 18:16:34
package stmp

import (
	"encoding/binary"
	"errors"
	"github.com/gorilla/websocket"
	"io"
	"net"
	"time"
)

type writeEvent struct {
	p *Packet
	r chan error
}

type Conn struct {
	// router to dispatch actions
	*Router
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
	nc net.Conn
	// content-type codec
	media MediaCodec
	// write signal
	writeChan chan *writeEvent
	// close state
	closeChan chan struct{}
	// pending requests & responses
	requests map[uint16]chan *writeEvent
}

func (c *Conn) Request(options SendContext) error {
	panic("implement me")
}

func (c *Conn) Notify(options SendContext) error {
	panic("implement me")
}

func (c *Conn) Close(status Status, message string) error {
	return nil
}

func newConn(nc net.Conn) *Conn {
	return &Conn{
		Major:     1,
		Minor:     0,
		nc:        nc,
		writeChan: make(chan *writeEvent),
		closeChan: make(chan struct{}),
		requests:  map[uint16]chan *writeEvent{},
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

func (c *Conn) binaryReadChannel(r io.ReadCloser, timeout time.Duration) {
	b1 := make([]byte, 1)
	b2 := make([]byte, 2)
	p := Packet{}
	for {
		c.nc.SetReadDeadline(time.Now().Add(timeout))
		err := p.ReadHead(r, b1)
		if err == nil {
			err = p.ReadMid(r, b2)
		}
		if err == nil {
			err = p.ReadAction(r, b1)
		}
		if err == nil {
			err = p.ReadStatus(r, b1)
		}
		if err == nil {
			err = p.ReadPayload(r, b1)
		}
		if err != nil {
			break
		}
		switch p.Kind {
		case MessageKindPing:
			c.handlePing()
		case MessageKindRequest:
			c.handleRequest(p.Mid, p.Action, p.Payload)
		case MessageKindNotify:
			c.handleNotify(p.Action, p.Payload)
		case MessageKindResponse:
			c.handleResponse(p.Mid, p.Status, p.Payload)
		case MessageKindClose:
			c.handleClose(p.Status, string(p.Payload))
		}
	}
}

func (c *Conn) binaryWriteChannel(w EncodingWriter, timeout time.Duration) {
	b1 := make([]byte, 1)
	b2 := make([]byte, 2)
	b10 := make([]byte, 10)
	var nvi int
	for {
		e, ok := <-c.writeChan
		if !ok {
			// active closed connection
			w.Close()
			return
		}
		c.nc.SetWriteDeadline(time.Now().Add(timeout))
		b1[0] = MessageKindPing << OffsetKind
		if isFin(e.p.Kind) || e.p.Fin {
			b1[0] |= MaskFin
		}
		if isHead(e.p.Kind) || len(e.p.Payload) == 0 {
			b1[0] |= MaskHead
		}
		w.Write(b1)
		switch e.p.Kind {
		case MessageKindClose:
			b1[0] = byte(e.p.Status)
			w.Write(b1)
		case MessageKindRequest:
			binary.LittleEndian.PutUint16(b2, e.p.Mid)
			w.Write(b2)
			nvi = binary.PutUvarint(b10, e.p.Action)
			w.Write(b10[0:nvi])
		case MessageKindNotify:
			if !e.p.Fin {
				binary.LittleEndian.PutUint16(b2, e.p.Mid)
				w.Write(b2)
			}
			nvi = binary.PutUvarint(b10, e.p.Action)
			w.Write(b10[0:nvi])
		case MessageKindResponse:
			binary.LittleEndian.PutUint16(b2, e.p.Mid)
			w.Write(b2)
			b1[0] = byte(e.p.Status)
			w.Write(b1)
		}
		if !isFin(e.p.Kind) && len(e.p.Payload) > 0 {
			nvi = binary.PutUvarint(b10, uint64(len(e.p.Payload)))
			w.Write(b10[0:nvi])
			w.Write(e.p.Payload)
		}
		e.r <- w.Flush()
	}
}

func (c *Conn) websocketBinaryReadChannel(wc *websocket.Conn) {
}

func (c *Conn) websocketBinaryWriteChannel(wc *websocket.Conn) {
}

func (c *Conn) websocketTextReadChannel(wc *websocket.Conn) {
}

func (c *Conn) websocketTextWriteChannel(wc *websocket.Conn) {
}

func (c *Conn) handlePing() {
}

func (c *Conn) handleClose(status Status, message string) {

}

func (c *Conn) handleNotify(action uint64, payload []byte) {
}

func (c *Conn) handleRequest(mid uint16, action uint64, payload []byte) {
}

func (c *Conn) handleResponse(mid uint16, status Status, payload []byte) {
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
		rw := plainEncoding{Conn: c.nc}
		r = rw
		w = rw
		return
	} else {
		r, err = ec.Reader(c.nc)
		if err != nil {
			return
		}
		w, err = ec.Writer(c.nc, compressLevel)
		if err != nil {
			r.Close()
			r = nil
			return
		}
		return
	}
}
