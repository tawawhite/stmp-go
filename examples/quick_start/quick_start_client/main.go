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
	conn  *stmp.Conn
	room  *string
	users []string
}

func (r *RoomScene) Join(ctx context.Context, in *pb.JoinEvent) (out *empty.Empty, err error) {
	panic("implement me")
}

func (r *RoomScene) Exit(ctx context.Context, in *pb.ExitEvent) (out *empty.Empty, err error) {
	panic("implement me")
}

func (r *RoomScene) PrintTip() {
	if r.room == nil {
		log.Print("Please enter a room name to join or create: ")
	} else {
		log.Printf("Please enter ^D to exit %s\n", *r.room)
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
				log.Printf("Read input error: %s\n", err)
				continue
			}
			if strings.TrimSpace(roomInput) != "" {
				out, err := r.rsc.Join(context.Background(), &pb.JoinInput{Name: roomInput})
				if err != nil {
					log.Printf("Join room %s error: %s\n", roomInput, err)
				} else {
					log.Printf("Welcome to room: %s\n", out.Name)
					log.Printf("Users in this room: %s\n", strings.Join(out.Users, ", "))
					r.mu.Lock()
					r.room = &out.Name
					r.users = out.Users
					r.mu.Unlock()
					pb.STMPRegisterRoomEventsServer(r.conn, r)
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
			_, err := r.rsc.Exit(context.Background(), &pb.ExitInput{Name: *r.room})
			if err != nil {
				log.Printf("Exit %s error: %s\n", *r.room, err)
			} else {
				log.Printf("Exit %s done\n", *r.room)
				r.PrintTip()
				r.room = nil
				r.users = nil
			}
		}
		r.mu.Unlock()
	}
}

func NewRoomScene(rsc pb.STMPRoomServiceClient, conn *stmp.Conn) *RoomScene {
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
