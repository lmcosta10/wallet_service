package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Currently using this slice to simulate users
var users []User = []User{{"Alex"}, {"Bruno"}, {"Carlos"}}
type User struct {
	Username string `json:"username" binding:"required,min=3,max=15"`
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

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	
	if id <= len(users) {
		user := users[id]

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}
