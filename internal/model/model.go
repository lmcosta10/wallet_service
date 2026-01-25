package model

type User struct {
	Id       int
	Username string `json:"username" binding:"required,min=3,max=15"`
}
