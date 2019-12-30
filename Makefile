build: build-proto build-gen-stmp

build-proto:
	protoc --proto_path=vendor --proto_path=. --go_out=$$GOPATH/src ./stmp/*.proto

build-gen-stmp:
	go build -o ./out/protoc-gen-stmp ./protoc-gen-stmp

test:
	go test ./stmp

bench:
	go test -bench=. ./stmp

init:
	go mod download
	go mod vendor

build-example-proto: build-gen-stmp
	protoc --proto_path=vendor --proto_path=. \
		--plugin=protoc-gen-stmp=$$PWD/out/protoc-gen-stmp \
		--gogofast_out=$$GOPATH/src \
		--stmp_out=lang=go+esm+dts:$$GOPATH/src \
		./examples/room/room_proto/*.proto

all: init build build-example-proto

run-example-room-server:
	go run examples/room/room_server/main.go

run-example-room-client:
	go run examples/room/room_client/main.go
