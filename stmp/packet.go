// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-30 15:24:09
package stmp

import (
	"errors"
	"io"
)

type Packet struct {
	Fin  bool
	Kind byte
	Head bool
	// except ping, close
	Mid uint16
	// request
	Action uint64
	// response, close
	Status Status
	// except ping
	Payload []byte
	// 0 means it is not in pool
	// offset 1 with pool index
	pos int
}

var (
	invalidReservedHeadBits = errors.New("invalid reserved head bits")
	invalidHeadFlags        = errors.New("invalid head flags")
	invalidMessageKind      = errors.New("invalid message Kind")
)

func (p *Packet) ReadHead(r io.Reader, b1 []byte) error {
	_, err := r.Read(b1)
	if err != nil {
		return err
	}
	h := b1[0]
	if h&0b111 != 0 {
		return invalidReservedHeadBits
	}
	p.Fin = h&0x80 != 0
	p.Kind = (h >> 4) & 0b111
	p.Head = h&0b1000 != 0
	if !isKind(p.Kind) {
		return invalidMessageKind
	}
	if isFin(p.Kind) && !p.Fin {
		return invalidHeadFlags
	}
	if isHead(p.Kind) && !p.Head {
		return invalidHeadFlags
	}
	return nil
}

func (p *Packet) ReadMid(r io.Reader, b2 []byte) error {
	switch p.Kind {
	case MessageKindRequest, MessageKindResponse:
	default:
		return nil
	}
	var err error
	p.Mid, err = readUint16(r, b2)
	return err
}

func (p *Packet) ReadAction(r io.Reader, b1 []byte) error {
	switch p.Kind {
	case MessageKindRequest, MessageKindNotify:
	default:
		return nil
	}
	var err error
	p.Action, err = readUvarint(r, b1)
	return err
}

func (p *Packet) ReadStatus(r io.Reader, b1 []byte) error {
	switch p.Kind {
	case MessageKindResponse, MessageKindClose:
	default:
		return nil
	}
	_, err := r.Read(b1)
	if err != nil {
		return err
	}
	p.Status = Status(b1[0])
	return nil
}

func (p *Packet) ReadPayload(r io.Reader, b1 []byte) error {
	if p.Head || p.Kind == MessageKindPing {
		return nil
	}
	ps, err := readUvarint(r, b1)
	if err != nil {
		return err
	}
	p.Payload = make([]byte, ps)
	_, err = r.Read(p.Payload)
	return err
}
