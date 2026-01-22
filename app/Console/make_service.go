package console

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	Register(Command{
		Name:        "make:service",
		Description: "Create a new service file",
		Execute:     handleMakeService,
	})
}

func handleMakeService(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("service name is required. Usage: go askar make:service <ServiceName>")
	}

	serviceName := args[0]

	// Validate service name
	if err := validateModelName(serviceName); err != nil {
		return err
	}

	// Ensure first letter is uppercase (Go convention)
	serviceName = capitalizeFirst(serviceName)

	// Create service file
	if err := createServiceFile(serviceName); err != nil {
		return err
	}

	fmt.Printf("\n✅ Service created successfully!\n")
	fmt.Printf("   📄 app/Services/%s_service.go\n\n", strings.ToLower(serviceName))

	return nil
}

func createServiceFile(serviceName string) error {
	servicesDir := "app/Services"
	fileName := filepath.Join(servicesDir, strings.ToLower(serviceName)+"_service.go")

	// Check if file already exists
	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("service file already exists: %s", fileName)
	}

	// Ensure directory exists
	if err := os.MkdirAll(servicesDir, 0755); err != nil {
		return fmt.Errorf("failed to create services directory: %w", err)
	}

	template := fmt.Sprintf(`package Services

import (
	"github.com/y-as-7/go-askar/app/Models"
	"github.com/y-as-7/go-askar/config"
)

type %sService struct{}

// NewService creates a new %s service instance
func New%sService() *%sService {
	return &%sService{}
}

// Create creates a new %s
func (s *%sService) Create(data *Models.%s) error {
	return config.DB.Create(data).Error
}

// GetByID retrieves a %s by ID
func (s *%sService) GetByID(id uint) (*Models.%s, error) {
	var item Models.%s
	err := config.DB.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetAll retrieves all %s records
func (s *%sService) GetAll() ([]Models.%s, error) {
	var items []Models.%s
	err := config.DB.Find(&items).Error
	return items, err
}

// Update updates a %s
func (s *%sService) Update(id uint, data *Models.%s) error {
	return config.DB.Model(&Models.%s{}).Where("id = ?", id).Updates(data).Error
}

// Delete deletes a %s by ID
func (s *%sService) Delete(id uint) error {
	return config.DB.Delete(&Models.%s{}, id).Error
}
`, 
		serviceName, // Type definition
		serviceName, // NewService comment
		serviceName, // NewService function name
		serviceName, // NewService return type
		serviceName, // NewService struct
		serviceName, // Create comment
		serviceName, // Create receiver
		serviceName, // Create parameter type
		serviceName, // GetByID comment
		serviceName, // GetByID receiver
		serviceName, // GetByID return type
		serviceName, // GetByID variable type
		serviceName, // GetAll comment (plural)
		serviceName, // GetAll receiver
		serviceName, // GetAll return type
		serviceName, // GetAll slice type
		serviceName, // Update comment
		serviceName, // Update receiver
		serviceName, // Update parameter type
		serviceName, // Update Model type
		serviceName, // Delete comment
		serviceName, // Delete receiver
		serviceName, // Delete Model type
	)

	if err := os.WriteFile(fileName, []byte(template), 0644); err != nil {
		return fmt.Errorf("failed to create service file: %w", err)
	}

	return nil
}
