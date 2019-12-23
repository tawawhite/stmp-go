// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:34:33
package stmp

type MessageKind byte

const (
	MessageKindPing MessageKind = 0
	MessageKindRequest = 1
	MessageKindNotify = 2
	MessageKindResponse = 3
	MessageKindFollowing = 4
	MessageKindClose = 5
)

type EncodingKind byte

const (
	EncodingKindBinary EncodingKind = 0
	// ascii is a subset of utf8
	EncodingKindUTF8 = 1
	EncodingKindUTF16 = 2
	// payload is empty
	EncodingKindNone = 3
)
