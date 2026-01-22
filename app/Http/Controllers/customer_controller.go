package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/app/DTO"
)

type CustomerController struct{}

// Index returns a list of resources
func (ctrl *CustomerController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Index method for Customer",
	})
}

// Show returns a single resource
func (ctrl *CustomerController) Show(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Show method for Customer with ID " + id,
	})
}

// Store creates a new resource
func (ctrl *CustomerController) Store(c *gin.Context) {
	c.JSON(http.StatusCreated, DTO.Response{
		Success: true,
		Message: "Store method for Customer",
	})
}

// Update updates an existing resource
func (ctrl *CustomerController) Update(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Update method for Customer with ID " + id,
	})
}

// Destroy deletes a resource
func (ctrl *CustomerController) Destroy(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Destroy method for Customer with ID " + id,
	})
}
