// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-08 22:12:23

/*
pointer cast has no performance effect, interface need query iface table to check in runtime

goos: darwin
goarch: amd64
pkg: github.com/acrazing/stmp-go/debug
BenchmarkTyped-4            	270218812	         4.45 ns/op
BenchmarkUntyped-4          	253607560	         4.74 ns/op
BenchmarkTypedInterface-4   	69845611	        17.4 ns/op
PASS
ok  	github.com/acrazing/stmp-go/debug	4.754s
 */

package main

import "testing"

type Foo struct {
}

func (*Foo) Debug() {
}

type typedMap = map[int]*Foo
type untypedMap = map[int]interface{}

func BenchmarkTyped(b *testing.B) {
	state := typedMap{0: new(Foo)}
	for i := 0; i < b.N; i++ {
		state[0].Debug()
	}
}

func BenchmarkUntyped(b *testing.B) {
	state := untypedMap{0: new(Foo)}
	for i := 0; i < b.N; i++ {
		state[0].(*Foo).Debug()
	}
}

type Debug interface {
	Debug()
}

func BenchmarkUntypedInterface(b *testing.B) {
	state := untypedMap{0: new(Foo)}
	for i := 0; i < b.N; i++ {
		state[0].(Debug).Debug()
	}
}

type typedInterface map[int]Debug

func BenchmarkTypedInterface(b *testing.B) {
	state := typedInterface{0: new(Foo)}
	for i := 0; i < b.N; i++ {
		state[0].Debug()
	}
}