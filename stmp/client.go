// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-28 15:15:32
package stmp

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/xtaci/kcp-go"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net"
	"net/url"
	"strings"
	"time"
)

func init() {
	RegisterMediaCodec(NewProtobufCodec(), NewJsonCodec())
}

type dialOptions struct {
	*connOptions
	// the headers for handshake
	header Header
	// works only if encoding is not empty
	// range is [1, 9], default is 6
	compressLevel int
	// the server name, if addr is ip & with tls, this is required
	serverName string
	// if with tls and server name is empty, this should be true
	insecureSkipVerify bool
}

func (o *dialOptions) WithLogger(logger *zap.Logger) *dialOptions {
	o.connOptions.WithLogger(logger)
	return o
}

func (o *dialOptions) WithRouter(r *Router) *dialOptions {
	o.connOptions.WithRouter(r)
	return o
}

func (o *dialOptions) WithWriteQueueLimit(max int) *dialOptions {
	o.connOptions.WithWriteQueueLimit(max)
	return o
}

func (o *dialOptions) WithPacketSizeLimit(max uint64) *dialOptions {
	o.connOptions.WithPacketSizeLimit(max)
	return o
}

func (o *dialOptions) WithTimeout(handshake, read, write time.Duration) *dialOptions {
	o.connOptions.WithTimeout(handshake, read, write)
	return o
}

func (o *dialOptions) WithHeader(key string, value ...string) *dialOptions {
	o.header.Set(key, value...)
	return o
}

func (o *dialOptions) WithPacketFormat(format string) *dialOptions {
	o.header.Set(AcceptPacketFormat, format)
	return o
}

func (o *dialOptions) WithEncoding(name string, level int) *dialOptions {
	o.header.Set(AcceptEncoding, name)
	o.compressLevel = level
	return o
}

func (o *dialOptions) WithContentType(typ string) *dialOptions {
	o.header.Set(AcceptContentType, typ)
	return o
}

func (o *dialOptions) WithServerName(name string) *dialOptions {
	o.serverName = name
	return o
}

func (o *dialOptions) WithInsecure() *dialOptions {
	o.insecureSkipVerify = true
	return o
}

func (o *dialOptions) applyDefault(addr string) *dialOptions {
	no := o
	if no == nil {
		no = NewDialOptions()
	}
	no.connOptions = no.connOptions.applyDefault()
	if no.header == nil {
		no.header = NewHeader()
	}
	if no.header.Has(AcceptEncoding) {
		encoding := no.header.Get(AcceptEncoding)
		if GetEncodingCodec(encoding) == nil {
			panic("invalid encoding: " + encoding + ", please register it first")
		}
		if no.compressLevel == 0 {
			no.compressLevel = 6
		}
	}
	format := no.header.Get(AcceptPacketFormat)
	if format != "" && format != "text" && format != "binary" {
		panic("invalid packet format: " + format + ", accepted values is text, binary")
	}
	if !no.header.Has(AcceptContentType) {
		if format == "text" {
			no.header.Set(AcceptContentType, "application/json")
		} else {
			no.header.Set(AcceptContentType, "application/protobuf")
		}
	}
	contentType := no.header.Get(AcceptContentType)
	if GetMediaCodec(contentType) == nil {
		panic("invalid content-type: " + contentType + ", please register it first")
	}
	if no.serverName == "" && strings.Contains(addr, "://") {
		URL, err := url.Parse(addr)
		if err != nil {
			panic(err)
		}
		host := URL.Hostname()
		ip := net.ParseIP(host)
		if ip == nil {
			no.serverName = host
		}
	}
	return no
}

func NewDialOptions() *dialOptions {
	return &dialOptions{
		connOptions: NewConnOptions(),
		header:      NewHeader(),
	}
}

type ClientConn struct {
	*Router
	*Conn
	opts *dialOptions
}

func LoadTLSClientConfig(certFile string, opts *dialOptions) (*tls.Config, error) {
	if opts.serverName == "" && !opts.insecureSkipVerify {
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
		ServerName:         opts.serverName,
		InsecureSkipVerify: opts.insecureSkipVerify,
	}, nil
}

// Private: opts should be inited with NewDialOptions().applyDefault()
func newClientConn(nc net.Conn, opts *dialOptions) *ClientConn {
	cc := &ClientConn{
		Router: opts.router,
		Conn:   NewConn(nc, opts.connOptions),
		opts:   opts,
	}
	cc.ClientHeader = opts.header
	return cc
}

