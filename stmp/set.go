// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-03 16:22:18
package stmp

type ConnSet map[*Conn]struct{}

func NewSet() ConnSet {
	return map[*Conn]struct{}{}
}

func (s ConnSet) Has(conn *Conn) bool {
	_, ok := s[conn]
	return ok
}

func (s ConnSet) Add(conn *Conn) {
	s[conn] = struct{}{}
}

func (s ConnSet) Del(conn *Conn) {
	delete(s, conn)
}

func (s ConnSet) Size() int {
	return len(s)
}

func (s ConnSet) Slice() []*Conn {
	var v []*Conn
	for k := range s {
		v = append(v, k)
	}
	return v
}
