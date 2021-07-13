client = ./client 
csrc = $(wildcard ./internal/client/*.go)

server = ./server
ssrc = $(wildcard ./internal/server/*.go)

#Standart rule (first)
all: run
		
run: build
	$(server) & #Running server in background
	sleep 2 
	$(client)

build: $(client) $(server)

$(client): $(csrc)
	go build -o $(client) cmd/client/main.go

$(server): $(ssrc)
	go build -o $(server) cmd/server/main.go

fclean:
	rm -rf $(server) $(client)

get:
	go get ./... 

fmt:
	go fmt ./...

protoc:
	protoc --go_out=. --go-grpc_out=.  internal/protobuf/note.proto
