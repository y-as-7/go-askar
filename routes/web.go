package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterWebRoutes registers all web routes
func RegisterWebRoutes(router *gin.Engine) {
	// Web routes (HTML pages)
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "pages/home", gin.H{
			"title": "Home",
		})
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(200, "pages/about", gin.H{
			"title": "About Us",
		})
	})
}
