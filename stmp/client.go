package stmp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/xtaci/kcp-go"
	"go.uber.org/zap"
	"io/ioutil"
	"net"
	"net/url"
	"strings"
	"sync/atomic"
	"time"
)

type ClientOptions struct {
	*ConnOptions
	header     Header
	reconnect  Backoff
	tlsConfig  *tls.Config
	serverName string
	certFile   string
	skipVerify bool
	resolver   *net.Resolver
	major      byte
	minor      byte
}

// set logger for client, default is zap.NewProduction()
func (o *ClientOptions) WithLogger(logger *zap.Logger) *ClientOptions {
	o.ConnOptions.WithLogger(logger)
	return o
}

// the max pending write packet size, default is 8
func (o *ClientOptions) WithWriteQueueLimit(max int) *ClientOptions {
	o.ConnOptions.WithWriteQueueLimit(max)
	return o
}

// the max message packet size, include handshake, exchange, close message, default is 16mb
func (o *ClientOptions) WithPacketSizeLimit(max uint64) *ClientOptions {
	o.ConnOptions.WithPacketSizeLimit(max)
	return o
}

// set timeout for handshake and exchange
// the read means at least one packet should be received in this time
// client will auto send ping message with the interval, if client do not receive pong from
// server in this time, client will close with StatusNetworkError
func (o *ClientOptions) WithTimeout(handshake, read, write time.Duration) *ClientOptions {
	o.ConnOptions.WithTimeout(handshake, read, write)
	return o
}

// set a custom handshake header
func (o *ClientOptions) WithHeader(key string, value ...string) *ClientOptions {
	o.header.Set(key, value...)
	return o
}

// packet format, only for websocket, could be text or binary, default is binary
func (o *ClientOptions) WithPacketFormat(format string) *ClientOptions {
	o.header.Set(AcceptPacketFormat, format)
	return o
}

// enable compress, only available algorithm is gzip, you can implement your custom
// algorithm with EncodingCodec interface, and register it with RegisterEncodingCodec
// (both server and client side need to register it)
// only supports tcp/kcp connections, websocket cannot use this
func (o *ClientOptions) WithEncoding(name string) *ClientOptions {
	o.header.Set(AcceptEncoding, name)
	return o
}

func (o *ClientOptions) WithCompress(level int) *ClientOptions {
	o.ConnOptions.WithCompress(level)
	return o
}

// set the payload content type, available values are application/json, application/protobuf, application/msgpack
// you can implement custom MediaCodec and register it both server and client side to activate it.
func (o *ClientOptions) WithContentType(typ string) *ClientOptions {
	o.header.Set(AcceptContentType, typ)
	return o
}

// Se necessary elements for build a custom tls.Config
//
// If skipVerify is true, serverName is omitted,
// else if serverName is empty, the input addr of DialXxxWithTLS must be a host:port rather than ip:port
func (o *ClientOptions) WithTLS(certFile string, serverName string, skipVerify bool) *ClientOptions {
	o.certFile = certFile
	o.serverName = serverName
	o.skipVerify = skipVerify
	return o
}

// set custom tls config, if this is set, certFile and skipVerify will be omitted
func (o *ClientOptions) WithTLSConfig(config *tls.Config) *ClientOptions {
	o.tlsConfig = config
	return o
}

func (o *ClientOptions) WithProtocolVersion(major byte, minor byte) *ClientOptions {
	if major != 1 && minor != 0 {
		panic("unsupported stmp protocol version: " + hexFormatProtocolVersion(major, minor))
	}
	o.major = major
	o.minor = minor
	return o
}

// set custom resolver if addr is host:port
func (o *ClientOptions) WithResolver(r *net.Resolver) *ClientOptions {
	o.resolver = r
	return o
}

// enable or disable reconnect
// if backoff is nil, will not reconnect
// else will retry when network error delay with backoff
func (o *ClientOptions) WithReconnect(backoff Backoff) *ClientOptions {
	o.reconnect = backoff
	return o
}

