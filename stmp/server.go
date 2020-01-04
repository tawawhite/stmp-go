// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:31:13
package stmp

import (
	"bytes"
	"context"
	"crypto/tls"
	"github.com/gorilla/websocket"
	"github.com/xtaci/kcp-go"
	"go.uber.org/zap"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

type ServerConfig struct {
	Logger           *zap.Logger
	LogAccess        []string
	CompressLevel    int
	Authenticate     func(c *Conn) error
	MaxPacketSize    uint64
	HandshakeTimeout time.Duration
	WriteTimeout     time.Duration
	ReadTimeout      time.Duration
}

func (config *ServerConfig) ApplyDefault() *ServerConfig {
	out := config
	if out == nil {
		out = NewServerConfig()
	}
	if out.Logger == nil {
		var err error
		out.Logger, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	}
	return out
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		LogAccess:        []string{"host", "user-agent", "referer"},
		CompressLevel:    6,
		Authenticate:     func(c *Conn) error { return nil },
		MaxPacketSize:    1 << 24,
		HandshakeTimeout: time.Minute,
		WriteTimeout:     time.Minute,
		ReadTimeout:      time.Minute,
	}
}

type Server struct {
	*router
	config    *ServerConfig
	mu        sync.RWMutex
	listeners map[io.Closer]struct{}
	conns     ConnSet
	done      chan error
}

func NewServer(config *ServerConfig) *Server {
	config = config.ApplyDefault()
	return &Server{
		router:    NewRouter(),
		config:    config,
		listeners: map[io.Closer]struct{}{},
		conns:     NewConnSet(),
		done:      make(chan error),
	}
}

type ConnFilter func(conn *Conn) bool

var filterAll ConnFilter = func(conn *Conn) bool {
	return true
}

