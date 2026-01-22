package Middleware

import (
	"github.com/gin-gonic/gin"
)

// CheckRole middleware
func CheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic before request
		
		c.Next()
		
		// Logic after request
	}
}
