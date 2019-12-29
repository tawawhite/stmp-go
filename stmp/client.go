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
	// the headers for writeHandshakeResponse
	Header Header
	// writeHandshakeResponse timeout, each writeHandshakeResponse packet timeout
	// which means the final timeout is double
	HandshakeTimeout time.Duration
	// write timeout
	WriteTimeout time.Duration
	// read timeout
	ReadTimeout time.Duration
	// could be text or binary for websocket
	// else only could be binary
	// default is binary
	PacketFormat string
	// could be gzip, br, deflate
	// omitted for websocket
	// default is empty
	Encoding string
	// works only if encoding is not empty
	// range is [1, 9], default is 6
	CompressLevel int
	// could be application/json, application/protobuf, application/msgpack
	// default is application/protobuf
	ContentType string
	// the writeHandshakeResponse message
	Message string
	// the server name, if addr is ip & with tls, this is required
	ServerName string
	// if with tls and server name is empty, this should be true
	InsecureSkipVerify bool
}

func dialOptionsDefaulter(addr string, opts *DialOptions) {
	if opts.ServerName == "" {
		URL, err := url.Parse(addr)
		if err != nil {
			panic(err)
		}
		host := URL.Hostname()
		ip := net.ParseIP(host)
		if ip == nil {
			opts.ServerName = host
		}
	}
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
		if opts.CompressLevel == 0 {
			opts.CompressLevel = 6
		}
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

func loadTLSClientConfig(certFile string, opts *DialOptions) (*tls.Config, error) {
	if opts.ServerName == "" && !opts.InsecureSkipVerify {
		return nil, errors.New("opts.ServerName is required for tls connection")
	}
	b, err := ioutil.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		return nil, errors.New("tls failed to append certificates")
	}
	return &tls.Config{
		RootCAs:            cp,
		ServerName:         opts.ServerName,
		InsecureSkipVerify: opts.InsecureSkipVerify,
	}, nil
}

func newClientConn(nc net.Conn, opts *DialOptions) (c *Conn) {
	c = newConn(nc)
	c.Router = NewRouter()
	c.ClientHeader = opts.Header
	c.handshakeTimeout = opts.HandshakeTimeout
	c.readTimeout = opts.ReadTimeout
	c.writeTimeout = opts.WriteTimeout
	c.ClientHeader = opts.Header
	c.ClientMessage = opts.Message
	return
}

// create a client conn from nc, will writeHandshakeResponse automatically
func Client(nc net.Conn, opts *DialOptions) (c *Conn, err error) {
	c = newClientConn(nc, opts)
	defer func() {
		if err != nil {
			nc.Close()
		}
	}()
	input := make([]byte, 6)
	copy(input, "STMP")
	input[4] = c.Major
	input[5] = c.Minor
	rawHeaders := c.ClientHeader.Marshal()
	if len(rawHeaders) > 0 {
		input = append(input, '\n')
		input = append(input, rawHeaders...)
	}
	if len(c.ClientMessage) > 0 {
		input = append(input, "\n\n"...)
		input = append(input, c.ClientMessage...)
	}
	c.nc.SetWriteDeadline(time.Now().Add(c.handshakeTimeout))
	_, err = c.nc.Write(input)
	if err != nil {
		err = NewStatusError(StatusNetworkError, err)
		return
	}
	_, err = c.nc.Read(input[:5])
	if err != nil {
		err = NewStatusError(StatusNetworkError, err)
		return
	}
	size, err := c.readUvarint()
	if err != nil {
		err = NewStatusError(StatusNetworkError, err)
		return
	}
	input = make([]byte, size)
	_, err = c.nc.Read(input)
	if err != nil {
		err = NewStatusError(StatusNetworkError, err)
		return
	}
	sep := bytes.Index(input, []byte("\n\n"))
	if sep == -1 {
		sep = len(input)
	} else {
		c.ServerMessage = string(input[sep+2:])
	}
	c.ServerHeader = NewHeader()
	err = c.ServerHeader.Unmarshal(input[0:sep])
	if err != nil {
		return
	}
	r, w, err := c.initEncoding(opts.CompressLevel)
	if err != nil {
		return
	}
	go c.binaryReadChannel(r)
	go c.binaryWriteChannel(w)
	return
}

func WebSocketClient(wc *websocket.Conn, opts *DialOptions) (c *Conn, err error) {
	c = newClientConn(wc.UnderlyingConn(), opts)
	defer func() {
		if err != nil {
			wc.Close()
		}
	}()
	kind, data, err := wc.ReadMessage()
	// read header
	if err != nil {
		err = NewStatusError(StatusNetworkError, err)
		return
	}
	if len(data) < 6 || bytes.Equal(data[0:4], []byte("STMP")) {
		// must container headers, with STMP<STATUS>\n
		err = StatusProtocolError
		return
	}
	var sep int
	if kind == websocket.TextMessage {
		sep = bytes.IndexByte(data, '\n')
		if sep == -1 {
			// no new line for header
			err = StatusProtocolError
			return
		}
		status, err := strconv.ParseUint(string(data[4:sep]), 16, 8)
		if err != nil {
			// invalid status
			err = NewStatusError(StatusProtocolError, err)
			return
		}
		if status != 0 {
			// bad status
			err = Status(status)
			return
		}
		data = data[sep+1:]
	} else {
		if data[4] != 0 {
			// bad status
			err = Status(data[4])
			return
		}
		data = data[5:]
	}
	sep = bytes.Index(data, []byte("\n\n"))
	if sep == -1 {
		sep = len(data)
	} else {
		c.ServerMessage = string(data[sep+2])
	}
	c.ServerHeader = NewHeader()
	err = c.ServerHeader.Unmarshal(data[0:sep])
	if err != nil {
		return
	}
	// ws do not process encoding
	c.media = GetMediaCodec(c.ServerHeader.Get(DetermineContentType))
	if c.ServerHeader.Get(DeterminePacketFormat) == "text" {
		go c.websocketTextReadChannel(wc)
		go c.websocketTextWriteChannel(wc)
	} else {
		go c.websocketBinaryReadChannel(wc)
		go c.websocketBinaryWriteChannel(wc)
	}
	return
}

func DialTCP(addr string, opts *DialOptions) (*Conn, error) {
	dialOptionsDefaulter(addr, opts)
	nc, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return Client(nc, opts)
}

func DialTCPWithTLS(addr, certFile string, opts *DialOptions) (*Conn, error) {
	dialOptionsDefaulter(addr, opts)
	tc, err := loadTLSClientConfig(certFile, opts)
	if err != nil {
		return nil, err
	}
	nc, err := tls.Dial("tcp", addr, tc)
	if err != nil {
		return nil, err
	}
	return Client(nc, opts)
}

func DialKCP(addr string, opts *DialOptions) (*Conn, error) {
	dialOptionsDefaulter(addr, opts)
	nc, err := kcp.Dial(addr)
	if err != nil {
		return nil, err
	}
	return Client(nc, opts)
}

func DialKCPWithTLS(addr, certFile string, opts *DialOptions) (*Conn, error) {
	dialOptionsDefaulter(addr, opts)
	tc, err := loadTLSClientConfig(certFile, opts)
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
	return Client(tlsConn, opts)
}

// the gorilla/websocket only accepts public verified tls config for tls connection
func DialWebSocket(urlStr string, opts *DialOptions) (*Conn, error) {
	dialOptionsDefaulter(urlStr, opts)
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
		return nil, err
	}
	return WebSocketClient(wc, opts)
}
