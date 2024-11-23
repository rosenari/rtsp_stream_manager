package main

import (
	"fmt"
	"net/http"
	"rtsp_stream_manager/handlers"
)

func main() {
	// 핸들러 등록
	http.HandleFunc("/", handlers.HandleRequest)

	// 서버 시작
	port := "8080"
	fmt.Printf("Server running on http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
