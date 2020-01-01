// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-27 13:27:58
package room_client

import (
	"context"
	"github.com/acrazing/stmp-go/examples/room/room_proto"
	"github.com/acrazing/stmp-go/stmp"
	"github.com/golang/protobuf/ptypes/empty"
)

type RoomScene struct {
	rc   room_proto.STMPRoomServiceClient
	room *room_proto.RoomModel
}

func NewRoomEventsListener() room_proto.STMPRoomEventsServer {
	return &RoomScene{}
}

func (re *RoomScene) UserEnter(ctx context.Context, in *room_proto.UserEnterEvent) (out *empty.Empty, err error) {
	return
}

func (re *RoomScene) UserExit(ctx context.Context, in *room_proto.UserExitEvent) (out *empty.Empty, err error) {
	return
}

func main() {
	conn, err := stmp.DialTCP("127.0.0.1:5001", nil)
	if err != nil {
		panic(err)
	}
	rc := room_proto.STMPNewRoomServiceClient(conn)

	re := NewRoomEventsListener(rc)
}
