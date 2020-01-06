package stmp

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/xtaci/kcp-go"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"net"
	"net/url"
	"strings"
	"time"
)

type DialOptions struct {
	*ConnOptions
	header     Header
	serverName string
	insecure   bool
	reconnect  Backoff
}

// set logger for client, default is zap.NewProduction()
func (o *DialOptions) WithLogger(logger *zap.Logger) *DialOptions {
	o.ConnOptions.WithLogger(logger)
	return o
}

// set custom router for client, you do not need this
func (o *DialOptions) WithRouter(r *Router) *DialOptions {
	o.ConnOptions.WithRouter(r)
	return o
}

// the max pending write packet size, default is 8
func (o *DialOptions) WithWriteQueueLimit(max int) *DialOptions {
	o.ConnOptions.WithWriteQueueLimit(max)
	return o
}

// the max message packet size, include handshake, exchange, close message, default is 16mb
func (o *DialOptions) WithPacketSizeLimit(max uint64) *DialOptions {
	o.ConnOptions.WithPacketSizeLimit(max)
	return o
}

// set timeout for handshake and exchange
// the read means at least one packet should be received in this time
// client will auto send ping message with the interval, if client do not receive pong from
// server in this time, client will close with StatusNetworkError
func (o *DialOptions) WithTimeout(handshake, read, write time.Duration) *DialOptions {
	o.ConnOptions.WithTimeout(handshake, read, write)
	return o
}

// set a custom handshake header
func (o *DialOptions) WithHeader(key string, value ...string) *DialOptions {
	o.header.Set(key, value...)
	return o
}

// packet format, only for websocket, could be text or binary, default is binary
func (o *DialOptions) WithPacketFormat(format string) *DialOptions {
	o.header.Set(AcceptPacketFormat, format)
	return o
}

// enable compress, only available algorithm is gzip, you can implement your custom
// algorithm with EncodingCodec interface, and register it with RegisterEncodingCodec
// (both server and client side need to register it)
// only supports tcp/kcp connections, websocket cannot use this
func (o *DialOptions) WithEncoding(name string) *DialOptions {
	o.header.Set(AcceptEncoding, name)
	return o
}

func (o *DialOptions) WithCompress(level int) *DialOptions {
	o.ConnOptions.WithCompress(level)
	return o
}

// set the payload content type, available values are application/json, application/protobuf, application/msgpack
// you can implement custom MediaCodec and register it both server and client side to activate it.
func (o *DialOptions) WithContentType(typ string) *DialOptions {
	o.header.Set(AcceptContentType, typ)
	return o
}

// set the server name for tls connection, this is required for SNI
func (o *DialOptions) WithServerName(name string) *DialOptions {
	o.serverName = name
	return o
}

// do not check the certificate for tls connection
func (o *DialOptions) WithInsecure() *DialOptions {
	o.insecure = true
	return o
}

// enable or disable reconnect
// if backoff is nil, will not reconnect
// else will retry when network error delay with backoff
func (o *DialOptions) WithReconnect(backoff Backoff) *DialOptions {
	o.reconnect = backoff
	return o
}

