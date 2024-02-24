package main

import (
	"fmt"
	"net/http"
)

// handlerFunc handles incoming HTTP requests for /ingest endpoint
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// Message here to the response writer.
	fmt.Println("handlerFunc called")
}
