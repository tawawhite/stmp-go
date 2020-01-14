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

get:
	go get -u ./...

proto-gomoku: gen-stmp
	protoc --proto_path=vendor --proto_path=. \
		--plugin=protoc-gen-stmp=$$PWD/out/protoc-gen-stmp \
		--gogofast_out=$$GOPATH/src \
		--validate_out=lang=gogo:$$GOPATH/src \
		--stmp_out=lang=go:$$GOPATH/src \
		./examples/gomoku/gomoku_pb/*.proto
	yarn pbjs -t static-module -w commonjs -p ./vendor -p . \
		--no-create --no-verify \
		--no-convert --no-delimited --keep-case --sparse \
		-o ./examples/gomoku/gomoku_pb/gomoku.pb.js ./examples/gomoku/gomoku_pb/*.proto
	yarn pbts -n pb --no-comments \
		-o ./examples/gomoku/gomoku_pb/gomoku.pb.d.ts ./examples/gomoku/gomoku_pb/gomoku.pb.js
	protoc --proto_path=vendor --proto_path=. \
		--plugin=protoc-gen-stmp=$$PWD/out/protoc-gen-stmp \
		--stmp_out=lang=js,js.module=cjs,js.pb=./examples/gomoku/gomoku_pb/gomoku.pb.js,js.out=./examples/gomoku/gomoku_pb/gomoku.stmp.js:. \
		./examples/gomoku/gomoku_pb/*.proto

proto-quick-start: gen-stmp
	protoc --proto_path=vendor --proto_path=. \
		--plugin=protoc-gen-stmp=$$PWD/out/protoc-gen-stmp \
		--gogofast_out=$$GOPATH/src \
		--stmp_out=lang=go:$$GOPATH/src \
		./examples/quick_start/quick_start_pb/*.proto

all: init build proto-gomoku proto-quick-start

run-gomoku-server:
	go run ./examples/gomoku/gomoku_server

run-gomoku-client:
	go run examples/gomoku/gomoku_client

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