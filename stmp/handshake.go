package stmp

import (
	"bytes"
	"encoding/binary"
	"io"
	"strconv"
)

const (
	HandshakeKindServer byte = 0x01
	HandshakeKindClient byte = 0x02
)

type Handshake struct {
	Kind    byte
	Major   byte
	Minor   byte
	Status  Status
	Header  Header
	Message string
}

func NewServerHandshake(status Status, header Header, message string) *Handshake {
	if header == nil {
		header = NewHeader()
	}
	return &Handshake{
		Kind:    HandshakeKindServer,
		Status:  status,
		Header:  header,
		Message: message,
	}
}

func NewClientHandshake(major, minor byte, header Header, message string) *Handshake {
	if header == nil {
		header = NewHeader()
	}
	return &Handshake{
		Kind:    HandshakeKindClient,
		Major:   major,
		Minor:   minor,
		Header:  header,
		Message: message,
	}
}

func (h *Handshake) MarshalHead() (title []byte, header []byte) {
	header = h.Header.Marshal()
	title = make([]byte, 15, 15)
	copy(title, "STMP")
	if h.Kind == HandshakeKindServer {
		title[4] = byte(h.Status)
	}
	if h.Kind == HandshakeKindClient {
		title[4] = h.Major<<4 | h.Minor&0xF
	}
	ps := len(header)
	if len(h.Message) > 0 {
		ps += 2 + len(h.Message)
	}
	i := binary.PutUvarint(title[5:], uint64(ps))
	return title[:i+5], header
}

func (h *Handshake) Read(r io.Reader, limit uint64) StatusError {
	title := make([]byte, 5)
	if _, err := r.Read(title); err != nil {
		return NewStatusError(StatusNetworkError, "read handshake title: "+err.Error())
	}
	if !bytes.Equal(title[:4], []byte("STMP")) {
		return NewStatusError(StatusProtocolError, "handshake magic "+string(title)+" is not STMP")
	}
	if h.Kind == HandshakeKindServer {
		h.Status = Status(title[4])
	}
	if h.Kind == HandshakeKindClient {
		h.Major = title[4] >> 4
		h.Minor = title[4] & 0xF
	}
	var err error
	var hs uint64
	if hs, err = readUvarint(r, title[:1]); err != nil {
		return NewStatusError(StatusNetworkError, "read handshake header size: "+err.Error())
	}
	if hs > limit {
		return NewStatusError(StatusRequestEntityTooLarge, "handshake size "+strconv.Itoa(int(hs))+" too large, limit is "+strconv.Itoa(int(limit)))
	}
	if hs > 0 {
		hb := make([]byte, hs)
		if _, err = r.Read(hb); err != nil {
			return NewStatusError(StatusNetworkError, "read handshake header: "+err.Error())
		}
		if err = h.Header.Unmarshal(hb); err != nil {
			return NewStatusError(StatusProtocolError, "parse handshake header: "+err.Error())
		}
	}
	var ms uint64
	if ms, err = readUvarint(r, title[:1]); err != nil {
		return NewStatusError(StatusNetworkError, "read handshake message size: "+err.Error())
	}
	if ms+hs > limit {
		return NewStatusError(StatusRequestEntityTooLarge, "handshake size "+strconv.Itoa(int(hs+ms))+" too large, limit is "+strconv.Itoa(int(limit)))
	}
	if ms > 0 {
		mb := make([]byte, ms)
		if _, err = r.Read(mb); err != nil {
			return NewStatusError(StatusNetworkError, "read handshake message: "+err.Error())
		}
		h.Message = string(mb)
	}
	return nil
}

func (h *Handshake) Write(w io.Writer) StatusError {
	buf := h.MarshalBinary()
	if _, err := w.Write(buf); err != nil {
		return NewStatusError(StatusNetworkError, "write handshake: "+err.Error())
	}
	return nil
}

// server only
func (h *Handshake) MarshalText() []byte {
	buf := []byte("STMP")
	if h.Kind == HandshakeKindServer {
		buf = append(buf, hexFormatUint64(uint64(h.Status))...)
	} else {
		buf = append(buf, hexDigits[h.Major], hexDigits[h.Minor])
	}
	header := h.Header.Marshal()
	if len(header) > 0 {
		buf = append(buf, '\n', 'H')
		buf = append(buf, header...)
	}
	if len(h.Message) > 0 {
		buf = append(buf, '\n', 'M')
		buf = append(buf, h.Message...)
	}
	return buf
}

