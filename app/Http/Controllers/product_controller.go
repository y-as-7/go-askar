package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/app/DTO"
)

type ProductController struct{}

// Index returns a list of resources
func (ctrl *ProductController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Index method for Product",
	})
}

// Show returns a single resource
func (ctrl *ProductController) Show(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Show method for Product with ID " + id,
	})
}

// Store creates a new resource
func (ctrl *ProductController) Store(c *gin.Context) {
	c.JSON(http.StatusCreated, DTO.Response{
		Success: true,
		Message: "Store method for Product",
	})
}

// Update updates an existing resource
func (ctrl *ProductController) Update(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Update method for Product with ID " + id,
	})
}

// Destroy deletes a resource
func (ctrl *ProductController) Destroy(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Destroy method for Product with ID " + id,
	})
}
