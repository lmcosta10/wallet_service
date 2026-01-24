package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lmcosta10/wallet_service/internal/handler"
)

func InitializeServer(db *pgxpool.Pool) {
	r := gin.Default()
	
	userHandler := handler.NewUserHandler(db)
	
	r.GET("/", handler.InitialPage)
	r.POST("/users", handler.PostUser) // change: userHandler here
	r.GET("/users/:id", userHandler.GetUserById)


	r.Run(":7878")
}
