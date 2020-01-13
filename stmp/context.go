package stmp

import "context"

type ctxStmpValue struct {
	packet *Packet
	conn   *Conn
	method *Method
	header Header
	host   interface{}
	async  func() (out interface{}, err error)
}

type ctxStmpKey struct{}

func withStmp(ctx context.Context, p *Packet, c *Conn, m *Method, h interface{}) context.Context {
	return context.WithValue(ctx, ctxStmpKey{}, ctxStmpValue{
		packet: p,
		conn:   c,
		method: m,
		host:   h,
		header: NewHeader(),
	})
}

func selectStmp(ctx context.Context) ctxStmpValue {
	return ctx.Value(ctxStmpKey{}).(ctxStmpValue)
}

func SelectPacket(ctx context.Context) *Packet {
	return ctx.Value(ctxStmpKey{}).(ctxStmpValue).packet
}

func SelectConn(ctx context.Context) *Conn {
	return ctx.Value(ctxStmpKey{}).(ctxStmpValue).conn
}

func SelectMethod(ctx context.Context) *Method {
	return ctx.Value(ctxStmpKey{}).(ctxStmpValue).method
}

func SelectServer(ctx context.Context) *Server {
	return ctx.Value(ctxStmpKey{}).(ctxStmpValue).host.(*Server)
}

func SelectClient(ctx context.Context) *Client {
	return ctx.Value(ctxStmpKey{}).(ctxStmpValue).host.(*Client)
}

func SelectOutputHeader(ctx context.Context) Header {
	return ctx.Value(ctxStmpKey{}).(ctxStmpValue).header
}

func SelectInputHeader(ctx context.Context) Header {
	return ctx.Value(ctxStmpKey{}).(ctxStmpValue).packet.Header
}

// TODO: dispatch handlers in read channel, and go async when user call Async(ctx, fn) method
func Async(ctx context.Context, fn func() (out interface{}, err error)) {
	s := selectStmp(ctx)
	if s.async != nil {
		panic("current handler is async already for action: " + string(hexFormatUint64(s.packet.Action)))
	}
	s.async = fn
}
