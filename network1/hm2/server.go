package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", ":8081")
	onErrPanic(err)

	conn, err := ln.Accept()
	onErrPanic(err)

	for {

		message, err := bufio.NewReader(conn).ReadString('\n')
		onErrPanic(err)

		fmt.Print("Message Received: ", message)
		message = message[:len(message)-2]
		newmessage := ""
		i, err := strconv.Atoi(message)
		if err != nil {
			newmessage = strings.ToUpper(message)
		} else {
			newmessage = strconv.Itoa(i * 2)
		}

		conn.Write([]byte(newmessage + "\n"))
	}
}

func onErrPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}
