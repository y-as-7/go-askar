package console

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	Register(Command{
		Name:        "make:controller",
		Description: "Create a new controller file",
		Execute:     handleMakeController,
	})
}

func handleMakeController(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("controller name is required. Usage: go askar make:controller <ControllerName>")
	}

	name := args[0]
	// Ensure it ends with Controller
	if !strings.HasSuffix(strings.ToLower(name), "controller") {
		name = capitalizeFirst(name) + "Controller"
	} else {
		name = capitalizeFirst(name)
	}

	// Create controller file
	if err := createControllerFile(name); err != nil {
		return err
	}

	fmt.Printf("\n✅ Controller created successfully!\n")
	fmt.Printf("   📄 app/Http/Controllers/%s.go\n\n", toSnakeCase(name))

	return nil
}

func createControllerFile(name string) error {
	controllersDir := "app/Http/Controllers"
	fileName := filepath.Join(controllersDir, toSnakeCase(name)+".go")

	// Check if file already exists
	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("controller file already exists: %s", fileName)
	}

	// Ensure directory exists
	if err := os.MkdirAll(controllersDir, 0755); err != nil {
		return fmt.Errorf("failed to create controllers directory: %w", err)
	}

	// Base name without Controller suffix for model/service references
	baseName := strings.TrimSuffix(name, "Controller")

	template := fmt.Sprintf(`package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-as-7/go-askar/app/DTO"
)

type %s struct{}

// Index returns a list of resources
func (ctrl *%s) Index(c *gin.Context) {
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Index method for %s",
	})
}

// Show returns a single resource
func (ctrl *%s) Show(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Show method for %s with ID " + id,
	})
}

// Store creates a new resource
func (ctrl *%s) Store(c *gin.Context) {
	c.JSON(http.StatusCreated, DTO.Response{
		Success: true,
		Message: "Store method for %s",
	})
}

// Update updates an existing resource
func (ctrl *%s) Update(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Update method for %s with ID " + id,
	})
}

// Destroy deletes a resource
func (ctrl *%s) Destroy(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, DTO.Response{
		Success: true,
		Message: "Destroy method for %s with ID " + id,
	})
}
`, name, name, baseName, name, baseName, name, baseName, name, baseName, name, baseName)

	if err := os.WriteFile(fileName, []byte(template), 0644); err != nil {
		return fmt.Errorf("failed to create controller file: %w", err)
	}

	return nil
}
