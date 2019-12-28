// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 15:31:13
package stmp

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
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

type AuthenticateFunc func(reqHeader Header, resHeader Header) (status Status, message string, err error)

type Server struct {
	mu        *sync.Mutex
	listeners []io.Closer
	conns     map[Conn]bool
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

var noAuth AuthenticateFunc = func(reqHeader Header, resHeader Header) (status Status, message string, err error) {
	return StatusOk, "OK", nil
}

type ServerOption func(srv *Server)

func NewServer() *Server {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return &Server{
		mu:               &sync.Mutex{},
		listeners:        []io.Closer{},
		conns:            map[Conn]bool{},
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

type readonlyHeader interface {
	Get(key string) string
}

func (s *Server) logAccess(event string, localAddr string, remoteAddr string, h readonlyHeader) {
	if s.LogAccessFields == nil {
		return
	}
	fields := []zap.Field{
		zap.String("server", localAddr),
		zap.String("addr", remoteAddr),
	}
	for _, k := range s.LogAccessFields {
		fields = append(fields, zap.String(k, h.Get(k)))
	}
	if event == "" {
		event = "new connection"
	}
	s.Log.Info(event, fields...)
}

func (s *Server) writeConn(nc *netConn, w EncodingWriter) {
	for {
		buf := <-nc.writeChan
		if buf == nil {
			break
		}
		nc.SetWriteDeadline(time.Now().Add(s.WriteTimeout))
		_, err := w.Write(buf)
		if err != nil {
			// TODO
			break
		}
		err = w.Flush()
		if err != nil {
			// TODO
			break
		}
	}
}

func (s *Server) handleRequest(mid uint16, action uint64, payload []byte) ([]byte, error) {
	return nil, nil
}

func (s *Server) handleNotify(action uint64, payload []byte) {
}

func (s *Server) handleResponse(status Status, payload []byte) {
}

func (s *Server) handleClose(status byte, message string) {

}

func (s *Server) readConn(nc *netConn, r io.ReadCloser) {
	var action uint64
	var status byte
	var err error
	var mid uint16
	var h byte
	var fin bool
	var pure bool
	var kind MessageKind
	var ps uint64
	var payload []byte
	for {
		nc.SetReadDeadline(time.Now().Add(s.ReadTimeout))
		h, err = nc.ReadByte()
		if err != nil {
			nc.Terminate(StatusProtocolError, "")
			break
		}
		fin, kind, pure, err = ParseHead(h)
		if err != nil {
			nc.Terminate(StatusProtocolError, err.Error())
			return
		}
		switch kind {
		case MessageKindPing:
			// TODO handle ping
		case MessageKindClose:
			// receive close message, must response
			status, err = nc.ReadByte()
			payload = nil
			if !pure && err == nil {
				ps, err = binary.ReadUvarint(nc)
				if err == nil {
					payload = make([]byte, ps)
					_, err = nc.Read(payload)
				}
			}
			// TODO handle close safe
			s.handleClose(status, string(payload))
			return
		case MessageKindRequest:
			mid, err = nc.ReadInt16()
			if err != nil {
				nc.Terminate(StatusProtocolError, err.Error())
				return
			}
			action, err = binary.ReadUvarint(nc)
			if err != nil {
				nc.Terminate(StatusProtocolError, err.Error())
				return
			}
			payload = nil
			if !pure {
				ps, err = binary.ReadUvarint(nc)
				if err != nil {
					nc.Terminate(StatusProtocolError, err.Error())
					return
				}
				payload = make([]byte, ps)
				_, err = nc.Read(payload)
				if err != nil {
					nc.Terminate(StatusProtocolError, err.Error())
					return
				}
			}
			if fin {
				// TODO this should be handled by conn, and server should implement
				// 	a interface for adapt client and server side
				s.handleRequest(mid, action, payload)
			} else {
				if payload == nil {
					payload = make([]byte, 0)
				}
				nc.pendingIncoming[mid] = &incomingEvent{kind: kind, payload: payload, action: action}
				nc.CheckPendingVolume()
			}
		case MessageKindNotify:
			if !fin {
				mid, err = nc.ReadInt16()
				if err != nil {
					nc.Terminate(StatusProtocolError, "")
					return
				}
			}
			action, err = binary.ReadUvarint(nc)
			if err != nil {
				nc.Terminate(StatusProtocolError, "")
				return
			}
			if pure {
				payload = nil
			} else {
				ps, err = binary.ReadUvarint(nc)
				if err != nil {
					nc.Terminate(StatusProtocolError, "")
					return
				}
				payload = make([]byte, ps)
			}
			if fin {
				s.handleNotify(action, payload)
			} else {
				nc.pendingIncoming[mid] = &incomingEvent{kind: kind, payload: payload, action: action}
			}
		case MessageKindResponse:
		}
	}
}

func (s *Server) HandleConn(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(s.HandshakeTimeout))
	fixHead := make([]byte, 6)
	_, err := conn.Read(fixHead)
	if err != nil {
		conn.Close()
		return
	}
	if !bytes.Equal(fixHead[0:4], []byte("STMP")) {
		conn.Close()
		return
	}
	nc := newNetConn(conn)
	nc.major = fixHead[4]
	nc.minor = fixHead[5]
	// length
	n, err := binary.ReadVarint(nc)
	if err != nil {
		conn.Close()
		return
	}
	rawHeader := make([]byte, n)
	resHeader := NewHeader()
	_, err = conn.Read(rawHeader)
	if err != nil {
		nc.Handshake(StatusBadRequest, resHeader, "read handshake header error")
		return
	}
	err = nc.header.Unmarshal(rawHeader)
	if err != nil {
		nc.Handshake(StatusProtocolError, resHeader, "parse handshake header error")
		return
	}
	mediaInput := nc.header.Get(AcceptContentType)
	inputLength := 0
	var inputValue string
	for inputLength < len(mediaInput) {
		inputValue, inputLength = ReadNegotiate(mediaInput)
		if nc.media = GetMediaCodec(inputValue); nc.media != nil {
			resHeader.Set(DetermineContentType, inputValue)
			break
		}
	}
	if nc.media == nil {
		nc.Handshake(StatusUnsupportedContentType, resHeader, "")
		return
	}
	encodingInput := nc.header.Get(AcceptEncoding)
	inputLength = 0
	var encoding EncodingCodec
	for inputLength < len(encodingInput) {
		inputValue, inputLength = ReadNegotiate(encodingInput)
		if encoding = GetEncodingCodec(inputValue); encoding != nil {
			resHeader.Set(DetermineEncoding, inputValue)
			break
		}
	}
	var wc EncodingWriter
	var rc io.ReadCloser
	if encoding != nil {
		wc, err = encoding.Writer(conn, s.CompressLevel)
		if err != nil {
			nc.Handshake(StatusInternalServerError, resHeader, "")
			return
		}
		rc, err = encoding.Reader(conn)
		if err != nil {
			nc.Handshake(StatusInternalServerError, resHeader, "")
			return
		}
	} else {
		wc = nc
		rc = nc
	}
	status, message, err := s.Authenticate(nc.header, resHeader)
	if err != nil {
		if status == StatusOk {
			if se, ok := err.(*StatusError); ok {
				status = se.code
				message = se.err.Error()
			}
		}
		nc.Handshake(status, resHeader, message)
		return
	}
	go s.writeConn(nc, wc)
	go s.readConn(nc, rc)
}

func (s *Server) HandleWebsocketConn(conn *websocket.Conn, req *http.Request) {
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

func (s *Server) Serve(lis net.Listener) {
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
			s.logAccess("404 not found", addr, "ws://"+q.RemoteAddr+q.RequestURI, q.Header)
			w.WriteHeader(404)
			return
		}
		conn, err := up.Upgrade(w, q, nil)
		if err != nil {
			s.logAccess("500 upgrade error: "+err.Error(), addr, "ws://"+q.RemoteAddr+q.RequestURI, q.Header)
			return
		}
		s.HandleWebsocketConn(conn, q)
	}
	hs := &http.Server{Addr: addr, Handler: handler}
	s.mu.Lock()
	s.listeners = append(s.listeners, hs)
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
