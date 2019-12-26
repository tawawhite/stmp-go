// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:34:33
package stmp

type MessageKind byte

const (
	MessageKindPing      MessageKind = 0x0
	MessageKindRequest               = 0x1
	MessageKindNotify                = 0x2
	MessageKindResponse              = 0x3
	MessageKindFollowing             = 0x4
	MessageKindClose                 = 0x5
)

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
