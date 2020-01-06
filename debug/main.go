package main

import (
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
}

func castMapInterface(in interface{}) interface{} {
	v := in.(map[string]interface{})
	return v["ok"]
}

type debugMethod struct {
}

type debugMethodWithField struct {
	v int
}

func (*debugMethod) echo() {
}

func usage() {
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
			usage()
			os.Exit(0)
		}
	}
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	cmd := cmds[os.Args[1]]
	cmd(nil)
}