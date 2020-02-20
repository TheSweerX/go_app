package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/fabiokung/shm"
)

// only needed below for sample processing

func main() {

	fmt.Println("Launching server...")

	file, err := shm.Open("myShm", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	fmt.Print(file)

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// will listen for message to process ending in newline (\n)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	// output message received
	fmt.Print("Message Received:", string(message))
	// sample process for string received
	newmessage := strings.ToUpper(message)
	// send new string back to client
	conn.Write([]byte(newmessage + "\n"))

	conn.Close()
	defer file.Close()
	shm.Unlink(file.Name())
}
