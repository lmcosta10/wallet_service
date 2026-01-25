package server

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lmcosta10/wallet_service/internal/handler"
	"github.com/lmcosta10/wallet_service/internal/service"
)

func InitializeServer(db *pgxpool.Pool) {
	r := gin.Default()

	root, _ := os.Getwd()
	faviconPath := filepath.Join(root, "resources", "go_logo.png")

	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(userService)

	r.StaticFile("/favicon.ico", faviconPath)
	r.GET("/", handler.InitialPage)
	r.POST("/users", handler.PostUser) // change: userHandler here
	r.GET("/users/:id", userHandler.GetUserById)

	r.Run(":7878")
}
