// Code generated by protoc-gen-stmp. DO NOT EDIT.
// source: examples/room/room_proto/room.proto
package room_proto

import (
	"context"
	"github.com/acrazing/stmp-go/stmp"
	empty "github.com/golang/protobuf/ptypes/empty"
)


func init() {
	stmp.RegisterMethodAction("stmp.examples.room.UserService.ListUser", 0x1001, func() interface{} { return &ListUserInput{} }, func() interface{} { return &ListUserOutput{} })
}

type STMPUserServiceServer interface {
	ListUser(ctx context.Context, in *ListUserInput) (out *ListUserOutput, err error)
}

func STMPRegisterUserServiceServer(r stmp.Router, s STMPUserServiceServer) {
	r.Register(s, "stmp.examples.room.UserService.ListUser", func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error) { return inst.(STMPUserServiceServer).ListUser(ctx, in.(*ListUserInput)) })
}

func STMPUnregisterUserServiceServer(r stmp.Router, s STMPUserServiceServer) {
	r.Unregister(s, "stmp.examples.room.UserService.ListUser")
}

type STMPUserServiceBroadcaster interface {
	ListUserToOne(ctx context.Context, in *ListUserInput, conn *stmp.Conn, opts ...stmp.CallOption) (*ListUserOutput, error)
	ListUserToList(ctx context.Context, in *ListUserInput, conns ...*stmp.Conn) error
	ListUserToSet(ctx context.Context, in *ListUserInput, conns stmp.ConnSet) error
	ListUserToAll(ctx context.Context, in *ListUserInput, srv *stmp.Server, filter stmp.ConnFilter) error
}

type stmpUserServiceBroadcaster struct{}

func (s stmpUserServiceBroadcaster) ListUserToOne(ctx context.Context, in *ListUserInput, conn *stmp.Conn, opts ...stmp.CallOption) (*ListUserOutput, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.UserService.ListUser", in, stmp.NewCallOptions(opts...))
	return out.(*ListUserOutput), err
}

func (s stmpUserServiceBroadcaster) ListUserToList(ctx context.Context, in *ListUserInput, conns ...*stmp.Conn) error {
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

func (s stmpUserServiceBroadcaster) ListUserToSet(ctx context.Context, in *ListUserInput, conns stmp.ConnSet) error {
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

func (s stmpUserServiceBroadcaster) ListUserToAll(ctx context.Context, in *ListUserInput, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.UserService.ListUser", in, filter)
}

func STMPNewUserServiceBroadcaster() STMPUserServiceBroadcaster {
	return stmpUserServiceBroadcaster{}
}

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

func init() {
	stmp.RegisterMethodAction("stmp.examples.room.UserEvents.StatusUpdated", 0x1101, func() interface{} { return &UserModel{} }, func() interface{} { return &empty.Empty{} })
}

type STMPUserEventsServer interface {
	StatusUpdated(ctx context.Context, in *UserModel) (out *empty.Empty, err error)
}

func STMPRegisterUserEventsServer(r stmp.Router, s STMPUserEventsServer) {
	r.Register(s, "stmp.examples.room.UserEvents.StatusUpdated", func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error) { return inst.(STMPUserEventsServer).StatusUpdated(ctx, in.(*UserModel)) })
}

func STMPUnregisterUserEventsServer(r stmp.Router, s STMPUserEventsServer) {
	r.Unregister(s, "stmp.examples.room.UserEvents.StatusUpdated")
}

type STMPUserEventsBroadcaster interface {
	StatusUpdatedToOne(ctx context.Context, in *UserModel, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error)
	StatusUpdatedToList(ctx context.Context, in *UserModel, conns ...*stmp.Conn) error
	StatusUpdatedToSet(ctx context.Context, in *UserModel, conns stmp.ConnSet) error
	StatusUpdatedToAll(ctx context.Context, in *UserModel, srv *stmp.Server, filter stmp.ConnFilter) error
}

type stmpUserEventsBroadcaster struct{}

func (s stmpUserEventsBroadcaster) StatusUpdatedToOne(ctx context.Context, in *UserModel, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.UserEvents.StatusUpdated", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s stmpUserEventsBroadcaster) StatusUpdatedToList(ctx context.Context, in *UserModel, conns ...*stmp.Conn) error {
	payloads := stmp.NewPayloadMap(in)
	for _, conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.UserEvents.StatusUpdated", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpUserEventsBroadcaster) StatusUpdatedToSet(ctx context.Context, in *UserModel, conns stmp.ConnSet) error {
	payloads := stmp.NewPayloadMap(in)
	for conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.UserEvents.StatusUpdated", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpUserEventsBroadcaster) StatusUpdatedToAll(ctx context.Context, in *UserModel, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.UserEvents.StatusUpdated", in, filter)
}

