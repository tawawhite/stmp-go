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
	if [ ! -d googleapis ]; then \
		git clone https://github.com/googleapis/googleapis --depth 1; \
	else \
		cd googleapis; \
		git pull -r; \
	fi

build-example-proto: build-gen-stmp
	protoc --proto_path=vendor --proto_path=. \
		--proto_path=googleapis \
		--plugin=protoc-gen-stmp=$$PWD/out/protoc-gen-stmp \
		--gogofast_out=$$GOPATH/src \
		--stmp_out=lang=go:$$GOPATH/src \
		./examples/room/room_proto/*.proto
	pbjs -t static-module -w es6 -p ./vendor -p ./googleapis \
		-o ./examples/room/room_proto/room.pb.js ./examples/room/room_proto/*.proto
	pbts -n pb \
		-o ./examples/room/room_proto/room.pb.d.ts ./examples/room/room_proto/room.pb.js
	protoc --proto_path=vendor --proto_path=. \
		--proto_path=googleapis \
		--plugin=protoc-gen-stmp=$$PWD/out/protoc-gen-stmp \
		--stmp_out=lang=js,js.pb=./examples/room/room_proto/room.pb.js,js.out=./examples/room/room_proto/room.stmp.js:. \
		./examples/room/room_proto/*.proto

all: init build build-example-proto

run-example-room-server:
	go run examples/room/room_server/main.go

run-example-room-client:
	go run examples/room/room_client/main.go
