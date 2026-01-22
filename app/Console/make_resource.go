package console

import (
	"fmt"
	"strings"
)

func init() {
	Register(Command{
		Name:        "make:resource",
		Description: "Generate a full CRUD resource (Model, DTO, Service, Controller, Migration)",
		Execute:     handleMakeResource,
	})
}

func handleMakeResource(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("resource name is required. Usage: go askar make:resource <ResourceName>")
	}

	name := capitalizeFirst(args[0])

	fmt.Printf("\n🚀 Generating resource: %s\n", name)

	// 1. Generate Model and DTO
	fmt.Printf("📦 Generating Model and DTO...\n")
	tableName := pluralize(strings.ToLower(name))
	if err := createModelFile(name, tableName); err != nil {
		fmt.Printf("⚠️  %v\n", err)
	}
	if err := createDTOFile(name); err != nil {
		fmt.Printf("⚠️  %v\n", err)
	}

	// 2. Generate Service
	fmt.Printf("🛠️  Generating Service...\n")
	if err := createServiceFile(name); err != nil {
		fmt.Printf("⚠️  %v\n", err)
	}

	// 3. Generate Controller
	fmt.Printf("🎮 Generating Controller...\n")
	controllerName := name + "Controller"
	if err := createControllerFile(controllerName); err != nil {
		fmt.Printf("⚠️  %v\n", err)
	}

	// 4. Generate Migration
	fmt.Printf("📜 Generating Migration...\n")
	migrationName := fmt.Sprintf("create_%s_table", pluralize(toSnakeCase(name)))
	if _, err := createMigrationFile(migrationName); err != nil {
		fmt.Printf("⚠️  %v\n", err)
	}

	fmt.Printf("\n✅ Resource %s generated successfully!\n", name)
	fmt.Printf("   Please update your routes and database config.\n\n")

	return nil
}
