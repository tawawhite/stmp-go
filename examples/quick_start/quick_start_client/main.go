package main

import (
	"context"
	"fmt"
	pb "github.com/acrazing/stmp-go/examples/quick_start/quick_start_pb"
	"github.com/acrazing/stmp-go/stmp"
	"go.uber.org/zap"
	"log"
	"strings"
	"sync"
)

type RoomScene struct {
	mu    sync.Mutex
	rsc   pb.STMPRoomServiceClient
	conn  *stmp.ClientConn
	room  *string
	users []string
}

func (r *RoomScene) HandleUserJoin(ctx context.Context, in *pb.UserJoinEvent) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.room == nil {
		return
	}
	log.Printf("\nUser %s entered.", in.User)
	r.users = append(r.users, in.User)
	log.Printf("Users in %s: %s.", *r.room, strings.Join(r.users, ", "))
	r.PrintTip()
}

func (r *RoomScene) HandleUserExit(ctx context.Context, in *pb.UserExitEvent) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.room == nil {
		return
	}
	log.Printf("\nUser %s exited.", in.User)
	for pos, u := range r.users {
		if u == in.User {
			r.users[pos] = r.users[len(r.users)-1]
			r.users = r.users[:len(r.users)-1]
			break
		}
	}
	log.Printf("Users in %s: %s.", *r.room, strings.Join(r.users, ", "))
	r.PrintTip()
}

func (r *RoomScene) PrintTip() {
	if r.room == nil {
		print("Enter a room name to join: ")
	} else {
		print("Enter `exit` to exit " + *r.room + ": ")
	}
}

func (r *RoomScene) Run() {
	for {
		var roomInput string
		r.mu.Lock()
		r.PrintTip()
		r.mu.Unlock()
		_, err := fmt.Scanln(&roomInput)
		if err != nil {
			log.Printf("Read input error: %s.", err)
			continue
		}
		roomInput = strings.TrimSpace(roomInput)
		r.mu.Lock()
		if r.room != nil && roomInput == "exit" {
			_, err := r.rsc.ExitRoom(context.Background(), &pb.ExitRoomInput{Name: *r.room})
			if err != nil {
				log.Printf("Exit room %s error: %s.", *r.room, err)
			} else {
				log.Printf("Exited room %s.", *r.room)
				r.room = nil
				r.users = nil
				r.PrintTip()
				pb.STMPUnregisterRoomEventsListener(r.conn, r)
			}
		} else if r.room == nil {
			if roomInput != "" {
				out, err := r.rsc.JoinRoom(context.Background(), &pb.JoinRoomInput{Name: roomInput})
				if err != nil {
					log.Printf("Join room %s error: %s.", roomInput, err)
				} else {
					log.Printf("Welcome to room: %s.", out.Name)
					log.Printf("Users in this room: %s.", strings.Join(out.Users, ", "))
					r.room = &out.Name
					r.users = out.Users
					pb.STMPRegisterRoomEventsListener(r.conn, r)
				}
			}
		}
		r.mu.Unlock()
	}
}

func NewRoomScene(rsc pb.STMPRoomServiceClient, conn *stmp.ClientConn) *RoomScene {
	return &RoomScene{
		rsc:   rsc,
		conn:  conn,
		room:  nil,
		users: nil,
	}
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("init logger error: %s", err)
	}
	conn, err := stmp.DialTCP("127.0.0.1:5001", stmp.NewDialOptions().WithLogger(logger).WithEncoding("gzip"))
	if err != nil {
		log.Fatalf("dial error: %s", err)
	}
	rsc := pb.STMPNewRoomServiceClient(conn)
	scene := NewRoomScene(rsc, conn)
	scene.Run()
}
