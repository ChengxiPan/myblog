package middleware

import (
	"net/http"
	"strings"

	"chris.com/common"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header
		token := c.Request.Header.Get("Authorization")
		// Check token
		if token == "" || strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			c.Abort()
			return
		}

		token = token[7:]
		claims, err := common.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
