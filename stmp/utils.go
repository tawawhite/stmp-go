// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-23 19:55:29
package stmp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/gorilla/websocket"
	"io"
	"strconv"
	"strings"
)

var (
	ErrPayloadTooLarge = errors.New("handshake size too large")
	ErrInvalidMessage  = errors.New("invalid message")
)

func ReadVarint(r io.Reader) (n uint64, err error) {
	buf := make([]byte, 1)
	var b byte
	i := 0
	for {
		_, err = r.Read(buf)
		if err != nil {
			return n, err
		}
		b = buf[0]
		if i == 8 {
			n |= uint64(b) << (i * 7)
			return
		} else {
			n |= uint64(b&0x7F) << (i * 7)
		}
		if b&0x80 == 0 {
			return
		}
		i += 1
	}
}

func ParseHeaders(buf []byte) (Headers, error) {
	h := NewHeaders()
	for _, l := range strings.Split(string(buf), "\n") {
		sepIndex := strings.IndexByte(l, ':')
		if sepIndex == -1 {
			return nil, ErrInvalidMessage
		}
		key := strings.ToLower(strings.TrimSpace(l[0:sepIndex]))
		value := strings.TrimSpace(l[:sepIndex+1])
		h.Add(key, value)
	}
	return h, nil
}

type HandshakeRequest struct {
	Major   byte
	Minor   byte
	Headers Headers
}

func ReadHandshakeRequest(r io.Reader, maxSize uint64) (*HandshakeRequest, error) {
	size, err := ReadVarint(r)
	if err != nil {
		return nil, err
	}
	if size > maxSize {
		return nil, ErrPayloadTooLarge
	}
	buf := make([]byte, size+2)
	_, err = r.Read(buf)
	if err != nil {
		return nil, err
	}
	headers, err := ParseHeaders(buf[2:])
	if err != nil {
		return nil, err
	}
	return &HandshakeRequest{
		Major:   buf[0],
		Minor:   buf[1],
		Headers: headers,
	}, nil
}

type HandshakeResponse struct {
	Status  HandshakeStatus
	Message string
}

func ReadHandshakeResponse(r io.Reader, maxSize uint64) (*HandshakeResponse, error) {
	buf := make([]byte, 1)
	_, err := r.Read(buf)
	if err != nil {
		return nil, err
	}
	status := buf[0]
	if status&0x80 == 0 {
		return &HandshakeResponse{
			Status:  HandshakeStatus(status),
			Message: "",
		}, nil
	}
	size, err := ReadVarint(r)
	if err != nil {
		return nil, err
	}
	if size > maxSize {
		return nil, ErrPayloadTooLarge
	}
	buf = make([]byte, size)
	_, err = r.Read(buf)
	if err != nil {
		return nil, err
	}
	return &HandshakeResponse{
		Status:  HandshakeStatus(status | 0x7F),
		Message: string(buf),
	}, nil
}

func ParseHead(r byte) (format FormatKind, kind MessageKind, fin bool, encoding EncodingKind) {
	format = FormatKind(r >> 7)
	if format == FormatKindText {
		kind = MapTextKind[r]
		fin = true
		encoding = EncodingKindUTF8
		return
	}
	kind = MessageKind(r & 0b01110000)
	fin = (r & 0b00001000) != 0
	encoding = EncodingKind(r & 0b00000110)
	return
}

func ParseVarint(d []byte) (n uint64, s int, err error) {
	max := len(d)
	var b byte
	for {
		b = d[s]
		if s == 8 {
			n |= uint64(b) << (s * 7)
			return
		} else if b < 0x80 {
			n |= uint64(b) << (s * 7)
			return
		} else {
			n |= uint64(b&0x7F) << (s * 7)
			s += 1
			if max == s {
				err = ErrInvalidMessage
				return
			}
		}
	}
}

func ParseRequestMessage(d []byte, useActionId bool, format FormatKind, encoding EncodingKind) (id uint16, actionId uint64, actionName string, payload []byte, err error) {
	if format == FormatKindText {
		idSeq := bytes.IndexByte(d, ',')
		if idSeq < 0 {
			err = ErrInvalidMessage
			return
		}
		_id, err := strconv.ParseUint(string(d[:idSeq]), 16, 16)
		if err != nil {
			return
		}
		id = uint16(_id)
		actionSeq := 0
		if useActionId {
			actionSeq = bytes.IndexByte(d[idSeq+1:idSeq+1+17], '\n')
			if actionSeq == -1 {
				if len(d) <= idSeq+1+17 {
					err = ErrInvalidMessage
					return
				}
				actionSeq = len(d)
			}
			actionId, err = strconv.ParseUint(string(d[idSeq+1:actionSeq]), 16, 64)
			if err != nil {
				return
			}
		} else {
			actionSeq = bytes.IndexByte(d[idSeq+1:idSeq+1+256], '\n')
			if actionSeq == -1 {
				if len(d) < idSeq+1+256 {
					err = ErrInvalidMessage
				}
				actionSeq = len(d)
			}
			actionName = string(d[idSeq+1 : actionSeq])
		}
		payload = d[actionSeq+1:]
		return
	} else if len(d) < 3 {
		err = ErrInvalidMessage
		return
	} else {
		id = binary.LittleEndian.Uint16(d)
		if useActionId {
			var offset int
			actionId, offset, err = ParseVarint(d[2:])
			if err == nil {
				payload = d[2+offset:]
			}
		} else {
			offset := bytes.IndexByte(d[2:258], '\n')
			if offset == -1 {
				if len(d) <= 258 {
					actionName = string(d[2:])
				} else {
					err = ErrInvalidMessage
				}
			} else {
				actionName = string(d[2:offset])
				payload = d[offset+1:]
			}
		}
		return
	}
}

func ReadRequestMessage(r io.Reader, useActionId bool, format FormatKind, encoding EncodingKind) (id uint16, actionId uint64, actionName string, payload []byte, err error) {
	websocket.Upgrade()
}