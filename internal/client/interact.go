package client

import (
	"bufio"
	"context"
	"fmt"
	"github.com/inancgumus/screen"
	. "github.com/xasai/todogo/internal/cli"
	pb "github.com/xasai/todogo/internal/protobuf"
	"io"
	"log"
	"os"
	"time"
)

//Output strings
const (
	USAGE = "Write a number of action or [Q]uit: \n\t" +
		"1 - create new todo-note\n\t" +
		"2 - get note by Id\n\t" +
		"3 - get all existing notes\n\t" +
		"4 - change a content of note\n\t" +
		"5 - mark done\n\t" +
		"6 - delete note\n" +
		"action: "
)

func runInteractiveMod(cli pb.TodoServiceClient) {
	log.Println(YELL + "Running interactive mod" + RES)
	time.Sleep(time.Second)
	fmt.Println(RED + "Hello friend." + RES)
	time.Sleep(time.Second)
	interact(cli)
}

func interact(client pb.TodoServiceClient) {

	var (
		err     error
		res     *pb.Response
		reqId   pb.RequestId
		reqNote pb.Note
	)

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(GREEN + USAGE + CYAN)
		buf, _, _ := r.ReadLine()
		fmt.Println(RES)
		switch string(buf) {
		case "1":
			//CreateNewNote()
			//Creating new note with Title and Description read from keyboard.
			reqNote, err = readNote(r)
			if err != nil {
				fmt.Printf(RED+"%s\n"+RES, err.Error())
				break
			}
			res, err = client.CreateNewNote(context.Background(), &reqNote)
			pprintResponse(res, err)
		case "2":
			//GetNoteById()
			//Reading Id from keyboard and request Note with such Id.
			reqId.NoteId, err = readInt32(r)
			if err != nil {
				fmt.Printf(RED+"%s\n"+RES, err.Error())
				break
			}
			res, err = client.GetNoteById(context.Background(), &reqId)
			pprintResponse(res, err)
		case "3":
			//GetAllNotes()
			//Calling this method and printing all notes received in returned stream.
			stream, err := client.GetAllNotes(context.Background(), &pb.RequestId{})
			if err != nil {
				fmt.Printf(RED+"%s\n"+RES, err.Error())
				break
			}
			screen.Clear()
			screen.MoveTopLeft()
			for {
				n, err := stream.Recv()
				if err == io.EOF {
					break
				} else if err != nil {
					log.Fatalf("%v.GetAllNotes(_) = _, %v", client, err)
				}
				pprintNote(n)
				fmt.Println("\n")
			}
		case "4":
			//ChangeNoteContent()
			//Read note's id, title, description. Then ask if user sure to change note's content.
			//Send this new content to the grpc server.
			id, err := readInt32(r)
			if err != nil {
				fmt.Printf(RED+"%s\n"+RES, err.Error())
				break
			}
			reqNote, _ = readNote(r)
			reqNote.Id = id
			if sure(r) {
				res, err = client.ChangeNoteContent(context.Background(), &reqNote)
				pprintResponse(res, err)
			}
		case "5":
			//MarkDoneById()
			//Read id from keyboard and send it to grpc server.
			reqId.NoteId, err = readInt32(r)
			if err != nil {
				fmt.Printf(RED+"%s\n"+RES, err.Error())
				break
			}
			res, err = client.MarkDoneById(context.Background(), &reqId)
			pprintResponse(res, err)
		case "6":
			//DelNoteById()
			//Same as "5", but asking if client sure.
			reqId.NoteId, err = readInt32(r)
			if err != nil {
				fmt.Printf(RED+"%s\n"+RES, err.Error())
				break
			}
			if sure(r) {
				res, err = client.DelNoteById(context.Background(), &reqId)
				pprintResponse(res, err)
			}
		case "q", "Q":
			fmt.Println(RED + "See you soon" + RES)
			os.Exit(0)
		default:
			screen.Clear()
			screen.MoveTopLeft()
			fmt.Println(RED + "Insert only one of 1-6 or Q to quit!" + RES)
		}
	}
}
