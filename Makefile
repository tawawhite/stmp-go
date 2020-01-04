build: proto gen-stmp

proto:
	protoc --proto_path=vendor --proto_path=. --go_out=$$GOPATH/src ./stmp/*.proto

gen-stmp:
	go build -o ./out/protoc-gen-stmp ./protoc-gen-stmp

test:
	go test ./stmp

bench:
	go test -bench=. ./stmp

init:
	go mod download
	go mod vendor

proto-room: gen-stmp
	protoc --proto_path=vendor --proto_path=. \
		--plugin=protoc-gen-stmp=$$PWD/out/protoc-gen-stmp \
		--gogofast_out=$$GOPATH/src \
		--validate_out=lang=gogo:$$GOPATH/src \
		--stmp_out=lang=go:$$GOPATH/src \
		./examples/room/room_proto/*.proto
	pbjs -t static-module -w commonjs -p ./vendor \
		--no-create --no-verify \
		--no-convert --no-delimited --keep-case --sparse \
		-o ./examples/room/room_proto/room.pb.js ./examples/room/room_proto/*.proto
	pbts -n pb --no-comments \
		-o ./examples/room/room_proto/room.pb.d.ts ./examples/room/room_proto/room.pb.js
	protoc --proto_path=vendor --proto_path=. \
		--plugin=protoc-gen-stmp=$$PWD/out/protoc-gen-stmp \
		--stmp_out=lang=js,js.module=cjs,js.pb=./examples/room/room_proto/room.pb.js,js.out=./examples/room/room_proto/room.stmp.js:. \
		./examples/room/room_proto/*.proto

proto-quick-start: gen-stmp
	protoc --proto_path=vendor --proto_path=. \
		--plugin=protoc-gen-stmp=$$PWD/out/protoc-gen-stmp \
		--gogofast_out=$$GOPATH/src \
		--stmp_out=lang=go:$$GOPATH/src \
		./examples/quick_start/quick_start_pb/*.proto

all: init build proto-room proto-quick-start

run-room-server:
	go run ./examples/room/room_server

run-room-client:
	go run examples/room/room_client

run-quick-start-server:
	go run ./examples/quick_start/quick_start_server

run-quick-start-client:
	go run ./examples/quick_start/quick_start_client
