package main

import (
	"context"
	"fmt"
	pb "github.com/acrazing/stmp-go/examples/quick_start/quick_start_pb"
	"github.com/acrazing/stmp-go/stmp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strings"
	"sync"
	"time"
)

type RoomScene struct {
	mu    sync.Mutex
	rsc   pb.STMPRoomServiceClient
	sc    *stmp.Client
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
				pb.STMPDetachRoomEventsListener(r.sc, r)
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
					pb.STMPAttachRoomEventsListener(r.sc, r)
				}
			}
		}
		r.mu.Unlock()
	}
}

func NewRoomScene(rsc pb.STMPRoomServiceClient, sc *stmp.Client) *RoomScene {
	return &RoomScene{
		rsc:   rsc,
		sc:    sc,
		room:  nil,
		users: nil,
	}
}

func main() {
	logConfig := zap.NewProductionConfig()
	logConfig.DisableCaller = true
	logConfig.EncoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
	}
	logConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, err := logConfig.Build()
	if err != nil {
		log.Fatalf("init logger error: %q.", err)
	}
	sc := stmp.NewClient(stmp.NewClientOptions().WithLogger(logger).WithEncoding("gzip").WithPacketFormat("text"))
	sc.HandleConnected(func(header stmp.Header, message string) {
		logger.Info("stmp client connected", zap.String("message", message))
	})
	sc.HandleDisconnected(func(reason stmp.StatusError, willRetry bool, retryCount int, retryWait time.Duration) {
		logger.Warn("stmp client disconnected",
			zap.String("reason", reason.Error()),
			zap.Bool("willRetry", willRetry),
			zap.Int("retryCount", retryCount),
			zap.Duration("retryWait", retryWait))
	})
	go sc.DialWebsocket("ws://127.0.0.1:5002/quick_start")
	rsc := pb.STMPNewRoomServiceClient(sc)
	scene := NewRoomScene(rsc, sc)
	scene.Run()
}
