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

type incomingEvent struct {
	kind    MessageKind
	action  uint64
	payload []byte
}

type responseEvent struct {
	status  Status
	payload []byte
}

type wsConn interface {
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
}

type Conn struct {
	*Router

	Major byte
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

	compressLevel int

	// send a nil well stop write
	writeChan chan []byte

	b1 []byte
	b2 []byte

	handshakeTimeout time.Duration
	readTimeout      time.Duration
	writeTimeout     time.Duration

	// pending requests & responses
	requests        map[uint16]chan *responseEvent
	pendingIncoming map[uint16]*incomingEvent
	pendingResponse map[uint16]*responseEvent
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
		Major:           1,
		Minor:           0,
		nc:              nc,
		b1:              make([]byte, 1),
		b2:              make([]byte, 2),
		writeChan:       make(chan []byte),
		requests:        map[uint16]chan *responseEvent{},
		pendingIncoming: map[uint16]*incomingEvent{},
		pendingResponse: map[uint16]*responseEvent{},
	}
}

// for binary.ReadUvarint
func (c *Conn) readUvarint() (uint64, error) {
	var x uint64
	var s uint
	for i := 0; ; i++ {
		_, err := c.nc.Read(c.b1)
		if err != nil {
			return x, err
		}
		b := c.b1[0]
		if b < 0x80 {
			if i > 9 || i == 9 && b > 1 {
				return x, errors.New("uint64 overflow")
			}
			return x | uint64(b)<<s, nil
		}
		x |= uint64(b&0x7f) << s
		s += 7
	}
}

func (c *Conn) readUint16() (v uint16, err error) {
	_, err = c.nc.Read(c.b2)
	if err != nil {
		return
	}
	v = binary.LittleEndian.Uint16(c.b2)
	return
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

func (c *Conn) binaryReadChannel(r io.ReadCloser) {
}

func (c *Conn) binaryWriteChannel(w EncodingWriter) {
}

func (c *Conn) websocketBinaryReadChannel(wc *websocket.Conn) {
}

func (c *Conn) websocketBinaryWriteChannel(wc *websocket.Conn) {
}

func (c *Conn) websocketTextReadChannel(wc *websocket.Conn) {
}

func (c *Conn) websocketTextWriteChannel(wc *websocket.Conn) {
}

func (c *Conn) handleNotify(action uint64, payload []byte) {
}

func (c *Conn) handleRequest(mid uint16, action uint64, payload []byte) {
}

func (c *Conn) handleResponse(mid uint16, status byte, payload []byte) {
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
