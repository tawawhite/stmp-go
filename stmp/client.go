// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-28 15:15:32
package stmp

import "time"

type DialOptions struct {
	Header           Header
	HandshakeTimeout time.Duration
	WriteTimeout     time.Duration
	ReadTimeout      time.Duration
}

func DialTCP(addr string, opts *DialOptions) (Conn, *StatusError) {
	panic("not implemented")
}

func DialTCPWithTLS(addr, certFile, keyFile string, opts *DialOptions) (Conn, *StatusError) {
	panic("not implemented")
}

func DialKCP(addr string, opts *DialOptions) (Conn, *StatusError) {
	panic("not implemented")
}

func DialKCPWithTLS(addr, certFile, keyFile string, opts *DialOptions) (Conn, *StatusError) {
	panic("not implemented")
}

func DialWebSocket(addr string, opts *DialOptions) (Conn, *StatusError) {
	panic("not implemented")
}

func DialWebSocketWithTLS(addr, certFile, keyFile string, opts *DialOptions) (Conn, *StatusError) {
	panic("not implemented")
}
