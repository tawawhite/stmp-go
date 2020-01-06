package stmp

type ConnSet map[*Conn]struct{}

func NewConnSet() ConnSet {
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
