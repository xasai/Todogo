package client

import (
	"bufio"
	"errors"
	"fmt"
	. "github.com/xasai/todogo/internal/cli"
	pb "github.com/xasai/todogo/internal/protobuf"
	s "strings"
)

func sure(r *bufio.Reader) bool {
	fmt.Print(YELL + "You sure you want to do that? [y]es / [n]o: " + CYAN)
	buf, _, _ := r.ReadLine()
	str := s.ToLower(string(buf))
	if str == "yes" || str == "y" {
		return true
	}
	fmt.Println(RES)
	return false
}

func readInt32(r *bufio.Reader) (int32, error) {
	var i int

	fmt.Print(GREEN + "Input Id: " + CYAN)
	buf, _, _ := r.ReadLine()
	n, err := fmt.Sscanf(string(buf), "%d\n", &i)
	fmt.Print(RES)
	if n != 1 || err != nil {
		return 0, errors.New("Wrong input Id")
	}
	return int32(i), nil
}

func readNote(r *bufio.Reader) (pb.Note, error) {
	note := pb.Note{}
	fmt.Print(GREEN + "Title: " + CYAN)
	buf, _, _ := r.ReadLine()
	note.Title = string(buf)
	if note.Title == "" {
		fmt.Print(RED + "Empty title!\n" + CYAN)
	}
	note.Description = readDescription(r)
	fmt.Print(RES)
	return note, nil
}

func readDescription(r *bufio.Reader) string {
	fmt.Print(GREEN + "Description (Empty Line to exit/double enter):\n" + CYAN)
	var d string
	buf, _, _ := r.ReadLine()
	for ; string(buf) != ""; buf, _, _ = r.ReadLine() {
		d += string(buf) + "\n"
	}
	fmt.Print(RES)
	return d
}
