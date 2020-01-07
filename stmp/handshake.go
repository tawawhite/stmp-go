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
	size, err := readUvarint(r, title[:1])
	if err != nil {
		return NewStatusError(StatusNetworkError, "read handshake header size: "+err.Error())
	}
	if size == 0 {
		return nil
	}
	if size > limit {
		return NewStatusError(StatusRequestEntityTooLarge, "handshake header size "+strconv.Itoa(int(size))+" too large, limit is "+strconv.Itoa(int(limit)))
	}
	head := make([]byte, size)
	if _, err := r.Read(head); err != nil {
		return NewStatusError(StatusNetworkError, "read handshake header: "+err.Error())
	}
	sep := bytes.Index(head, []byte("\n\n"))
	if sep > -1 {
		h.Message = string(head[sep+2:])
		head = head[:sep]
	}
	if err := h.Header.Unmarshal(head); err != nil {
		return NewStatusError(StatusProtocolError, "parse handshake header: "+err.Error())
	}
	return nil
}

func (h *Handshake) Write(r io.Writer) StatusError {
	title, header := h.MarshalHead()
	if _, err := r.Write(title); err != nil {
		return NewStatusError(StatusNetworkError, "write handshake title: "+err.Error())
	}
	if _, err := r.Write(header); err != nil {
		return NewStatusError(StatusNetworkError, "write handshake header: "+err.Error())
	}
	if len(h.Message) == 0 {
		return nil
	}
	if _, err := r.Write(append([]byte("\n\n"), h.Message...)); err != nil {
		return NewStatusError(StatusNetworkError, "write handshake message: "+err.Error())
	}
	return nil
}

// server only
func (h *Handshake) MarshalText() []byte {
	title := make([]byte, 6)
	copy(title, "STMP")
	if h.Status > 0xF {
		title[4] = hexDigits[h.Status>>4]
		title[5] = hexDigits[h.Status&0xF]
		title = title[:6]
	} else {
		title[4] = hexDigits[h.Status]
		title = title[:5]
	}
	header := h.Header.Marshal()
	if len(header) > 0 {
		title = append(title, '\n')
		title = append(title, header...)
	}
	if len(h.Message) > 0 {
		title = append(title, "\n\n"...)
		title = append(title, h.Message...)
	}
	return title
}

func (h *Handshake) MarshalBinary() []byte {
	title := []byte{'S', 'T', 'M', 'P', byte(h.Status)}
	header := h.Header.Marshal()
	if len(header) > 0 {
		title = append(title, '\n')
		title = append(title, header...)
	}
	if len(h.Message) > 0 {
		title = append(title, "\n\n"...)
		title = append(title, h.Message...)
	}
	return title
}

// client only
func (h *Handshake) UnmarshalText(data []byte) error {
	if len(data) < 4 || !bytes.Equal(data[:4], []byte("STMP")) {
		return NewStatusError(StatusProtocolError, "invalid magic: "+string(data[0:min(4, len(data))]))
	}
	data = data[4:]
	sep := bytes.IndexByte(data, '\n')
	if sep == -1 {
		sep = len(data)
	}
	var err error
	h.Status, err = hexParseStatus(data[:sep])
	if err != nil {
		return NewStatusError(StatusProtocolError, "invalid status: "+err.Error())
	}
	if len(data) == sep {
		return nil
	}
	data = data[sep:]
	sep = bytes.Index(data, []byte("\n\n"))
	if sep != -1 {
		h.Message = string(data[sep+2:])
	}
	if sep > 0 {
		data = data[1:sep]
	} else {
		data = data[1:]
	}
	err = h.Header.Unmarshal(data)
	if err != nil {
		return NewStatusError(StatusProtocolError, "invalid header: "+err.Error())
	}
	return nil
}

// client only
func (h *Handshake) UnmarshalBinary(data []byte) error {
	if len(data) < 4 || !bytes.Equal(data[:4], []byte("STMP")) {
		return NewStatusError(StatusProtocolError, "invalid magic: "+string(data[0:min(4, len(data))]))
	}
	data = data[4:]
	if len(data) == 0 {
		return NewStatusError(StatusProtocolError, "empty status")
	}
	h.Status = Status(data[4])
	data = data[5:]
	if len(data) == 0 {
		return nil
	}
	sep := bytes.Index(data, []byte("\n\n"))
	if sep != -1 {
		h.Message = string(data[sep+2:])
	}
	if sep > 0 {
		data = data[1:sep]
	} else {
		data = data[1:]
	}
	err := h.Header.Unmarshal(data)
	if err != nil {
		return NewStatusError(StatusProtocolError, "invalid header: "+err.Error())
	}
	return nil
}
