// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-27 13:55:03
package main

import (
	"encoding/json"
	"flag"
	"log"
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
