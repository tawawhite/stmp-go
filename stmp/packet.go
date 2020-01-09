package stmp

import (
	"bytes"
	"encoding/binary"
	"io"
	"strconv"
)

type Packet struct {
	Kind         byte
	WithPayload  bool
	WithHeader   bool
	StringAction bool
	Status       Status
	Mid          uint16
	Action       uint64
	Method       string
	Header       Header
	Payload      []byte
}

func NewClosePacket(status Status, message string) *Packet {
	return &Packet{Kind: MessageKindClose, Status: status, Payload: []byte(message)}
}

var (
	invalidReservedHeadBits = NewStatusError(StatusProtocolError, "invalid reserved head bits")
	invalidHeadFlags        = NewStatusError(StatusProtocolError, "invalid head flags")
	invalidMessageKind      = NewStatusError(StatusProtocolError, "invalid message kind")
	invalidPacketHead       = NewStatusError(StatusProtocolError, "invalid packet head")
	invalidPacketMid        = NewStatusError(StatusProtocolError, "invalid packet mid")
	invalidPacketAction     = NewStatusError(StatusProtocolError, "invalid packet action")
	invalidPacketStatus     = NewStatusError(StatusProtocolError, "invalid packet status")
	invalidPacketPayload    = NewStatusError(StatusProtocolError, "invalid packet payload")
)

func (p *Packet) HasMid() bool {
	return p.Kind == MessageKindRequest || p.Kind == MessageKindResponse
}

func (p *Packet) HasAction() bool {
	return p.Kind == MessageKindRequest || p.Kind == MessageKindNotify
}

func (p *Packet) HasStatus() bool {
	return p.Kind == MessageKindResponse || p.Kind == MessageKindClose
}

func (p *Packet) HasHeader() bool {
	return p.Kind == MessageKindRequest || p.Kind == MessageKindNotify || p.Kind == MessageKindResponse
}

func (p *Packet) HasPayload() bool {
	return p.Kind != MessageKindPing && p.Kind != MessageKindPong
}

func (p *Packet) BinarySize(ps bool) (header []byte, size int) {
	size = 1
	if p.HasMid() {
		size += 2
	}
	if p.HasStatus() {
		size += 1
	}
	if p.HasAction() {
		if p.StringAction {
			size += uvarintSize(uint64(len(p.Method))) + len(p.Method)
		} else {
			size += uvarintSize(p.Action)
		}
	}
	if p.HasHeader() {
		header = p.Header.Marshal()
		if len(header) > 0 {
			size += uvarintSize(uint64(len(header))) + len(header)
		}
	}
	if ps && p.HasPayload() && len(p.Payload) > 0 {
		size += uvarintSize(uint64(len(p.Payload)))
	}
	return
}

func (p *Packet) MarshalHead(ps bool) []byte {
	header, size := p.BinarySize(ps)
	head := 0x80 | (p.Kind << 4)
	if p.HasPayload() && len(p.Payload) > 0 {
		head |= 0b1000
	}
	if p.HasHeader() && len(header) > 0 {
		head |= 0b100
	}
	if p.HasAction() && p.StringAction {
		head |= 0b10
	}
	data := make([]byte, size, size)
	data[0] = head
	n := 1
	if p.HasMid() {
		binary.LittleEndian.PutUint16(data[n:], p.Mid)
		n += 2
	}
	if p.HasStatus() {
		data[n] = byte(p.Status)
		n += 1
	}
	if p.HasAction() {
		if p.StringAction {
			n += binary.PutUvarint(data[n:], uint64(len(p.Method)))
			copy(data[n:], p.Method)
			n += len(p.Method)
		} else {
			n += binary.PutUvarint(data[n:], p.Action)
		}
	}
	if len(header) > 0 {
		n += binary.PutUvarint(data[n:], uint64(len(header)))
		copy(data[n:], header)
		n += len(header)
	}
	if ps && p.HasPayload() && len(p.Payload) > 0 {
		binary.PutUvarint(data[n:], uint64(len(p.Payload)))
		n += len(p.Payload)
	}
	return data
}

func (p *Packet) UnmarshalHead(h byte) StatusError {
	if h&0x80 == 0 || h&1 != 0 {
		return invalidReservedHeadBits
	}
	p.Kind = (h >> 4) & 0b111
	if !isValidKind(p.Kind) {
		return invalidMessageKind
	}
	p.WithPayload = h&0b1000 != 0
	if !p.HasPayload() && p.WithPayload {
		return invalidHeadFlags
	}
	p.WithHeader = h&0b100 != 0
	if !p.HasHeader() && p.WithHeader {
		return invalidHeadFlags
	}
	p.StringAction = h&0b10 != 0
	if !p.HasAction() && p.StringAction {
		return invalidHeadFlags
	}
	return nil
}

