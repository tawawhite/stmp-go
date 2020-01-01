// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-30 15:24:09
package stmp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

type packet struct {
	fin  bool
	kind byte
	// except ping, close
	mid uint16
	// request
	action uint64
	// response, close
	status Status
	// except ping
	payload []byte
}

var (
	invalidReservedHeadBits = errors.New("invalid reserved head bits")
	invalidHeadFlags        = errors.New("invalid head flags")
	invalidMessageKind      = errors.New("invalid message kind")
	invalidPacketHead       = errors.New("invalid packet head")
	invalidPacketMid        = errors.New("invalid packet mid")
	invalidPacketAction     = errors.New("invalid packet action")
	invalidPacketStatus     = errors.New("invalid packet status")
	invalidPacketPayload    = errors.New("invalid packet payload")
)

var emptyBytes = make([]byte, 0, 0)

func (p *packet) unmarshalHead(h byte) error {
	if h&0b111 != 0 {
		return invalidReservedHeadBits
	}
	p.fin = h&0x80 != 0
	p.kind = (h >> 4) & 0b111
	head := h&0b1000 != 0
	if !isKind(p.kind) {
		return invalidMessageKind
	}
	if isFin(p.kind) && !p.fin {
		return invalidHeadFlags
	}
	if isHead(p.kind) && !head {
		return invalidHeadFlags
	}
	if head {
		p.payload = nil
	} else {
		p.payload = emptyBytes
	}
	return nil
}

func (p *packet) marshalHead(buf []byte, ps bool) []byte {
	n := 1
	buf[0] = p.kind << OffsetKind
	if isFin(p.kind) || p.fin {
		buf[0] |= MaskFin
	}
	if isHead(p.kind) || len(p.payload) == 0 {
		buf[0] |= MaskHead
	}
	if isMid(p.kind) {
		binary.LittleEndian.PutUint16(buf[n:], p.mid)
		n += 2
	}
	if isAction(p.kind) {
		n += binary.PutUvarint(buf[n:], p.action)
	}
	if isStatus(p.kind) {
		buf[n] = byte(p.status)
		n += 1
	}
	if ps && !isHead(p.kind) && len(p.payload) > 0 {
		n += binary.PutUvarint(buf[n:], uint64(len(p.payload)))
	}
	return buf[:n]
}

func (p *packet) write(w EncodingWriter, buf []byte) (err error) {
	buf = p.marshalHead(buf, true)
	if _, err = w.Write(buf); err != nil {
		return
	}
	if !isHead(p.kind) && len(p.payload) > 0 {
		if _, err = w.Write(p.payload); err != nil {
			return
		}
	}
	return w.Flush()
}

// PING/PONG: <HEAD> -> 1
// REQUEST: <HEAD><MID><ACTION><PS><P> 		-> 23
// NOTIFY: <HEAD><ACTION><PS><P> 			-> 21
// RESPONSE: <HEAD><MID><STATUS><PS><P> 	-> 22
// CLOSE: <HEAD><STATUS><PS><P>				-> 12

const maxStreamHeadSize = 23
const maxBinaryHeadSize = 13
const maxTextHeadSize = 21

func (p *packet) read(r io.ReadCloser, buf []byte) (err error) {
	if _, err = r.Read(buf[:1]); err != nil {
		return
	}
	if err = p.unmarshalHead(buf[0]); err != nil {
		return
	}
	if isMid(p.kind) {
		if p.mid, err = ReadUint16(r, buf[:2]); err != nil {
			return
		}
	}
	if isAction(p.kind) {
		if p.action, err = ReadUvarint(r, buf[:1]); err != nil {
			return
		}
	}
	if isStatus(p.kind) {
		if _, err = r.Read(buf[:1]); err != nil {
			return
		}
		p.status = Status(buf[0])
	}
	if p.payload != nil {
		var ps uint64
		if ps, err = ReadUvarint(r, buf[:1]); err != nil {
			return
		}
		p.payload = make([]byte, ps)
		_, err = r.Read(p.payload)
	}
	return
}

