// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:34:33
package stmp

import (
	"strconv"
)

const (
	MessageKindPing     = 0x0
	MessageKindPong     = 0x1
	MessageKindRequest  = 0x2
	MessageKindNotify   = 0x3
	MessageKindResponse = 0x4
	MessageKindClose    = 0x5
)

func ShouldAlwaysFinal(kind byte) bool {
	// always true for following is not implemented
	return true
}

func ShouldHeadOnly(kind byte) bool {
	return kind == MessageKindPing || kind == MessageKindPong
}

func HasMid(kind byte) bool {
	return kind == MessageKindRequest || kind == MessageKindResponse
}

func HasAction(kind byte) bool {
	return kind == MessageKindRequest || kind == MessageKindNotify
}

func HasStatus(kind byte) bool {
	return kind == MessageKindResponse || kind == MessageKindClose
}

func IsValidKind(kind byte) bool {
	switch kind {
	case MessageKindPing, MessageKindPong, MessageKindRequest, MessageKindNotify, MessageKindResponse, MessageKindClose:
		return true
	default:
		return false
	}
}

const MaskFin = 0x80
const MaskHead = 0b1000
const OffsetKind = 4

var MapTextKind = map[byte]byte{
	'I': MessageKindPing,
	'O': MessageKindPong,
	'Q': MessageKindRequest,
	'N': MessageKindNotify,
	'S': MessageKindResponse,
	'C': MessageKindClose,
	//'F': MessageKindFollowing,
}

var MapKindText = map[byte]byte{
	MessageKindPing:     'I',
	MessageKindPong:     'O',
	MessageKindRequest:  'Q',
	MessageKindNotify:   'N',
	MessageKindResponse: 'S',
	MessageKindClose:    'C',
}

type Status byte

const (
	StatusOk                         Status = 0x00
	StatusNetworkError               Status = 0x01
	StatusProtocolError              Status = 0x02
	StatusUnsupportedProtocolVersion Status = 0x03
	StatusUnsupportedContentType     Status = 0x04
	StatusUnsupportedFormat          Status = 0x05
	StatusUnknown                    Status = 0x06
	StatusUnmarshalError             Status = 0x07
	StatusMarshalError               Status = 0x08
	StatusCancelled                  Status = 0x09
	StatusBadRequest                 Status = 0x20
	StatusUnauthorized               Status = 0x21
	StatusNotFound                   Status = 0x22
	StatusRequestTimeout             Status = 0x23
	StatusRequestEntityTooLarge      Status = 0x24
	StatusTooManyRequests            Status = 0x25
	StatusClientClosed               Status = 0x26
	StatusClientCancelled            Status = 0x27
	StatusInternalServerError        Status = 0x40
	StatusServerShutdown             Status = 0x41
)

var MapStatus = map[Status]string{
	StatusOk:                         "Ok",
	StatusNetworkError:               "Network error",
	StatusProtocolError:              "Protocol error",
	StatusUnsupportedProtocolVersion: "Unsupported protocol version",
	StatusUnsupportedContentType:     "Unsupported content type",
	StatusUnsupportedFormat:          "Unsupported format",
	StatusUnknown:                    "Unknown",
	StatusUnmarshalError:             "Unmarshal error",
	StatusMarshalError:               "Marshal error",
	StatusCancelled:                  "Cancelled",
	StatusBadRequest:                 "Bad request",
	StatusUnauthorized:               "Unauthorized",
	StatusNotFound:                   "Not found",
	StatusRequestTimeout:             "Request timeout",
	StatusRequestEntityTooLarge:      "Request entity too large",
	StatusTooManyRequests:            "Too many requests",
	StatusClientClosed:               "Client closed",
	StatusClientCancelled:            "Client cancelled",
	StatusInternalServerError:        "Internal server error",
	StatusServerShutdown:             "Server shutdown",
}

func (s Status) Error() string {
	if m, ok := MapStatus[s]; ok {
		return m
	}
	return "Unknown (0x" + strconv.FormatUint(uint64(s), 16) + ")"
}

func (s Status) Spread() (Status, []byte) {
	return s, []byte(s.Error())
}

type StatusError struct {
	code    Status
	message string
}

func (e *StatusError) Error() string {
	return "STMP " + e.code.Error() + ": " + e.message
}

func (e *StatusError) Spread() (Status, []byte) {
	return e.code, []byte(e.message)
}

func NewStatusError(code Status, data interface{}) *StatusError {
	se := new(StatusError)
	se.code = code
	if m, ok := data.(string); ok {
		se.message = m
	} else if e, ok := data.(error); ok {
		se.message = e.Error()
	}
	if se.message == "" {
		se.message = se.code.Error()
	}
	return se
}

func DetectError(err error, rollbackStatus Status) *StatusError {
	if se, ok := err.(*StatusError); ok {
		return se
	}
	if sc, ok := err.(Status); ok {
		return NewStatusError(sc, "")
	}
	return NewStatusError(rollbackStatus, err.Error())
}

const AcceptContentType = "Accept"
const AcceptEncoding = "Accept-Encoding"
const AcceptPacketFormat = "Accept-Packet-Format"

const DetermineContentType = "Content-Type"
const DetermineEncoding = "Content-Encoding"
const DeterminePacketFormat = "Packet-Format"
const DetermineStmpVersion = "Stmp-Version"
