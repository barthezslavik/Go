package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Listen for incoming connections on port 8000
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Listening on port 8000...")

	// Accept incoming connections and handle them in a new goroutine
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Use a scanner to read the incoming data line by line
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// Echo the data back to the client
		fmt.Fprintln(conn, scanner.Text())
	}

	// Close the connection when the loop is done
	conn.Close()
}
