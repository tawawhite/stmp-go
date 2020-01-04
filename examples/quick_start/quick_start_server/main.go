// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-04 23:06:08
package main

import (
	"context"
	pb "github.com/acrazing/stmp-go/examples/quick_start/quick_start_pb"
	"github.com/acrazing/stmp-go/stmp"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
	"sync"
)

type RoomService struct {
	reb   pb.STMPRoomEventsBroadcaster
	mu    sync.RWMutex
	rooms map[string]stmp.ConnSet
	users map[string][]string
}

func (r *RoomService) Join(ctx context.Context, in *pb.JoinInput) (out *pb.RoomModel, err error) {
	conn := stmp.SelectConn(ctx)
	// user remote address as the user name
	user := conn.RemoteAddr().String()
	out = new(pb.RoomModel)
	out.Name = in.Name
	r.mu.Lock()
	room, ok := r.rooms[in.Name]
	if !ok {
		room = stmp.NewConnSet()
		r.rooms[in.Name] = room
	}
	if !room.Has(conn) {
		defer func() {
			r.mu.RLock()
			r.reb.JoinToSet(context.Background(), &pb.JoinEvent{User: user}, r.rooms[in.Name], conn)
			r.mu.RUnlock()
		}()
		room.Add(conn)
		r.users[in.Name] = append(r.users[in.Name], user)
	}
	out.Users = r.users[in.Name]
	r.mu.Unlock()
	return
}

func (r *RoomService) Exit(ctx context.Context, in *pb.ExitInput) (out *empty.Empty, err error) {
	conn := stmp.SelectConn(ctx)
	user := conn.RemoteAddr().String()
	r.mu.Lock()
	room, ok := r.rooms[in.Name]
	if ok {
		defer func() {
			r.mu.RLock()
			r.reb.ExitToSet(context.Background(), &pb.ExitEvent{User: user}, r.rooms[in.Name], conn)
			r.mu.RUnlock()
		}()
		room.Del(conn)
		users := r.users[in.Name]
		for i, u := range users {
			if u == user {
				users[i] = users[len(users)-1]
				r.users[in.Name] = users[:len(users)-1]
				break
			}
		}
	}
	r.mu.Unlock()
	return
}

func NewRoomServiceServer() pb.STMPRoomServiceServer {
	return &RoomService{
		rooms: map[string]stmp.ConnSet{},
		users: map[string][]string{},
	}
}

func main() {
	srv := stmp.NewServer()
	pb.STMPRegisterRoomServiceServer(srv, NewRoomServiceServer())
	go srv.ListenAndServeTCP("127.0.0.1:5001")
	log.Println("server is listening at tcp://127.0.0.1:5001")
	go srv.ListenAndServeWebSocket("127.0.0.1:5002", "/quick_start")
	log.Println("server is listening at  ws://127.0.0.1:5002/quick_start")
	err := srv.Wait()
	if err != nil {
		log.Fatalf("server listen error: %s", err)
	}
}
