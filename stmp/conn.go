// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-26 18:16:34
package stmp

import (
	"encoding/binary"
	"errors"
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
	Major            byte
	Minor            byte
	ClientHeader     Header
	ServerHeader     Header
	HandshakeMessage string
	nc               net.Conn
	wc               wsConn
	media            MediaCodec
	// send a nil well stop write
	writeChan chan []byte
	b1        []byte
	b2        []byte

	handshakeTimeout time.Duration
	readTimeout      time.Duration
	writeTimeout     time.Duration

	// pending requests & responses
	requests        map[uint16]chan *responseEvent
	pendingIncoming map[uint16]*incomingEvent
	pendingResponse map[uint16]*responseEvent
}

func (n *Conn) Request(options SendContext) error {
	panic("implement me")
}

func (n *Conn) Notify(options SendContext) error {
	panic("implement me")
}

func (n *Conn) Close(status Status, message string) error {
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
func (n *Conn) readUvarint() (uint64, error) {
	var x uint64
	var s uint
	for i := 0; ; i++ {
		_, err := n.nc.Read(n.b1)
		if err != nil {
			return x, err
		}
		b := n.b1[0]
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

func (n *Conn) readUint16() (v uint16, err error) {
	_, err = n.nc.Read(n.b2)
	if err != nil {
		return
	}
	v = binary.LittleEndian.Uint16(n.b2)
	return
}
func (n *Conn) handshake(status Status, header Header, message string) error {
	// TODO
	return nil
}

func (n *Conn) checkPendingVolume() {
}
