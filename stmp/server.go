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
	"strings"
	"sync"
	"time"
)

type SendContext struct {
	Action   int64
	Data     proto.Message
	Payloads map[string][]byte
}

type AuthenticateFunc func(reqHeader Header, resHeader Header) (status Status, message string, err error)

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
	listeners []io.Closer
	conns     map[Conn]bool
	groups    map[string]*Group
	done      chan error
	Id        string
	Log       *zap.Logger
	// [1, 9]
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
		groups:           map[string]*Group{},
		done:             make(chan error),
		Id:               "",
		Log:              log.With(zap.String("source", "stmp")),
		CompressLevel:    6,
		Authenticate:     noAuth,
		MaxPacketSize:    1 << 24, // 16Mb
		HandshakeTimeout: time.Minute,
		WriteTimeout:     time.Minute,
		ReadTimeout:      time.Minute,
	}
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
	if encoding != nil {
		wc, err = encoding.Writer(conn, s.CompressLevel)
		if err != nil {
			nc.Handshake(StatusInternalServerError, resHeader, "")
			return
		}
	} else {
		wc = nc
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
	for {
		nc.SetReadDeadline(time.Now().Add(s.ReadTimeout))
		h, err := nc.ReadByte()
		if err != nil {
			nc.Terminate(StatusProtocolError, "")
			break
		}
		fin, kind, pure, err := ParseHead(h)
		if err != nil {
			nc.Terminate(StatusProtocolError, err.Error())
			return
		}
		switch kind {
		case MessageKindPing:
			nc.writeChan <- PingMessage
		case MessageKindClose:
			// TODO
			nc.Close()
		case MessageKindRequest:
			mid, err := nc.ReadUint16()
			if err != nil {
				nc.Terminate(StatusProtocolError, err.Error())
				return
			}
			action, err := binary.ReadUvarint(nc)
			if err != nil {
				nc.Terminate(StatusProtocolError, err.Error())
				return
			}
			var payload []byte
			if !pure {
				ps, err := binary.ReadUvarint(nc)
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
				s.handleRequest(mid, action, payload)
			} else {
				if payload == nil {
					payload = make([]byte, 0)
				}
				nc.pendingRequest[mid] = &incomingEvent{payload: payload, action: action}
				nc.CheckPendingVolume()
			}
		case MessageKindNotify:
		case MessageKindResponse:
		case MessageKindFollowing:
		}
	}
}

func (s *Server) handleRequest(mid uint16, action uint64, payload []byte) {
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
	upgrader := &websocket.Upgrader{}
	mu := http.NewServeMux()
	mu.HandleFunc("/"+strings.TrimPrefix(path, "/"), func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			s.shutdown(err)
			return
		}
		s.HandleWebsocketConn(conn, request)
	})
	hs := &http.Server{Addr: addr, Handler: mu}
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
