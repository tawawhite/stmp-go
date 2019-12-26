// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:31:13
package stmp

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"net"
	"sync"
)

type SendContext struct {
	Action   int64
	Data     proto.Message
	Payloads map[string][]byte
}

type Listener interface {
	Accept() (Conn, error)
	Close() error
}

type AuthenticateFunc func(conn Conn, resHeaders Headers) (status Status, message string, err error)

type Group struct {
	name  string
	mu    *sync.RWMutex
	conns map[Conn]bool
}

func NewGroup(name string) *Group {
	return &Group{name: name, mu: &sync.RWMutex{}, conns: map[Conn]bool{}}
}

type Server struct {
	mu        *sync.Mutex
	auth      AuthenticateFunc
	listeners []Listener
	conns     map[Conn]bool
	groups    map[string]*Group
	done      chan error
}

var noAuth AuthenticateFunc = func(conn Conn, resHeaders Headers) (status Status, message string, err error) {
	return StatusOk, "OK", nil
}

type ServerOption func(srv *Server)

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		mu:        &sync.Mutex{},
		auth:      noAuth,
		listeners: nil,
		conns:     map[Conn]bool{},
		groups:    map[string]*Group{},
		done:      make(chan error),
	}
	for _, o := range opts {
		o(srv)
	}
	return srv
}

func WithAuthenticate(authFn AuthenticateFunc) ServerOption {
	return func(srv *Server) {
		srv.auth = authFn
	}
}

func (s *Server) handleConn(conn Conn) {

}

func (s *Server) shutdown(err error) {
	s.mu.Lock()
	lis := s.listeners
	s.listeners = nil
	s.mu.Unlock()
	if lis == nil {
		return
	}
	for _, l := range lis {
		err := l.Close()
		if err != nil {
			panic(err)
		}
	}
	s.done <- err
}

func (s *Server) serve(lis Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				continue
			}
			s.shutdown(err)
			break
		}
		go s.handleConn(conn)
	}
}

func (s *Server) Serve(lis ...Listener) error {
	if len(lis) == 0 {
		return errors.New("must serve 1 listener at least")
	}
	s.mu.Lock()
	if s.listeners != nil {
		s.mu.Unlock()
		return errors.New("server is listening already")
	}
	s.listeners = lis
	s.mu.Unlock()
	for _, l := range lis {
		go s.serve(l)
	}
	return <-s.done
}

func (s *Server) Close() {
	s.shutdown(nil)
}

func (s *Server) JoinGroup(group string, conn Conn) {

}
