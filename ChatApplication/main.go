package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Send messages to the server and print the responses
	go sendMessages(conn)
	printMessages(conn)
}

func sendMessages(conn net.Conn) {
	// Use a scanner to read messages from the terminal
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Send the message to the server and exit if it's "quit"
		text := scanner.Text()
		fmt.Fprintln(conn, text)
		if text == "quit" {
			break
		}
	}
}

func printMessages(conn net.Conn) {
	// Use a scanner to read messages from the server
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// Print the message and exit if it's "quit"
		text := scanner.Text()
		fmt.Println(text)
		if strings.TrimSpace(text) == "quit" {
			break
		}
	}
}
