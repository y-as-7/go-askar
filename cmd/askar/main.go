package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/y-as-7/go-askar/app/console"
	"github.com/y-as-7/go-askar/foundation"
	"github.com/y-as-7/go-askar/pkg/ui"
)

func main() {
	// Display logo using shared UI package
	ui.DisplayLogo(foundation.Version)
	time.Sleep(500 * time.Millisecond)

	// Context Detection: Are we in an askar project?
	isInsideProject := false
	if _, err := os.Stat("go.mod"); err == nil {
		content, _ := os.ReadFile("go.mod")
		if strings.Contains(string(content), "github.com/y-as-7/go-askar") {
			isInsideProject = true
		}
	}

	// Get subcommand
	if len(os.Args) < 2 {
		showUsage(isInsideProject)
		os.Exit(0)
	}

	command := os.Args[1]

	// Handle Global Project Creation Commands
	if command == "create/project" || command == "create" || command == "new" {
		handleCreateCommand(command)
		return
	}

	// Handle Framework-level Artisan Commands (Artisan Mode)
	if isInsideProject {
		err := console.Run(command, os.Args[1:])
		if err != nil {
			if strings.HasPrefix(err.Error(), "unknown command") {
				ui.PrintError(err.Error())
				showUsage(true)
			} else {
				ui.PrintError(err.Error())
			}
			os.Exit(1)
		}
		return
	}

	// Default fallback for unknown commands outside project
	ui.PrintError("Unknown command: " + command)
	showUsage(false)
	os.Exit(1)
}

func showUsage(isInsideProject bool) {
	fmt.Printf("\n%sUsage:%s\n", ui.Bold, ui.Reset)
	
	if !isInsideProject {
		fmt.Printf("  go askar create/project %s<project-name>%s\n", ui.Cyan, ui.Reset)
		fmt.Printf("  go askar new %s<project-name>%s\n\n", ui.Cyan, ui.Reset)
	} else {
		fmt.Printf("  go askar %s<command>%s [options]\n\n", ui.Cyan, ui.Reset)
		fmt.Printf("%sAvailable commands:%s\n", ui.Bold, ui.Reset)
		
		cmds := console.GetCommands()
		for name, cmd := range cmds {
			if name == "run" { continue }
			fmt.Printf("  %-15s %s\n", ui.Green+name+ui.Reset, cmd.Description)
		}
		fmt.Println("")
	}

	fmt.Printf("%sExample:%s\n", ui.Bold, ui.Reset)
	if !isInsideProject {
		fmt.Printf("  go askar create/project my-shop\n")
	} else {
		fmt.Printf("  go askar serve --watch\n")
	}
	fmt.Println("")
}

func handleCreateCommand(command string) {
	var projectName string

	if command == "create/project" {
		if len(os.Args) < 3 {
			ui.PrintError("Please provide a project name")
			showUsage(false)
			os.Exit(1)
		}
		projectName = strings.TrimSpace(os.Args[2])
	} else if command == "create" {
		if len(os.Args) < 3 {
			ui.PrintError("Please provide a subcommand (e.g., project)")
			showUsage(false)
			os.Exit(1)
		}
		subcommand := os.Args[2]
		if subcommand == "project" {
			if len(os.Args) < 4 {
				ui.PrintError("Please provide a project name")
				showUsage(false)
				os.Exit(1)
			}
			projectName = strings.TrimSpace(os.Args[3])
		} else {
			ui.PrintError("Unknown subcommand: " + subcommand)
			showUsage(false)
			os.Exit(1)
		}
	} else if command == "new" {
		if len(os.Args) < 3 {
			ui.PrintError("Please provide a project name")
			showUsage(false)
			os.Exit(1)
		}
		projectName = strings.TrimSpace(os.Args[2])
	}

	// Validate project name
	if !isValidProjectName(projectName) {
		ui.PrintError(fmt.Sprintf("Invalid project name '%s'. Use only letters, numbers, hyphens, and underscores", projectName))
		os.Exit(1)
	}

	ui.PrintStep("🚀 Creating new askar project: " + projectName)
	time.Sleep(300 * time.Millisecond)

	// Clone repository
	ui.PrintStep("📦 Downloading framework...")
	if err := cloneFramework(projectName); err != nil {
		ui.PrintError("Failed to clone framework: " + err.Error())
		os.Exit(1)
	}

	// Setup project
	if err := setupProject(projectName); err != nil {
		ui.PrintError("Failed to setup project: " + err.Error())
		os.Exit(1)
	}

	// Success message
	printSuccessBox(projectName)
}

