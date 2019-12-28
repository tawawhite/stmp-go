// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-26 18:16:34
package stmp

import (
	"encoding/binary"
	"net"
)

type Conn interface {
	Header() Header
	RemoteAddr() net.Addr
	LocalAddr() net.Addr
	Request(options SendContext) error
	Notify(options SendContext) error
	Terminate(status Status, message string) error
}

type incomingEvent struct {
	kind    MessageKind
	action  uint64
	payload []byte
}

type responseEvent struct {
	status  Status
	payload []byte
}

type netConn struct {
	net.Conn
	header Header
	media  MediaCodec
	major  byte
	minor  byte
	// send a nil well stop write
	writeChan       chan []byte
	b1              []byte
	b2              []byte
	requests        map[uint16]chan *responseEvent
	pendingIncoming map[uint16]*incomingEvent
	pendingResponse map[uint16]*responseEvent
}

func (n *netConn) Flush() error {
	return nil
}

func (n *netConn) Header() Header {
	return n.header
}

func (n *netConn) Request(options SendContext) error {
	panic("implement me")
}

func (n *netConn) Notify(options SendContext) error {
	panic("implement me")
}

func (n *netConn) ReadByte() (byte, error) {
	_, err := n.Read(n.b1)
	return n.b1[0], err
}

func (n *netConn) ReadInt16() (v uint16, err error) {
	_, err = n.Read(n.b2)
	if err != nil {
		return
	}
	v = binary.LittleEndian.Uint16(n.b2)
	return
}

func (n *netConn) Handshake(status Status, header Header, message string) error {
	// TODO
	return nil
}

func (n *netConn) Terminate(status Status, message string) error {
	return nil
}

func (n *netConn) CheckPendingVolume() {
}

func newNetConn(conn net.Conn) *netConn {
	return &netConn{
		Conn:   conn,
		b1:     make([]byte, 1),
		b2:     make([]byte, 2),
		header: NewHeader(),
	}
}
