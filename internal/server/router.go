package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lmcosta10/wallet_service/internal/handler"
)

func InitializeServer() {
	r := gin.Default()
	
	r.GET("/", handler.InitialPage)
    r.POST("/users", handler.PostUser)
	r.GET("/users/:id", handler.GetUserById)


	r.Run("localhost:7878")
}
