package handler

import (
	"fmt"
	"net/http"
)

func InitialHandler(w http.ResponseWriter, req *http.Request) {
	resp := fmt.Sprintf("Request: %v", req)
	fmt.Fprintln(w, "OK") // to browser
	fmt.Fprintln(w, resp)
}