func (p *Packet) Write(w EncodingWriter) StatusError {
	var err error
	head := p.MarshalHead(true)
	if _, err = w.Write(head); err != nil {
		return NewStatusError(StatusNetworkError, "write packet error: "+err.Error())
	}
	if p.HasPayload() && len(p.Payload) > 0 {
		if _, err = w.Write(p.Payload); err != nil {
			return NewStatusError(StatusNetworkError, "write packet payload error: "+err.Error())
		}
	}
	if err = w.Flush(); err != nil {
		return NewStatusError(StatusNetworkError, "flush packet error: "+err.Error())
	}
	return nil
}

// PING/PONG: <HEAD> -> 1
// REQUEST: <HEAD><MID><ACTION><PS><P> 		-> 23
// NOTIFY: <HEAD><ACTION><PS><P> 			-> 21
// RESPONSE: <HEAD><MID><STATUS><PS><P> 	-> 22
// CLOSE: <HEAD><STATUS><PS><P>				-> 12

func (p *Packet) Read(r io.Reader, maxPacketSize uint64) StatusError {
	var err error
	buf := make([]byte, 2)
	if _, err = r.Read(buf[:1]); err != nil {
		return NewStatusError(StatusNetworkError, "read packet header: "+err.Error())
	}
	if se := p.UnmarshalHead(buf[0]); se != nil {
		return se
	}
	if p.HasMid() {
		if p.Mid, err = readUint16(r, buf[:2]); err != nil {
			return NewStatusError(StatusNetworkError, "read packet mid: "+err.Error())
		}
	}
	if p.HasStatus() {
		if _, err = r.Read(buf[:1]); err != nil {
			return NewStatusError(StatusNetworkError, "read packet status: "+err.Error())
		}
		p.Status = Status(buf[0])
	}
	var ms uint64
	if p.StringAction {
		if ms, err = readUvarint(r, buf[:1]); err != nil {
			return NewStatusError(StatusNetworkError, "read packet method size: "+err.Error())
		}
		if ms > maxPacketSize {
			return NewStatusError(StatusRequestEntityTooLarge, "packet size "+strconv.Itoa(int(ms))+" is large than "+strconv.Itoa(int(maxPacketSize)))
		}
		mb := make([]byte, ms)
		if _, err = r.Read(mb); err != nil {
			return NewStatusError(StatusNetworkError, "read packet method: "+err.Error())
		}
		p.Method = string(mb)
	} else if p.HasAction() {
		if p.Action, err = readUvarint(r, buf[:1]); err != nil {
			return NewStatusError(StatusNetworkError, "read packet action: "+err.Error())
		}
	}
	var hs uint64
	if p.WithHeader {
		if hs, err = readUvarint(r, buf[:1]); err != nil {
			return NewStatusError(StatusNetworkError, "read packet header size: "+err.Error())
		}
		if hs+ms > maxPacketSize {
			return NewStatusError(StatusRequestEntityTooLarge, "packet size "+strconv.Itoa(int(ms+hs))+" is large than "+strconv.Itoa(int(maxPacketSize)))
		}
		hb := make([]byte, hs)
		if _, err = r.Read(hb); err != nil {
			return NewStatusError(StatusNetworkError, "read packet header: "+err.Error())
		}
		p.Header = NewHeader()
		if err = p.Header.Unmarshal(hb); err != nil {
			return NewStatusError(StatusBadRequest, "parse packet header: "+err.Error())
		}
	}
	var ps uint64
	if p.WithPayload {
		if ps, err = readUvarint(r, buf[:1]); err != nil {
			return NewStatusError(StatusNetworkError, "read packet payload size: "+err.Error())
		}
		if hs+ms+ps > maxPacketSize {
			return NewStatusError(StatusRequestEntityTooLarge, "packet size "+strconv.Itoa(int(ms+hs+ps))+" is large than "+strconv.Itoa(int(maxPacketSize)))
		}
		p.Payload = make([]byte, ps)
		if _, err = r.Read(p.Payload); err != nil {
			return NewStatusError(StatusNetworkError, "read packet payload: "+err.Error())
		}
	}
	return nil
}

func (p *Packet) MarshalBinary() []byte {
	head := p.MarshalHead(false)
	data := make([]byte, len(head)+len(p.Payload), len(head)+len(p.Payload))
	copy(data, head)
	copy(data[len(head):], p.Payload)
	return data
}

