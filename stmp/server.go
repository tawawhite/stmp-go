// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:31:13
package stmp

import (
	"context"
	"crypto/tls"
	"github.com/gorilla/websocket"
	"github.com/xtaci/kcp-go"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

func init() {
	RegisterMediaCodec(NewMsgpackCodec(), NewJsonCodec())
	RegisterEncodingCodec(NewGzipCodec())
}

type ServerOptions struct {
	*ConnOptions
	LogAccess     []string
	CompressLevel int
	Authenticate  func(c *Conn) error
}

func (o *ServerOptions) WithLogAccess(fields ...string) *ServerOptions {
	o.LogAccess = fields
	return o
}

func (o *ServerOptions) WithCompress(level int) *ServerOptions {
	o.CompressLevel = level
	return o
}

func (o *ServerOptions) WithAuthenticate(fn func(c *Conn) error) *ServerOptions {
	o.Authenticate = fn
	return o
}

func (o *ServerOptions) ApplyDefault() *ServerOptions {
	no := o
	if no == nil {
		no = NewServerOptions()
	}
	no.ConnOptions = no.ConnOptions.ApplyDefault()
	return no
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		ConnOptions:   NewConnOptions(),
		LogAccess:     []string{"host", "user-agent", "referer"},
		CompressLevel: 6,
		Authenticate:  func(c *Conn) error { return nil },
	}
}

type Server struct {
	*Router
	opts      *ServerOptions
	mu        sync.RWMutex
	listeners map[io.Closer]struct{}
	conns     ConnSet
	done      chan error
}

func NewServer(opts *ServerOptions) *Server {
	opts = opts.ApplyDefault()
	return &Server{
		Router:    opts.Router,
		opts:      opts,
		listeners: map[io.Closer]struct{}{},
		conns:     NewConnSet(),
		done:      make(chan error, 1),
	}
}

type ConnFilter func(conn *Conn) bool

var AllowAll ConnFilter = func(conn *Conn) bool { return true }

func (s *Server) Broadcast(ctx context.Context, method string, in interface{}, filter ConnFilter) error {
	if filter == nil {
		filter = AllowAll
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
	c := NewConn(nc, s.opts.ConnOptions)
	c.ClientHeader = NewHeader()
	c.ServerHeader = NewHeader()
	return c
}

func (s *Server) HandleConn(nc net.Conn) (err error) {
	c := s.newClient(nc)
	defer func() {
		if err != nil {
			// TODO write handshake response
			c.Conn.Close()
		}
	}()
	ch := NewClientHandshake(0, 0, c.ClientHeader, "")
	nc.SetReadDeadline(time.Now().Add(s.opts.HandshakeTimeout))
	err = ch.Read(nc, s.opts.MaxPacketSize)
	if err != nil {
		return
	}
	err = c.negotiate()
	if err != nil {
		return
	}
	err = s.opts.Authenticate(c)
	if err != nil {
		err = DetectError(err, StatusUnauthorized)
		return
	}
	r, w, err := c.initEncoding(s.opts.CompressLevel)
	if err != nil {
		err = NewStatusError(StatusProtocolError, "init encoding error: "+err.Error())
		return
	}
	sh := NewServerHandshake(StatusOk, c.ServerHeader, "")
	err = sh.Write(nc)
	if err != nil {
		return
	}
	s.mu.Lock()
	s.conns.Add(c)
	s.mu.Unlock()
	go c.binaryWriteChannel(w)
	go func() {
		c.binaryReadChannel(r)
		// TODO close connection
	}()
	return
}

func (s *Server) HandleWebsocketConn(wc *websocket.Conn, req *http.Request) (status *StatusError) {
	c := s.newClient(wc.UnderlyingConn())
	defer func() {
		if status != nil {
			// TODO write handshake response
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
	err := s.opts.Authenticate(c)
	if err != nil {
		if se, ok := err.(*StatusError); ok {
			status = se
		} else {
			status = NewStatusError(StatusInternalServerError, "authenticate error: "+err.Error())
		}
		return
	}
	sh := NewServerHandshake(StatusOk, c.ServerHeader, "")
	var data []byte
	format := c.ServerHeader.Get(DeterminePacketFormat)
	var typ int
	if format == "text" {
		typ = websocket.TextMessage
		data = sh.MarshalText()
	} else {
		typ = websocket.BinaryMessage
		data = sh.MarshalBinary()
	}
	err = wc.WriteMessage(typ, data)
	if err != nil {
		return
	}
	s.mu.Lock()
	s.conns.Add(c)
	s.mu.Unlock()
	if c.ServerHeader.Get(DeterminePacketFormat) == "text" {
		go c.websocketTextWriteChannel(wc)
		go func() {
			c.websocketTextReadChannel(wc)
		}()
	} else {
		go c.websocketBinaryWriteChannel(wc)
		go func() {
			c.websocketBinaryReadChannel(wc)
		}()
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