func STMPNewUserEventsBroadcaster() STMPUserEventsBroadcaster {
	return stmpUserEventsBroadcaster{}
}

type STMPUserEventsClient interface {
	StatusUpdated(ctx context.Context, in *UserModel, opts ...stmp.CallOption) (*empty.Empty, error)
}

type stmpUserEventsClient struct {
	c *stmp.Conn
}

func (s *stmpUserEventsClient) StatusUpdated(ctx context.Context, in *UserModel, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := s.c.Invoke(ctx, "stmp.examples.room.UserEvents.StatusUpdated", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func STMPNewUserEventsClient(c *stmp.Conn) STMPUserEventsClient {
	return &stmpUserEventsClient{c: c}
}

func init() {
	stmp.RegisterMethodAction("stmp.examples.room.RoomService.CreateRoom", 0x1201, func() interface{} { return &CreateRoomInput{} }, func() interface{} { return &RoomModel{} })
	stmp.RegisterMethodAction("stmp.examples.room.RoomService.ListRoom", 0x1202, func() interface{} { return &ListRoomInput{} }, func() interface{} { return &ListRoomOutput{} })
	stmp.RegisterMethodAction("stmp.examples.room.RoomService.JoinRoom", 0x1203, func() interface{} { return &JoinRoomInput{} }, func() interface{} { return &RoomModel{} })
	stmp.RegisterMethodAction("stmp.examples.room.RoomService.ExitRoom", 0x1204, func() interface{} { return &ExitRoomInput{} }, func() interface{} { return &empty.Empty{} })
	stmp.RegisterMethodAction("stmp.examples.room.RoomService.SendMessage", 0x1205, func() interface{} { return &SendMessageInput{} }, func() interface{} { return &empty.Empty{} })
}

type STMPRoomServiceServer interface {
	CreateRoom(ctx context.Context, in *CreateRoomInput) (out *RoomModel, err error)
	ListRoom(ctx context.Context, in *ListRoomInput) (out *ListRoomOutput, err error)
	JoinRoom(ctx context.Context, in *JoinRoomInput) (out *RoomModel, err error)
	ExitRoom(ctx context.Context, in *ExitRoomInput) (out *empty.Empty, err error)
	SendMessage(ctx context.Context, in *SendMessageInput) (out *empty.Empty, err error)
}

func STMPRegisterRoomServiceServer(r stmp.Router, s STMPRoomServiceServer) {
	r.Register(s, "stmp.examples.room.RoomService.CreateRoom", func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error) { return inst.(STMPRoomServiceServer).CreateRoom(ctx, in.(*CreateRoomInput)) })
	r.Register(s, "stmp.examples.room.RoomService.ListRoom", func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error) { return inst.(STMPRoomServiceServer).ListRoom(ctx, in.(*ListRoomInput)) })
	r.Register(s, "stmp.examples.room.RoomService.JoinRoom", func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error) { return inst.(STMPRoomServiceServer).JoinRoom(ctx, in.(*JoinRoomInput)) })
	r.Register(s, "stmp.examples.room.RoomService.ExitRoom", func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error) { return inst.(STMPRoomServiceServer).ExitRoom(ctx, in.(*ExitRoomInput)) })
	r.Register(s, "stmp.examples.room.RoomService.SendMessage", func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error) { return inst.(STMPRoomServiceServer).SendMessage(ctx, in.(*SendMessageInput)) })
}

