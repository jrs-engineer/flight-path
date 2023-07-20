package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/C001-developer/flight-path/src/handler"
)

func main() {
	// Register the handler function for "/calculate" endpoint
	http.HandleFunc("/calculate", handler.FlightPathHandler)

	// Start the HTTP server
	port := 8080 // Change this to the desired port number
	fmt.Printf("Server listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
