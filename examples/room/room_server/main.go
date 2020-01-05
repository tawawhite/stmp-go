// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-27 13:27:58
package main

import (
	"context"
	"github.com/acrazing/stmp-go/examples/room/room"
	roompb "github.com/acrazing/stmp-go/examples/room/room_pb"
	"github.com/acrazing/stmp-go/stmp"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type UserService struct {
	ut *room.UserTable
}

func (u *UserService) Login(ctx context.Context, in *roompb.LoginInput) (out *roompb.UserModel, err error) {
	panic("implement me")
}

func (u *UserService) ListUser(ctx context.Context, in *roompb.ListInput) (out *roompb.ListUserOutput, err error) {
	out = new(roompb.ListUserOutput)
	out.Total, out.Users = u.ut.List(in.Limit, in.Offset)
	return
}

func NewUserService(ut *room.UserTable) roompb.STMPUserServiceServer {
	return &UserService{ut: ut}
}

func main() {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	srvConfig := stmp.NewServerConfig()
	srvConfig.Logger = log
	stmp.RegisterMediaCodec(stmp.NewJsonCodec(), stmp.NewProtobufCodec())
	srv := stmp.NewServer(srvConfig)

	ut := room.NewUserTable()
	us := NewUserService(ut)
	roompb.STMPRegisterUserServiceServer(srv, us)

	go srv.ListenAndServeTCP("127.0.0.1:5001")
	log.Info("room server is listening at tcp://127.0.0.1:5001")
	go srv.ListenAndServeWebSocket("127.0.0.1:5002", "/ws")
	log.Info("room server is listening at ws://127.0.0.1:5002/ws")
	go func() {
		killSignal := make(chan os.Signal, 1)
		signal.Notify(killSignal, syscall.SIGINT, syscall.SIGTERM)
		sig := <-killSignal
		log.Info("receiving kill signal, server will shutdown", zap.String("signal", sig.String()))
		srv.Close()
	}()
	err = srv.Wait()
	if err != nil {
		log.Error("room server listen error", zap.Error(err))
		os.Exit(1)
	} else {
		log.Info("room server shutdown")
	}
}