func (s *Server) Broadcast(ctx context.Context, method string, in interface{}, filter ConnFilter) error {
	if filter == nil {
		filter = filterAll
	}
	payloads := NewPayloadMap(in)
	s.mu.RLock()
	defer s.mu.RUnlock()
	for conn := range s.conns {
		if filter(conn) {
			payload, err := payloads.Marshal(conn)
			if err != nil {
				return err
			}
			_, err = conn.Call(ctx, method, payload, NotifyOptions)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Server) newClient(nc net.Conn) *Conn {
	c := newConn(nc)
	c.router = s.router
	c.ClientHeader = NewHeader()
	c.ServerHeader = NewHeader()
	return c
}

func (s *Server) HandleConn(nc net.Conn) (status *StatusError) {
	c := s.newClient(nc)
	defer func() {
		if status != nil {
			c.ServerMessage = status.err.Error()
			c.writeHandshakeResponse(status.code)
			c.Conn.Close()
		}
	}()
	nc.SetReadDeadline(time.Now().Add(s.config.HandshakeTimeout))
	fixHead := make([]byte, 6)
	_, err := nc.Read(fixHead)
	if err != nil {
		status = NewStatusError(StatusBadRequest, "Read request error: "+err.Error())
		return
	}
	if !bytes.Equal(fixHead[0:4], []byte("STMP")) {
		status = NewStatusError(StatusProtocolError, "magic header is not STMP")
		return
	}
	c.Major = fixHead[4]
	c.Minor = fixHead[5]
	if c.Major != 1 || c.Minor != 0 {
		status = NewStatusError(StatusUnsupportedProtocolVersion, "unsupported STMP version: "+string([]byte{c.Major + '0', '.', c.Minor + '0'}))
	}
	// length
	n, err := ReadUvarint(nc, fixHead[:1])
	if err != nil {
		status = NewStatusError(StatusBadRequest, "Read header length error: "+err.Error())
		return
	}
	rawHeader := make([]byte, n)
	_, err = nc.Read(rawHeader)
	if err != nil {
		status = NewStatusError(StatusBadRequest, "Read header error: "+err.Error())
		return
	}
	err = c.ClientHeader.Unmarshal(rawHeader)
	if err != nil {
		status = NewStatusError(StatusBadRequest, "parse header error: "+err.Error())
		return
	}
	status = c.negotiate()
	if status != nil {
		return
	}
	err = s.config.Authenticate(c)
	if err != nil {
		if se, ok := err.(*StatusError); ok {
			status = se
		} else {
			status = NewStatusError(StatusInternalServerError, "authenticate error: "+err.Error())
		}
		return
	}
	r, w, err := c.initEncoding(s.config.CompressLevel)
	if err != nil {
		status = NewStatusError(StatusProtocolError, "init encoding error: "+err.Error())
		return
	}
	err = c.writeHandshakeResponse(StatusOk)
	if err != nil {
		c.Conn.Close()
		return
	}
	s.mu.Lock()
	s.conns.Add(c)
	s.mu.Unlock()
	go c.binaryReadChannel(r, s.config.ReadTimeout)
	go c.binaryWriteChannel(w, s.config.WriteTimeout)
	return
}

func (s *Server) HandleWebsocketConn(wc *websocket.Conn, req *http.Request) (status *StatusError) {
	c := s.newClient(wc.UnderlyingConn())
	defer func() {
		if status != nil {
			c.ServerMessage = status.err.Error()
			c.websocketWriteHandshakeResponse(status.code)
			wc.Close()
		}
	}()
	for k, v := range req.Header {
		c.ClientHeader.Set(k, v...)
	}
	for k, v := range req.URL.Query() {
		c.ClientHeader.Set(k, v...)
	}
	rawVersion := c.ClientHeader.Get(DetermineStmpVersion)
	if len(rawVersion) != 3 {
		status = NewStatusError(StatusUnsupportedProtocolVersion, "unsupported STMP version: "+rawVersion)
		return
	}
	c.Major = rawVersion[0] - '0'
	c.Minor = rawVersion[2] - '0'
	if c.Major != 1 || c.Minor != 0 {
		status = NewStatusError(StatusUnsupportedProtocolVersion, "unsupported STMP version: "+rawVersion)
		return
	}
	status = c.negotiate()
	if status != nil {
		return
	}
	err := s.config.Authenticate(c)
	if err != nil {
		if se, ok := err.(*StatusError); ok {
			status = se
		} else {
			status = NewStatusError(StatusInternalServerError, "authenticate error: "+err.Error())
		}
		return
	}
	err = c.websocketWriteHandshakeResponse(StatusOk)
	if err != nil {
		wc.Close()
		return
	}
	s.mu.Lock()
	s.conns.Add(c)
	s.mu.Unlock()
	if c.ServerHeader.Get(DeterminePacketFormat) == "text" {
		go c.websocketTextReadChannel(wc, s.config.ReadTimeout)
		go c.websocketTextWriteChannel(wc, s.config.WriteTimeout)
	} else {
		go c.websocketBinaryReadChannel(wc, s.config.ReadTimeout)
		go c.websocketBinaryWriteChannel(wc, s.config.WriteTimeout)
	}
	return
}

func (s *Server) shutdown(err error) {
	s.mu.Lock()
	for l := range s.listeners {
		l.Close()
		delete(s.listeners, l)
	}
	for c := range s.conns {
		c.Close(StatusServerShutdown, "")
		delete(s.conns, c)
	}
	s.mu.Unlock()
	s.done <- err
}

func (s *Server) Wait() error {
	return <-s.done
}

func (s *Server) Close() {
	s.shutdown(nil)
}

func (s *Server) Serve(lis net.Listener) {
	if _, ok := s.listeners[lis]; ok {
		panic("the listener " + lis.Addr().Network() + ":" + lis.Addr().String() + " is listening already")
	}
	s.mu.Lock()
	s.listeners[lis] = struct{}{}
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
		go s.HandleConn(conn)
	}
}

func (s *Server) ListenAndServeTCP(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		s.shutdown(err)
		return
	}
	s.Serve(lis)
}

func (s *Server) ListenAndServeTCPWithTLS(addr string, certFile, keyFile string) {
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
	s.Serve(lis)
}

func (s *Server) ListenAndServeKCP(addr string) {
	lis, err := kcp.Listen(addr)
	if err != nil {
		s.shutdown(err)
		return
	}
	s.Serve(lis)
}

func (s *Server) ListenAndServeKCPWithTLS(addr string, certFile, keyFile string) {
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
	s.Serve(lis)
}

func (s *Server) newWsServer(addr, path string) *http.Server {
	if path == "" {
		path = "/"
	}
	up := &websocket.Upgrader{}
	var handler http.HandlerFunc = func(w http.ResponseWriter, q *http.Request) {
		if q.URL.Path != path {
			w.WriteHeader(404)
			return
		}
		wc, err := up.Upgrade(w, q, nil)
		if err != nil {
			return
		}
		s.HandleWebsocketConn(wc, q)
	}
	hs := &http.Server{Addr: addr, Handler: handler}
	s.mu.Lock()
	s.listeners[hs] = struct{}{}
	s.mu.Unlock()
	return hs
}

func (s *Server) ListenAndServeWebSocket(addr, path string) {
	hs := s.newWsServer(addr, path)
	err := hs.ListenAndServe()
	if err != nil {
		s.shutdown(err)
	}
}

func (s *Server) ListenAndServeWebSocketWithTLS(addr, path, certFile, keyFile string) {
	hs := s.newWsServer(addr, path)
	err := hs.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		s.shutdown(err)
	}
}
