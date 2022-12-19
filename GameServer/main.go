package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
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
		// Split the input on spaces and store it in a slice
		input := strings.Split(scanner.Text(), " ")

		// Handle the command based on the first word in the input
		switch input[0] {
		case "move":
			// Handle the "move" command
			handleMove(input[1:], conn)
		case "attack":
			// Handle the "attack" command
			handleAttack(input[1:], conn)
		case "quit":
			// Handle the "quit" command
			handleQuit(conn)
			return
		default:
			// Send an error message for unknown commands
			fmt.Fprintln(conn, "Error: unknown command")
		}
	}

	// Close the connection when the loop is done
	conn.Close()
}

func handleMove(input []string, conn net.Conn) {
	// Validate the input and move the player
	if len(input) != 2 {
		fmt.Fprintln(conn, "Error: invalid input")
		return
	}

	x, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Fprintln(conn, "Error: invalid x coordinate")
		return
	}

	y, err := strconv.Atoi(input[1])
	if err != nil {
		fmt.Fprintln(conn, "Error: invalid y coordinate")
		return
	}

	// Move the player to the specified coordinates
	fmt.Fprintln(conn, "Player moved to", x, y)
}

func handleAttack(input []string, conn net.Conn) {
	// Validate the input and attack the target
	if len(input) != 1 {
		fmt.Fprintln(conn, "Error: invalid input")
		return
	}
	target := input[0]

	// Attack the target
	fmt.Fprintln(conn, "Player attacked", target)
}

func handleQuit(conn net.Conn) {
	// Disconnect the player from the server
	fmt.Fprintln(conn, "Goodbye!")
}
