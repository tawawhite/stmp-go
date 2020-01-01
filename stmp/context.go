// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-31 18:55:07
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
