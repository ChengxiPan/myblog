package main

import "github.com/gin-gonic/gin"

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

func main() {
	ctx := gin.Default()
	ctx.POST("/register", register)
}

func register() {

}
