package console

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	Register(Command{
		Name:        "make:request",
		Description: "Create a new request validation file",
		Execute:     handleMakeRequest,
	})
}

func handleMakeRequest(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("request name is required. Usage: go askar make:request <RequestName>")
	}

	name := args[0]
	// Ensure it ends with Request
	if !strings.HasSuffix(strings.ToLower(name), "request") {
		name = capitalizeFirst(name) + "Request"
	} else {
		name = capitalizeFirst(name)
	}

	// Create request file
	if err := createRequestFile(name); err != nil {
		return err
	}

	fmt.Printf("\n✅ Request created successfully!\n")
	fmt.Printf("   📄 app/Http/Requests/%s.go\n\n", toSnakeCase(name))

	return nil
}

func createRequestFile(name string) error {
	requestsDir := "app/Http/Requests"
	fileName := filepath.Join(requestsDir, toSnakeCase(name)+".go")

	// Check if file already exists
	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("request file already exists: %s", fileName)
	}

	// Ensure directory exists
	if err := os.MkdirAll(requestsDir, 0755); err != nil {
		return fmt.Errorf("failed to create requests directory: %w", err)
	}

	template := fmt.Sprintf(`package Requests

type %s struct {
	// Add your fields here with validation tags
	// Example:
	// Name  string `+"`json:\"name\" binding:\"required\"`"+`
	// Email string `+"`json:\"email\" binding:\"required,email\"`"+`
}
`, name)

	if err := os.WriteFile(fileName, []byte(template), 0644); err != nil {
		return fmt.Errorf("failed to create request file: %w", err)
	}

	return nil
}