// apply default configuration, you do not to call this when call DialXXX(addr, opts)
// only when you want to create ClientConn by call NewClientConn(net.Conn, opts)
// you should call this.
func (o *ClientOptions) ApplyDefault() *ClientOptions {
	no := o
	if no == nil {
		no = NewClientOptions()
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
	return no
}

// create dial options
func NewClientOptions() *ClientOptions {
	return &ClientOptions{
		ConnOptions: NewConnOptions(),
		header:      NewHeader(),
		reconnect:   NewBackoff(time.Second, 200*time.Millisecond, 20, 7),
		major:       1,
		minor:       0,
	}
}

type ClientStatus int32

const (
	StatusDisconnected ClientStatus = 1
	StatusConnecting   ClientStatus = 2
	StatusConnected    ClientStatus = 3
	StatusClosing      ClientStatus = 4
)

// client conn
type Client struct {
	*Router
	*Conn
	opts *ClientOptions

	state *ClientStatus

	// callbacks

	// the handshake success callback
	connectedHandlers []func(header Header, message string)

	// the handshake failed, closed callback
	disconnectedHandlers []func(reason StatusError, willRetry bool, retryCount int, retryWait time.Duration)
}

func NewClient(opts *ClientOptions) *Client {
	opts = opts.ApplyDefault()
	cc := &Client{opts: opts, Conn: &Conn{State: NewState()}}
	cc.opts.router = NewRouter(cc)
	cc.Router = cc.opts.router
	return cc
}

func (c *Client) HandleConnected(fn func(header Header, message string)) {
	c.connectedHandlers = append(c.connectedHandlers, fn)
}

func (c *Client) HandleDisconnected(fn func(reason StatusError, willRetry bool, retryCount int, retryWait time.Duration)) {
	c.disconnectedHandlers = append(c.disconnectedHandlers, fn)
}

func (c *Client) Status() ClientStatus {
	return ClientStatus(atomic.LoadInt32((*int32)(c.state)))
}

func (c *Client) Handshake(agent interface{}) (sh *Handshake, se StatusError) {
	c.SetWriteDeadline(time.Now().Add(c.opts.handshakeTimeout))
	ch := NewClientHandshake(c.opts.major, c.opts.minor, c.ClientHeader, "")
	if se = ch.Write(c); se != nil {
		return
	}
	sh = NewServerHandshake(0, c.ServerHeader, "")
	se = sh.Read(c, c.opts.maxPacketSize)
	return
}

// send and wait for handshake response
func (c *Client) handshakeWebsocket(agent interface{}) (sh *Handshake, se StatusError) {
	wc := agent.(*websocket.Conn)
	wc.SetReadLimit(int64(c.opts.maxPacketSize))
	var kind int
	var data []byte
	var err error
	kind, data, err = wc.ReadMessage()
	// Read header
	if err != nil {
		se = NewStatusError(StatusNetworkError, err)
		return
	}
	sh = NewServerHandshake(0, nil, "")
	if kind == websocket.TextMessage {
		se = sh.UnmarshalText(data)
	} else {
		se = sh.UnmarshalBinary(data)
	}
	return
}

// IPv4[:port]
// IPv6[:port]
// host[:port]
func (c *Client) resolveTLSConfig(addr string) (*tls.Config, error) {
	if c.opts.tlsConfig != nil {
		return c.opts.tlsConfig, nil
	}
	config := &tls.Config{ServerName: c.opts.serverName, InsecureSkipVerify: c.opts.skipVerify}
	if c.opts.certFile != "" {
		b, err := ioutil.ReadFile(c.opts.certFile)
		if err != nil {
			return nil, err
		}
		cp := x509.NewCertPool()
		if !cp.AppendCertsFromPEM(b) {
			return nil, errors.New("tls failed to append certificates")
		}
		config.RootCAs = cp
	}
	if config.ServerName != "" || config.InsecureSkipVerify {
		return config, nil
	}
	lbracket := strings.LastIndexByte(addr, ']')
	if lbracket > -1 {
		return nil, errors.New("cannot resolve server name from " + addr)
	}
	lcolon := strings.LastIndexByte(addr, ':')
	if lcolon == -1 {
		lcolon = len(addr)
	}
	ldot := strings.LastIndexByte(addr, '.')
	if ldot == -1 {
		return nil, errors.New("cannot resolve server name from " + addr)
	}
	for ; ldot < lcolon; ldot++ {
		if addr[ldot] < '0' || addr[ldot] > '9' {
			config.ServerName = addr[:lcolon]
			return config, nil
		}
	}
	return nil, errors.New("cannot resolve server name from " + addr)
}

func (c *Client) resetConn(nc net.Conn, sh *Handshake) {
	c.Conn = &Conn{
		Conn:         nc,
		State:        c.State,
		mid:          c.mid,
		opts:         c.opts.ConnOptions,
		Major:        c.opts.major,
		Minor:        c.opts.minor,
		writeChan:    make(chan writeEvent, c.opts.writeQueueSize),
		requests:     make(map[uint16]chan *Packet),
		ClientHeader: c.opts.header,
		ServerHeader: sh.Header,
	}
}

func (c *Client) read(agent interface{}) {
	c.Conn.read()
}

func (c *Client) write(agent interface{}) StatusError {
	return c.Conn.write()
}

func (c *Client) readWebsocket(agent interface{}) {
	if c.ServerHeader.Get(DeterminePacketFormat) == "text" {
		c.Conn.readTextWebsocket(agent.(*websocket.Conn))
	} else {
		c.Conn.readBinaryWebsocket(agent.(*websocket.Conn))
	}
}

func (c *Client) writeWebsocket(agent interface{}) StatusError {
	if c.ServerHeader.Get(DeterminePacketFormat) == "text" {
		return c.Conn.writeTextWebsocket(agent.(*websocket.Conn))
	} else {
		return c.Conn.writeBinaryWebsocket(agent.(*websocket.Conn))
	}
}

func (c *Client) dial(makeConn func() (
	nc net.Conn, agent interface{}, err error),
	handshake func(agent interface{}) (*Handshake, StatusError),
	read func(agent interface{}),
	write func(agent interface{}) StatusError,
) {
	for {
		var se StatusError
		var sh *Handshake
		nc, agent, err := makeConn()
		if err != nil {
			se = NewStatusError(StatusNetworkError, "dial error: "+err.Error())
			goto RETRY
		}
		if sh, se = handshake(agent); se != nil {
			goto RETRY
		}
		if sh.Status != StatusOk {
			se = NewStatusError(sh.Status, sh.Message)
			goto RETRY
		}
		c.resetConn(nc, sh)
		if se = c.initEncoding(); se != nil {
			goto RETRY
		}
		c.opts.reconnect.Reset()
		for _, fn := range c.connectedHandlers {
			fn(c.ServerHeader, sh.Message)
		}
		go read(agent)
		se = write(agent)
	RETRY:
		var retryWait time.Duration
		var retryCount int
		var willRetry bool
		if c.opts.reconnect != nil && se.Code() == StatusNetworkError {
			retryWait, retryCount, willRetry = c.opts.reconnect.Next()
		}
		for _, fn := range c.disconnectedHandlers {
			fn(se, willRetry, retryCount, retryWait)
		}
		if willRetry {
			time.Sleep(retryWait)
		} else {
			break
		}
	}
}

// dial method will create *Conn, and setup, and handshake, and serve
func (c *Client) DialTCP(addr string) {
	dialer := net.Dialer{Resolver: c.opts.resolver, Timeout: c.opts.handshakeTimeout}
	c.dial(func() (conn net.Conn, agent interface{}, err error) {
		conn, err = dialer.DialContext(context.Background(), "tcp", addr)
		return
	}, c.Handshake, c.read, c.write)
}

// dial method will create *Conn, and setup, and handshake, and serve
func (c *Client) DialTCPWithTLS(addr string) {
	config, err := c.resolveTLSConfig(addr)
	if err != nil {
		panic(err)
	}
	dialer := &net.Dialer{Resolver: c.opts.resolver, Timeout: c.opts.handshakeTimeout}
	c.dial(func() (conn net.Conn, agent interface{}, err error) {
		conn, err = tls.DialWithDialer(dialer, "tcp", addr, config)
		return
	}, c.Handshake, c.read, c.write)
}

// dial method will create *Conn, and setup, and handshake, and serve
func (c *Client) DialKCP(addr string) {
	c.dial(func() (conn net.Conn, agent interface{}, err error) {
		conn, err = kcp.Dial(addr)
		return
	}, c.Handshake, c.read, c.write)
}

func (c *Client) DialKCPWithTLS(addr string) {
	config, err := c.resolveTLSConfig(addr)
	if err != nil {
		panic(err)
	}
	c.dial(func() (conn net.Conn, agent interface{}, err error) {
		conn, err = kcp.Dial(addr)
		if err != nil {
			return
		}
		tlsConn := tls.Client(conn, config)
		if err = tlsConn.Handshake(); err != nil {
			return
		}
		return tlsConn, nil, nil
	}, c.Handshake, c.read, c.write)
}

func (c *Client) DialWebsocket(urlstr string) {
	var config *tls.Config
	var err error
	if strings.HasPrefix(urlstr, "wss://") {
		addr := urlstr[6:]
		slash := strings.IndexByte(addr, '/')
		if slash > -1 {
			addr = addr[:slash]
		}
		config, err = c.resolveTLSConfig(addr)
		if err != nil {
			panic(err)
		}
	} else if !strings.HasPrefix(urlstr, "ws://") {
		panic("invalid websocket address: " + urlstr)
	}
	c.opts.WithHeader(DetermineStmpVersion, hexFormatProtocolVersion(c.opts.major, c.opts.minor))
	headStr := url.Values(c.opts.header).Encode()
	if strings.IndexByte(urlstr, '?') > 0 {
		urlstr += "&" + headStr
	} else {
		urlstr += "?" + headStr
	}
	netDialer := net.Dialer{
		Timeout:  c.opts.handshakeTimeout,
		Resolver: c.opts.resolver,
	}
	dialer := websocket.Dialer{
		HandshakeTimeout: c.opts.handshakeTimeout,
		TLSClientConfig:  config,
		NetDial:          netDialer.Dial,
		NetDialContext:   netDialer.DialContext,
	}
	c.dial(func() (nc net.Conn, agent interface{}, err error) {
		var wc *websocket.Conn
		wc, _, err = dialer.Dial(urlstr, nil)
		if err != nil {
			return
		}
		return wc.UnderlyingConn(), wc, nil
	}, c.handshakeWebsocket, c.readWebsocket, c.writeWebsocket)
}