func STMPUnregisterRoomServiceServer(r stmp.Router, s STMPRoomServiceServer) {
	r.Unregister(s, "stmp.examples.room.RoomService.CreateRoom")
	r.Unregister(s, "stmp.examples.room.RoomService.ListRoom")
	r.Unregister(s, "stmp.examples.room.RoomService.JoinRoom")
	r.Unregister(s, "stmp.examples.room.RoomService.ExitRoom")
	r.Unregister(s, "stmp.examples.room.RoomService.SendMessage")
}

type STMPRoomServiceBroadcaster interface {
	CreateRoomToOne(ctx context.Context, in *CreateRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*RoomModel, error)
	CreateRoomToList(ctx context.Context, in *CreateRoomInput, conns ...*stmp.Conn) error
	CreateRoomToSet(ctx context.Context, in *CreateRoomInput, conns stmp.ConnSet) error
	CreateRoomToAll(ctx context.Context, in *CreateRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error
	ListRoomToOne(ctx context.Context, in *ListRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*ListRoomOutput, error)
	ListRoomToList(ctx context.Context, in *ListRoomInput, conns ...*stmp.Conn) error
	ListRoomToSet(ctx context.Context, in *ListRoomInput, conns stmp.ConnSet) error
	ListRoomToAll(ctx context.Context, in *ListRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error
	JoinRoomToOne(ctx context.Context, in *JoinRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*RoomModel, error)
	JoinRoomToList(ctx context.Context, in *JoinRoomInput, conns ...*stmp.Conn) error
	JoinRoomToSet(ctx context.Context, in *JoinRoomInput, conns stmp.ConnSet) error
	JoinRoomToAll(ctx context.Context, in *JoinRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error
	ExitRoomToOne(ctx context.Context, in *ExitRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error)
	ExitRoomToList(ctx context.Context, in *ExitRoomInput, conns ...*stmp.Conn) error
	ExitRoomToSet(ctx context.Context, in *ExitRoomInput, conns stmp.ConnSet) error
	ExitRoomToAll(ctx context.Context, in *ExitRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error
	SendMessageToOne(ctx context.Context, in *SendMessageInput, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error)
	SendMessageToList(ctx context.Context, in *SendMessageInput, conns ...*stmp.Conn) error
	SendMessageToSet(ctx context.Context, in *SendMessageInput, conns stmp.ConnSet) error
	SendMessageToAll(ctx context.Context, in *SendMessageInput, srv *stmp.Server, filter stmp.ConnFilter) error
}

type stmpRoomServiceBroadcaster struct{}

func (s stmpRoomServiceBroadcaster) CreateRoomToOne(ctx context.Context, in *CreateRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*RoomModel, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomService.CreateRoom", in, stmp.NewCallOptions(opts...))
	return out.(*RoomModel), err
}

func (s stmpRoomServiceBroadcaster) CreateRoomToList(ctx context.Context, in *CreateRoomInput, conns ...*stmp.Conn) error {
	payloads := stmp.NewPayloadMap(in)
	for _, conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomService.CreateRoom", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomServiceBroadcaster) CreateRoomToSet(ctx context.Context, in *CreateRoomInput, conns stmp.ConnSet) error {
	payloads := stmp.NewPayloadMap(in)
	for conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomService.CreateRoom", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomServiceBroadcaster) CreateRoomToAll(ctx context.Context, in *CreateRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomService.CreateRoom", in, filter)
}

func (s stmpRoomServiceBroadcaster) ListRoomToOne(ctx context.Context, in *ListRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*ListRoomOutput, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomService.ListRoom", in, stmp.NewCallOptions(opts...))
	return out.(*ListRoomOutput), err
}

func (s stmpRoomServiceBroadcaster) ListRoomToList(ctx context.Context, in *ListRoomInput, conns ...*stmp.Conn) error {
	payloads := stmp.NewPayloadMap(in)
	for _, conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomService.ListRoom", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomServiceBroadcaster) ListRoomToSet(ctx context.Context, in *ListRoomInput, conns stmp.ConnSet) error {
	payloads := stmp.NewPayloadMap(in)
	for conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomService.ListRoom", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomServiceBroadcaster) ListRoomToAll(ctx context.Context, in *ListRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomService.ListRoom", in, filter)
}

func (s stmpRoomServiceBroadcaster) JoinRoomToOne(ctx context.Context, in *JoinRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*RoomModel, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomService.JoinRoom", in, stmp.NewCallOptions(opts...))
	return out.(*RoomModel), err
}

func (s stmpRoomServiceBroadcaster) JoinRoomToList(ctx context.Context, in *JoinRoomInput, conns ...*stmp.Conn) error {
	payloads := stmp.NewPayloadMap(in)
	for _, conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomService.JoinRoom", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomServiceBroadcaster) JoinRoomToSet(ctx context.Context, in *JoinRoomInput, conns stmp.ConnSet) error {
	payloads := stmp.NewPayloadMap(in)
	for conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomService.JoinRoom", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomServiceBroadcaster) JoinRoomToAll(ctx context.Context, in *JoinRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomService.JoinRoom", in, filter)
}

func (s stmpRoomServiceBroadcaster) ExitRoomToOne(ctx context.Context, in *ExitRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomService.ExitRoom", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s stmpRoomServiceBroadcaster) ExitRoomToList(ctx context.Context, in *ExitRoomInput, conns ...*stmp.Conn) error {
	payloads := stmp.NewPayloadMap(in)
	for _, conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomService.ExitRoom", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomServiceBroadcaster) ExitRoomToSet(ctx context.Context, in *ExitRoomInput, conns stmp.ConnSet) error {
	payloads := stmp.NewPayloadMap(in)
	for conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomService.ExitRoom", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomServiceBroadcaster) ExitRoomToAll(ctx context.Context, in *ExitRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomService.ExitRoom", in, filter)
}

