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

type WalletHandler struct {
	WalletService *service.WalletService
}

type TransactionData struct {
	FromWalletID int
	ToWalletID   int
	Amount       float64
}

// ====================
// User handlers
// ====================

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

// ====================
// Wallet handlers
// ====================

func NewWalletHandler(walletService *service.WalletService) *WalletHandler {
	return &WalletHandler{
		WalletService: walletService,
	}
}

func (h *WalletHandler) Transfer(c *gin.Context) {
	var req TransactionData

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.WalletService.Transfer(
		c.Request.Context(),
		req.FromWalletID,
		req.ToWalletID,
		req.Amount,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
