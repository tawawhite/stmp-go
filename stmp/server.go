// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:31:13
package stmp

import (
	"crypto/tls"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/xtaci/kcp-go"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
)

type SendContext struct {
	Action   int64
	Data     proto.Message
	Payloads map[string][]byte
}

type AuthenticateFunc func(conn Conn, resHeaders Header) (status Status, message string, err error)

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
	listeners []io.Closer
	conns     map[Conn]bool
	groups    map[string]*Group
	done      chan error
}

var noAuth AuthenticateFunc = func(conn Conn, resHeaders Header) (status Status, message string, err error) {
	return StatusOk, "OK", nil
}

type ServerOption func(srv *Server)

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		mu:        &sync.Mutex{},
		auth:      noAuth,
		listeners: []io.Closer{},
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

func (s *Server) handleNetConn(conn net.Conn) {

}

func (s *Server) handleWsConn(conn *websocket.Conn, req *http.Request) {
}

func (s *Server) shutdown(err error) {
	s.mu.Lock()
	lis := s.listeners
	s.listeners = nil
	s.mu.Unlock()
	if lis == nil {
		// shutdown already
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

func (s *Server) Wait() error {
	return <-s.done
}

func (s *Server) Close() {
	s.shutdown(nil)
}

func (s *Server) Accept(lis net.Listener) {
	s.mu.Lock()
	s.listeners = append(s.listeners, lis)
	s.mu.Unlock()
	for {
		conn, err := lis.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				continue
			}
			s.shutdown(err)
			break
		}
		go s.handleNetConn(conn)
	}
}

func (s *Server) ServeTCP(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		s.shutdown(err)
		return
	}
	s.Accept(lis)
}

func (s *Server) ServeTCPWithTLS(addr string, certFile, keyFile string) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		s.shutdown(err)
		return
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	lis, err := tls.Listen("tcp", addr, cfg)
	if err != nil {
		s.shutdown(err)
		return
	}
	s.Accept(lis)
}

func (s *Server) ServeKCP(addr string) {
	lis, err := kcp.Listen(addr)
	if err != nil {
		s.shutdown(err)
		return
	}
	s.Accept(lis)
}

func (s *Server) ServeKCPWithTLS(addr string, certFile, keyFile string) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		s.shutdown(err)
		return
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	lis, err := kcp.Listen(addr)
	if err != nil {
		s.shutdown(err)
		return
	}
	lis = tls.NewListener(lis, cfg)
	s.Accept(lis)
}

func (s *Server) newWsServer(addr, path string) *http.Server {
	upgrader := &websocket.Upgrader{}
	mu := http.NewServeMux()
	mu.HandleFunc("/"+strings.TrimPrefix(path, "/"), func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			s.shutdown(err)
			return
		}
		s.handleWsConn(conn, request)
	})
	hs := &http.Server{Addr: addr, Handler: mu}
	s.mu.Lock()
	s.listeners = append(s.listeners, hs)
	s.mu.Unlock()
	return hs
}

func (s *Server) ServeWS(addr, path string) {
	hs := s.newWsServer(addr, path)
	err := hs.ListenAndServe()
	if err != nil {
		s.shutdown(err)
	}
}

func (s *Server) ServeWSWithTLS(addr, path, certFile, keyFile string) {
	hs := s.newWsServer(addr, path)
	err := hs.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		s.shutdown(err)
	}
}
