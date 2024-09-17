package main

import (
	"chris.com/common"
	"chris.com/config"
	"chris.com/controller"
	"chris.com/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	db := common.InitDB()
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Tag{})

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/api/auth/register", controller.UserRegister)
	r.POST("/api/auth/login", controller.UserLogin)
	// r.POST("/login", controller.UserLogin)

	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port) // listen to 8080
}