func (h *Handshake) MarshalBinary() []byte {
	buf := []byte("STMP")
	if h.Kind == HandshakeKindClient {
		buf = append(buf, h.Major<<4|h.Minor)
	} else {
		buf = append(buf, byte(h.Status))
	}
	header := h.Header.Marshal()
	chunk := make([]byte, 10, 10)
	buf = append(buf, chunk[:binary.PutUvarint(chunk, uint64(len(header)))]...)
	buf = append(buf, header...)
	buf = append(buf, chunk[:binary.PutUvarint(chunk, uint64(len(h.Message)))]...)
	buf = append(buf, h.Message...)
	return buf
}

// client only
func (h *Handshake) UnmarshalText(data []byte) StatusError {
	if len(data) < 4 || !bytes.Equal(data[:4], []byte("STMP")) {
		return NewStatusError(StatusProtocolError, "invalid magic: "+string(data[0:min(4, len(data))]))
	}
	data = data[4:]
	i := bytes.IndexByte(data, '\n')
	if i == -1 {
		i = len(data)
	}
	var err error
	if h.Kind == HandshakeKindServer {
		if h.Status, err = hexParseStatus(data[0:i]); err != nil {
			return NewStatusError(StatusProtocolError, "invalid status: "+err.Error())
		}
	} else if i != 2 {
		return NewStatusError(StatusProtocolError, "invalid protocol version: "+string(data[0:i]))
	} else {
		h.Major = hexChunks[data[0]]
		h.Minor = hexChunks[data[1]]
		if h.Major > 0xF || h.Minor > 0xF {
			return NewStatusError(StatusProtocolError, "invalid protocol version: "+string(data[:2]))
		}
	}
	if i == len(data) {
		return nil
	}
	data = data[i+1:]
	if len(data) == 0 {
		return nil
	}
	if data[0] == 'H' {
		i = bytes.IndexByte(data, '\n')
		if i == -1 {
			i = len(data)
		}
		if err = h.Header.Unmarshal(data[1:i]); err != nil {
			return NewStatusError(StatusProtocolError, "invalid header: "+err.Error())
		}
		if i == len(data) {
			return nil
		}
		data = data[i+1:]
	}
	if len(data) == 0 {
		return nil
	}
	if data[0] == 'M' {
		h.Message = string(data[1:])
	}
	return nil
}

// client only
func (h *Handshake) UnmarshalBinary(data []byte) StatusError {
	if len(data) < 4 || !bytes.Equal(data[:4], []byte("STMP")) {
		return NewStatusError(StatusProtocolError, "invalid magic: "+string(data[0:min(4, len(data))]))
	}
	data = data[4:]
	if h.Kind == HandshakeKindServer {
		if len(data) == 0 {
			return NewStatusError(StatusProtocolError, "invalid status")
		}
		h.Status = Status(data[0])
		data = data[1:]
	} else {
		if len(data) == 0 {
			return NewStatusError(StatusProtocolError, "invalid protocol version")
		}
		h.Major = data[0] >> 4
		h.Minor = data[0] & 0xF
		data = data[1:]
	}
	hs, n := binary.Uvarint(data)
	if n <= 0 {
		return NewStatusError(StatusProtocolError, "invalid header size")
	}
	data = data[n:]
	if len(data) < int(hs) {
		return NewStatusError(StatusProtocolError, "invalid header")
	}
	if err := h.Header.Unmarshal(data[:hs]); err != nil {
		return NewStatusError(StatusProtocolError, "invalid header: "+err.Error())
	}
	data = data[hs:]
	ms, n := binary.Uvarint(data)
	if n <= 0 {
		return NewStatusError(StatusProtocolError, "invalid message size")
	}
	data = data[n:]
	if len(data) != int(ms) {
		return NewStatusError(StatusProtocolError, "invalid message")
	}
	h.Message = string(data)
	return nil
}
