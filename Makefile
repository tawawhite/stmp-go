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
		./examples/room/room_pb/*.proto
	pbjs -t static-module -w commonjs -p ./vendor \
		--no-create --no-verify \
		--no-convert --no-delimited --keep-case --sparse \
		-o ./examples/room/room_pb/room.pb.js ./examples/room/room_pb/*.proto
	pbts -n pb --no-comments \
		-o ./examples/room/room_pb/room.pb.d.ts ./examples/room/room_pb/room.pb.js
	protoc --proto_path=vendor --proto_path=. \
		--plugin=protoc-gen-stmp=$$PWD/out/protoc-gen-stmp \
		--stmp_out=lang=js,js.module=cjs,js.pb=./examples/room/room_pb/room.pb.js,js.out=./examples/room/room_pb/room.stmp.js:. \
		./examples/room/room_pb/*.proto

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
	CompileDaemon -build "go build -o ./out/quick_start_server ./examples/quick_start/quick_start_server" \
		--exclude-dir node_modules --exclude-dir vendor --exclude-dir .git --exclude-dir .idea \
		-command "./out/quick_start_server"

run-quick-start-client:
	CompileDaemon -build "go build -o ./out/quick_start_client ./examples/quick_start/quick_start_client" \
		--exclude-dir node_modules --exclude-dir vendor --exclude-dir .git --exclude-dir .idea \
		-command "./out/quick_start_client"

quick-start-server:
	go build -o ./out/quick_start_server ./examples/quick_start/quick_start_server

quick-start-client:
	go build -o ./out/quick_start_client ./examples/quick_start/quick_start_client