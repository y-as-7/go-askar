package Middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is a custom logging middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		// Color codes for terminal output
		var statusColor string
		switch {
		case statusCode >= 200 && statusCode < 300:
			statusColor = "\033[32m" // Green
		case statusCode >= 300 && statusCode < 400:
			statusColor = "\033[36m" // Cyan
		case statusCode >= 400 && statusCode < 500:
			statusColor = "\033[33m" // Yellow
		default:
			statusColor = "\033[31m" // Red
		}
		reset := "\033[0m"

		fmt.Printf("[ASKAR] %s %3d %s| %13v | %s %s\n",
			statusColor, statusCode, reset,
			duration,
			method,
			path,
		)
	}
}