// apply default configuration, you do not to call this when call DialXXX(addr, opts)
// only when you want to create ClientConn by call NewClientConn(net.Conn, opts)
// you should call this.
func (o *DialOptions) ApplyDefault(addr string) *DialOptions {
	no := o
	if no == nil {
		no = NewDialOptions()
	}
	no.ConnOptions = no.ConnOptions.ApplyDefault()
	if no.header == nil {
		no.header = NewHeader()
	}
	if no.header.Has(AcceptEncoding) {
		encoding := no.header.Get(AcceptEncoding)
		if GetEncodingCodec(encoding) == nil {
			panic("invalid encoding: " + encoding + ", please register it first")
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

// create dial options
func NewDialOptions() *DialOptions {
	return &DialOptions{
		ConnOptions: NewConnOptions(),
		header:      NewHeader(),
		reconnect:   NewBackoff(time.Second, 200*time.Millisecond, 20, 7),
	}
}

// client conn
type ClientConn struct {
	*Router
	*Conn
	opts *DialOptions
}

// load tls client certificate from file system
func loadTLSClientConfig(certFile string, opts *DialOptions) (*tls.Config, error) {
	if opts.serverName == "" && !opts.insecure {
		return nil, errors.New("serverName is required for tls connection")
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
		InsecureSkipVerify: opts.insecure,
	}, nil
}

// create a client conn, the opts must not be nil, and should called ApplyDefault()
func NewClientConn(nc net.Conn, opts *DialOptions) *ClientConn {
	cc := &ClientConn{
		Router: opts.router,
		Conn:   NewConn(nc, opts.ConnOptions),
		opts:   opts,
	}
	cc.ClientHeader = opts.header
	cc.ServerHeader = NewHeader()
	return cc
}

// send & accept handshake packet, and run the read/write channel,
// use net.Conn as the transport.
// if status is not StatusOk, or do not receive the correct handshake response,
// will emit an error, the error should be StatusError
func (c *ClientConn) Handshake() (err error) {
	defer func() {
		if err != nil {
			c.Close()
		}
	}()
	c.SetWriteDeadline(time.Now().Add(c.opts.handshakeTimeout))
	ch := NewClientHandshake(c.Major, c.Minor, c.ClientHeader, "")
	err = ch.Write(c)
	if err != nil {
		return
	}
	log.Println("client handshake sent")
	sh := NewServerHandshake(0, c.ServerHeader, "")
	if err = sh.Read(c, c.opts.maxPacketSize); err != nil {
		return NewStatusError(StatusNetworkError, err)
	}
	var ec EncodingCodec
	ec, err = c.initEncoding()
	if err != nil {
		return
	}
	go c.binaryReadChannel(ec)
	go func() {
		c.binaryWriteChannel(ec)
	}()
	return
}

// send & accept handshake packet, and run the read/write channel
// use websocket.Conn as the transport.
// if status is not StatusOk, or do not receive the correct handshake response,
// will emit an error, the error should be StatusError
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
		go c.websocketTextReadChannel(wc)
		go func() {
			c.websocketTextWriteChannel(wc)
			// TODO reconnection
		}()
	} else {
		go c.websocketBinaryReadChannel(wc)
		go func() {
			c.websocketBinaryWriteChannel(wc)
			// TODO reconnection
		}()
	}
	return
}

// create a tcp connection and auto handshake with addr and opts
func DialTCP(addr string, opts *DialOptions) (*ClientConn, error) {
	opts = opts.ApplyDefault(addr)
	nc, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	cc := NewClientConn(nc, opts)
	return cc, cc.Handshake()
}

// create a tls tcp connection and auto handshake with addr and opts
func DialTCPWithTLS(addr, certFile string, opts *DialOptions) (*ClientConn, error) {
	opts = opts.ApplyDefault(addr)
	tc, err := loadTLSClientConfig(certFile, opts)
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

// create a kcp connection and auto handshake with addr and opts
func DialKCP(addr string, opts *DialOptions) (*ClientConn, error) {
	opts = opts.ApplyDefault(addr)
	nc, err := kcp.Dial(addr)
	if err != nil {
		return nil, err
	}
	cc := NewClientConn(nc, opts)
	return cc, cc.Handshake()
}

// create a tls kcp connection and auto handshake with addr and opts
func DialKCPWithTLS(addr, certFile string, opts *DialOptions) (*ClientConn, error) {
	opts = opts.ApplyDefault(addr)
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
	cc := NewClientConn(nc, opts)
	return cc, cc.Handshake()
}

// create a websocket connection and auto handshake with urlstr and opts
func DialWebSocket(urlstr string, opts *DialOptions) (*ClientConn, error) {
	opts = opts.ApplyDefault(urlstr)
	if len(opts.header) > 0 {
		headStr := url.Values(opts.header).Encode()
		if strings.IndexByte(urlstr, '?') > 0 {
			urlstr += "&" + headStr
		} else {
			urlstr += "?" + headStr
		}
	}
	dialer := websocket.Dialer{HandshakeTimeout: opts.handshakeTimeout}
	wc, _, err := dialer.Dial(urlstr, nil)
	if err != nil {
		return nil, err
	}
	cc := NewClientConn(wc.UnderlyingConn(), opts)
	return cc, cc.WebSocketHandshake(wc)
}
