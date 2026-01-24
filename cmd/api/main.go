package main

import (
	"fmt"
	
	"github.com/lmcosta10/wallet_service/internal/db"
	"github.com/lmcosta10/wallet_service/internal/server"
)

func main() {
	pool, err := db.InitializeConnectionToDB()
	if err != nil {
		fmt.Println("DB connection failed: %v", err)
	}
	defer pool.Close()
	
	server.InitializeServer(pool)
}