func (s stmpRoomServiceBroadcaster) SendMessageToOne(ctx context.Context, in *SendMessageInput, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomService.SendMessage", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s stmpRoomServiceBroadcaster) SendMessageToList(ctx context.Context, in *SendMessageInput, conns ...*stmp.Conn) error {
	payloads := stmp.NewPayloadMap(in)
	for _, conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomService.SendMessage", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomServiceBroadcaster) SendMessageToSet(ctx context.Context, in *SendMessageInput, conns stmp.ConnSet) error {
	payloads := stmp.NewPayloadMap(in)
	for conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomService.SendMessage", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomServiceBroadcaster) SendMessageToAll(ctx context.Context, in *SendMessageInput, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomService.SendMessage", in, filter)
}

func STMPNewRoomServiceBroadcaster() STMPRoomServiceBroadcaster {
	return stmpRoomServiceBroadcaster{}
}

type STMPRoomServiceClient interface {
	CreateRoom(ctx context.Context, in *CreateRoomInput, opts ...stmp.CallOption) (*RoomModel, error)
	ListRoom(ctx context.Context, in *ListRoomInput, opts ...stmp.CallOption) (*ListRoomOutput, error)
	JoinRoom(ctx context.Context, in *JoinRoomInput, opts ...stmp.CallOption) (*RoomModel, error)
	ExitRoom(ctx context.Context, in *ExitRoomInput, opts ...stmp.CallOption) (*empty.Empty, error)
	SendMessage(ctx context.Context, in *SendMessageInput, opts ...stmp.CallOption) (*empty.Empty, error)
}

type stmpRoomServiceClient struct {
	c *stmp.Conn
}

func (s *stmpRoomServiceClient) CreateRoom(ctx context.Context, in *CreateRoomInput, opts ...stmp.CallOption) (*RoomModel, error) {
	out, err := s.c.Invoke(ctx, "stmp.examples.room.RoomService.CreateRoom", in, stmp.NewCallOptions(opts...))
	return out.(*RoomModel), err
}

func (s *stmpRoomServiceClient) ListRoom(ctx context.Context, in *ListRoomInput, opts ...stmp.CallOption) (*ListRoomOutput, error) {
	out, err := s.c.Invoke(ctx, "stmp.examples.room.RoomService.ListRoom", in, stmp.NewCallOptions(opts...))
	return out.(*ListRoomOutput), err
}

