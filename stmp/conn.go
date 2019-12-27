// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-26 18:16:34
package stmp

import (
	"net"
)

type Conn interface {
	Headers() Header
	RemoteAddr() net.TCPAddr
	LocalAddr() net.TCPAddr
	Request(options SendContext) error
	Notify(options SendContext) error
	Groups() map[*Group]bool
}
