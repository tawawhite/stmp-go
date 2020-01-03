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
	ListUser(ctx context.Context, in *ListUserInput, conn *stmp.Conn, opts ...stmp.CallOption) (*ListUserOutput, error)
	ListUserForList(ctx context.Context, in *ListUserInput, conns ...*stmp.Conn) error
	ListUserForKeys(ctx context.Context, in *ListUserInput, conns stmp.ConnSet) error
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

func (s stmpUserServiceBroadcaster) ListUserForKeys(ctx context.Context, in *ListUserInput, conns stmp.ConnSet) error {
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
	StatusUpdated(ctx context.Context, in *UserModel, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error)
	StatusUpdatedForList(ctx context.Context, in *UserModel, conns ...*stmp.Conn) error
	StatusUpdatedForKeys(ctx context.Context, in *UserModel, conns stmp.ConnSet) error
	BroadcastStatusUpdated(ctx context.Context, in *UserModel, srv *stmp.Server, filter stmp.ConnFilter) error
	StatusUpdatedMethod() string
	StatusUpdatedAction() uint64
}

type stmpUserEventsBroadcaster struct{}

func (s stmpUserEventsBroadcaster) StatusUpdated(ctx context.Context, in *UserModel, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.UserEvents.StatusUpdated", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s stmpUserEventsBroadcaster) StatusUpdatedForList(ctx context.Context, in *UserModel, conns ...*stmp.Conn) error {
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

func (s stmpUserEventsBroadcaster) StatusUpdatedForKeys(ctx context.Context, in *UserModel, conns stmp.ConnSet) error {
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

func (s stmpUserEventsBroadcaster) BroadcastStatusUpdated(ctx context.Context, in *UserModel, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.UserEvents.StatusUpdated", in, filter)
}

func (s stmpUserEventsBroadcaster) StatusUpdatedMethod() string {
	return "stmp.examples.room.UserEvents.StatusUpdated"
}

func (s stmpUserEventsBroadcaster) StatusUpdatedAction() uint64 {
	return 0x1101
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
	CreateRoom(ctx context.Context, in *CreateRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*RoomModel, error)
	CreateRoomForList(ctx context.Context, in *CreateRoomInput, conns ...*stmp.Conn) error
	CreateRoomForKeys(ctx context.Context, in *CreateRoomInput, conns stmp.ConnSet) error
	BroadcastCreateRoom(ctx context.Context, in *CreateRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error
	CreateRoomMethod() string
	CreateRoomAction() uint64
	ListRoom(ctx context.Context, in *ListRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*ListRoomOutput, error)
	ListRoomForList(ctx context.Context, in *ListRoomInput, conns ...*stmp.Conn) error
	ListRoomForKeys(ctx context.Context, in *ListRoomInput, conns stmp.ConnSet) error
	BroadcastListRoom(ctx context.Context, in *ListRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error
	ListRoomMethod() string
	ListRoomAction() uint64
	JoinRoom(ctx context.Context, in *JoinRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*RoomModel, error)
	JoinRoomForList(ctx context.Context, in *JoinRoomInput, conns ...*stmp.Conn) error
	JoinRoomForKeys(ctx context.Context, in *JoinRoomInput, conns stmp.ConnSet) error
	BroadcastJoinRoom(ctx context.Context, in *JoinRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error
	JoinRoomMethod() string
	JoinRoomAction() uint64
	ExitRoom(ctx context.Context, in *ExitRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error)
	ExitRoomForList(ctx context.Context, in *ExitRoomInput, conns ...*stmp.Conn) error
	ExitRoomForKeys(ctx context.Context, in *ExitRoomInput, conns stmp.ConnSet) error
	BroadcastExitRoom(ctx context.Context, in *ExitRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error
	ExitRoomMethod() string
	ExitRoomAction() uint64
	SendMessage(ctx context.Context, in *SendMessageInput, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error)
	SendMessageForList(ctx context.Context, in *SendMessageInput, conns ...*stmp.Conn) error
	SendMessageForKeys(ctx context.Context, in *SendMessageInput, conns stmp.ConnSet) error
	BroadcastSendMessage(ctx context.Context, in *SendMessageInput, srv *stmp.Server, filter stmp.ConnFilter) error
	SendMessageMethod() string
	SendMessageAction() uint64
}

type stmpRoomServiceBroadcaster struct{}

func (s stmpRoomServiceBroadcaster) CreateRoom(ctx context.Context, in *CreateRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*RoomModel, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomService.CreateRoom", in, stmp.NewCallOptions(opts...))
	return out.(*RoomModel), err
}

func (s stmpRoomServiceBroadcaster) CreateRoomForList(ctx context.Context, in *CreateRoomInput, conns ...*stmp.Conn) error {
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

func (s stmpRoomServiceBroadcaster) CreateRoomForKeys(ctx context.Context, in *CreateRoomInput, conns stmp.ConnSet) error {
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

func (s stmpRoomServiceBroadcaster) BroadcastCreateRoom(ctx context.Context, in *CreateRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomService.CreateRoom", in, filter)
}

func (s stmpRoomServiceBroadcaster) CreateRoomMethod() string {
	return "stmp.examples.room.RoomService.CreateRoom"
}

func (s stmpRoomServiceBroadcaster) CreateRoomAction() uint64 {
	return 0x1201
}

func (s stmpRoomServiceBroadcaster) ListRoom(ctx context.Context, in *ListRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*ListRoomOutput, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomService.ListRoom", in, stmp.NewCallOptions(opts...))
	return out.(*ListRoomOutput), err
}

func (s stmpRoomServiceBroadcaster) ListRoomForList(ctx context.Context, in *ListRoomInput, conns ...*stmp.Conn) error {
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

func (s stmpRoomServiceBroadcaster) ListRoomForKeys(ctx context.Context, in *ListRoomInput, conns stmp.ConnSet) error {
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

func (s stmpRoomServiceBroadcaster) BroadcastListRoom(ctx context.Context, in *ListRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomService.ListRoom", in, filter)
}

func (s stmpRoomServiceBroadcaster) ListRoomMethod() string {
	return "stmp.examples.room.RoomService.ListRoom"
}

func (s stmpRoomServiceBroadcaster) ListRoomAction() uint64 {
	return 0x1202
}

func (s stmpRoomServiceBroadcaster) JoinRoom(ctx context.Context, in *JoinRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*RoomModel, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomService.JoinRoom", in, stmp.NewCallOptions(opts...))
	return out.(*RoomModel), err
}

func (s stmpRoomServiceBroadcaster) JoinRoomForList(ctx context.Context, in *JoinRoomInput, conns ...*stmp.Conn) error {
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

func (s stmpRoomServiceBroadcaster) JoinRoomForKeys(ctx context.Context, in *JoinRoomInput, conns stmp.ConnSet) error {
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

func (s stmpRoomServiceBroadcaster) BroadcastJoinRoom(ctx context.Context, in *JoinRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomService.JoinRoom", in, filter)
}

func (s stmpRoomServiceBroadcaster) JoinRoomMethod() string {
	return "stmp.examples.room.RoomService.JoinRoom"
}

func (s stmpRoomServiceBroadcaster) JoinRoomAction() uint64 {
	return 0x1203
}

func (s stmpRoomServiceBroadcaster) ExitRoom(ctx context.Context, in *ExitRoomInput, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomService.ExitRoom", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s stmpRoomServiceBroadcaster) ExitRoomForList(ctx context.Context, in *ExitRoomInput, conns ...*stmp.Conn) error {
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

func (s stmpRoomServiceBroadcaster) ExitRoomForKeys(ctx context.Context, in *ExitRoomInput, conns stmp.ConnSet) error {
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

func (s stmpRoomServiceBroadcaster) BroadcastExitRoom(ctx context.Context, in *ExitRoomInput, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomService.ExitRoom", in, filter)
}

func (s stmpRoomServiceBroadcaster) ExitRoomMethod() string {
	return "stmp.examples.room.RoomService.ExitRoom"
}

func (s stmpRoomServiceBroadcaster) ExitRoomAction() uint64 {
	return 0x1204
}

func (s stmpRoomServiceBroadcaster) SendMessage(ctx context.Context, in *SendMessageInput, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomService.SendMessage", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s stmpRoomServiceBroadcaster) SendMessageForList(ctx context.Context, in *SendMessageInput, conns ...*stmp.Conn) error {
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

func (s stmpRoomServiceBroadcaster) SendMessageForKeys(ctx context.Context, in *SendMessageInput, conns stmp.ConnSet) error {
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

func (s stmpRoomServiceBroadcaster) BroadcastSendMessage(ctx context.Context, in *SendMessageInput, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomService.SendMessage", in, filter)
}

func (s stmpRoomServiceBroadcaster) SendMessageMethod() string {
	return "stmp.examples.room.RoomService.SendMessage"
}

func (s stmpRoomServiceBroadcaster) SendMessageAction() uint64 {
	return 0x1205
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
	UserEnter(ctx context.Context, in *UserEnterEvent, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error)
	UserEnterForList(ctx context.Context, in *UserEnterEvent, conns ...*stmp.Conn) error
	UserEnterForKeys(ctx context.Context, in *UserEnterEvent, conns stmp.ConnSet) error
	BroadcastUserEnter(ctx context.Context, in *UserEnterEvent, srv *stmp.Server, filter stmp.ConnFilter) error
	UserEnterMethod() string
	UserEnterAction() uint64
	UserExit(ctx context.Context, in *UserExitEvent, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error)
	UserExitForList(ctx context.Context, in *UserExitEvent, conns ...*stmp.Conn) error
	UserExitForKeys(ctx context.Context, in *UserExitEvent, conns stmp.ConnSet) error
	BroadcastUserExit(ctx context.Context, in *UserExitEvent, srv *stmp.Server, filter stmp.ConnFilter) error
	UserExitMethod() string
	UserExitAction() uint64
	NewMessage(ctx context.Context, in *ChatMessageModel, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error)
	NewMessageForList(ctx context.Context, in *ChatMessageModel, conns ...*stmp.Conn) error
	NewMessageForKeys(ctx context.Context, in *ChatMessageModel, conns stmp.ConnSet) error
	BroadcastNewMessage(ctx context.Context, in *ChatMessageModel, srv *stmp.Server, filter stmp.ConnFilter) error
	NewMessageMethod() string
	NewMessageAction() uint64
}

type stmpRoomEventsBroadcaster struct{}

func (s stmpRoomEventsBroadcaster) UserEnter(ctx context.Context, in *UserEnterEvent, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomEvents.UserEnter", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s stmpRoomEventsBroadcaster) UserEnterForList(ctx context.Context, in *UserEnterEvent, conns ...*stmp.Conn) error {
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

func (s stmpRoomEventsBroadcaster) UserEnterForKeys(ctx context.Context, in *UserEnterEvent, conns stmp.ConnSet) error {
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

func (s stmpRoomEventsBroadcaster) BroadcastUserEnter(ctx context.Context, in *UserEnterEvent, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomEvents.UserEnter", in, filter)
}

func (s stmpRoomEventsBroadcaster) UserEnterMethod() string {
	return "stmp.examples.room.RoomEvents.UserEnter"
}

func (s stmpRoomEventsBroadcaster) UserEnterAction() uint64 {
	return 0x1301
}

func (s stmpRoomEventsBroadcaster) UserExit(ctx context.Context, in *UserExitEvent, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomEvents.UserExit", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s stmpRoomEventsBroadcaster) UserExitForList(ctx context.Context, in *UserExitEvent, conns ...*stmp.Conn) error {
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

func (s stmpRoomEventsBroadcaster) UserExitForKeys(ctx context.Context, in *UserExitEvent, conns stmp.ConnSet) error {
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

func (s stmpRoomEventsBroadcaster) BroadcastUserExit(ctx context.Context, in *UserExitEvent, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomEvents.UserExit", in, filter)
}

func (s stmpRoomEventsBroadcaster) UserExitMethod() string {
	return "stmp.examples.room.RoomEvents.UserExit"
}

func (s stmpRoomEventsBroadcaster) UserExitAction() uint64 {
	return 0x1302
}

func (s stmpRoomEventsBroadcaster) NewMessage(ctx context.Context, in *ChatMessageModel, conn *stmp.Conn, opts ...stmp.CallOption) (*empty.Empty, error) {
	out, err := conn.Invoke(ctx, "stmp.examples.room.RoomEvents.NewMessage", in, stmp.NewCallOptions(opts...))
	return out.(*empty.Empty), err
}

func (s stmpRoomEventsBroadcaster) NewMessageForList(ctx context.Context, in *ChatMessageModel, conns ...*stmp.Conn) error {
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

func (s stmpRoomEventsBroadcaster) NewMessageForKeys(ctx context.Context, in *ChatMessageModel, conns stmp.ConnSet) error {
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

func (s stmpRoomEventsBroadcaster) BroadcastNewMessage(ctx context.Context, in *ChatMessageModel, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "stmp.examples.room.RoomEvents.NewMessage", in, filter)
}

func (s stmpRoomEventsBroadcaster) NewMessageMethod() string {
	return "stmp.examples.room.RoomEvents.NewMessage"
}

func (s stmpRoomEventsBroadcaster) NewMessageAction() uint64 {
	return 0x1303
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
