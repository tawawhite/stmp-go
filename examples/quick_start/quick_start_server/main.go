package main

import (
	"context"
	pb "github.com/acrazing/stmp-go/examples/quick_start/quick_start_pb"
	"github.com/acrazing/stmp-go/stmp"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"log"
	"sync"
)

type RoomService struct {
	reb   pb.STMPRoomEventsBroadcaster
	mu    sync.RWMutex
	rooms map[string]stmp.ConnSet
	users map[string][]string
}

func (r *RoomService) JoinRoom(ctx context.Context, in *pb.JoinRoomInput) (out *pb.RoomModel, err error) {
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
			r.reb.UserJoinToSet(context.Background(), &pb.UserJoinEvent{User: user}, r.rooms[in.Name], conn)
			r.mu.RUnlock()
		}()
		room.Add(conn)
		r.users[in.Name] = append(r.users[in.Name], user)
	}
	out.Users = r.users[in.Name]
	r.mu.Unlock()
	return
}

func (r *RoomService) ExitRoom(ctx context.Context, in *pb.ExitRoomInput) (out *empty.Empty, err error) {
	conn := stmp.SelectConn(ctx)
	user := conn.RemoteAddr().String()
	r.mu.Lock()
	room, ok := r.rooms[in.Name]
	if ok {
		defer func() {
			r.mu.RLock()
			r.reb.UserExitToSet(context.Background(), &pb.UserExitEvent{User: user}, r.rooms[in.Name], conn)
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
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("init logger error: %s", err)
	}
	srv := stmp.NewServer(stmp.NewServerOptions().WithLogger(logger))
	pb.STMPRegisterRoomServiceServer(srv, NewRoomServiceServer())
	go srv.ListenAndServeTCP("127.0.0.1:5001")
	log.Println("server is listening at tcp://127.0.0.1:5001")
	go srv.ListenAndServeWebSocket("127.0.0.1:5002", "/quick_start")
	log.Println("server is listening at  ws://127.0.0.1:5002/quick_start")
	err = srv.Wait()
	if err != nil {
		log.Fatalf("server listen error: %s", err)
	}
}
