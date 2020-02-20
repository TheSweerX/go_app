package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"syscall"

	"github.com/fabiokung/shm"
)

// only needed below for sample processing

func main() {

	expected := ([]byte)("a test")

	fmt.Println("Launching server...")

	file, err := shm.Open("myShm", os.O_RDWR|os.O_CREATE, 0600)
	defer file.Close()
	defer shm.Unlink(file.Name())
	if err != nil {
		panic(err)
	}

	if err := syscall.Ftruncate(
		int(file.Fd()), int64(len(expected)),
	); err != nil {
		panic(err)
	}

	if _, err := file.Write(expected); err != nil {
		panic(err)
	}

	if err := file.Sync(); err != nil {
		panic(err)
	}

	fmt.Println(file)

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// will listen for message to process ending in newline (\n)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	// output message received
	fmt.Println("Message Received:", string(message))
	// sample process for string received
	newmessage := strings.ToUpper(message)
	// send new string back to client
	conn.Write([]byte(newmessage + "\n"))

	conn.Close()
}
