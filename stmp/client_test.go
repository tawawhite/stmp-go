package stmp_test

import (
	"crypto/tls"
	"github.com/acrazing/stmp-go/stmp"
	"log"
)

func createClient(cc *stmp.Client) {
}

func ExampleNewClientConn() {
	nc, err := tls.Dial("tcp", "github.com:443", nil)
	if err != nil {
		log.Fatalf("dial error: %q.", err)
	}
	cc := stmp.NewClientConn(nc, stmp.NewDialOptions().ApplyDefault())
	err = cc.Handshake()
	if err != nil {
		log.Fatalf("handshake error: %q.", err.Error())
	}
	createClient(cc)
}

func ExampleDialTCP() {
	cc, err := stmp.DialTCP("github.com:443", nil)
	if err != nil {
		log.Fatalf("dial error: %q.", err)
	}
	createClient(cc)
}

func ExampleDialTCP_insecure() {
	cc, err := stmp.DialTCP("github.com:80", stmp.NewDialOptions().WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %q.", err)
	}
	createClient(cc)
}

func ExampleDialTCP_customCertFile() {
	cc, err := stmp.DialTCP("my.example.com:80", stmp.NewDialOptions().WithCert("./example.crt", true))
	if err != nil {
		log.Fatalf("dial error: %q.", err)
	}
	createClient(cc)
}
