// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-08 17:14:13
package stmp_test

import (
	"github.com/acrazing/stmp-go/stmp"
	"log"
)

func ExampleServer_listenMultiple() {
	srv := stmp.NewServer(stmp.NewServerOptions())
	go srv.ListenAndServeTCP("127.0.0.1:9991")
	log.Printf("stmp server is listening at %q.", "tcp://127.0.0.1:9991")
	go srv.ListenAndServeWebsocket("127.0.0.1:9992", "/stmp")
	log.Printf("stmp server is listening at %q.", "ws://127.0.0.1:9992/stmp")
	go srv.ListenAndServeKCPWithTLS("127.0.0.1:9993", "./example.crt", "./example.key")
	log.Printf("stmp server is listening at %q.", "kcp+tls://127.0.0.1:9992")
	err := srv.Wait()
	if err != nil {
		log.Fatalf("stmp server listen error: %q.", err)
	} else {
		log.Println("stmp server shutdown.")
	}
}
