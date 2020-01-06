package stmp

import (
	"context"
	"crypto/tls"
	"github.com/gorilla/websocket"
	"github.com/xtaci/kcp-go"
	"go.uber.org/zap"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type ServerOptions struct {
	*ConnOptions
	logAccess    []string
	authenticate func(c *Conn) error
}

func (o *ServerOptions) WithLogger(logger *zap.Logger) *ServerOptions {
	o.ConnOptions.WithLogger(logger)
	return o
}

func (o *ServerOptions) WithRouter(r *Router) *ServerOptions {
	o.ConnOptions.WithRouter(r)
	return o
}

func (o *ServerOptions) WithCompress(level int) *ServerOptions {
	o.ConnOptions.WithCompress(level)
	return o
}

func (o *ServerOptions) WithWriteQueueLimit(max int) *ServerOptions {
	o.ConnOptions.WithWriteQueueLimit(max)
	return o
}

func (o *ServerOptions) WithPacketSizeLimit(max uint64) *ServerOptions {
	o.ConnOptions.WithPacketSizeLimit(max)
	return o
}

func (o *ServerOptions) WithTimeout(handshake, read, write time.Duration) *ServerOptions {
	o.ConnOptions.WithTimeout(handshake, read, write)
	return o
}

func (o *ServerOptions) WithLogAccess(fields ...string) *ServerOptions {
	o.logAccess = fields
	return o
}

func (o *ServerOptions) WithAuthenticate(fn func(c *Conn) error) *ServerOptions {
	o.authenticate = fn
	return o
}

func (o *ServerOptions) applyDefault() *ServerOptions {
	no := o
	if no == nil {
		no = NewServerOptions()
	}
	no.ConnOptions = no.ConnOptions.ApplyDefault()
	return no
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		ConnOptions:  NewConnOptions(),
		logAccess:    []string{"host", "user-agent", "referer", AcceptContentType},
		authenticate: func(c *Conn) error { return nil },
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
	opts = opts.applyDefault()
	return &Server{
		Router:    opts.router,
		opts:      opts,
		listeners: map[io.Closer]struct{}{},
		conns:     NewConnSet(),
		done:      make(chan error, 1),
	}
}

type ConnFilter func(conn *Conn) bool

var AllowAll ConnFilter = func(conn *Conn) bool { return true }

func (s *Server) Broadcast(ctx context.Context, method string, in interface{}, filters ...ConnFilter) error {
	var filter ConnFilter
	if len(filters) == 0 {
		filter = AllowAll
	} else {
		filter = filters[0]
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

func (s *Server) newConn(nc net.Conn) *Conn {
	c := NewConn(nc, s.opts.ConnOptions)
	c.ClientHeader = NewHeader()
	c.ServerHeader = NewHeader()
	return c
}

func (s *Server) CloseConn(c *Conn, status Status, message string) error {
	err := c.close(status, message)
	s.mu.Lock()
	s.conns.Del(c)
	s.mu.Unlock()
	return err
}

func (s *Server) logAccess(c *Conn, event string, sh *handshake, err error) {
	if s.opts.logAccess == nil {
		return
	}
	fields := []zap.Field{
		zap.String("event", event),
		zap.String("addr", c.RemoteAddr().String()),
		zap.String("server", c.LocalAddr().String()),
		zap.Uint8("status", uint8(sh.Status)),
		zap.String("message", sh.Message),
	}
	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	if len(s.opts.logAccess) == 1 && s.opts.logAccess[0] == "*" {
		for k, v := range c.ClientHeader {
			if len(v) == 1 {
				fields = append(fields, zap.String(k, v[0]))
			} else {
				fields = append(fields, zap.Strings(k, v))
			}
		}
	} else {
		for _, k := range s.opts.logAccess {
			v := c.ClientHeader.GetAll(k)
			if len(v) == 1 {
				fields = append(fields, zap.String(k, v[0]))
			} else if v != nil {
				fields = append(fields, zap.Strings(k, v))
			}
		}
	}
	s.opts.logger.Info("access", fields...)
}

func (s *Server) prepare(c *Conn) (err error) {
	if c.Major != 1 || c.Minor != 0 {
		err = NewStatusError(StatusUnsupportedProtocolVersion, "")
		return
	}
	err = c.negotiate()
	if err != nil {
		return
	}
	err = s.opts.authenticate(c)
	if err != nil {
		err = DetectError(err, StatusUnauthorized)
		return
	}
	return
}

func (s *Server) HandleConn(nc net.Conn) (err error) {
	c := s.newConn(nc)
	closeSent := false
	var ec EncodingCodec
	ch := NewClientHandshake(0, 0, c.ClientHeader, "")
	sh := NewServerHandshake(StatusOk, c.ServerHeader, "")
	defer func() {
		// defer log access, close error connection
		if !closeSent && err != nil {
			se := DetectError(err, StatusInternalServerError)
			sh.Status = se.Code()
			sh.Message = se.Message()
			err = sh.Write(nc)
			nc.Close()
		} else if err != nil {
			nc.Close()
		}
		s.logAccess(c, "connect", sh, err)
	}()
	nc.SetReadDeadline(time.Now().Add(s.opts.handshakeTimeout))
	if err = ch.Read(nc, s.opts.maxPacketSize); err != nil {
		return
	}
	if err = s.prepare(c); err != nil {
		return
	}
	if ec, err = c.initEncoding(); err != nil {
		return
	}
	closeSent = true
	if err = sh.Write(nc); err != nil {
		return
	}

	// fine
	s.mu.Lock()
	s.conns.Add(c)
	s.mu.Unlock()
	go c.binaryWriteChannel(ec)
	go func() {
		c.binaryReadChannel(ec)
		// TODO close connection
	}()
	return
}

func (s *Server) writeWebsocketHandshake(wc *websocket.Conn, c *Conn, sh *handshake) error {
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
	return wc.WriteMessage(typ, data)
}

func (s *Server) HandleWebsocketConn(wc *websocket.Conn, req *http.Request) (err error) {
	c := s.newConn(wc.UnderlyingConn())
	sh := NewServerHandshake(StatusOk, c.ServerHeader, "")
	closeSent := false
	defer func() {
		if err != nil {
			sh := NewServerHandshake(StatusOk, c.ServerHeader, "")
			defer func() {
				// defer log access, close error connection
				if !closeSent && err != nil {
					se := DetectError(err, StatusInternalServerError)
					sh.Status = se.Code()
					sh.Message = se.Message()
					err = s.writeWebsocketHandshake(wc, c, sh)
					wc.Close()
				} else if err != nil {
					wc.Close()
				}
				s.logAccess(c, "connect", sh, err)
			}()
		}
	}()

	// transfer headers to conn
	for k, v := range req.Header {
		c.ClientHeader.Set(k, v...)
	}
	for k, v := range req.URL.Query() {
		c.ClientHeader.Set(k, v...)
	}
	rawVersion := c.ClientHeader.Get(DetermineStmpVersion)
	if len(rawVersion) != 3 {
		err = NewStatusError(StatusProtocolError, "invalid protocol version: "+rawVersion)
		return
	}
	c.Major = hexChunks[rawVersion[0]]
	c.Minor = hexChunks[rawVersion[2]]

	if err = s.prepare(c); err != nil {
		return
	}
	closeSent = true
	if err = s.writeWebsocketHandshake(wc, c, sh); err != nil {
		return
	}

	// fine
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
	defer s.mu.Unlock()
	if s.listeners == nil {
		// closed already
		return
	}
	for c := range s.conns {
		c.close(StatusServerShutdown, "")
		s.conns.Del(c)
	}
	for l := range s.listeners {
		l.Close()
		delete(s.listeners, l)
	}
	s.listeners = nil
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
