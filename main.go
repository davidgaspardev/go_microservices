package main

import (
    "fmt"
    "net/http"
)

func main() {
    const message = "Hello, I'm using Golang";

    // Creating a server mux
    mux := http.NewServeMux();
    mux.HandleFunc("/", func(w http.ResponseWriter, r * http.Request) {
        w.Header().Set("Content-Type", "text/plain; charset=utf8");
      	w.WriteHeader(http.StatusOK);
 	// Conveting message to byte array
        w.Write([]byte(message));
    });

    err := http.ListenAndServe(":8080", mux);
    if err != nil {
	log.Fatalf("Server failed to start: %v", err);
    }
}
