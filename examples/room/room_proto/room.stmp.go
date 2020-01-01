// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-01 22:10:09
package room_proto

import (
	"context"
	"github.com/acrazing/stmp-go/stmp"
	"github.com/golang/protobuf/ptypes/empty"
)

type STMPUserServiceServer interface {
	ListUser(ctx context.Context, in *ListUserInput) (out *ListUserOutput, err error)
}

type STMPUserEventsServer interface {
	StatusUpdated(ctx context.Context, in *UserModel) (out *empty.Empty, err error)
}

type STMPRoomServiceServer interface {
	CreateRoom(ctx context.Context, in *CreateRoomInput) (out *RoomModel, err error)
	ListRoom(ctx context.Context, in *ListRoomInput) (out *ListRoomOutput, err error)
	JoinRoom(ctx context.Context, in *JoinRoomInput) (out *empty.Empty, err error)
	ExitRoom(ctx context.Context, in *ExitRoomInput) (out *empty.Empty, err error)
	SendMessage(ctx context.Context, in *SendMessageInput) (out *empty.Empty, err error)
}

type STMPRoomEventsServer interface {
	UserEnter(ctx context.Context, in *UserEnterEvent) (out *empty.Empty, err error)
	UserExit(ctx context.Context, in *UserExitEvent) (out *empty.Empty, err error)
	NewMessage(ctx context.Context, in *NewMessageEvent) (out *empty.Empty, err error)
}

func init() {
	stmp.RegisterMethodActions(
		"stmp.examples.room.UserService.ListUser", 0x1001,
		"stmp.examples.room.UserEvents.StatusUpdated", 0x1101,
		"stmp.examples.room.RoomService.CreateRoom", 0x1201,
		"stmp.examples.room.RoomService.ListRoom", 0x1202,
		"stmp.examples.room.RoomService.JoinRoom", 0x1203,
		"stmp.examples.room.RoomService.ExitRoom", 0x1204,
		"stmp.examples.room.RoomService.SendMessage", 0x1205,
		"stmp.examples.room.RoomEvents.UserEnter", 0x1301,
		"stmp.examples.room.RoomEvents.UserExit", 0x1302,
		"stmp.examples.room.RoomEvents.NewMessage", 0x1303,
	)
}

func STMPRegisterUserServiceServer(r *stmp.Router, s STMPUserServiceServer) {
	r.Register("stmp.examples.room.UserService.ListUser", func() interface{} {
		return &CreateRoomInput{}
	}, func(ctx context.Context, in interface{}) (out interface{}, err error) {
		return s.ListUser(ctx, in.(*ListUserInput))
	})
}
