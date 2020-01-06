package stmp

// A connection set util to hold as group
type ConnSet map[*Conn]struct{}

// create a new connection set
func NewConnSet() ConnSet {
	return map[*Conn]struct{}{}
}

// check has conn or not
func (s ConnSet) Has(conn *Conn) bool {
	_, ok := s[conn]
	return ok
}

// add conn to set
func (s ConnSet) Add(conn *Conn) {
	s[conn] = struct{}{}
}

// remove a conn
func (s ConnSet) Del(conn *Conn) {
	delete(s, conn)
}

// the size of the set
func (s ConnSet) Size() int {
	return len(s)
}

// transform to a slice
func (s ConnSet) Slice() []*Conn {
	var v []*Conn
	for k := range s {
		v = append(v, k)
	}
	return v
}
