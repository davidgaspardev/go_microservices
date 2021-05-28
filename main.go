package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	const message = "Hello, I'm using Golang"

	// Creating a server mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf8")
		w.WriteHeader(http.StatusOK)
		// Conveting message to byte array
		w.Write([]byte(message))

		// Loop over header names
		fmt.Println("=====> http header (", time.Now().Format("2006.01.02 15:04:05"), ")")
		for name, values := range r.Header {
			// Loop over all values for the name.
			for _, value := range values {
				fmt.Println(" - ", name, ": ", value)
			}
		}
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
