package console

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func init() {
	Register(Command{
		Name:        "make:model",
		Description: "Create a new model and DTO files",
		Execute:     handleMakeModel,
	})
}

func handleMakeModel(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("model name is required. Usage: go askar make:model <ModelName>")
	}

	modelName := args[0]

	// Validate model name
	if err := validateModelName(modelName); err != nil {
		return err
	}

	// Ensure first letter is uppercase (Go convention)
	modelName = capitalizeFirst(modelName)

	// Generate table name (lowercase, pluralized)
	tableName := pluralize(strings.ToLower(modelName))

	// Create model file
	if err := createModelFile(modelName, tableName); err != nil {
		return err
	}

	// Create DTO file
	if err := createDTOFile(modelName); err != nil {
		return err
	}

	fmt.Printf("\n✅ Model created successfully!\n")
	fmt.Printf("   📄 app/Models/%s.go\n", strings.ToLower(modelName))
	fmt.Printf("   📄 app/DTO/%s_dto.go\n\n", strings.ToLower(modelName))

	return nil
}

func validateModelName(name string) error {
	if name == "" {
		return fmt.Errorf("model name cannot be empty")
	}

	// Check if first character is a letter
	firstChar := rune(name[0])
	if !unicode.IsLetter(firstChar) {
		return fmt.Errorf("model name must start with a letter")
	}

	// Check if all characters are alphanumeric
	for _, char := range name {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return fmt.Errorf("model name must contain only letters and numbers")
		}
	}

	return nil
}

func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func pluralize(word string) string {
	// Simple pluralization rules
	if strings.HasSuffix(word, "y") && len(word) > 1 {
		// Check if the letter before 'y' is a consonant
		beforeY := rune(word[len(word)-2])
		if !isVowel(beforeY) {
			return word[:len(word)-1] + "ies"
		}
	}

	if strings.HasSuffix(word, "s") || strings.HasSuffix(word, "x") ||
		strings.HasSuffix(word, "ch") || strings.HasSuffix(word, "sh") {
		return word + "es"
	}

	return word + "s"
}

func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, r)
}

func createModelFile(modelName, tableName string) error {
	modelsDir := "app/Models"
	fileName := filepath.Join(modelsDir, strings.ToLower(modelName)+".go")

	// Check if file already exists
	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("model file already exists: %s", fileName)
	}

	// Ensure directory exists
	if err := os.MkdirAll(modelsDir, 0755); err != nil {
		return fmt.Errorf("failed to create models directory: %w", err)
	}

	template := fmt.Sprintf(`package Models

import (
	"time"

	"gorm.io/gorm"
)

type %s struct {
	ID        uint           `+"`gorm:\"primarykey\" json:\"id\"`"+`
	CreatedAt time.Time      `+"`json:\"created_at\"`"+`
	UpdatedAt time.Time      `+"`json:\"updated_at\"`"+`
	DeletedAt gorm.DeletedAt `+"`gorm:\"index\" json:\"-\"`"+`
	
	// Add your fields here
}

// TableName specifies table name
func (%s) TableName() string {
	return "%s"
}
`, modelName, modelName, tableName)

	if err := os.WriteFile(fileName, []byte(template), 0644); err != nil {
		return fmt.Errorf("failed to create model file: %w", err)
	}

	return nil
}

func createDTOFile(modelName string) error {
	dtoDir := "app/DTO"
	fileName := filepath.Join(dtoDir, strings.ToLower(modelName)+"_dto.go")

	// Check if file already exists
	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("DTO file already exists: %s", fileName)
	}

	// Ensure directory exists
	if err := os.MkdirAll(dtoDir, 0755); err != nil {
		return fmt.Errorf("failed to create DTO directory: %w", err)
	}

	template := fmt.Sprintf(`package DTO

type Create%sRequest struct {
	// Add your fields here with validation tags
	// Example:
	// Name string `+"`json:\"name\" binding:\"required\"`"+`
}

type Update%sRequest struct {
	// Add your fields here with validation tags
	// Example:
	// Name string `+"`json:\"name\" binding:\"required\"`"+`
}
`, modelName, modelName)

	if err := os.WriteFile(fileName, []byte(template), 0644); err != nil {
		return fmt.Errorf("failed to create DTO file: %w", err)
	}

	return nil
}
