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
	pb.UnimplementedTodoServiceServer
}

var Notes = make(map[int32]*pb.Note) //Here I stored all Notes

func Run() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTodoServiceServer(grpcServer, &todoServer{})
	log.Printf("\nServer now listening at %v\tPress Ctrl-C to exit\n", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Server Failure%v", err)
	}
}

func init() {
	fmt.Print("Launching server")
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second / 2)
		fmt.Print(".")
	}
	fmt.Println()
}
