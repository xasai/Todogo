#Standart rule (first)
all: run 
		
run: build
	@./todo-server 

build: todo-server todo-cli

todo-cli:
	go build -o todo-cli cmd/client/main.go

todo-server:
	go build -o todo-server cmd/server/main.go
	
test: #TODO

fmt:
	go fmt ./...

fclean:
	rm -rf todo-server todo-cli		

protoc:
	@bash -c "protoc --go_out=. --go-grpc_out=.  internal/protobuf/list.proto 1>/dev/null"


.PHONY: todo-server todo-cli
