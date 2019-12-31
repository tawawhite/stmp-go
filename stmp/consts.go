// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:34:33
package stmp

import (
	"errors"
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

func isFin(kind byte) bool {
	// always true for following is not implemented
	return true
}

func isHead(kind byte) bool {
	return kind == MessageKindPing || kind == MessageKindPong
}

func isMid(kind byte) bool {
	return kind == MessageKindRequest || kind == MessageKindResponse
}

func isAction(kind byte) bool {
	return kind == MessageKindRequest || kind == MessageKindNotify
}

func isStatus(kind byte) bool {
	return kind == MessageKindResponse || kind == MessageKindClose
}

func isKind(kind byte) bool {
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

type StatusError struct {
	code Status
	err  error
}

func (e *StatusError) Error() string {
	return "STMP " + e.code.Error() + ": " + e.err.Error()
}

func NewStatusError(code Status, err interface{}) *StatusError {
	if err == nil {
		return &StatusError{code: code, err: errors.New(MapStatus[code])}
	}
	if str, ok := err.(string); ok {
		return &StatusError{code: code, err: errors.New(str)}
	}
	return &StatusError{code: code, err: err.(error)}
}

func DetectError(err error, rollbackStatus Status) (Status, []byte) {
	if se, ok := err.(*StatusError); ok {
		return se.code, []byte(se.err.Error())
	}
	if sc, ok := err.(Status); ok {
		return sc, []byte(sc.Error())
	}
	return rollbackStatus, []byte(err.Error())
}

const AcceptContentType = "Accept"
const AcceptEncoding = "Accept-Encoding"
const AcceptPacketFormat = "Accept-packet-Format"

const DetermineContentType = "Content-Type"
const DetermineEncoding = "Content-Encoding"
const DeterminePacketFormat = "packet-Format"
const DetermineStmpVersion = "Stmp-Version"
