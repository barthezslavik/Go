package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		// Perform the task and get the result
		result := doTask()

		// Return the result to the client
		fmt.Fprintf(w, result)
	})

	http.ListenAndServe(":8081", nil)
}

func doTask() string {
	// Perform the task and return the result as a string
	return "Task complete"
}
