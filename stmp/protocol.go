// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-23 19:55:29
package stmp

import (
	"errors"
	"strings"
)

func ReadNegotiate(input string) (v string, n int) {
	n = strings.IndexByte(input, ',')
	if n == -1 {
		n = len(input)
	}
	seg := strings.IndexByte(input[0:n], ';')
	if seg == -1 {
		seg = n
	}
	v = strings.TrimSpace(v[0:seg])
	return
}

var (
	invalidReservedHeadBits = errors.New("invalid reserved head bits")
	invalidHeadFlags        = errors.New("invalid head flags")
	invalidMessageKind      = errors.New("invalid message kind")
)

func parseHead(h byte) (fin bool, kind MessageKind, pure bool, err error) {
	if h&0b111 != 0 {
		err = invalidReservedHeadBits
		return
	}
	fin = h&0x80 != 0
	kind = MessageKind((h >> 4) & 0b111)
	pure = h&0b1000 != 0
	switch kind {
	case MessageKindResponse, MessageKindRequest, MessageKindNotify:
	case MessageKindPing:
		if !fin || !pure {
			err = invalidHeadFlags
		}
	case MessageKindClose:
		if !fin {
			err = invalidHeadFlags
		}
	default:
		err = invalidMessageKind
	}
	return
}
