package Http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/app/DTO"
)

// Success returns a standard success response
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error returns a standard error response
func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, DTO.Response{
		Success: false,
		Error:   message,
	})
}

// ValidationError returns a validation error response
func ValidationError(c *gin.Context, errors interface{}) {
	c.JSON(http.StatusBadRequest, DTO.Response{
		Success: false,
		Error:   "Validation failed",
		Data:    errors,
	})
}

// NotFound returns a 404 response
func NotFound(c *gin.Context, message string) {
	if message == "" {
		message = "Resource not found"
	}
	c.JSON(http.StatusNotFound, DTO.Response{
		Success: false,
		Error:   message,
	})
}

// Unauthorized returns a 401 response
func Unauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "Unauthorized"
	}
	c.JSON(http.StatusUnauthorized, DTO.Response{
		Success: false,
		Error:   message,
	})
}

// Forbidden returns a 403 response
func Forbidden(c *gin.Context, message string) {
	if message == "" {
		message = "Forbidden"
	}
	c.JSON(http.StatusForbidden, DTO.Response{
		Success: false,
		Error:   message,
	})
}
