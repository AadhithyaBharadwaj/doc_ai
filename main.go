package main

import (
	handler "doc_ai/api"
	"net/http"
)

func main() {
	// Handle requests using the Handler function from your handler package
	http.HandleFunc("/", handler.Handler)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
