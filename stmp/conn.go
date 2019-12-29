// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-26 18:16:34
package stmp

import (
	"encoding/binary"
	"errors"
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
	// client handshake request header
	ClientHeader Header
	// server handshake response header
	ServerHeader Header
	// client handshake request message
	ClientMessage string
	// server handshake response message
	ServerMessage string

	// network conn
	nc net.Conn
	// websocket conn, use for websocket connection
	wc wsConn

	reader io.ReadCloser
	writer EncodingWriter

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

func newConn() *Conn {
	return &Conn{
		Major:           1,
		Minor:           0,
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
func (c *Conn) handshake(status Status, header Header, message string) error {
	// TODO
	return nil
}

func (c *Conn) checkPendingVolume() {
}

func (c *Conn) initEncoding(compressLevel int) error {
	c.media = GetMediaCodec(c.ServerHeader.Get(DetermineContentType))
	if c.media == nil {
		return errors.New("cannot find the codec for content-type: " + c.ServerHeader.Get(DetermineContentType) + ", please register it at first")
	}
	ec := GetEncodingCodec(c.ServerHeader.Get(DetermineEncoding))
	if ec == nil {
		rw := plainEncoding{Conn: c.nc}
		c.reader = rw
		c.writer = rw
		return nil
	} else {
		r, err := ec.Reader(c.nc)
		if err != nil {
			return err
		}
		w, err := ec.Writer(c.nc, compressLevel)
		if err != nil {
			r.Close()
			return err
		}
		c.reader = r
		c.writer = w
		return nil
	}
}
