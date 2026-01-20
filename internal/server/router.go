package server

import (
	"net/http"

	"github.com/lmcosta10/wallet_service/internal/handler"
)

func InitializeRouter() {
	http.Handle("/", http.HandlerFunc(handler.InitialHandler))
}
