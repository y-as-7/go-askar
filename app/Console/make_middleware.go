package console

import (
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	Register(Command{
		Name:        "make:middleware",
		Description: "Create a new middleware file",
		Execute:     handleMakeMiddleware,
	})
}

func handleMakeMiddleware(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("middleware name is required. Usage: go askar make:middleware <MiddlewareName>")
	}

	name := args[0]
	name = capitalizeFirst(name)

	// Create middleware file
	if err := createMiddlewareFile(name); err != nil {
		return err
	}

	fmt.Printf("\n✅ Middleware created successfully!\n")
	fmt.Printf("   📄 app/Http/Middleware/%s.go\n\n", toSnakeCase(name))

	return nil
}

func createMiddlewareFile(name string) error {
	middlewareDir := "app/Http/Middleware"
	fileName := filepath.Join(middlewareDir, toSnakeCase(name)+".go")

	// Check if file already exists
	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("middleware file already exists: %s", fileName)
	}

	// Ensure directory exists
	if err := os.MkdirAll(middlewareDir, 0755); err != nil {
		return fmt.Errorf("failed to create middleware directory: %w", err)
	}

	template := fmt.Sprintf(`package Middleware

import (
	"github.com/gin-gonic/gin"
)

// %s middleware
func %s() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic before request
		
		c.Next()
		
		// Logic after request
	}
}
`, name, name)

	if err := os.WriteFile(fileName, []byte(template), 0644); err != nil {
		return fmt.Errorf("failed to create middleware file: %w", err)
	}

	return nil
}
