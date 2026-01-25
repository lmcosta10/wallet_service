package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lmcosta10/wallet_service/internal/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func InitialPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
	fmt.Println("OK")
}

func PostUser(c *gin.Context) {
	// TODO

	// c.IndentedJSON(http.StatusCreated, newUser)
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	h.UserService.FetchUserByID(c, id)
}
