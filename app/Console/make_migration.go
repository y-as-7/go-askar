package console

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func init() {
	Register(Command{
		Name:        "make:migration",
		Description: "Create a new migration file",
		Execute:     handleMakeMigration,
	})
}

func handleMakeMigration(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("migration name is required. Usage: go askar make:migration <migration_name>")
	}

	migrationName := args[0]

	// Convert to snake_case if needed
	migrationName = toSnakeCase(migrationName)

	// Create migration file
	fileName, err := createMigrationFile(migrationName)
	if err != nil {
		return err
	}

	fmt.Printf("\n✅ Migration created successfully!\n")
	fmt.Printf("   📄 %s\n\n", fileName)

	return nil
}

func createMigrationFile(migrationName string) (string, error) {
	migrationsDir := "database/migrations"
	
	// Generate timestamp prefix (YYYYMMDDHHMMSS format)
	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("%s_%s.go", timestamp, migrationName)
	filePath := filepath.Join(migrationsDir, fileName)

	// Ensure directory exists
	if err := os.MkdirAll(migrationsDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create migrations directory: %w", err)
	}

	// Convert snake_case to PascalCase for struct name
	structName := toPascalCase(migrationName)

	template := fmt.Sprintf(`package migrations

import (
	"gorm.io/gorm"
)

// %s migration
type %s struct{}

// Up runs the migration
func (m *%s) Up(db *gorm.DB) error {
	// Add your migration logic here
	// Example:
	// return db.Exec("CREATE TABLE example (id INT PRIMARY KEY)").Error
	return nil
}

// Down reverses the migration
func (m *%s) Down(db *gorm.DB) error {
	// Add your rollback logic here
	// Example:
	// return db.Exec("DROP TABLE example").Error
	return nil
}
`, structName, structName, structName, structName)

	if err := os.WriteFile(filePath, []byte(template), 0644); err != nil {
		return "", fmt.Errorf("failed to create migration file: %w", err)
	}

	return filePath, nil
}


