// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:34:33
package stmp

import (
	"errors"
	"strconv"
)

type MessageKind byte

const (
	MessageKindPing      MessageKind = 0x0
	MessageKindRequest               = 0x1
	MessageKindNotify                = 0x2
	MessageKindResponse              = 0x3
	MessageKindFollowing             = 0x4
	MessageKindClose                 = 0x5
)

const FlagFin = 0x80
const FlagPure = 0b1000
const KindOffset = 4

var PingMessage = []byte{byte(FlagFin | MessageKindPing<<KindOffset | FlagPure)}

var MapTextKind = map[byte]MessageKind{
	'P': MessageKindPing,
	'Q': MessageKindRequest,
	'N': MessageKindNotify,
	'S': MessageKindResponse,
	'F': MessageKindFollowing,
	'C': MessageKindClose,
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
const AcceptPacketFormat = "Accept-Packet-Format"

const DetermineContentType = "Content-Type"
const DetermineEncoding = "Content-Encoding"
const DeterminePacketFormat = "Packet-Format"
const DetermineStmpVersion = "Stmp-Version"