func isValidProjectName(name string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", name)
	return matched && name != ""
}

func cloneFramework(projectName string) error {
	if _, err := os.Stat(projectName); err == nil {
		return fmt.Errorf("directory '%s' already exists", projectName)
	}

	cmd := exec.Command("git", "clone", "https://github.com/y-as-7/go-askar.git", projectName)
	return cmd.Run()
}

func setupProject(projectName string) error {
	projectPath, _ := filepath.Abs(projectName)
	
	if err := os.Chdir(projectPath); err != nil {
		return err
	}

	ui.PrintStep("📝 Configuring project...")
	if err := updateGoMod(projectName); err != nil {
		return err
	}

	ui.PrintStep("🔧 Updating imports...")
	if err := updateImports(projectName); err != nil {
		return err
	}

	ui.PrintStep("⚙️  Creating environment file...")
	if err := createEnvFile(); err != nil {
		return err
	}

	ui.PrintStep("📥 Installing dependencies...")
	cmd := exec.Command("go", "mod", "tidy")
	if err := cmd.Run(); err != nil {
		return err
	}

	os.RemoveAll(".git")
	os.Remove("install.sh")
	os.Remove("install.ps1")

	return nil
}

func updateGoMod(projectName string) error {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return err
	}

	updated := strings.ReplaceAll(string(content), "github.com/y-as-7/go-askar", projectName)
	return os.WriteFile("go.mod", []byte(updated), 0644)
}

func updateImports(projectName string) error {
	return filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			updated := strings.ReplaceAll(string(content), "github.com/y-as-7/go-askar/", projectName+"/")
			return os.WriteFile(path, []byte(updated), 0644)
		}

		return nil
	})
}

func createEnvFile() error {
	if _, err := os.Stat(".env"); err == nil {
		return nil
	}

	content, err := os.ReadFile(".env.example")
	if err != nil {
		return err
	}

	jwtSecret := generateSecret(64)
	updated := strings.ReplaceAll(string(content), "CHANGE_THIS_TO_RANDOM_SECRET", jwtSecret)

	return os.WriteFile(".env", []byte(updated), 0644)
}

func generateSecret(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	secret := make([]byte, length)
	for i := range secret {
		secret[i] = charset[time.Now().UnixNano()%int64(len(charset))]
		time.Sleep(1 * time.Nanosecond)
	}
	return string(secret)
}

func printSuccessBox(projectName string) {
	fmt.Printf("\n%s%s✓ Project created successfully!%s\n\n", ui.Green, ui.Bold, ui.Reset)
	
	box := `
  ╔════════════════════════════════════════════════════════════╗
  ║                                                            ║
  ║  ` + ui.Green + `✓` + ui.Reset + ` Your go-askar project is ready!                      ║
  ║                                                            ║
  ║  ` + ui.Cyan + `Next steps:` + ui.Reset + `                                            ║
  ║                                                            ║
  ║    ` + ui.Yellow + `cd ` + projectName + ui.Reset + `                                           ║
  ║    ` + ui.Yellow + `cp .env.example .env` + ui.Reset + `                                ║
  ║    ` + ui.Yellow + `go run main.go` + ui.Reset + `                                     ║
  ║                                                            ║
  ║  ` + ui.Magenta + `Documentation:` + ui.Reset + ` docs/                                   ║
  ║  ` + ui.Magenta + `API:` + Reset + ` http://localhost:8080                            ║
  ║                                                            ║
  ╚════════════════════════════════════════════════════════════╝
`
	fmt.Println(box)
	fmt.Printf("%sHappy coding! 🎉%s\n\n", ui.Bold, ui.Reset)
}
