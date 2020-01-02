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

// broadcaster, use for server side send request to all client

type STMPUserServiceBroadcaster interface {
	ListUser(ctx context.Context, in *ListUserInput, conn *stmp.Conn, opts ...stmp.CallOption) (*ListUserOutput, error)
	ListUserForList(ctx context.Context, in *ListUserInput, conns ...*stmp.Conn) error
	ListUserForKeys(ctx context.Context, in *ListUserInput, conns map[*stmp.Conn]bool) error
	BroadcastListUser(ctx context.Context, in *ListUserInput, srv *stmp.Server, filter stmp.ConnFilter) error
	ListUserMethod() string
	ListUserAction() uint64
}

type stmpUserServiceBroadcaster struct{}

func (s stmpUserServiceBroadcaster) ListUser(ctx context.Context, in *ListUserInput, conn *stmp.Conn, opts ...stmp.CallOption) (*ListUserOutput, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.UserService.ListUser", in, stmp.NewCallOptions(opts...))
	return out.(*ListUserOutput), err
}

func (s stmpUserServiceBroadcaster) ListUserForList(ctx context.Context, in *ListUserInput, conns ...*stmp.Conn) error {
	payloads := stmp.NewPayloadMap(in)
	for _, conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.UserService.ListUser", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpUserServiceBroadcaster) ListUserForKeys(ctx context.Context, in *ListUserInput, conns map[*stmp.Conn]bool) error {
	payloads := stmp.NewPayloadMap(in)
	for conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.UserService.ListUser", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpUserServiceBroadcaster) BroadcastListUser(ctx context.Context, in *ListUserInput, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.UserService.ListUser", in, filter)
}

func (s stmpUserServiceBroadcaster) ListUserMethod() string {
	return "stmp.examples.room.UserService.ListUser"
}

func (s stmpUserServiceBroadcaster) ListUserAction() uint64 {
	return 0x1001
}

func STMPNewUserServiceBroadcaster() STMPUserServiceBroadcaster {
	return stmpUserServiceBroadcaster{}
}

// client

type STMPUserServiceClient interface {
	ListUser(ctx context.Context, in *ListUserInput, opts ...stmp.CallOption) (*ListUserOutput, error)
}

type stmpUserServiceClient struct {
	c *stmp.Conn
}

func (s *stmpUserServiceClient) ListUser(ctx context.Context, in *ListUserInput, opts ...stmp.CallOption) (*ListUserOutput, error) {
	out, err := s.c.Invoke(ctx, "stmp.examples.room.UserService.ListUser", in, stmp.NewCallOptions(opts...))
	return out.(*ListUserOutput), err
}

func STMPNewUserServiceClient(c *stmp.Conn) STMPUserServiceClient {
	return &stmpUserServiceClient{c: c}
}
