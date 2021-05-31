package main

import (
	"fmt"

	service "github.com/davidgaspardev/go_microservices/src/service"
)

func main() {
	var testService service.Service
	// Printing test service
	fmt.Printf("%+v\n", testService)
	// Testing service
	testService.Run()
}
