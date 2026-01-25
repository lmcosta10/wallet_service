package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lmcosta10/wallet_service/internal/model"
)

type UserService struct {
	DB *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) FetchUserByID(c *gin.Context, id int) {
	var user model.User

	user.Id = id

	err := s.DB.QueryRow(
		c.Request.Context(),
		`SELECT username FROM users WHERE id = $1`,
		id,
	).Scan(&user.Username)

	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
