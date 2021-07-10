package client

import (
	"fmt"
	pb "github.com/xasai/todogo/internal/protobuf"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	PORT = ":4242"
	DONE bool = true
	TODO bool = false
)

func Run() {
	conn , err := grpc.Dial(PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %v\n", err)
	}
	defer conn.Close()
	RunInteractiveMod(conn)
}

func RunInteractiveMod(connection *grpc.ClientConn) {
	log.Println("Running interactive mod")
	log.Println("Successfully connected to grpc server . . . press Q to terminate")
	var (
		action string
		id int
		ticket pb.Ticket
	)
	for {
		fmt.Println("Write a number of action: \n\t" +
					"1 - create new todo-note\n\t" +
					"2 - get note\n\t" +
					"3 - get list of all notes\n\t" +
					"4 - change a content of note\n\t" +
					"5 - change state of a todo note\n\t" +
					"6 - delete note")

		fmt.Scanf("%s", &action)
		switch action {
		case "1":

		case "q", "Q":
			fmt.Println("See you soon")
			os.Exit(0)
		default:
			fmt.Println("Insert only on of 1-6 or Q to quit")
			continue
		}
	}
}


func printTicket(t *pb.Ticket) {
	fmt.Println("ID:", t.Id)
	fmt.Println("Title:", t.Title)
	fmt.Println("Body:", t.Body)
	//converting ticket's state to message
	state := map[bool]string{TODO: "To do", DONE: "Done"}[t.State == DONE]
	fmt.Println("State:", state)
}