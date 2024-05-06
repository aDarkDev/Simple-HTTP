package main

import (
	"bytes"
	"fmt"
	"net"
)

func splitHost(buf *bytes.Buffer) string {
	ByteHost := bytes.Split(buf.Bytes(), []byte("Host: "))[1]
	ByteHost = bytes.Split(ByteHost, []byte("\r\n"))[0]
	var Host string
	if bytes.Contains(ByteHost, []byte(":")) {
		Host = string(ByteHost)
	} else {
		Host = string(ByteHost) + ":80"
	}

	return Host
}

func sendTcp(buf *bytes.Buffer) []byte {
	Host := splitHost(buf)
	conn, err := net.Dial("tcp", Host)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return nil
	}
	defer conn.Close()

	// Send the data from the buffer to the server
	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("Error sending message:", err)
		return nil
	}

	// Read the response from the server
	var response bytes.Buffer
	data := make([]byte, 1024)
	for {
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("Error reading response:", err)
			break
		}
		response.Write(data[:n])

		// Check if all data has been read
		if n < len(data) {
			break
		}
	}

	return response.Bytes()
}

func connectHandler(conn net.Conn, buf *bytes.Buffer) {
	conn.Write([]byte(
		"HTTP/1.1 200 Connection Established\r\nProxy-agent: SIMPLEHTTP/0.1\r\n\r\n",
	))
}

func RequestHandler(conn net.Conn, buf *bytes.Buffer) {
	// close connection
	defer conn.Close()

	// split method from buffer
	methodBytes := bytes.Split(buf.Bytes(), []byte(" "))[0]
	method := string(methodBytes)

	if method == "CONNECT" {
		connectHandler(conn, buf)
		return
	}

	result := sendTcp(buf)
	if result == nil{
		conn.Close()
		return
	}
	
	fmt.Println("Response: "+string(result))
	conn.Write(result)
}