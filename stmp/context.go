package stmp

import "context"

type ctxStmpValue struct {
	packet *Packet
	conn   *Conn
	method *Method
	router *Router
	header Header
}

type ctxStmpKey struct{}

func withStmp(ctx context.Context, p *Packet, c *Conn, m *Method, r *Router) context.Context {
	return context.WithValue(ctx, ctxStmpKey{}, ctxStmpValue{packet: p, conn: c, method: m, router: r, header: NewHeader()})
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

func SelectRouter(ctx context.Context) *Router {
	return ctx.Value(ctxStmpKey{}).(ctxStmpValue).router
}

func SelectServer(ctx context.Context) *Server {
	return ctx.Value(ctxStmpKey{}).(ctxStmpValue).router.host.(*Server)
}

func SelectClient(ctx context.Context) *Client {
	return ctx.Value(ctxStmpKey{}).(ctxStmpValue).router.host.(*Client)
}

func SelectOutputHeader(ctx context.Context) Header {
	return ctx.Value(ctxStmpKey{}).(ctxStmpValue).header
}

func SelectInputHeader(ctx context.Context) Header {
	return ctx.Value(ctxStmpKey{}).(ctxStmpValue).packet.Header
}
