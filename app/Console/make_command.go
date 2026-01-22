package console

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	Register(Command{
		Name:        "make:command",
		Description: "Create a new console command",
		Execute:     handleMakeCommand,
	})
}

func handleMakeCommand(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("command name is required. Usage: go askar make:command <CommandName>")
	}

	name := args[0]
	name = capitalizeFirst(name)

	// Create command file
	if err := createCommandFile(name); err != nil {
		return err
	}

	fmt.Printf("\n✅ Command created successfully!\n")
	fmt.Printf("   📄 app/Console/%s.go\n\n", toSnakeCase(name))

	return nil
}

func createCommandFile(name string) error {
	consoleDir := "app/Console"
	fileName := filepath.Join(consoleDir, toSnakeCase(name)+".go")

	// Check if file already exists
	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("command file already exists: %s", fileName)
	}

	// Generate command name (e.g. TestCommand -> test:command or just test)
	cmdTrigger := strings.ToLower(strings.TrimSuffix(name, "Command"))

	template := fmt.Sprintf(`package console

import (
	"fmt"
)

func init() {
	Register(Command{
		Name:        "%s",
		Description: "Command description",
		Execute:     handle%s,
	})
}

func handle%s(args []string) error {
	fmt.Println("Hello from %s command!")
	return nil
}
`, cmdTrigger, name, name, cmdTrigger)

	if err := os.WriteFile(fileName, []byte(template), 0644); err != nil {
		return fmt.Errorf("failed to create command file: %w", err)
	}

	return nil
}
