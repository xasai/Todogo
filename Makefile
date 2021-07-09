gosrc := $(shell find . -type f -name *.go)

all:
	go run server.go


run:
	go run ./cmd/server/main.go

build:
	

test:

install:
	go install ./...

fmt:
	go fmt $(gosrc)
