// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 16:21:44
package main

import (
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:4000")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				panic(err)
			}
			println("new connection from %s", conn.RemoteAddr().String())
			conn.Close()
		}
	}()
	_, _ = net.Dial("tcp", "127.0.0.1:4000")
	time.Sleep(time.Second)
	_ = lis.Close()
	time.Sleep(time.Microsecond)
}