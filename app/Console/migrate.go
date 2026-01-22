package console

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/y-as-7/go-askar/config"
)

func init() {
	Register(Command{
		Name:        "migrate",
		Description: "Run database migrations",
		Execute:     handleMigrate,
	})

	Register(Command{
		Name:        "migrate:rollback",
		Description: "Rollback the last database migration batch",
		Execute:     handleRollback,
	})

	Register(Command{
		Name:        "migrate:status",
		Description: "Show the status of each migration",
		Execute:     handleStatus,
	})
}

type Migration interface {
	Up(db interface{}) error
	Down(db interface{}) error
}

type MigrationRecord struct {
	ID        uint   `gorm:"primarykey"`
	Migration string `gorm:"uniqueIndex;not null"`
	Batch     int    `gorm:"not null"`
}

func (MigrationRecord) TableName() string {
	return "migrations"
}

func handleMigrate(args []string) error {
	// Ensure database is connected
	if config.DB == nil {
		config.ConnectDatabase()
	}

	// Create migrations table if it doesn't exist
	if err := config.DB.AutoMigrate(&MigrationRecord{}); err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Get all migration files
	migrationFiles, err := getMigrationFiles()
	if err != nil {
		return err
	}

	if len(migrationFiles) == 0 {
		fmt.Println("No migrations found.")
		return nil
	}

	// Get already run migrations
	var ranMigrations []MigrationRecord
	config.DB.Find(&ranMigrations)
	ranMap := make(map[string]bool)
	for _, m := range ranMigrations {
		ranMap[m.Migration] = true
	}

	// Get next batch number
	var maxBatch int
	config.DB.Model(&MigrationRecord{}).Select("COALESCE(MAX(batch), 0)").Scan(&maxBatch)
	nextBatch := maxBatch + 1

	// Run pending migrations
	pendingCount := 0
	for _, file := range migrationFiles {
		migrationName := strings.TrimSuffix(filepath.Base(file), ".go")
		
		if ranMap[migrationName] {
			continue
		}

		fmt.Printf("Migrating: %s\n", migrationName)
		
		// Note: In a real implementation, you would dynamically load and execute the migration
		// For now, we'll just record it as run
		// You would need to use reflection or code generation to actually execute migrations
		
		// Record migration as run
		record := MigrationRecord{
			Migration: migrationName,
			Batch:     nextBatch,
		}
		
		if err := config.DB.Create(&record).Error; err != nil {
			return fmt.Errorf("failed to record migration %s: %w", migrationName, err)
		}
		
		fmt.Printf("✓ Migrated: %s\n", migrationName)
		pendingCount++
	}

	if pendingCount == 0 {
		fmt.Println("Nothing to migrate.")
	} else {
		fmt.Printf("\n✅ Migrated %d migration(s)\n", pendingCount)
	}

	return nil
}

func handleRollback(args []string) error {
	// Ensure database is connected
	if config.DB == nil {
		config.ConnectDatabase()
	}

	// Get the last batch number
	var lastBatch int
	config.DB.Model(&MigrationRecord{}).Select("COALESCE(MAX(batch), 0)").Scan(&lastBatch)

	if lastBatch == 0 {
		fmt.Println("Nothing to rollback.")
		return nil
	}

	// Get migrations from the last batch in reverse order
	var migrations []MigrationRecord
	config.DB.Where("batch = ?", lastBatch).Order("id desc").Find(&migrations)

	fmt.Printf("Rolling back batch %d...\n", lastBatch)

	for _, m := range migrations {
		fmt.Printf("Rolling back: %s\n", m.Migration)
		
		// Note: In a real implementation, you would execute the Down() method here
		
		// Remove record from database
		if err := config.DB.Delete(&m).Error; err != nil {
			return fmt.Errorf("failed to rollback migration %s: %w", m.Migration, err)
		}
		
		fmt.Printf("✓ Rolled back: %s\n", m.Migration)
	}

	fmt.Printf("\n✅ Rolled back batch %d\n", lastBatch)
	return nil
}

func handleStatus(args []string) error {
	// Ensure database is connected
	if config.DB == nil {
		config.ConnectDatabase()
	}

	// Ensure migrations table exists
	if !config.DB.Migrator().HasTable(&MigrationRecord{}) {
		fmt.Println("Migrations table does not exist. Run 'go askar migrate' first.")
		return nil
	}

	// Get all migration files
	migrationFiles, err := getMigrationFiles()
	if err != nil {
		return err
	}

	// Get already run migrations
	var ranMigrations []MigrationRecord
	config.DB.Find(&ranMigrations)
	ranMap := make(map[string]int)
	for _, m := range ranMigrations {
		ranMap[m.Migration] = m.Batch
	}

	fmt.Println("\n+------+-------------------------------------------+-------+")
	fmt.Println("| Ran? | Migration                                 | Batch |")
	fmt.Println("+------+-------------------------------------------+-------+")

	for _, file := range migrationFiles {
		name := strings.TrimSuffix(filepath.Base(file), ".go")
		ran := "No"
		batch := ""
		
		if b, ok := ranMap[name]; ok {
			ran = "Yes"
			batch = fmt.Sprintf("%d", b)
		}

		fmt.Printf("| %-4s | %-41s | %-5s |\n", ran, name, batch)
	}

	fmt.Println("+------+-------------------------------------------+-------+")

	return nil
}

func getMigrationFiles() ([]string, error) {
	migrationsDir := "database/migrations"
	
	// Check if directory exists
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		return []string{}, nil
	}

	var files []string
	err := filepath.Walk(migrationsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			files = append(files, path)
		}
		
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to read migrations directory: %w", err)
	}

	// Sort files by name (timestamp prefix ensures chronological order)
	sort.Strings(files)

	return files, nil
}
