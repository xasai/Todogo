package client

import (
	"fmt"
	"github.com/inancgumus/screen"
	. "github.com/xasai/todogo/internal/cli"
	pb "github.com/xasai/todogo/internal/protobuf"
	"google.golang.org/grpc"
	"log"
	"time"
	"strings"
)

const (
	PORT = ":4242"
	DONE = true
	TODO = false
)

func Run() {
	screen.Clear()
	screen.MoveTopLeft()
	cc, err := grpc.Dial(PORT,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.FailOnNonTempDialError(true))
	if err != nil {
		log.Fatalf(RED+"Failed to establish connection with gRPC: %s\n"+RES, err.Error())
	}
	defer cc.Close()
	log.Println(YELL + "Successfully connected to grpc server" + RES)
	client := pb.NewTodoServiceClient(cc)
	runInteractiveMod(client)
}

func pprintResponse(r *pb.Response, err error) {
	if err != nil {
		log.Panic(err)
	}
	screen.Clear()
	screen.MoveTopLeft()

	w, _ := screen.Size()
	template := fmt.Sprintf("%0*d", w, 0)
	delim := strings.ReplaceAll(template, "0", "+")

	fmt.Println(GREEN + "\n" + delim + "\n" +
		YELL + "\n" + r.Text + "\n" +
		GREEN + "\n" + delim + RES)
	if r.Note != nil {
		pprintNote(r.Note)
	}
}

func pprintNote(n *pb.Note) {
	w, _ := screen.Size()
	template := fmt.Sprintf("%0*d", w, 0)
	delim := strings.ReplaceAll(template, "0", "-")

	//C like ternary: status = n.Status == DONE ? "DONE" : "TODO"
	status := (map[bool]string{
		false: YELL + "TODO", true: RED + "DONE"})[n.Status == DONE]

	//Time converting 
	created := ((n.WasCreated.AsTime()).Local()).Format(time.Stamp)
	updated := ((n.LastUpdated.AsTime()).Local()).Format(time.Stamp)

	fmt.Printf(
			PINK+delim+"\n"+
			GREEN+"ID: "+CYAN+"%v\n"+
			GREEN+"Title: "+CYAN+"%s \n"+
			GREEN+"Was created: "+CYAN+"%s "+
			GREEN+"| Last updated: "+CYAN+"%s\n"+
			PINK+delim+"\n"+
			GREEN+"Status: %s\n"+
			CYAN+"%s"+RES+
			PINK+delim+"\n"+RES,
			n.Id, n.Title, created, updated, status, n.Description)
}
