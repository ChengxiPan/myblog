package controller

import (
	"net/http"

	"chris.com/models"
	"chris.com/util"
	"github.com/gin-gonic/gin"
)

// Method: UserRegister
// Params: c *gin.Context, bindjson
// Restrictions:
//   - Post request
//   - email must in correct form
//   - password must be at least 6 characters
func UserRegister(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	// check password
	if len(user.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 6 characters"})
		return
	}
	//check if Username exists
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username is required"})
		return
	}
	// check email
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email is required"})
		return
	}
	// check if email is in correct form
	if !util.IsValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email is not valid"})
		return
	}

	//successfully registered
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Successfully registered",
	})
}
