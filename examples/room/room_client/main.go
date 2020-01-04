// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-27 13:27:58
package room_client

import (
	"context"
	"fmt"
	"github.com/acrazing/stmp-go/examples/room/room"
	"github.com/acrazing/stmp-go/examples/room/room_proto"
	"github.com/acrazing/stmp-go/stmp"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type LobbyScene struct {
	us *room.UserTable
	uc room_proto.STMPUserServiceClient
}

func (ls *LobbyScene) Mount() {
	log.Println("Entering lobby scene...")
	out, err := ls.uc.ListUser(context.Background(), &room_proto.ListUserInput{Limit: 10})
	if err != nil {
		log.Println("list user error:", err)
		return
	}
	ls.us.Push(out.Users)
}

func (ls *LobbyScene) Unmount() {
	log.Println("Exiting lobby scene...")
}

func NewLobbyScene(uc room_proto.STMPUserServiceClient) *LobbyScene {
	return &LobbyScene{uc: uc}
}

func main() {
	log.Println("\x1b[1mWelcome to stmp chat room!\x1b[0m")
	log.Println("\x1b[1mPlease enter your name: \x1b[0m")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatalln("\nERROR: ", err)
	}
	var addr = "ws://127.0.0.1:5001/ws"
	log.Print("\x1b[1mPlease enter server addr [" + addr + "]: \x1b[0m")
	_, err = fmt.Scanln(&addr)
	if err != nil {
		log.Fatalln("\nERROR:", err)
	}
	var conn *stmp.Conn
	var dialOptions = &stmp.DialOptions{Header: stmp.NewHeader()}
	dialOptions.Header.Set("X-User-Name", name)
	if strings.HasPrefix(addr, "ws://") || strings.HasPrefix(addr, "wss://") {
		log.Println("dialing", addr)
		conn, err = stmp.DialWebSocket(addr, dialOptions)
	} else if strings.HasPrefix(addr, "kcp://") {
		log.Println("dialing", addr)
		conn, err = stmp.DialKCP(addr[6:], dialOptions)
	} else if strings.HasPrefix(addr, "tcp://") || strings.Index(addr, "://") == -1 {
		log.Println("dialing", addr)
		conn, err = stmp.DialTCP(strings.TrimPrefix(addr, "tcp://"), dialOptions)
	} else {
		log.Println("ERROR: unsupported address format:", addr)
		os.Exit(1)
	}
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	usc := room_proto.STMPNewUserServiceClient(conn)
	ls := NewLobbyScene(usc)
	ls.Mount()

	killSignal := make(chan os.Signal)
	signal.Notify(killSignal, syscall.SIGINT, syscall.SIGTERM)
	sig := <-killSignal
	log.Printf("exiting for receiving signal: %s\n", sig.String())
}
