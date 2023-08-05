package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func respond(connection net.Conn) {
	request_bytes := make([]byte, 16)
	_, err := connection.Read(request_bytes)

	if err != nil {
		fmt.Println("Error reading connection: ", err.Error())
		os.Exit(1)
	}

	request_string := strings.Trim(string(request_bytes), "\x00")

	if len(request_string) > 0 {
		fmt.Println("Received PING")
		connection.Write([]byte("+PONG\r\n"))
		fmt.Println("Sent PONG")
	}

	connection.Close()
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6379")

	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		respond(connection)
	}
}
