package main

import (
	"context"
	"fmt"
	"github.com/acrazing/stmp-go/examples/room/room"
	roompb "github.com/acrazing/stmp-go/examples/room/room_pb"
	"github.com/acrazing/stmp-go/stmp"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type LobbyScene struct {
	us *room.UserTable
	uc roompb.STMPUserServiceClient
}

func (ls *LobbyScene) Mount() {
	log.Println("Entering lobby scene...")
	out, err := ls.uc.ListUser(context.Background(), &roompb.ListInput{Limit: 20})
	if err != nil {
		log.Println("list user error:", err)
		return
	}
	ls.us.Push(out.Users)
}

func (ls *LobbyScene) Unmount() {
	log.Println("Exiting lobby scene...")
}

func NewLobbyScene(uc roompb.STMPUserServiceClient) *LobbyScene {
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
	var conn *stmp.Client
	var dialOptions = stmp.NewDialOptions().WithHeader("X-User-Name", name)
	if strings.HasPrefix(addr, "ws://") || strings.HasPrefix(addr, "wss://") {
		log.Println("dialing", addr)
		conn, err = stmp.DialWebsocket(addr, dialOptions)
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

	usc := roompb.STMPNewUserServiceClient(conn)
	ls := NewLobbyScene(usc)
	ls.Mount()

	killSignal := make(chan os.Signal)
	signal.Notify(killSignal, syscall.SIGINT, syscall.SIGTERM)
	sig := <-killSignal
	log.Printf("exiting for receiving signal: %s\n", sig.String())
}
