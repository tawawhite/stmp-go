package stmp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

type Packet struct {
	Fin     bool
	Kind    byte
	Mid     uint16
	Action  uint64
	Status  Status
	Payload []byte
}

var (
	invalidReservedHeadBits = errors.New("invalid reserved head bits")
	invalidHeadFlags        = errors.New("invalid head flags")
	invalidMessageKind      = errors.New("invalid message Kind")
	invalidPacketHead       = errors.New("invalid Packet head")
	invalidPacketMid        = errors.New("invalid Packet Mid")
	invalidPacketAction     = errors.New("invalid Packet Action")
	invalidPacketStatus     = errors.New("invalid Packet Status")
	invalidPacketPayload    = errors.New("invalid Packet Payload")
)

var emptyBytes = make([]byte, 0, 0)

func (p *Packet) UnmarshalHead(h byte) error {
	if h&0b111 != 0 {
		return invalidReservedHeadBits
	}
	p.Fin = h&0x80 != 0
	p.Kind = (h >> 4) & 0b111
	head := h&0b1000 != 0
	if !isValidKind(p.Kind) {
		return invalidMessageKind
	}
	if shouldAlwaysFinal(p.Kind) && !p.Fin {
		return invalidHeadFlags
	}
	if shouldHeadOnly(p.Kind) && !head {
		return invalidHeadFlags
	}
	if head {
		p.Payload = nil
	} else {
		p.Payload = emptyBytes
	}
	return nil
}

func (p *Packet) MarshalHead(buf []byte, ps bool) []byte {
	n := 1
	buf[0] = p.Kind << offsetKind
	if shouldAlwaysFinal(p.Kind) || p.Fin {
		buf[0] |= maskFin
	}
	if shouldHeadOnly(p.Kind) || len(p.Payload) == 0 {
		buf[0] |= maskHead
	}
	if hasMid(p.Kind) {
		binary.LittleEndian.PutUint16(buf[n:], p.Mid)
		n += 2
	}
	if hasAction(p.Kind) {
		n += binary.PutUvarint(buf[n:], p.Action)
	}
	if hasStatus(p.Kind) {
		buf[n] = byte(p.Status)
		n += 1
	}
	if ps && !shouldHeadOnly(p.Kind) && len(p.Payload) > 0 {
		n += binary.PutUvarint(buf[n:], uint64(len(p.Payload)))
	}
	return buf[:n]
}

func (p *Packet) Write(w EncodingWriter, buf []byte) (err error) {
	buf = p.MarshalHead(buf, true)
	if _, err = w.Write(buf); err != nil {
		return
	}
	if !shouldHeadOnly(p.Kind) && len(p.Payload) > 0 {
		if _, err = w.Write(p.Payload); err != nil {
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

func (p *Packet) Read(r io.ReadCloser, buf []byte) (err error) {
	if _, err = r.Read(buf[:1]); err != nil {
		return
	}
	if err = p.UnmarshalHead(buf[0]); err != nil {
		return
	}
	if hasMid(p.Kind) {
		if p.Mid, err = readUint16(r, buf[:2]); err != nil {
			return
		}
	}
	if hasAction(p.Kind) {
		if p.Action, err = readUvarint(r, buf[:1]); err != nil {
			return
		}
	}
	if hasStatus(p.Kind) {
		if _, err = r.Read(buf[:1]); err != nil {
			return
		}
		p.Status = Status(buf[0])
	}
	if p.Payload != nil {
		var ps uint64
		if ps, err = readUvarint(r, buf[:1]); err != nil {
			return
		}
		p.Payload = make([]byte, ps)
		_, err = r.Read(p.Payload)
	}
	return
}

func (p *Packet) MarshalBinary(buf []byte) []byte {
	buf = p.MarshalHead(buf, false)
	data := make([]byte, len(buf)+len(p.Payload), len(buf)+len(p.Payload))
	copy(data, buf)
	copy(data[len(buf):], p.Payload)
	return data
}

func (p *Packet) UnmarshalBinary(data []byte) (err error) {
	if len(data) < 1 {
		return invalidPacketHead
	}
	if err = p.UnmarshalHead(data[0]); err != nil {
		return
	}
	var n int
	data = data[1:]
	// Mid
	if hasMid(p.Kind) {
		if len(data) < 2 {
			return invalidPacketMid
		}
		p.Mid = binary.LittleEndian.Uint16(data)
		data = data[3:]
	}
	if hasAction(p.Kind) {
		if len(data) < 1 {
			return invalidPacketAction
		}
		p.Action, n = binary.Uvarint(data)
		if n <= 0 {
			return invalidPacketAction
		}
		data = data[n:]
	}
	if hasStatus(p.Kind) {
		if len(data) < 1 {
			return invalidPacketStatus
		}
		p.Status = Status(data[0])
		data = data[1:]
	}
	if shouldHeadOnly(p.Kind) || p.Payload == nil && len(data) > 0 {
		return invalidPacketPayload
	}
	p.Payload = data
	return nil
}

// PING: 		I
// PONG: 		O
// REQUEST: 	Q<MID>:<ACTION>[\nP] 	-> 1 + 4 + 1 + 16 + 1 -> 21
// NOTIFY: 		N<ACTION>[\nP]			-> 18
// RESPONSE: 	S<MID>:<STATUS>[\nP]	-> 9
// CLOSE:		C<STATUS>[\nP]			-> 4
func (p *Packet) MarshalText(buf []byte) []byte {
	n := 0
	buf[n] = mapKindText[p.Kind]
	n += 1
	if hasMid(p.Kind) {
		n += appendHex(uint64(p.Mid), buf[n:])
		buf[n] = ':'
		n += 1
	}
	if hasAction(p.Kind) {
		n += appendHex(p.Action, buf[n:])
	}
	if hasStatus(p.Kind) {
		n += appendHex(uint64(p.Status), buf[n:])
	}
	if !shouldHeadOnly(p.Kind) && len(p.Payload) > 0 {
		buf[n] = '\n'
		n += 1
	}
	data := make([]byte, n+len(p.Payload), n+len(p.Payload))
	copy(data, buf[:n])
	copy(data, p.Payload)
	return data
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (p *Packet) UnmarshalText(data []byte) (err error) {
	if len(data) == 0 {
		return invalidPacketHead
	}
	var ok bool
	if p.Kind, ok = mapTextKind[data[0]]; !ok {
		return invalidMessageKind
	}
	data = data[1:]
	var i int
	if hasMid(p.Kind) {
		i = bytes.IndexByte(data[:min(17, len(data))], ':')
		if i == -1 {
			return invalidPacketMid
		}
		if p.Mid, err = parseHexUint16(data[:i]); err != nil {
			err = errors.New(invalidPacketMid.Error() + ": " + err.Error())
			return
		}
		data = data[i+1:]
	}
	if hasAction(p.Kind) {
		i = bytes.IndexByte(data[:min(17, len(data))], '\n')
		if i == -1 {
			i = len(data)
		}
		if p.Action, err = parseHexUint64(data[:i]); err != nil {
			err = errors.New(invalidPacketAction.Error() + ": " + err.Error())
			return
		}
		data = data[i:]
	}
	if hasStatus(p.Kind) {
		if len(data) < 1 {
			return invalidPacketStatus
		}
		p.Status = Status(data[0])
		data = data[1:]
	}
	if shouldHeadOnly(p.Kind) && len(data) > 0 {
		return invalidPacketPayload
	}
	p.Payload = data[1:]
	return
}
