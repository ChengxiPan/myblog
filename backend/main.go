package main

import (
	"chris.com/common"
	"chris.com/config"
	"chris.com/controller"
	"chris.com/middleware"
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

	// Public routes that do not require authentication
	r.POST("/api/auth/register", controller.UserRegister)
	r.POST("/api/auth/login", controller.UserLogin)

	// Protected route group
	protected := r.Group("/api/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		// Add routes that require authentication
		protected.GET("/profile", func(c *gin.Context) {
			// Get claims from the context
			claims, _ := c.Get("claims")
			c.JSON(200, gin.H{"message": "This is a protected endpoint", "claims": claims})
		})
		// You can add more protected routes here
	}

	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port) // Listen on the specified port
}
