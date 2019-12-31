// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-31 18:55:07
package stmp

import "context"

type ctxType int

const ctxConn ctxType = 0

func NewContext(conn *Conn) context.Context {
	ctx := context.Background()
	return context.WithValue(ctx, ctxConn, conn)
}

func SelectConn(ctx context.Context) *Conn {
	return ctx.Value(ctxConn).(*Conn)
}
