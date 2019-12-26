// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:34:33
package stmp

type MessageKind byte

const (
	MessageKindPing      MessageKind = 0
	MessageKindRequest               = 1
	MessageKindNotify                = 2
	MessageKindResponse              = 3
	MessageKindFollowing             = 4
	MessageKindClose                 = 5
)

var MapTextKind = map[byte]MessageKind{
	'p': MessageKindPing,
	'q': MessageKindRequest,
	'n': MessageKindNotify,
	's': MessageKindResponse,
	'f': MessageKindFollowing,
	'c': MessageKindClose,
}

type EncodingKind byte

const (
	EncodingKindNone EncodingKind = 0
	// ascii is a subset of utf8
	EncodingKindUTF8  = 1
	EncodingKindUTF16 = 2
	// payload is empty
	EncodingKindBinary = 3
)

type HandshakeStatus byte

const (
	HandshakeStatusOk                 HandshakeStatus = 0
	HandshakeStatusAuthenticateFailed                 = 1
	HandshakeStatusProtocolError                      = 2
	HandshakeStatusNetworkError                       = 3
)

type ResponseStatus byte

const (
	ResponseStatusOk           ResponseStatus = 0
	ResponseStatusServerError                 = 1
	ResponseStatusClientError                 = 2
	ResponseStatusNetworkError                = 3
	ResponseStatusServerClosed                = 4
)

type CloseStatus byte

const (
	CloseStatusOk            CloseStatus = 0
	CloseStatusClientError               = 1
	CloseStatusServerError               = 2
	CloseStatusServerClosed              = 3
	CloseStatusProtocolError             = 4
	CloseStatusNetworkError              = 5
)

var MapWsCloseStatus = map[int]CloseStatus{
	1000: CloseStatusOk,
	1001: CloseStatusServerClosed,
	1002: CloseStatusProtocolError,
	1006: CloseStatusNetworkError,
	1011: CloseStatusServerError,
	1009: CloseStatusClientError,
}

type FormatKind byte

const (
	FormatKindText = 0
	FormatKindBinary = 1
)