package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/fabiokung/shm"
)

func main() {

	expected := ([]byte)("notest")

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	fmt.Println(str)

	// connect to socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	fmt.Fprintf(conn, str+"\n")
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Message from server: " + message)

	file, err := shm.Open("myShm", os.O_RDONLY, 0600)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	buff := make([]byte, len(expected))
	if _, err := file.ReadAt(buff, 0); err != nil {
		panic(err)
	}

	fmt.Println(string(buff))

	// for i := 0; true; i++ {
	// 	time.Sleep(1)
	// }

	conn.Close()
}
