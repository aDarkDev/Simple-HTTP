package main

import (
	"bytes"
	"fmt"
	"net"
)

func ConnectionHandler(conn net.Conn) {
	// Create a buffer to store the received data
	var buf bytes.Buffer

	for {
		data := make([]byte, 4096)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}

		// Append the received data to the buffer
		buf.Write(data[:n])

		// // Check if the data received is less than the buffer size
		if n < len(data) {
			break // Exit the loop if all data has been read
		}
	}
	fmt.Println("\nRequest: ",buf.String())
	RequestHandler(conn,&buf)
}