func (p *Packet) UnmarshalBinary(data []byte) StatusError {
	if len(data) < 1 {
		return invalidPacketHead
	}
	if se := p.UnmarshalHead(data[0]); se != nil {
		return se
	}
	var n int
	data = data[1:]
	// Mid
	if p.HasMid() {
		if len(data) < 2 {
			return invalidPacketMid
		}
		p.Mid = binary.LittleEndian.Uint16(data)
		data = data[2:]
	}
	if p.HasStatus() {
		if len(data) < 1 {
			return invalidPacketStatus
		}
		p.Status = Status(data[0])
		data = data[1:]
	}
	if p.StringAction {
		var ms uint64
		if ms, n = binary.Uvarint(data); n <= 0 || len(data) < n+int(ms) {
			return NewStatusError(StatusProtocolError, "invalid packet method size")
		}
		p.Method = string(data[n : n+int(ms)])
		data = data[n+int(ms):]
	} else {
		if p.HasAction() {
			if len(data) < 1 {
				return invalidPacketAction
			}
			if p.Action, n = binary.Uvarint(data); n <= 0 {
				return NewStatusError(StatusProtocolError, "invalid packet action")
			}
			data = data[n:]
		}
	}
	if p.WithHeader {
		var hs uint64
		if hs, n = binary.Uvarint(data); n <= 0 || len(data) < n+int(hs) {
			return NewStatusError(StatusProtocolError, "invalid packet header size")
		}
		p.Header = NewHeader()
		if err := p.Header.Unmarshal(data[n : n+int(hs)]); err != nil {
			return NewStatusError(StatusProtocolError, "invalid packet header: "+err.Error())
		}
		data = data[n+int(hs):]
	}
	if p.WithPayload {
		p.Payload = data
		data = nil
	}
	if len(data) > 0 {
		return NewStatusError(StatusProtocolError, "invalid packet payload")
	}
	return nil
}

func (p *Packet) MarshalText() []byte {
	buf := []byte{mapKindText[p.Kind]}
	if p.HasMid() {
		buf = append(buf, hexFormatUint64(uint64(p.Mid))...)
		buf = append(buf, ':')
	}
	if p.HasStatus() {
		buf = append(buf, hexFormatUint64(uint64(p.Status))...)
	}
	if p.HasAction() {
		if p.StringAction {
			buf = append(buf, 'M', ':')
			buf = append(buf, escape(p.Method)...)
		} else {
			buf = append(buf, hexFormatUint64(p.Action)...)
		}
	}
	if p.HasHeader() {
		header := p.Header.Marshal()
		if len(header) > 0 {
			buf = append(buf, '\n', 'H')
			buf = append(buf, header...)
		}
	}
	if p.HasPayload() && len(p.Payload) > 0 {
		buf = append(buf, '\n', 'P')
		buf = append(buf, p.Payload...)
	}
	return buf
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (p *Packet) UnmarshalText(data []byte) StatusError {
	var err error
	if len(data) == 0 {
		return invalidPacketHead
	}
	var ok bool
	if p.Kind, ok = mapTextKind[data[0]]; !ok {
		return invalidMessageKind
	}
	data = data[1:]
	var i int
	if p.HasMid() {
		if i = bytes.IndexByte(data[:min(5, len(data))], ':'); i == -1 {
			return invalidPacketMid
		}
		if p.Mid, err = hexParseUint16(data[:i]); err != nil {
			return NewStatusError(StatusProtocolError, "invalid packet mid: "+err.Error())
		}
		data = data[i+1:]
	}
	if p.HasStatus() {
		i = bytes.IndexByte(data, '\n')
		if i == -1 {
			i = len(data)
		}
		if p.Status, err = hexParseStatus(data[:i]); err != nil {
			return NewStatusError(StatusProtocolError, "invalid packet status: "+err.Error())
		}
		if i == len(data) {
			data = data[0:0]
		} else {
			data = data[i+1:]
		}
	}
	if p.HasAction() {
		i = bytes.IndexByte(data, '\n')
		if i == -1 {
			i = len(data)
		}
		if len(data) > 2 && data[0] == 'M' && data[1] == ':' {
			p.Method = string(data[2:i])
			p.StringAction = true
		} else if p.Action, err = hexParseUint64(data[:i]); err != nil {
			return NewStatusError(StatusProtocolError, "invalid packet action: "+err.Error())
		}
		if i == len(data) {
			data = data[0:0]
		} else {
			data = data[i+1:]
		}
	}
	if p.HasHeader() && len(data) > 0 && data[0] == 'H' {
		p.WithHeader = true
		i = bytes.IndexByte(data, '\n')
		if i == -1 {
			i = len(data)
		}
		p.Header = NewHeader()
		if err = p.Header.Unmarshal(data[1:i]); err != nil {
			return NewStatusError(StatusProtocolError, "invalid packet header: "+err.Error())
		}
		if i == len(data) {
			data = data[0:0]
		} else {
			data = data[i+1:]
		}
	}
	if p.HasPayload() && len(data) > 0 && data[0] == 'P' {
		p.WithPayload = true
		p.Payload = data[1:]
		data = nil
	}
	if len(data) > 0 {
		return NewStatusError(StatusProtocolError, "invalid packet payload")
	}
	return nil
}
