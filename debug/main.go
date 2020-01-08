package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
	"unsafe"
)

type flagSet flag.FlagSet

func printJson(prefix string, v interface{}) {
	data, _ := json.MarshalIndent(v, "", "  ")
	println(prefix + ": " + string(data))
}

var cmds = map[string]func(flag *flagSet){
	"debugChannel": func(flag *flagSet) {
		ch := make(chan int)
		wait := make(chan int)
		go func() {
			for i := 0; i < 2; i++ {
				time.Sleep(time.Second)
				select {
				case v, ok := <-ch:
					log.Printf("receive ch %d, ok: %t at %d", v, ok, i)
				default:
					log.Printf("default case at %d", i)
				}
			}
			wait <- 1
		}()
		ch <- 0
		close(wait)
		d := <-wait
		time.Sleep(3 * time.Second)
		log.Printf("done with %d", d)
	},
	"debugSliceMap": func(flag *flagSet) {
		type foo struct {
			Foo string
		}
		slice := []foo{{"foo"}, {"bar"}}
		v1 := &slice[0]
		v1.Foo = "foo-2"
		printJson("changed by ref element", slice)
		mapVar := map[int]foo{0: {"foo"}, 1: {"bar"}}
		//v1 = &mapVar[0] cannot take the address, so map could not use struct directly
		printJson("map", mapVar)
	},
	"debugRange": func(flag *flagSet) {
		v1 := make([]byte, 2)
		_ = v1[2:]
		// should panic
		_ = v1[:3]
	},
	"debugSizeof": func(flag *flagSet) {
		type empty struct{}
		type emptyField struct {
			empty empty
		}
		type emptyFieldMore struct {
			empty  empty
			size   int
			empty2 empty
			empty3 empty
			size2  int
		}
		// 0, 0, 16, empty fields cost 0 space
		log.Printf("empty size: %d, emptyField size: %d, emptyFieldMore size: %d.", unsafe.Sizeof(empty{}), unsafe.Sizeof(emptyField{}), unsafe.Sizeof(emptyFieldMore{}))
	},
	"debugMethod": func(flag *flagSet) {
		d1 := &debugMethod{}
		d2 := &debugMethod{}
		// will panic
		//log.Println("d1.echo == d2.echo:", compare(d1.echo, d2.echo))
		// false
		log.Println("empty struct pointer:", d1 == d2)
		// true
		log.Println("empty struct:", debugMethod{} == debugMethod{})
		// false
		log.Println("struct pointer:", &debugMethodWithField{} == &debugMethodWithField{})
		// true
		log.Println("struct:", debugMethodWithField{} == debugMethodWithField{})
	},
	"debugCast": func(flag *flagSet) {
		//msb := map[string]bool{"ok": true}
		// panic: interface conversion: interface {} is map[string]bool, not map[string]interface {}
		//log.Println("map.<string, bool> cast:", castMapInterface(msb))
	},
	"debugWriteClosed": func(flag *flagSet) {
		defer func() {
			if e := recover(); e != nil {
				log.Printf("recover err: %s", e)
			}
		}()
		c := make(chan int)
		close(c)
		c <- 1
	},
	"debugAppend": func(flag *flagSet) {
		p1 := make([]byte, 2)
		p1[0] = 'a'
		p1[1] = 'b'
		p2 := p1[:1]
		p2 = append(p2, 'B')
		p4 := append(p2, 'C')
		log.Printf("p1: %s, p2: %s, p4: %s, cap(p1): %d, cap(p2): %d, cap(p4): %d", string(p1), string(p2), string(p4), cap(p1), cap(p2), cap(p4))
	},
	"debugSize": func(flag *flagSet) {
		type foo struct {
			net.Conn
			io.ReadWriteCloser
		}
		log.Printf("Sizeof foo with mutliple ReadWriteCloser: 32=%d.", unsafe.Sizeof(foo{}))
		var ifc io.ReadWriteCloser
		log.Printf("Sizeof single interface: 16=%d.", unsafe.Sizeof(ifc))
	},
	"debugCloseCh": func(flag *flagSet) {
		ch := make(chan int)
		go func() {
			time.Sleep(time.Second)
			close(ch)
		}()
		ch <- 1
	},
	"debugReadClosed": func(flag *flagSet) {
		srv, _ := net.Listen("tcp", "127.0.0.1:9991")
		go func() {
			for {
				conn, _ := srv.Accept()
				go func() {
					buf := make([]byte, 5)
					for {
						_, err := conn.Read(buf)
						if err != nil {
							// EOF
							log.Printf("read error: %s.", err.Error())
							break
						}
					}
				}()
			}
		}()
		time.Sleep(time.Second)
		conn, _ := net.Dial("tcp", "127.0.0.1:9991")
		_, err := conn.Write([]byte("12345"))
		if err != nil {
			log.Printf("write error: %s.", err)
		}
		conn.Close()
		time.Sleep(time.Second)
	},
	"debugNil": func(flag *flagSet) {
		type foo struct {
		}
		var a *foo
		log.Printf("*foo isNil: %t, nil==%t.", isNil(a), a == nil)
		a = nil
		log.Printf("a = nil, isNil: %t, nil==%t.", isNil(a), a == nil)
		log.Printf("(*foo)(nil), isNil: %t.", isNil((*foo)(nil)))
		var b *foo
		log.Printf("*foo == *foo: %t.", a == b)
		log.Printf("getMethod(): %t.", isNil(getMethod()))
		var c interface{}
		c = (*foo)(nil)
		log.Printf("assign nil pointer to interface type: false==%t.", isNil(c))
	},
	"dialTLS": func(flag *flagSet) {
		conn, err := tls.Dial("tcp", "github.com", nil)
		if err != nil {
			log.Fatalf("tls dial error: %q.", err)
		}
		_, err = conn.Write([]byte("GET /welcome HTTP/1.1\r\nHost: github.com\r\nUser-Agent: Google Chrome/78.0\r\nConnection: keep-alive\r\n\r\n"))
		if err != nil {
			log.Fatalf("tls write header error: %q.", err)
		}
		r := bufio.NewReader(conn)
		status, err := r.ReadString('\n')
		if err != nil {
			log.Fatalf("tls read status line error: %q.", err)
		}
		log.Printf("tls status line: %q.", status)
		nc, err := net.Dial("tcp", "github.com:443")
		if err != nil {
			log.Fatalf("net dial error: %q.", err)
		}
		tc := tls.Client(nc, &tls.Config{ServerName: "github.com"})
		_, err = tc.Write([]byte("GET /welcome HTTP/1.1\r\nHost: github.com\r\nUser-Agent: Google Chrome/78.0\r\nConnection: keep-alive\r\n\r\n"))
		if err != nil {
			log.Fatalf("net write header error: %q.", err)
		}
		r = bufio.NewReader(tc)
		status, err = r.ReadString('\n')
		if err != nil {
			log.Fatalf("net read status line error: %q.", err)
		}
		log.Printf("net status line: %q.", status)
	},
	"debugReadClosedBuffedChan": func(flag *flagSet) {
		ch := make(chan int, 2)
		ch <- 1
		close(ch)
		i, ok := <-ch
		log.Printf("should read 1==%d, true=%t.", i, ok)
		i, ok = <-ch
		log.Printf("should read 0==%d, false=%t.", i, ok)
		ch1 := make(chan int)
		close(ch1)
		<-ch1
		// Result: read on closed channel will not emit error, need ok to detect state
	},
	"debugSelectWrite": func(flag *flagSet) {
		ch := make(chan int)
		go func() {
			time.Sleep(time.Second)
			v := <-ch
			log.Printf("receive v: %d.", v)
		}()
		ctx, _ := context.WithTimeout(context.Background(), time.Millisecond)
		select {
		case ch <- 1:
			log.Printf("ch write 1 done")
		case <-ctx.Done():
			log.Printf("timeout: %s.", ctx.Err())
		}
		ch <- 2
		log.Printf("ch write 2 done")
		time.Sleep(time.Second)
	},
}

func isNil(v interface{}) bool {
	return v == nil
}

func castMapInterface(in interface{}) interface{} {
	v := in.(map[string]interface{})
	return v["ok"]
}

type debugMethod struct {
}

func getMethod() *debugMethod {
	return nil
}

type debugMethodWithField struct {
	v int
}

func (*debugMethod) echo() {
}

func Usage() {
	print("Available commands: ")
	keys := make([]string, 0, len(cmds))
	for k := range cmds {
		keys = append(keys, k)
	}
	println(strings.Join(keys, ", "))
}

func main() {
	for _, v := range os.Args {
		if v == "-h" || v == "--help" {
			Usage()
			os.Exit(0)
		}
	}
	if len(os.Args) < 2 || cmds[os.Args[1]] == nil {
		Usage()
		os.Exit(1)
	}
	cmd := cmds[os.Args[1]]
	cmd(nil)
}
