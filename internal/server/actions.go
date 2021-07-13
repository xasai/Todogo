package server

import (
	"context"
	pb "github.com/xasai/todogo/internal/protobuf"
	ts "google.golang.org/protobuf/types/known/timestamppb"
	_ "log" //TODO USE ME
	"strconv"
)

var NextId = int32(0)

func (s *todoServer) CreateNewNote(x context.Context, note *pb.Note) (*pb.Response, error) {
	note.Id = NextId
	NextId++
	note.WasCreated = ts.Now()
	note.LastUpdated = note.WasCreated
	note.Status = TODO
	Notes[note.Id] = note
	return &pb.Response{
		Text: "New note created and has id " + strconv.Itoa(int(note.Id)),
		Note: note}, nil
}

func (s *todoServer) GetNoteById(x context.Context, req *pb.RequestId) (*pb.Response, error) {
	id := req.NoteId
	if _, key := Notes[id]; !key {
		return &pb.Response{Text: "Note doesn't exist", Note: nil}, nil
	}
	return &pb.Response{Text: "Note found",
		Note: Notes[id]}, nil
}

func (s *todoServer) GetAllNotes(req *pb.RequestId, stream pb.TodoService_GetAllNotesServer) error {
	for _, n := range Notes {
		if err := stream.Send(n); err != nil {
			return err
		}
	}
	return nil
}

func (s *todoServer) ChangeNoteContent(x context.Context, req *pb.Note) (*pb.Response, error) {
	id := req.Id
	if _, key := Notes[id]; !key {
		return &pb.Response{Text: "Note doesn't exist", Note: nil}, nil
	}
	Notes[id].Title = req.Title
	Notes[id].Description = req.Description
	Notes[id].LastUpdated = ts.Now()
	return &pb.Response{
		Text: "Note " + strconv.Itoa(int(id)) + " successfully changed",
		Note: Notes[id]}, nil
}

func (s *todoServer) DelNoteById(x context.Context, req *pb.RequestId) (*pb.Response, error) {
	id := req.NoteId
	if _, key := Notes[id]; !key {
		return &pb.Response{Text: "Note doesn't exist", Note: nil}, nil
	}
	delete(Notes, id)
	return &pb.Response{
		Text: "Note " + strconv.Itoa(int(id)) + " successfully deleted",
		Note: nil}, nil
}

func (s *todoServer) MarkDoneById(x context.Context, req *pb.RequestId) (*pb.Response, error) {
	id := req.NoteId
	if _, key := Notes[id]; !key {
		return &pb.Response{Text: "Note doesn't exist", Note: nil}, nil
	}
	if Notes[id].Status == DONE {
		return &pb.Response{
			Text: "Note <" + Notes[id].Title + "> is already done",
			Note: nil}, nil
	}
	Notes[id].Status = DONE
	Notes[id].LastUpdated = ts.Now()
	return &pb.Response{
		Text: "Grats! Note <" + Notes[id].Title + "> is DONE",
		Note: nil}, nil
}
