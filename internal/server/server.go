package server

import (
	"context"
	"fmt"
	pb "github.com/xasai/todogo/internal/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const PORT = ":4242"

type Server struct {
	pb.UnimplementedAddTicketServer
}

func (s *Server) SendTicket(ctx context.Context, ticket *pb.Ticket) (*pb.Response, error) {
	log.Println("Received ticket with title", ticket.Title)
	return &pb.Response{Body: "success!", Success: true}, nil
}

func (s *Server) mustEmbedUnimplementedAddTicketServer() {
	panic("implement me")
}

func Run() {
	//Create new listener on port 4242
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := Server{}//Server instance
	gs := grpc.NewServer()
	pb.RegisterAddTicketServer(gs, &s)

	log.Printf("\n Server now listening at %v \n", listener.Addr())
	if err := gs.Serve(listener); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

func init() {
	fmt.Print("Launching server")
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second / 2)
		fmt.Print(".")
	}
}