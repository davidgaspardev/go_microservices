package service

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (service *Service) Run() {

	service.name = "Test"
	service.port = ":8080"

	mux := http.NewServeMux()

	// Handling with routes
	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "text/plain; charset=utf8")
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("Testing microservices with Golang"))

		fmt.Println("\n=> HTTP header (", time.Now().Local(), ")")
		fmt.Println("- Method:", request.Method)
		fmt.Println("- Path:", request.URL.Path)
		fmt.Println("- Query:", request.URL.RawQuery)
		fmt.Println("- ContentLength:", request.ContentLength)
		fmt.Println("- RemoteAddr:", request.RemoteAddr)
	})

	err := http.ListenAndServe(service.port, mux)
	if err != nil {
		log.Fatalf("Server dailed to start: %v", err)
	}
}
