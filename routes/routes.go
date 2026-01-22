package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/app/controllers"
	"github.com/y-as-7/go-askar/app/dto"
	"github.com/y-as-7/go-askar/app/middleware"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine) {
	// Initialize controllers
	authCtrl := &controllers.AuthController{}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, dto.Response{
			Success: true,
			Message: "API is running",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Public routes - Auth
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authCtrl.Register)
			auth.POST("/login", authCtrl.Login)
		}

		// Protected routes - require authentication
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// Auth profile
			protected.GET("/auth/profile", authCtrl.GetProfile)
		}
	}

	// 404 handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Error:   "Route not found",
		})
	})
}
