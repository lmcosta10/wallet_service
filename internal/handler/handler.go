package handler

import (
	"fmt"
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Currently using this slice to simulate users
var users []User = []User{{Id: 0, Username: "Alex"}, {Id: 1, Username: "Bruno"}, {Id: 2, Username: "Carlos"}}

type User struct {
	Id int
	Username string `json:"username" binding:"required,min=3,max=15"`
}

type UserHandler struct {
	DB *pgxpool.Pool
}

func NewUserHandler(db *pgxpool.Pool) *UserHandler {
	return &UserHandler{DB: db}
}

func InitialPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
	fmt.Println("OK")
}

func PostUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
        return
    }
	users = append(users, newUser)

	c.IndentedJSON(http.StatusCreated, newUser)

	fmt.Println("Posted user")
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	var user User;
	
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	
	user.Id = id;
	
	err = h.DB.QueryRow(
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

