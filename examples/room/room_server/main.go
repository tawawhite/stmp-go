// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-27 13:27:58
package main

import (
	"context"
	"github.com/acrazing/stmp-go/examples/room/room"
	"github.com/acrazing/stmp-go/examples/room/room_proto"
	"github.com/acrazing/stmp-go/stmp"
	"os"
	"time"
)

type UserService struct {
	us *room.UserStore
}

func (u *UserService) ListUser(ctx context.Context, in *room_proto.ListUserInput) (out *room_proto.ListUserOutput, err error) {
	u.us.Mu.RLock()
	defer u.us.Mu.RUnlock()
	out = &room_proto.ListUserOutput{
		Total: int64(len(u.us.Users)),
		Users: make([]*room_proto.UserModel, 0, in.Limit),
	}
	i := int64(0)
	max := in.Offset + in.Limit
	for _, u := range u.us.Users {
		if i < in.Offset {
			continue
		}
		i += 1
		out.Users = append(out.Users, u)
		if i > max {
			break
		}
	}
	return
}

func NewUserService(us *room.UserStore) room_proto.STMPUserServiceServer {
	return &UserService{us: us}
}

func main() {
	stmp.RegisterMediaCodec(stmp.NewJsonCodec(), stmp.NewProtobufCodec())
	srv := stmp.NewServer()
	srv.WriteTimeout = time.Minute
	srv.ReadTimeout = time.Minute
	srv.HandshakeTimeout = time.Minute

	userStore := room.NewUserStore()
	userService := NewUserService(userStore)
	room_proto.STMPRegisterUserServiceServer(srv, userService)

	conn, _ := stmp.DialTCP("", nil)
	usc := room_proto.STMPNewUserServiceClient(conn)
	usb := room_proto.STMPNewUserServiceBroadcaster()
	reb := room_proto.STMPNewRoomEventsBroadcaster()
	_ = reb.NewMessageToAll(nil, nil, nil, nil)
	_ = reb.UserEnterToSet(nil, nil, nil)
	_, _ = usc.ListUser(nil, &room_proto.ListUserInput{})
	_ = usb.ListUserToAll(nil, &room_proto.ListUserInput{}, srv, nil)

	go srv.ListenAndServeWebSocket("127.0.0.1:5002", "/ws")
	println("room server is listening at ws://127.0.0.1:5002/ws")
	err := srv.Wait()
	if err != nil {
		println("listen error", err.Error())
		os.Exit(1)
	}
}
