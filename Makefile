.PHONY: test bench init build-example-proto run-example-room-server run-example-room-client

test:
	go test

bench:
	go bench

init:
	go mod download
	go mod vendor

build-example-proto:
	protoc --proto_path=vendor --proto_path=. --gogofast_out=plugins=grpc:$$GOPATH/src --validate_out=lang=gogo:$$GOPATH/src ./examples/room/room_proto/*.proto

run-example-room-server:
	go run examples/room/room_server/main.go

run-example-room-client:
	go run examples/room/room_client/main.go
