package server

import (
	"context"
	pb "github.com/xasai/todogo/internal/protobuf"
	ts "google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

//Assign id of created note to this value
var NextId = int32(0)

//CreateNewNote gets *pb.Note, sets its WasCreated LastUpdated and Id fields.
//Responses with note Id message and Note.
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

//GetNoteById gets *pb.RequestId and checks if note with such id exists.
//If so, it returns this note. Otherwise it returns explicit message and note = nil.
func (s *todoServer) GetNoteById(x context.Context, req *pb.RequestId) (*pb.Response, error) {
	id := req.NoteId
	if _, key := Notes[id]; !key {
		return &pb.Response{Text: "Note doesn't exist", Note: nil}, nil
	}
	return &pb.Response{Text: "Note found",
		Note: Notes[id]}, nil
}

//GetAllNotes gets stub *pb.RequestId and server's side stream.
//All Notes found in Notes map sends to stream in not sorted sequence
func (s *todoServer) GetAllNotes(req *pb.RequestId, stream pb.TodoService_GetAllNotesServer) error {
	for _, n := range Notes {
		if err := stream.Send(n); err != nil {
			return err
		}
	}
	return nil
}

//ChangeNoteContent gets *pb.Note with changed Title and Description field.
//Changing happens by pb.Note.Id field and if note with such Id doesn't exists
//func returns corresponding message and note = nil.
//If note with such Id exists, it assign it's Title, Description and LastUpdated
//field to new values, and returns changed note
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

//DelNoteById gets *pb.RequestId.
//If no such Id in map Notes, function returns explicit message and nil.
//Otherwise, it deletes this note, and returns other message and nil
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

//MarkDoneById gets *pb.RequestId.
//If no such Id in map Notes, returns explicit message and nil.
//If there is such Note but it's already DONE, return another message and nil.
//If there is such Note and it's not done yet, change its Status and LastUpdated fields.
//Then returns Congratulation message and nil
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
