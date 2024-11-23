package handlers

import (
	"fmt"
	"net/http"
)

// HandleRequest는 모든 요청에 대해 동일한 응답을 반환합니다.
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// 요청 메서드와 URL 출력 (로깅)
	fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)

	// 클라이언트에 응답
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, this is a simple HTTP server!"))
}
