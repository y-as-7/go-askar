package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/app/DTO"
)

type OrderController struct{}

// Index returns a list of resources
func (ctrl *OrderController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Index method for Order",
	})
}

// Show returns a single resource
func (ctrl *OrderController) Show(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Show method for Order with ID " + id,
	})
}

// Store creates a new resource
func (ctrl *OrderController) Store(c *gin.Context) {
	c.JSON(http.StatusCreated, DTO.Response{
		Success: true,
		Message: "Store method for Order",
	})
}

// Update updates an existing resource
func (ctrl *OrderController) Update(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Update method for Order with ID " + id,
	})
}

// Destroy deletes a resource
func (ctrl *OrderController) Destroy(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Destroy method for Order with ID " + id,
	})
}
