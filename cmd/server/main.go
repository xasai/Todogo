package main

import (
	"fmt"
	"time"
	"todogo/cmd/server/internal/server"
)

func main() {
	server.Server();
	run()

}

func run() {
	fmt.Print("Lauching server server")
	for i:=0;i<3;i++ {
		time.Sleep(time.Second / 2)
		fmt.Print(".")
	}
	fmt.Println("\t\t\nStarted")
	for {
		time.Sleep(1 * time.Hour)
	}
}
