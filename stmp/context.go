package stmp

import "context"

type ctxConnKey struct{}

func WithConn(ctx context.Context, conn *Conn) context.Context {
	return context.WithValue(ctx, ctxConnKey{}, conn)
}

func SelectConn(ctx context.Context) *Conn {
	return ctx.Value(ctxConnKey{}).(*Conn)
}

type ctxServerKey struct{}

func WithServer(ctx context.Context, srv *Server) context.Context {
	return context.WithValue(ctx, ctxServerKey{}, srv)
}

func SelectServer(ctx context.Context) *Server {
	return ctx.Value(ctxServerKey{}).(*Server)
}

type ctxPacketKey struct{}

func WithPacket(ctx context.Context, packet *Packet) context.Context {
	return context.WithValue(ctx, ctxPacketKey{}, packet)
}

func SelectPacket(ctx context.Context) uint64 {
	return ctx.Value(ctxPacketKey{}).(uint64)
}
