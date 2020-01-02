// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-01 22:10:09
//
// generated stmp file contains for section: init, server, builder, client
package room_proto

import (
	"context"
	"github.com/acrazing/stmp-go/stmp"
)

// init

func init() {
	stmp.RegisterMethodAction("stmp.examples.room.UserService.ListUser", 0x1001, func() interface{} { return &ListUserInput{} }, func() interface{} { return &ListUserOutput{} })
}

// server

type STMPUserServiceServer interface {
	ListUser(ctx context.Context, in *ListUserInput) (out *ListUserOutput, err error)
}

func STMPRegisterUserServiceServer(r stmp.Router, s STMPUserServiceServer) {
	r.Register(s, "stmp.examples.room.UserService.ListUser", func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error) { return inst.(STMPUserServiceServer).ListUser(ctx, in.(*ListUserInput)) })
}

func STMPUnregisterUserServiceServer(r stmp.Router, s STMPUserServiceServer) {
	r.Unregister(s, "stmp.examples.room.UserService.ListUser")
}

// builder

type STMPUserServiceBuilder interface {
	ListUser(in *ListUserInput) *stmp.SendOptions
}

type stmpUserServiceBuilder struct{}

func (s stmpUserServiceBuilder) ListUser(in *ListUserInput) *stmp.SendOptions {
	return stmp.NewSendOptions("stmp.examples.room.UserService.ListUser", in)
}

func STMPNewUserServiceBuilder() STMPUserServiceBuilder {
	return &stmpUserServiceBuilder{}
}

// client

type STMPUserServiceClient interface {
	ListUser(ctx context.Context, in *ListUserInput, opts ...stmp.CallOption) (out *ListUserOutput, err error)
}

type stmpUserServiceClient struct {
	b stmpUserServiceBuilder
	c *stmp.Conn
}

func (s *stmpUserServiceClient) ListUser(ctx context.Context, in *ListUserInput, opts ...stmp.CallOption) (*ListUserOutput, error) {
	out, err := s.c.Invoke(ctx, s.b.ListUser(in), opts...)
	return out.(*ListUserOutput), err
}

func STMPNewUserServiceClient(c *stmp.Conn) STMPUserServiceClient {
	return &stmpUserServiceClient{c: c}
}
