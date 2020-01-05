// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-28 15:15:32
package stmp

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
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
	RegisterMediaCodec(NewMsgpackCodec(), NewJsonCodec())
}

type DialOptions struct {
	*ConnOptions
	// the headers for handshake
	Header Header
	// works only if encoding is not empty
	// range is [1, 9], default is 6
	CompressLevel int
	// the server name, if addr is ip & with tls, this is required
	ServerName string
	// if with tls and server name is empty, this should be true
	InsecureSkipVerify bool
}

func (o *DialOptions) WithHeader(key string, value ...string) *DialOptions {
	o.Header.Set(key, value...)
	return o
}

func (o *DialOptions) WithPacketFormat(format string) *DialOptions {
	o.Header.Set(AcceptPacketFormat, format)
	return o
}

func (o *DialOptions) WithEncoding(name string, level int) *DialOptions {
	o.Header.Set(AcceptEncoding, name)
	o.CompressLevel = level
	return o
}

func (o *DialOptions) WithContentType(typ string) *DialOptions {
	o.Header.Set(AcceptContentType, typ)
	return o
}

func (o *DialOptions) WithServerName(name string) *DialOptions {
	o.ServerName = name
	return o
}

func (o *DialOptions) WithInsecure() *DialOptions {
	o.InsecureSkipVerify = true
	return o
}

func (o *DialOptions) ApplyDefault(addr string) *DialOptions {
	no := o
	if no == nil {
		no = NewDialOptions()
	}
	no.ConnOptions = no.ConnOptions.ApplyDefault()
	if no.Header == nil {
		no.Header = NewHeader()
	}
	if no.Header.Has(AcceptEncoding) {
		encoding := no.Header.Get(AcceptEncoding)
		if GetEncodingCodec(encoding) == nil {
			panic("invalid encoding: " + encoding + ", please register it first")
		}
		if no.CompressLevel == 0 {
			no.CompressLevel = 6
		}
	}
	format := no.Header.Get(AcceptPacketFormat)
	if format != "" && format != "text" && format != "binary" {
		panic("invalid packet format: " + format + ", accepted values is text, binary")
	}
	if !no.Header.Has(AcceptContentType) {
		if format == "text" {
			no.Header.Set(AcceptContentType, "application/json")
		} else {
			no.Header.Set(AcceptContentType, "application/protobuf")
		}
	}
	contentType := no.Header.Get(AcceptContentType)
	if GetMediaCodec(contentType) == nil {
		panic("invalid content-type: " + contentType + ", please register it first")
	}
	if no.ServerName == "" && strings.Contains(addr, "://") {
		URL, err := url.Parse(addr)
		if err != nil {
			panic(err)
		}
		host := URL.Hostname()
		ip := net.ParseIP(host)
		if ip == nil {
			no.ServerName = host
		}
	}
	return no
}

func NewDialOptions() *DialOptions {
	return &DialOptions{
		ConnOptions: NewConnOptions(),
		Header:      NewHeader(),
	}
}

type ClientConn struct {
	*Router
	*Conn
	opts *DialOptions
}

func LoadTLSClientConfig(certFile string, opts *DialOptions) (*tls.Config, error) {
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

// Private: opts should be inited with NewDialOptions().ApplyDefault()
func NewClientConn(nc net.Conn, opts *DialOptions) *ClientConn {
	cc := &ClientConn{
		Router: opts.Router,
		Conn:   NewConn(nc, opts.ConnOptions),
		opts:   opts,
	}
	cc.ClientHeader = opts.Header
	return cc
}

func (c *ClientConn) MarshalHandshake() []byte {
	rawHeaders := c.ClientHeader.Marshal()
	input := make([]byte, 6+UvarintSize(uint64(len(rawHeaders))+1)+len(rawHeaders))
	copy(input, "STMP")
	input[4] = c.Major
	input[5] = c.Minor
	n := binary.PutUvarint(input[6:], uint64(len(rawHeaders))+1)
	input[n+6] = '\n'
	copy(input[n+7:], rawHeaders)
	return input
}

func (c *ClientConn) Handshake() (err error) {
	defer func() {
		if err != nil {
			c.Conn.Conn.Close()
		}
	}()
	c.SetWriteDeadline(time.Now().Add(c.opts.HandshakeTimeout))
	h := NewClientHandshake(c.Major, c.Minor, c.ClientHeader, "")
	err = h.Write(c)
	if err != nil {
		return
	}
	var r io.ReadCloser
	var w EncodingWriter
	r, w, err = c.initEncoding(c.opts.CompressLevel)
	if err != nil {
		err = NewStatusError(StatusUnsupportedContentType, err)
		return
	}
	go c.binaryWriteChannel(w)
	go func() {
		err := c.binaryReadChannel(r)
		c.opts.Logger.Error("read error, reconnecting...", zap.Error(err))
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
	wc.SetReadLimit(int64(c.opts.MaxPacketSize))
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

func DialTCP(addr string, opts *DialOptions) (*ClientConn, error) {
	opts = opts.ApplyDefault(addr)
	nc, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	cc := NewClientConn(nc, opts)
	return cc, cc.Handshake()
}

func DialTCPWithTLS(addr, certFile string, opts *DialOptions) (*ClientConn, error) {
	opts = opts.ApplyDefault(addr)
	tc, err := LoadTLSClientConfig(certFile, opts)
	if err != nil {
		return nil, err
	}
	nc, err := tls.Dial("tcp", addr, tc)
	if err != nil {
		return nil, err
	}
	cc := NewClientConn(nc, opts)
	return cc, cc.Handshake()
}

func DialKCP(addr string, opts *DialOptions) (*ClientConn, error) {
	opts = opts.ApplyDefault(addr)
	nc, err := kcp.Dial(addr)
	if err != nil {
		return nil, err
	}
	cc := NewClientConn(nc, opts)
	return cc, cc.Handshake()
}

func DialKCPWithTLS(addr, certFile string, opts *DialOptions) (*ClientConn, error) {
	opts = opts.ApplyDefault(addr)
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
	cc := NewClientConn(nc, opts)
	return cc, cc.Handshake()
}

// the gorilla/websocket only accepts public verified tls config for tls connection
func DialWebSocket(urlStr string, opts *DialOptions) (*ClientConn, error) {
	opts = opts.ApplyDefault(urlStr)
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
	cc := NewClientConn(wc.UnderlyingConn(), opts)
	return cc, cc.WebSocketHandshake(wc)
}
