// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-27 13:27:58
package main

import (
	"github.com/acrazing/stmp-go/stmp"
	"os"
)

func main() {
	stmp.RegisterMediaCodec(stmp.NewJsonCodec(), stmp.NewProtobufCodec())
	stmp.RegisterEncodingCodec(stmp.NewGzipCodec())
	srv := stmp.NewServer()
	go srv.ListenAndServeTCP("127.0.0.1:5000")
	go srv.ListenAndServeKCP("127.0.0.1:5001")
	go srv.ListenAndServeWebSocket("127.0.0.1:5002", "/ws")
	go srv.ListenAndServeTCPWithTLS("127.0.0.1:5003", "./examples/tls_key/example.crt", "./examples/tls_key/example.key")
	go srv.ListenAndServeKCPWithTLS("127.0.0.1:5004", "./examples/tls_key/example.crt", "./examples/tls_key/example.key")
	go srv.ListenAndServeWebSocketWithTLS("127.0.0.1:5005", "/ws", "./examples/tls_key/example.crt", "./examples/tls_key/example.key")
	println("room server is listening at     tcp://127.0.0.1:5000")
	println("                                kcp://127.0.0.1:5001")
	println("                                 ws://127.0.0.1:5002/ws")
	println("                            tcp+tls://127.0.0.1:5003")
	println("                            kcp+tls://127.0.0.1:5004")
	println("                                wss://127.0.0.1:5005/ws")
	err := srv.Wait()
	if err != nil {
		println("listen error", err.Error())
		os.Exit(1)
	}
}
