package Middleware

import (
	"github.com/gin-gonic/gin"
)

// AuthLog middleware
func AuthLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic before request
		
		c.Next()
		
		// Logic after request
	}
}
