package stmp

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/xtaci/kcp-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
	"net/url"
	"strings"
	"time"
)

type DialOptions struct {
	*ConnOptions
	header     Header
	reconnect  Backoff
	addr       string
	tlsConfig  *tls.Config
	certFile   string
	skipVerify bool
	insecure   bool
	resolver   *net.Resolver
	major      byte
	minor      byte
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

// use custom cert file to make tls connection
// if skipVerify is false, the dial addr must be a host:port rather than ip:port
func (o *DialOptions) WithCert(certFile string, skipVerify bool) *DialOptions {
	o.certFile = certFile
	o.skipVerify = skipVerify
	return o
}

// set custom tls config, if this is set, certFile and skipVerify will be omitted
func (o *DialOptions) WithTLSConfig(config *tls.Config) *DialOptions {
	o.tlsConfig = config
	return o
}

// do not use tls
func (o *DialOptions) WithInsecure() *DialOptions {
	o.insecure = true
	return o
}

func (o *DialOptions) WithProtocolVersion(major byte, minor byte) *DialOptions {
	if major != 1 && minor != 0 {
		panic("unsupported stmp protocol version: " + hexFormatProtocolVersion(major, minor))
	}
	o.major = major
	o.minor = minor
	return o
}

// set custom resolver if addr is host:port
func (o *DialOptions) WithResolver(r *net.Resolver) *DialOptions {
	o.resolver = r
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
func (o *DialOptions) ApplyDefault() *DialOptions {
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
	return no
}

// load tls client certificate from file system
func LoadTLSClientCert(certFile string) (*x509.CertPool, error) {
	b, err := ioutil.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		return nil, errors.New("tls failed to append certificates")
	}
	return cp, nil
}

// create dial options
func NewDialOptions() *DialOptions {
	return &DialOptions{
		ConnOptions: NewConnOptions(),
		header:      NewHeader(),
		reconnect:   NewBackoff(time.Second, 200*time.Millisecond, 20, 7),
		major:       1,
		minor:       0,
	}
}

// client conn
type Client struct {
	*Router
	*Conn
	opts *DialOptions
}

func NewClientConn(nc net.Conn, opts *DialOptions) *Client {
	return &Client{
		Router: opts.router,
		opts:   opts,
		Conn:   NewConn(nc, opts.ConnOptions),
	}
}

// handshake, works include: send handshake request, wait handshake response, check status, check
func (c *Client) Handshake() (err error) {
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
	sh := NewServerHandshake(0, c.ServerHeader, "")
	if err = sh.Read(c, c.opts.maxPacketSize); err != nil {
		return NewStatusError(StatusNetworkError, err)
	}
	if sh.Status != StatusOk {
		err = NewStatusError(sh.Status, sh.Message)
	} else {
		err = c.initEncoding()
	}
	return
}

// start read & write channel
func (c *Client) Serve(nc net.Conn) error {
	return nil
}

// send and wait for handshake response
func (c *Client) HandshakeWebsocket(wc *websocket.Conn) (err error) {
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
		err = NewStatusError(h.Status, h.Message)
	} else {
		err = c.initEncoding()
	}
	return
}

// start websocket read & write channel
func (c *Client) ServeWebsocket() {

}

func (c *Client) DialTCP() {
}

// create a tcp connection and auto handshake with addr and opts
func DialTCP(addr string, opts *DialOptions) (*Client, error) {
	grpc.Dial()
	opts = opts.ApplyDefault(addr)
	nc, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	cc := NewClientConn(nc, opts)
	return cc, cc.Handshake()
}

// create a tls tcp connection and auto handshake with addr and opts
func DialTCPWithTLS(addr, certFile string, opts *DialOptions) (*Client, error) {
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
func DialKCP(addr string, opts *DialOptions) (*Client, error) {
	opts = opts.ApplyDefault(addr)
	nc, err := kcp.Dial(addr)
	if err != nil {
		return nil, err
	}
	cc := NewClientConn(nc, opts)
	return cc, cc.Handshake()
}

// create a tls kcp connection and auto handshake with addr and opts
func DialKCPWithTLS(addr, certFile string, opts *DialOptions) (*Client, error) {
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
func DialWebsocket(urlstr string, opts *DialOptions) (*Client, error) {
	opts = opts.ApplyDefault()
	opts.WithHeader(DetermineStmpVersion, hexFormatProtocolVersion(opts.major, opts.minor))
	headStr := url.Values(opts.header).Encode()
	if strings.IndexByte(urlstr, '?') > 0 {
		urlstr += "&" + headStr
	} else {
		urlstr += "?" + headStr
	}
	dialer := websocket.Dialer{HandshakeTimeout: opts.handshakeTimeout}
	wc, _, err := dialer.Dial(urlstr, nil)
	if err != nil {
		return nil, err
	}
	cc := NewClientConn(wc.UnderlyingConn(), opts)
	return cc, cc.HandshakeWebsocket(wc)
}
