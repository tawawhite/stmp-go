package stmp

import "sync"

type State struct {
	sync.RWMutex
	state map[interface{}]interface{}
}

func NewState() State {
	return State{state: make(map[interface{}]interface{})}
}

func (s State) Set(key, value interface{}) {
	s.Lock()
	s.state[key] = value
	s.Unlock()
}

func (s State) Get(key interface{}) interface{} {
	s.RLock()
	v := s.state[key]
	s.RUnlock()
	return v
}

func (s State) Has(key interface{}) interface{} {
	s.RLock()
	ok := s.state[key]
	s.RUnlock()
	return ok
}

func (s State) Del(key interface{}) {
	s.Lock()
	delete(s.state, key)
	s.Unlock()
}

func (s State) SetUnsafe(key, value interface{}) {
	s.state[key] = value
}

func (s State) GetUnsafe(key interface{}) interface{} {
	v := s.state[key]
	return v
}

func (s State) HasUnsafe(key interface{}) interface{} {
	ok := s.state[key]
	return ok
}

func (s State) DelUnsafe(key interface{}) {
	delete(s.state, key)
}
