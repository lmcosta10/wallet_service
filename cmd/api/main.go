package main

import (
	"fmt"
	"net/http"

	"github.com/lmcosta10/wallet_service/internal/server"
)

func main() {
	addr := ":7878"

	server.InitializeRouter()

	err := http.ListenAndServe(addr, nil)
    if err != nil { // in case the server fails to start
        fmt.Println("Error starting server")
    }
}
