// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-04 23:06:03
package main

import (
	"context"
	"fmt"
	pb "github.com/acrazing/stmp-go/examples/quick_start/quick_start_pb"
	"github.com/acrazing/stmp-go/stmp"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

type RoomScene struct {
	mu    sync.Mutex
	rsc   pb.STMPRoomServiceClient
	conn  *stmp.ClientConn
	room  *string
	users []string
}

func (r *RoomScene) HandleUserJoin(ctx context.Context, in *pb.UserJoinEvent) (out *empty.Empty, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.room == nil {
		return
	}
	log.Printf("New user %s. entered.", in.User)
	r.users = append(r.users, in.User)
	log.Printf("%s users: %s.", *r.room, strings.Join(r.users, ", "))
	r.PrintTip()
	return
}

func (r *RoomScene) HandleUserExit(ctx context.Context, in *pb.UserExitEvent) (out *empty.Empty, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.room == nil {
		return
	}
	log.Printf("User %s exited.", in.User)
	for pos, u := range r.users {
		if u == in.User {
			r.users[pos] = r.users[len(r.users)-1]
			r.users = r.users[:len(r.users)-1]
			break
		}
	}
	log.Printf("%s users: %s.", *r.room, strings.Join(r.users, ", "))
	r.PrintTip()
	return
}

func (r *RoomScene) PrintTip() {
	if r.room == nil {
		log.Print("Enter a room name to join: ")
	} else {
		log.Printf("Enter ^D to exit %s.", *r.room)
	}
}

func (r *RoomScene) Run() {
	go func() {
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
			if strings.TrimSpace(roomInput) != "" {
				out, err := r.rsc.JoinRoom(context.Background(), &pb.JoinRoomInput{Name: roomInput})
				if err != nil {
					log.Printf("Join room %s error: %s.", roomInput, err)
				} else {
					log.Printf("Welcome to room: %s.", out.Name)
					log.Printf("Users in this room: %s.", strings.Join(out.Users, ", "))
					r.mu.Lock()
					r.room = &out.Name
					r.users = out.Users
					r.mu.Unlock()
					pb.STMPRegisterRoomEventsListener(r.conn, r)
				}
			}
		}
	}()
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT)
	for {
		<-exitSignal
		r.mu.Lock()
		if r.room != nil {
			_, err := r.rsc.ExitRoom(context.Background(), &pb.ExitRoomInput{Name: *r.room})
			if err != nil {
				log.Printf("Exit %s error: %s.", *r.room, err)
			} else {
				log.Printf("Exit %s.", *r.room)
				r.PrintTip()
				r.room = nil
				r.users = nil
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
	conn, err := stmp.DialTCP("127.0.0.1:5001", nil)
	if err != nil {
		log.Fatalf("dial error: %s", err)
	}
	rsc := pb.STMPNewRoomServiceClient(conn)
	scene := NewRoomScene(rsc, conn)
	scene.Run()
}