func (c *ClientConn) Handshake() (err error) {
	defer func() {
		if err != nil {
			c.Conn.Conn.Close()
		}
	}()
	c.SetWriteDeadline(time.Now().Add(c.opts.handshakeTimeout))
	h := NewClientHandshake(c.Major, c.Minor, c.ClientHeader, "")
	err = h.Write(c)
	if err != nil {
		return
	}
	var r io.ReadCloser
	var w EncodingWriter
	r, w, err = c.initEncoding(c.opts.compressLevel)
	if err != nil {
		err = NewStatusError(StatusUnsupportedContentType, err)
		return
	}
	go c.binaryWriteChannel(w)
	go func() {
		err := c.binaryReadChannel(r)
		c.opts.logger.Error("read error, reconnecting...", zap.Error(err))
		// TODO reconnect
	}()
	return
}

func (c *ClientConn) WebSocketHandshake(wc *websocket.Conn) (err error) {
	defer func() {
		if err != nil {
			wc.Close()
		}
	}()
	wc.SetReadLimit(int64(c.opts.maxPacketSize))
	var kind int
	var data []byte
	kind, data, err = wc.ReadMessage()
	// Read header
	if err != nil {
		err = NewStatusError(StatusNetworkError, err)
		return
	}
	h := NewServerHandshake(0, nil, "")
	if kind == websocket.TextMessage {
		err = h.UnmarshalText(data)
	} else {
		err = h.UnmarshalBinary(data)
	}
	if err != nil {
		return
	}
	if h.Status != StatusOk {
		return NewStatusError(h.Status, h.Message)
	}
	// ws do not process encoding
	c.Media = GetMediaCodec(c.ServerHeader.Get(DetermineContentType))
	if c.ServerHeader.Get(DeterminePacketFormat) == "text" {
		c.websocketTextWriteChannel(wc)
		go func() {
			go c.websocketTextReadChannel(wc)
			// TODO reconnection
		}()
	} else {
		go c.websocketBinaryWriteChannel(wc)
		go func() {
			c.websocketBinaryReadChannel(wc)
			// TODO reconnection
		}()
	}
	return
}

func DialTCP(addr string, opts *dialOptions) (*ClientConn, error) {
	opts = opts.applyDefault(addr)
	nc, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	cc := newClientConn(nc, opts)
	return cc, cc.Handshake()
}

func DialTCPWithTLS(addr, certFile string, opts *dialOptions) (*ClientConn, error) {
	opts = opts.applyDefault(addr)
	tc, err := LoadTLSClientConfig(certFile, opts)
	if err != nil {
		return nil, err
	}
	nc, err := tls.Dial("tcp", addr, tc)
	if err != nil {
		return nil, err
	}
	cc := newClientConn(nc, opts)
	return cc, cc.Handshake()
}

func DialKCP(addr string, opts *dialOptions) (*ClientConn, error) {
	opts = opts.applyDefault(addr)
	nc, err := kcp.Dial(addr)
	if err != nil {
		return nil, err
	}
	cc := newClientConn(nc, opts)
	return cc, cc.Handshake()
}

func DialKCPWithTLS(addr, certFile string, opts *dialOptions) (*ClientConn, error) {
	opts = opts.applyDefault(addr)
	tc, err := LoadTLSClientConfig(certFile, opts)
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
	cc := newClientConn(nc, opts)
	return cc, cc.Handshake()
}

// the gorilla/websocket only accepts public verified tls config for tls connection
func DialWebSocket(urlStr string, opts *dialOptions) (*ClientConn, error) {
	opts = opts.applyDefault(urlStr)
	if len(opts.header) > 0 {
		headStr := url.Values(opts.header).Encode()
		if strings.IndexByte(urlStr, '?') > 0 {
			urlStr += "&" + headStr
		} else {
			urlStr += "?" + headStr
		}
	}
	dialer := websocket.Dialer{HandshakeTimeout: opts.handshakeTimeout}
	wc, _, err := dialer.Dial(urlStr, nil)
	if err != nil {
		return nil, err
	}
	cc := newClientConn(wc.UnderlyingConn(), opts)
	return cc, cc.WebSocketHandshake(wc)
}
