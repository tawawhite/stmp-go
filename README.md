# stmp-go

A lightweight real-time bidirectional framework for Golang.

**Features**

- Fast
    - Supports `Protocol Buffers`, `MsgPack`
    - Extremely tidy packet format, only 1 byte header
    - Supports `KCP`
    - Supports connection level `gzip` compression
- Easy to use
    - Supports `WebSockets` and highly optimized for it
    - Supports `text` format packet for `WebSockets`
    - Supports `JSON`
    - Supports compose different listeners in one server
    - Most concepts are the same as `gRPC`
    - `.proto` based service definition
- Secure
    - Supports `TLS`

*Supported transport layer protocols*

- `KCP`
- `TCP`
- `WebSockets`

## Install

```bash
go get -u github.com/acrazing/stmp-go
```

## Examples

You can get full examples at [examples](./examples) directory, you can run the examples with source code in root
directory, the examples list:

- [quick start](./examples/quick_start): The quick start example

    ```bash
    # run server
    make run-quick-start-server
  
    # run client
    make run-quick-start-client
    ```

- [room](./examples/room): A complex chat room service, include golang server, golang terminal client, and typescript browser client.

    ```bash
    # run server
    make run-room-server
  
    # run client
    make run-room-client
    ```

## Quick Start

### Service definition

As indicated above, we use `.proto` as `IDL` of service, almost same to `gRPC`, we use `service` section to define
services. But there's a slight difference: a `service` could be implemented by a server, and called by client, at
the same time, another `service` may be called by server, and listened by client, aka `server push`. So we agreed
to use `Service` as the end for the services implemented by the server, and `Events` as the end for the services
listened on by the client.

Imaging we are implementing a real-time room service, a client can join or exit a room by call a server implemented
method, at the same time, the server well push the join or exit event to all the clients in the room. Then, the
service definition could look like this:

See the source at [quick_start.proto](./examples/quick_start/quick_start_pb/quick_start.proto).

```proto
syntax = "proto3";
package stmp.examples.quick_start;
option go_package = "github.com/acrazing/stmp-go/examples/quick_start/quick_start_pb;pb";
import "google/protobuf/empty.proto";

message RoomModel {
    string name = 1;
    repeated string users = 2;
}

message JoinInput {
    string name = 1;
}

message ExitInput {
    string name = 1;
}

service RoomService {
    rpc Join(JoinInput) returns (RoomModel) {}
    rpc Exit(ExitInput) returns (google.protobuf.Empty) {}
}

message JoinEvent {
    string user = 1;
}

message ExitEvent {
    string user = 1;
}

service RoomEvents {
    rpc Join(JoinEvent) returns (google.protobuf.Empty) {}
    rpc Exit(ExitEvent) returns (google.protobuf.Empty) {}
}
```

### Compile proto

*You may need to install `protoc` and `protoc-gen-go`(or `protoc-gen-gogo`) at first.*

You need to install `protoc-gen-stmp` at first:

```bash
go get github.com/acrazing/stmp-go/stmp/protoc-gen-stmp
```

And then compile the `.proto` files:

```bash
protoc --go_out=. --stmp_out=. pb/room.proto
```

*Or you can run `make proto-quick-start` in source code root of this project*.

### Server implementation

See the source at [quick_start_server](./examples/quick_start/quick_start_server/main.go).

```go
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
```

### Client implementation

See the source at [quick_start_client](./examples/quick_start/quick_start_client/main.go).

```go
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
```

## LICENSE

```text
The MIT License (MIT)

Copyright (c) 2019-2020 acrazing

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```