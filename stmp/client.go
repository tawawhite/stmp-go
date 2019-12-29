// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-28 15:15:32
package stmp

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/xtaci/kcp-go"
	"io/ioutil"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type DialOptions struct {
	Header           Header
	HandshakeTimeout time.Duration
	WriteTimeout     time.Duration
	ReadTimeout      time.Duration
	PacketFormat     string
	Encoding         string
	ContentType      string
}

func dialOptionsDefaulter(opts *DialOptions) {
	if opts.HandshakeTimeout == 0 {
		opts.HandshakeTimeout = time.Minute
	}
	if opts.WriteTimeout == 0 {
		opts.WriteTimeout = time.Minute
	}
	if opts.ReadTimeout == 0 {
		opts.ReadTimeout = time.Minute
	}
	if opts.Header == nil {
		opts.Header = NewHeader()
	}
	if opts.PacketFormat != "" {
		opts.Header.Set(AcceptPacketFormat, opts.PacketFormat)
	}
	if opts.Encoding != "" {
		opts.Header.Set(AcceptEncoding, opts.Encoding)
	}
	if opts.ContentType == "" {
		if opts.PacketFormat == "text" {
			opts.ContentType = "application/json"
		} else {
			opts.ContentType = "application/protobuf"
		}
	}
	opts.Header.Set(AcceptContentType, opts.ContentType)
}

func DialTCP(addr string, opts *DialOptions) (conn *Conn, err error) {
	nc, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	conn := newConn()
}

func loadTLSClientConfig(certFile string) (*tls.Config, error) {
	b, err := ioutil.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		return nil, errors.New("tls failed to append certificates")
	}
	return &tls.Config{RootCAs: cp}, nil
}

func DialTCPWithTLS(addr, certFile string, opts *DialOptions) (*Conn, error) {
	tc, err := loadTLSClientConfig(certFile)
	if err != nil {
		return nil, err
	}
	nc, err := tls.Dial("tcp", addr, tc)
	if err != nil {
		return nil, err
	}
	return newConn(nc, nil), nil
}

func DialKCP(addr string, opts *DialOptions) (*Conn, error) {
	nc, err := kcp.Dial(addr)
	if err != nil {
		return nil, err
	}
	return newConn(nc, nil), nil
}

func DialKCPWithTLS(addr, certFile string, opts *DialOptions) (*Conn, error) {
	tc, err := loadTLSClientConfig(certFile)
	if err != nil {
		return nil, err
	}
	nc, err := kcp.Dial(addr)
	if err != nil {
		return nil, err
	}
	tlsConn := tls.Client(nc, tc)
	if err = tlsConn.Handshake(); err != nil {
		return nil, err
	}
	return newConn(tlsConn, nil), nil
}

// the gorilla/websocket only accepts public verified tls config for tls connection
func DialWebSocket(urlStr string, opts *DialOptions) (conn *Conn, err error) {
	dialOptionsDefaulter(opts)
	defer func() {
		if err != nil && conn != nil {
			conn.Close(StatusProtocolError, "")
			conn = nil
		}
	}()
	if len(opts.Header) > 0 {
		headStr := url.Values(opts.Header).Encode()
		if strings.IndexByte(urlStr, '?') > 0 {
			urlStr += "&" + headStr
		} else {
			urlStr += "?" + headStr
		}
	}
	dialer := websocket.Dialer{HandshakeTimeout: opts.HandshakeTimeout}
	wc, _, err := dialer.Dial(urlStr, nil)
	if err != nil {
		return
	}
	conn = newConn()
	conn.ClientHeader = opts.Header
	conn.nc = wc.UnderlyingConn()
	conn.wc = wc
	kind, data, err := wc.ReadMessage()
	if err != nil {
		err = NewStatusError(StatusNetworkError, err)
		return
	}
	if len(data) < 6 || bytes.Equal(data[0:4], []byte("STMP")) {
		err = StatusProtocolError
		return
	}
	var sep int
	if kind == websocket.TextMessage {
		sep = bytes.IndexByte(data, '\n')
		if sep == -1 {
			err = StatusProtocolError
			return
		}
		status, err := strconv.ParseUint(string(data[4:sep]), 16, 8)
		if err != nil {
			err = NewStatusError(StatusProtocolError, err)
			return
		}
		if status != 0 {
			err = Status(status)
			return
		}
		data = data[sep+1:]
	} else {
		if data[4] != 0 {
			err = Status(data[4])
			return
		}
		data = data[5:]
	}
	sep = bytes.Index(data, []byte("\n\n"))
	if sep == -1 {
		sep = len(data)
	} else {
		conn.HandshakeMessage = string(data[sep+2])
	}
	conn.ServerHeader = NewHeader()
	err = conn.ServerHeader.Unmarshal(data[0:sep])
	if err != nil {
		return
	}
	// ws do not process encoding
	conn.media = GetMediaCodec(conn.ServerHeader.Get(DetermineContentType))
	return
}