func (s *stmpRoomServiceClient) JoinRoom(ctx context.Context, in *JoinRoomInput, opts ...stmp.CallOption) (*RoomModel, error) {
	out, err := s.c.Invoke(ctx, "stmp.examples.room.RoomService.JoinRoom", in, stmp.NewCallOptions(opts...))
	return out.(*RoomModel), err
}

func (s *stmpRoomServiceClient) ExitRoom(ctx context.Context, in *ExitRoomInput, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := s.c.Invoke(ctx, "stmp.examples.room.RoomService.ExitRoom", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s *stmpRoomServiceClient) SendMessage(ctx context.Context, in *SendMessageInput, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := s.c.Invoke(ctx, "stmp.examples.room.RoomService.SendMessage", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func STMPNewRoomServiceClient(c *stmp.Conn) STMPRoomServiceClient {
	return &stmpRoomServiceClient{c: c}
}

func init() {
	stmp.RegisterMethodAction("stmp.examples.room.RoomEvents.UserEnter", 0x1301, func() interface{} { return &UserEnterEvent{} }, func() interface{} { return &empty.Empty{} })
	stmp.RegisterMethodAction("stmp.examples.room.RoomEvents.UserExit", 0x1302, func() interface{} { return &UserExitEvent{} }, func() interface{} { return &empty.Empty{} })
	stmp.RegisterMethodAction("stmp.examples.room.RoomEvents.NewMessage", 0x1303, func() interface{} { return &ChatMessageModel{} }, func() interface{} { return &empty.Empty{} })
}

type STMPRoomEventsServer interface {
	UserEnter(ctx context.Context, in *UserEnterEvent) (out *empty.Empty, err error)
	UserExit(ctx context.Context, in *UserExitEvent) (out *empty.Empty, err error)
	NewMessage(ctx context.Context, in *ChatMessageModel) (out *empty.Empty, err error)
}

func STMPRegisterRoomEventsServer(r stmp.Router, s STMPRoomEventsServer) {
	r.Register(s, "stmp.examples.room.RoomEvents.UserEnter", func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error) { return inst.(STMPRoomEventsServer).UserEnter(ctx, in.(*UserEnterEvent)) })
	r.Register(s, "stmp.examples.room.RoomEvents.UserExit", func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error) { return inst.(STMPRoomEventsServer).UserExit(ctx, in.(*UserExitEvent)) })
	r.Register(s, "stmp.examples.room.RoomEvents.NewMessage", func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error) { return inst.(STMPRoomEventsServer).NewMessage(ctx, in.(*ChatMessageModel)) })
}

func STMPUnregisterRoomEventsServer(r stmp.Router, s STMPRoomEventsServer) {
	r.Unregister(s, "stmp.examples.room.RoomEvents.UserEnter")
	r.Unregister(s, "stmp.examples.room.RoomEvents.UserExit")
	r.Unregister(s, "stmp.examples.room.RoomEvents.NewMessage")
}

type STMPRoomEventsBroadcaster interface {
	UserEnterToOne(ctx context.Context, in *UserEnterEvent, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error)
	UserEnterToList(ctx context.Context, in *UserEnterEvent, conns ...*stmp.Conn) error
	UserEnterToSet(ctx context.Context, in *UserEnterEvent, conns stmp.ConnSet) error
	UserEnterToAll(ctx context.Context, in *UserEnterEvent, srv *stmp.Server, filter stmp.ConnFilter) error
	UserExitToOne(ctx context.Context, in *UserExitEvent, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error)
	UserExitToList(ctx context.Context, in *UserExitEvent, conns ...*stmp.Conn) error
	UserExitToSet(ctx context.Context, in *UserExitEvent, conns stmp.ConnSet) error
	UserExitToAll(ctx context.Context, in *UserExitEvent, srv *stmp.Server, filter stmp.ConnFilter) error
	NewMessageToOne(ctx context.Context, in *ChatMessageModel, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error)
	NewMessageToList(ctx context.Context, in *ChatMessageModel, conns ...*stmp.Conn) error
	NewMessageToSet(ctx context.Context, in *ChatMessageModel, conns stmp.ConnSet) error
	NewMessageToAll(ctx context.Context, in *ChatMessageModel, srv *stmp.Server, filter stmp.ConnFilter) error
}

