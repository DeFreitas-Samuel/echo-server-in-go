package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Connecting to the server")
	conn, err := net.Dial("tcp", "127.0.0.1:54321")
	if err != nil {
		fmt.Print("Connection failed: ", err.Error())
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		conn.Write([]byte(input))
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(message)

	}

}