func (p *packet) marshalBinary(buf []byte) []byte {
	buf = p.marshalHead(buf, false)
	data := make([]byte, len(buf)+len(p.payload), len(buf)+len(p.payload))
	copy(data, buf)
	copy(data[len(buf):], p.payload)
	return data
}

func (p *packet) unmarshalBinary(data []byte) (err error) {
	if len(data) < 1 {
		return invalidPacketHead
	}
	if err = p.unmarshalHead(data[0]); err != nil {
		return
	}
	var n int
	data = data[1:]
	// mid
	if isMid(p.kind) {
		if len(data) < 2 {
			return invalidPacketMid
		}
		p.mid = binary.LittleEndian.Uint16(data)
		data = data[3:]
	}
	if isAction(p.kind) {
		if len(data) < 1 {
			return invalidPacketAction
		}
		p.action, n = binary.Uvarint(data)
		if n <= 0 {
			return invalidPacketAction
		}
		data = data[n:]
	}
	if isStatus(p.kind) {
		if len(data) < 1 {
			return invalidPacketStatus
		}
		p.status = Status(data[0])
		data = data[1:]
	}
	if isHead(p.kind) || p.payload == nil && len(data) > 0 {
		return invalidPacketPayload
	}
	p.payload = data
	return nil
}

// PING: 		I
// PONG: 		O
// REQUEST: 	Q<MID>:<ACTION>[\nP] 	-> 1 + 4 + 1 + 16 + 1 -> 21
// NOTIFY: 		N<ACTION>[\nP]			-> 18
// RESPONSE: 	S<MID>:<STATUS>[\nP]	-> 9
// CLOSE:		C<STATUS>[\nP]			-> 4
func (p *packet) marshalText(buf []byte) []byte {
	n := 0
	buf[n] = MapKindText[p.kind]
	n += 1
	if isMid(p.kind) {
		n += AppendHex(uint64(p.mid), buf[n:])
		buf[n] = ':'
		n += 1
	}
	if isAction(p.kind) {
		n += AppendHex(p.action, buf[n:])
	}
	if isStatus(p.kind) {
		n += AppendHex(uint64(p.status), buf[n:])
	}
	if !isHead(p.kind) && len(p.payload) > 0 {
		buf[n] = '\n'
		n += 1
	}
	data := make([]byte, n+len(p.payload), n+len(p.payload))
	copy(data, buf[:n])
	copy(data, p.payload)
	return data
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (p *packet) unmarshalText(data []byte) (err error) {
	if len(data) == 0 {
		return invalidPacketHead
	}
	var ok bool
	if p.kind, ok = MapTextKind[data[0]]; !ok {
		return invalidMessageKind
	}
	data = data[1:]
	var i int
	if isMid(p.kind) {
		i = bytes.IndexByte(data[:min(17, len(data))], ':')
		if i == -1 {
			return invalidPacketMid
		}
		if p.mid, err = ParseHexUint16(data[:i]); err != nil {
			err = errors.New(invalidPacketMid.Error() + ": " + err.Error())
			return
		}
		data = data[i+1:]
	}
	if isAction(p.kind) {
		i = bytes.IndexByte(data[:min(17, len(data))], '\n')
		if i == -1 {
			i = len(data)
		}
		if p.action, err = ParseHexUint64(data[:i]); err != nil {
			err = errors.New(invalidPacketAction.Error() + ": " + err.Error())
			return
		}
		data = data[i:]
	}
	if isStatus(p.kind) {
		if len(data) < 1 {
			return invalidPacketStatus
		}
		p.status = Status(data[0])
		data = data[1:]
	}
	if isHead(p.kind) && len(data) > 0 {
		return invalidPacketPayload
	}
	p.payload = data[1:]
	return
}
