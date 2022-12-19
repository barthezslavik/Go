package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		// Call the second microservice to get the result
		result, err := getResult()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return the result to the client
		fmt.Fprintf(w, result)
	})

	http.ListenAndServe(":8080", nil)
}

func getResult() (string, error) {
	// Call the second microservice and get the result
	resp, err := http.Get("http://localhost:8081/result")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the result from the response body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	result := buf.String()

	return result, nil
}
