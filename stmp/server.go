// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:31:13
package stmp

import (
	"context"
	"errors"
	"github.com/acrazing/stmp-go/stmp/md"
	"io"
	"net"
	"sync"
)

type Message interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
}

type SendOptions struct {
	ActionId   uint32
	ActionName string
	Data       Message
}

type Context interface {
	context.Context
	Conn() Connection
	Server() *Server
	State() interface{}
	Async(func() error)
}

type Connection interface {
	io.ReadWriteCloser
	Headers() md.Metadata
	Addr() net.TCPAddr
	Notify(options SendOptions) error
}

type Listener interface {
	Accept() (Connection, error)
	Close() error
}

type AuthenticateFunc func(conn Connection) error

type Server struct {
	mu   *sync.Mutex
	auth AuthenticateFunc
	lis  []Listener
	done chan error
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) handleConn(conn Connection) {

}

func (s *Server) shutdown(err error) {
	s.mu.Lock()
	lis := s.lis
	s.lis = nil
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
			if ne, ok := err.(interface {
				Temporary() bool
			}); ok && ne.Temporary() {
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
	if s.lis != nil {
		s.mu.Unlock()
		return errors.New("server is listening already")
	}
	s.lis = lis
	s.mu.Unlock()
	for _, l := range lis {
		go s.serve(l)
	}
	return <-s.done
}

func (s *Server) Close() {
	s.shutdown(nil)
}
