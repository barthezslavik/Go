package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create a channel to receive data
	in := make(chan int)

	// Start the pipeline
	go pipeline(in)

	// Send some data to the pipeline
	for i := 0; i < 10; i++ {
		in <- i
	}

	// Close the channel to signal the end of the input
	close(in)
}

func pipeline(in <-chan int) {
	// Create a channel to receive the transformed data
	out := make(chan string)

	// Start the transform stage
	go transform(in, out)

	// Start the output stage
	output(out)
}

func transform(in <-chan int, out chan<- string) {
	// Transform the data and send it to the output channel
	for num := range in {
		out <- strconv.Itoa(num)
	}

	// Close the output channel when the input channel is closed
	close(out)
}

func output(out <-chan string) {
	// Print the transformed data
	for str := range out {
		fmt.Println(str)
	}
}
