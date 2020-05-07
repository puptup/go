package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	onErrPanic(err)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		onErrPanic(err)
		if len(text) == 6 {
			if text[:4] == "exit" {
				os.Exit(2)
			}
		}

		fmt.Fprintf(conn, text+"\n")
		message, err := bufio.NewReader(conn).ReadString('\n')
		onErrPanic(err)
		fmt.Print("Message from server: " + message)
	}
}

func onErrPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}
