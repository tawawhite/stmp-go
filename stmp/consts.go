// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:34:33
package stmp

import "errors"

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
	StatusNetworkError                      = 0x01
	StatusProtocolError                     = 0x02
	StatusUnsupportedProtocolVersion        = 0x03
	StatusUnsupportedContentType            = 0x04
	StatusUnsupportedFormat                 = 0x05
	StatusBadRequest                        = 0x20
	StatusUnauthorized                      = 0x21
	StatusNotFound                          = 0x22
	StatusRequestTimeout                    = 0x23
	StatusRequestEntityTooLarge             = 0x24
	StatusTooManyRequests                   = 0x25
	StatusClientClosed                      = 0x26
	StatusClientCancelled                   = 0x27
	StatusInternalServerError               = 0x40
	StatusServerShutdown                    = 0x41
)

var MapStatus = map[Status]string{
	StatusOk:                         "(0x00) Ok",
	StatusNetworkError:               "(0x01) NetworkError",
	StatusProtocolError:              "(0x02) ProtocolError",
	StatusUnsupportedProtocolVersion: "(0x03) UnsupportedProtocolVersion",
	StatusUnsupportedContentType:     "(0x04) UnsupportedContentType",
	StatusUnsupportedFormat:          "(0x05) UnsupportedFormat",
	StatusBadRequest:                 "(0x20) BadRequest",
	StatusUnauthorized:               "(0x21) Unauthorized",
	StatusNotFound:                   "(0x22) NotFound",
	StatusRequestTimeout:             "(0x23) RequestTimeout",
	StatusRequestEntityTooLarge:      "(0x24) RequestEntityTooLarge",
	StatusTooManyRequests:            "(0x25) TooManyRequests",
	StatusClientClosed:               "(0x26) ClientClosed",
	StatusClientCancelled:            "(0x27) ClientCancelled",
	StatusInternalServerError:        "(0x40) InternalServerError",
	StatusServerShutdown:             "(0x41) ServerShutdown",
}

type StatusError struct {
	code Status
	err  error
}

func (e *StatusError) Error() string {
	return "STMP" + MapStatus[e.code] + ": " + e.err.Error()
}

func NewStatusError(code Status, err interface{}) error {
	if str, ok := err.(string); ok {
		return &StatusError{code: code, err: errors.New(str)}
	}
	return &StatusError{code: code, err: err.(error)}
}

const AcceptContentType = "Accept"
const AcceptEncoding = "Accept-Encoding"
const AcceptPacketFormat = "Accept-Packet-Format"

const DetermineContentType = "Content-Type"
const DetermineEncoding = "Content-Encoding"
const DeterminePacketFormat = "Packet-Format"
