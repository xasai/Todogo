package server

import (
	"fmt"
	pb "github.com/xasai/todogo/internal/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	PORT = ":4242"
	DONE = true
	TODO = false
)

type todoServer struct {
	pb.UnimplementedTodoServServer
}

var Tickets []pb.Ticket

func (s *todoServer) TodoRequest(req *pb.Request, stream pb.TodoServ_TodoRequestServer)  error {
	if req.Method == "GET" {
		HandleGetRequest(req, stream)
	} else if req.Method == "PUT" {
		HandlePutRequest(req)
	} else if req.Method == "DEL" {
		HandleDelRequest(req)
	}
	return nil
}

func Run() {
	//Create new listener on port 4242
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServServer(grpcServer, &todoServer{})

	log.Printf("\n Server now listening at %v \n", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

func init() {
	fmt.Print("Launching server")
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second / 2)
		fmt.Print(".")
	}
	// initializing 0 id with empty Ticket
	Tickets = append(Tickets[0:], pb.Ticket{})
	fmt.Println()
}