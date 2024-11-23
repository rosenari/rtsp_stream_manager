package main

import (
	"fmt"
	"net/http"
	"websocket/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HandleRequest)

	port := "8080"
	fmt.Printf("Server running on http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