type stmpRoomEventsBroadcaster struct{}

func (s stmpRoomEventsBroadcaster) UserEnterToOne(ctx context.Context, in *UserEnterEvent, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomEvents.UserEnter", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s stmpRoomEventsBroadcaster) UserEnterToList(ctx context.Context, in *UserEnterEvent, conns ...*stmp.Conn) error {
	payloads := stmp.NewPayloadMap(in)
	for _, conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomEvents.UserEnter", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomEventsBroadcaster) UserEnterToSet(ctx context.Context, in *UserEnterEvent, conns stmp.ConnSet) error {
	payloads := stmp.NewPayloadMap(in)
	for conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomEvents.UserEnter", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomEventsBroadcaster) UserEnterToAll(ctx context.Context, in *UserEnterEvent, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomEvents.UserEnter", in, filter)
}

func (s stmpRoomEventsBroadcaster) UserExitToOne(ctx context.Context, in *UserExitEvent, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomEvents.UserExit", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s stmpRoomEventsBroadcaster) UserExitToList(ctx context.Context, in *UserExitEvent, conns ...*stmp.Conn) error {
	payloads := stmp.NewPayloadMap(in)
	for _, conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomEvents.UserExit", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomEventsBroadcaster) UserExitToSet(ctx context.Context, in *UserExitEvent, conns stmp.ConnSet) error {
	payloads := stmp.NewPayloadMap(in)
	for conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomEvents.UserExit", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomEventsBroadcaster) UserExitToAll(ctx context.Context, in *UserExitEvent, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomEvents.UserExit", in, filter)
}

func (s stmpRoomEventsBroadcaster) NewMessageToOne(ctx context.Context, in *ChatMessageModel, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomEvents.NewMessage", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s stmpRoomEventsBroadcaster) NewMessageToList(ctx context.Context, in *ChatMessageModel, conns ...*stmp.Conn) error {
	payloads := stmp.NewPayloadMap(in)
	for _, conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomEvents.NewMessage", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomEventsBroadcaster) NewMessageToSet(ctx context.Context, in *ChatMessageModel, conns stmp.ConnSet) error {
	payloads := stmp.NewPayloadMap(in)
	for conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "stmp.examples.room.RoomEvents.NewMessage", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stmpRoomEventsBroadcaster) NewMessageToAll(ctx context.Context, in *ChatMessageModel, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomEvents.NewMessage", in, filter)
}

func STMPNewRoomEventsBroadcaster() STMPRoomEventsBroadcaster {
	return stmpRoomEventsBroadcaster{}
}

type STMPRoomEventsClient interface {
	UserEnter(ctx context.Context, in *UserEnterEvent, opts ...stmp.CallOption) (*empty.Empty, error)
	UserExit(ctx context.Context, in *UserExitEvent, opts ...stmp.CallOption) (*empty.Empty, error)
	NewMessage(ctx context.Context, in *ChatMessageModel, opts ...stmp.CallOption) (*empty.Empty, error)
}

type stmpRoomEventsClient struct {
	c *stmp.Conn
}

func (s *stmpRoomEventsClient) UserEnter(ctx context.Context, in *UserEnterEvent, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := s.c.Invoke(ctx, "stmp.examples.room.RoomEvents.UserEnter", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s *stmpRoomEventsClient) UserExit(ctx context.Context, in *UserExitEvent, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := s.c.Invoke(ctx, "stmp.examples.room.RoomEvents.UserExit", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s *stmpRoomEventsClient) NewMessage(ctx context.Context, in *ChatMessageModel, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := s.c.Invoke(ctx, "stmp.examples.room.RoomEvents.NewMessage", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func STMPNewRoomEventsClient(c *stmp.Conn) STMPRoomEventsClient {
	return &stmpRoomEventsClient{c: c}
}
