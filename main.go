package main

import (
	"fmt"
	"net"
)

func main() {
	// Define the host and port to listen on
	HOST := "127.0.0.1"
	PORT := "12345"

	// Create a listener
	listener, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Listening on", HOST+":"+PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}

		fmt.Println("Connected by", conn.RemoteAddr())
		go ConnectionHandler(conn)
	}
}
