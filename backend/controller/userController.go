package controller

import (
	"net/http"

	"chris.com/common"
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
	var db = common.GetDB()
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
	// bycrypt password
	user.Password = util.HashAndSalt(user.Password)
	// check if email exists in db
	if err := db.Where("email = ?", user.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		return
	}

	//successfully registered
	//release token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to release token"})
		return
	}
	// insert user to db
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to register. Please be sure the Database is running"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Successfully registered",
		"token":   token,
		"user":    user,
	})
}

// Method: UserLogin
// Params: c *gin.Context, bindjson
func UserLogin(c *gin.Context) {
	var user models.User
	var db = common.GetDB()
	// BindJSON
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if the user exists
	var existingUser models.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not registered"})
		return
	}

	// Check the password
	if !util.ComparePasswords(existingUser.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	// Release Token
	token, err := common.ReleaseToken(existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to release token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user":    existingUser,
	})
}
