# stmp-go

A lightweight real-time bidirectional framework for Golang.

**Features**

- Supports multiple transport layer protocols, include `TCP`, `KCP`, `WebSockets`
- Supports `TLS`
- Supports `gzip`
- Supports listen multiple ports with different protocol at the same time
- `Content-Type` free, supports `Protocol Buffers`, `JSON`, `msgpack` internally
- `WebSockets` optimized specially
    - Supports text and binary format
    - Supports `UTF-8` and `UTF-16` for text payload
- `.proto` based service definition

## Install

```bash
go get -u github.com/acrazing/stmp-go
```

## Usage

### Service definition

A service could be invoked by both server and client, for client, it is one(client) to one(server), which means
same to `gRPC`, for server, it is one(server) to many(clients), the api is different.

All `.proto` is same to `gRPC`, the difference is the generated code.

For example, a `room` service, allows client to join or exit a room, and server should notify other clients the event.

```proto
// room/room/room.proto

syntax = "proto3";
package room;

message Empty {
}

message RoomModel {
    int64 id = 1;
    repeated int64 users = 2;
}

message JoinRoomInput {
    int64 roomId = 1;
}

message JoinRoomOutput {
    RoomModel room = 1;
}

message ExitRoomInput {
    int64 roomId = 1;
}

service RoomService {
    rpc Join(JoinRoomInput) returns (JoinRoomOutput);
    rpc Exit(ExitRoomInput) returns (Empty);
}

message JoinRoomEvent {
    int64 userId = 1;
}

message ExitRoomEvnet {
    int64 userId = 1;
}

service RoomEvent {
    rpc Join(JoinRoomEvent) returns (Empty);
    rpc Exit(ExitRoomEvent) returns (Empty);
}
```

### Compile proto

*You may need to install `protoc` and `gogo/protobuf` at first.*

```bash
protoc --gogofaster_out=. --stmp_out=. room/room/room.proto
```

### Server side

```go
// room/server/main.go

package main

import "github.com/acrazing/stmp-go/stmp"
import pb "room/room"

type roomService struct {
    re pb.StmpRoomEvent
}

func NewRoomService() pb.StmpRoomService {
    return &roomService{
        re: pb.StmpNewRoomEvent(),
    }
}

func (rs *roomService) Join(ctx stmp.Context, input *pb.JoinRoomInput) (*pb.JoinRoomOutput, error) {
}

func (rs *roomService) Exit(ctx stmp.Context, input *pb.ExitRoomInput) (*pb.Empty, error) {
}

func main() {
}
```