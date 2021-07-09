package client

import (
	"context"
	"fmt"
	pb "github.com/xasai/todogo/internal/protobuf"
	"google.golang.org/grpc"
	"log"
)

const (
	PORT = ":4242"
	DONE bool = true
	TODO bool = false
)

func RunInteractiveMod() {
	log.Println("Running interactive mod")

	conn , err := grpc.Dial(PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %v\n", err)
	}
	defer conn.Close()

	ticket := pb.Ticket{
		Id: 0,
		Title: "First ticket ever",
		Body: "-",
		State: TODO,
	}
	cli := pb.NewAddTicketClient(conn)
	resp, _ := cli.SendTicket(context.Background(), &ticket)
	log.Println("SERVER RESPONSE:", resp)
}


func printTicket(t *pb.Ticket) {
	fmt.Println("ID:", t.Id)
	fmt.Println("Title:", t.Title)
	fmt.Println("Body:", t.Body)
	//converting ticket's state to message
	state := map[bool]string{TODO: "To do", DONE: "Done"}[t.State == DONE]
	fmt.Println("State:", state)
}