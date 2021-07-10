package server

import (
	pb "github.com/xasai/todogo/internal/protobuf"
)



func HandleGetRequest(r *pb.Request, s pb.TodoServ_TodoRequestServer) {

	id := int(r.GetGet().Id)

	if id == 0 {
		id++;
		for ;id <= len(Tickets);id++ {
			s.Send(&pb.Response{Body: &Tickets[id]})
		}
	} else {
		s.Send(&pb.Response{Body: &Tickets[id]})
	}
}
func HandlePutRequest(r *pb.Request) {
	tick := r.GetPut()

	if tick.Id == 0 {
		Tickets = append(Tickets, *tick)
	} else {
		if tick.Title != "" {
			Tickets[tick.Id].Title = tick.Title
		}
		if tick.Body != "" {
			Tickets[tick.Id].Body = tick.Body
		}
		if tick.State == DONE {
			Tickets[tick.Id].State = DONE
		}
	}
}

func HandleDelRequest(r *pb.Request) {
	idx := r.GetDel().Id
	Tickets = append(Tickets[:idx], Tickets[idx + 1:]...)
}