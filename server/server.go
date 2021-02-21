package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Initializing the server")
	server, err := net.Listen("tcp", "127.0.0.1:54321")
	if err != nil {
		fmt.Print("Server could not start: ", err.Error())
		os.Exit(1)
	}

	for {
		x, err := server.Accept()
		if err != nil {
			fmt.Println("Connection interrupted: ", err.Error())
			return
		}
		fmt.Print("Client connected to the server")
		connectionHandler(x)

	}

}

func connectionHandler(conn net.Conn) {

	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		fmt.Println("Client disconected: ", err.Error())
		conn.Close()
		return
	}
	fmt.Println("Message received from the client")
	conn.Write(buffer)
	connectionHandler(conn)

}
