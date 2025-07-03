package main

import (
	"learng/httpgordle/internal/handlers"
	"net/http"
)

func main() {
	// Start the server.
	err := http.ListenAndServe(":8080", handlers.Mux())
	if err != nil {
		panic(err)
	}
}
