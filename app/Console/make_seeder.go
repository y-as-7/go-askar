package console

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	Register(Command{
		Name:        "make:seeder",
		Description: "Create a new seeder file",
		Execute:     handleMakeSeeder,
	})
}

func handleMakeSeeder(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("seeder name is required. Usage: go askar make:seeder <SeederName>")
	}

	name := args[0]
	// Ensure it ends with Seeder
	if !strings.HasSuffix(strings.ToLower(name), "seeder") {
		name = capitalizeFirst(name) + "Seeder"
	} else {
		name = capitalizeFirst(name)
	}

	// Create seeder file
	if err := createSeederFile(name); err != nil {
		return err
	}

	fmt.Printf("\n✅ Seeder created successfully!\n")
	fmt.Printf("   📄 database/seeders/%s.go\n\n", toSnakeCase(name))

	return nil
}

func createSeederFile(name string) error {
	seedersDir := "database/seeders"
	fileName := filepath.Join(seedersDir, toSnakeCase(name)+".go")

	// Check if file already exists
	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("seeder file already exists: %s", fileName)
	}

	// Ensure directory exists
	if err := os.MkdirAll(seedersDir, 0755); err != nil {
		return fmt.Errorf("failed to create seeders directory: %w", err)
	}

	template := fmt.Sprintf(`package seeders

import (
	"github.com/y-as-7/go-askar/app/Models"
	"gorm.io/gorm"
)

type %s struct{}

func (s *%s) Run(db *gorm.DB) error {
	// Add your seeding logic here
	// Example:
	// return db.Create(&Models.User{Name: "Admin", Email: "admin@example.com"}).Error
	return nil
}
`, name, name)

	if err := os.WriteFile(fileName, []byte(template), 0644); err != nil {
		return fmt.Errorf("failed to create seeder file: %w", err)
	}

	return nil
}
