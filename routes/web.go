package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterWebRoutes registers all web routes
func RegisterWebRoutes(router *gin.Engine) {
	// Web routes (HTML pages, if you add views later)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Askar Framework",
			"version": "1.4.0",
		})
	})
}
