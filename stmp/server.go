// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:31:13
package stmp

import (
	"bytes"
	"crypto/tls"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/xtaci/kcp-go"
	"go.uber.org/zap"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

type SendContext struct {
	Action   int64
	Data     proto.Message
	Payloads map[string][]byte
}

type AuthenticateFunc func(c *Conn) (err error)

type Server struct {
	*Router
	mu        *sync.Mutex
	listeners map[io.Closer]bool
	conns     map[*Conn]bool
	done      chan error
	Id        string
	Log       *zap.Logger
	// default is host, user-agent
	// if set as nil, will not log access
	LogAccessFields []string
	// [1, 9], default is 6
	CompressLevel    int
	Authenticate     AuthenticateFunc
	MaxPacketSize    uint64
	HandshakeTimeout time.Duration
	WriteTimeout     time.Duration
	ReadTimeout      time.Duration
}

var noAuth AuthenticateFunc = func(c *Conn) (err error) {
	return nil
}

func NewServer() *Server {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return &Server{
		Router:           NewRouter(),
		mu:               &sync.Mutex{},
		listeners:        map[io.Closer]bool{},
		conns:            map[*Conn]bool{},
		done:             make(chan error),
		Id:               "",
		Log:              log.With(zap.String("source", "stmp")),
		LogAccessFields:  []string{"user-agent", "host"},
		CompressLevel:    6,
		Authenticate:     noAuth,
		MaxPacketSize:    1 << 24, // 16Mb
		HandshakeTimeout: time.Minute,
		WriteTimeout:     time.Minute,
		ReadTimeout:      time.Minute,
	}
}

func (s *Server) newClient(nc net.Conn) *Conn {
	c := newConn(nc)
	c.Router = s.Router
	c.ClientHeader = NewHeader()
	c.ServerHeader = NewHeader()
	c.handshakeTimeout = s.HandshakeTimeout
	c.readTimeout = s.ReadTimeout
	c.writeTimeout = s.WriteTimeout
	return c
}

func (s *Server) HandleConn(nc net.Conn) (status *StatusError) {
	c := s.newClient(nc)
	defer func() {
		if status != nil {
			c.ServerMessage = status.err.Error()
			c.writeBinaryHandshakeResponse(status.code)
			c.nc.Close()
		}
	}()
	nc.SetReadDeadline(time.Now().Add(s.HandshakeTimeout))
	fixHead := make([]byte, 6)
	_, err := nc.Read(fixHead)
	if err != nil {
		status = NewStatusError(StatusBadRequest, "read request error: "+err.Error())
		return
	}
	if !bytes.Equal(fixHead[0:4], []byte("STMP")) {
		status = NewStatusError(StatusProtocolError, "magic header is not STMP")
		return
	}
	c.Major = fixHead[4]
	c.Minor = fixHead[5]
	// length
	n, err := c.readUvarint()
	if err != nil {
		status = NewStatusError(StatusBadRequest, "read header length error: "+err.Error())
		return
	}
	rawHeader := make([]byte, n)
	_, err = nc.Read(rawHeader)
	if err != nil {
		status = NewStatusError(StatusBadRequest, "read header error: "+err.Error())
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
	err = s.Authenticate(c)
	if err != nil {
		if se, ok := err.(*StatusError); ok {
			status = se
		} else {
			status = NewStatusError(StatusInternalServerError, "authenticate error: "+err.Error())
		}
		return
	}
	r, w, err := c.initEncoding(s.CompressLevel)
	if err != nil {
		status = NewStatusError(StatusProtocolError, "init encoding error: "+err.Error())
		return
	}
	err = c.writeBinaryHandshakeResponse(StatusOk)
	if err != nil {
		c.nc.Close()
		return
	}
	s.mu.Lock()
	s.conns[c] = true
	s.mu.Unlock()
	go c.binaryReadChannel(r)
	go c.binaryWriteChannel(w)
	return
}

func (s *Server) HandleWebsocketConn(conn *websocket.Conn, req *http.Request) {
}

func (s *Server) shutdown(err error) {
	for l := range s.listeners {
		l.Close()
		delete(s.listeners, l)
	}
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
	s.listeners[lis] = true
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
		conn, err := up.Upgrade(w, q, nil)
		if err != nil {
			return
		}
		s.HandleWebsocketConn(conn, q)
	}
	hs := &http.Server{Addr: addr, Handler: handler}
	s.listeners[hs] = true
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
