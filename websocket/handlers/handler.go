package handlers

import (
	"fmt"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("websocket server"))
}
