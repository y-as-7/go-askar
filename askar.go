package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// ANSI color codes
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Bold    = "\033[1m"
)

func main() {
	// Display logo
	displayLogo()
	time.Sleep(500 * time.Millisecond)

	// Get project name
	if len(os.Args) < 2 {
		printError("Please provide a project name")
		fmt.Printf("\n%sUsage:%s go-askar %s<project-name>%s\n\n", Bold, Reset, Cyan, Reset)
		fmt.Printf("%sExample:%s go-askar my-shop\n\n", Bold, Reset)
		os.Exit(1)
	}

	projectName := os.Args[1]

	// Validate project name
	if !isValidProjectName(projectName) {
		printError("Invalid project name. Use only letters, numbers, hyphens, and underscores")
		os.Exit(1)
	}

	printStep("🚀 Creating new go-askar project: " + projectName)
	time.Sleep(300 * time.Millisecond)

	// Clone repository
	printStep("📦 Downloading framework...")
	if err := cloneFramework(projectName); err != nil {
		printError("Failed to clone framework: " + err.Error())
		os.Exit(1)
	}

	// Setup project
	if err := setupProject(projectName); err != nil {
		printError("Failed to setup project: " + err.Error())
		os.Exit(1)
	}

	// Success message
	printSuccess(projectName)
}

func displayLogo() {
	logo := `
` + Cyan + Bold + `
   ██████╗  ██████╗        █████╗ ███████╗██╗  ██╗ █████╗ ██████╗ 
  ██╔════╝ ██╔═══██╗      ██╔══██╗██╔════╝██║ ██╔╝██╔══██╗██╔══██╗
  ██║  ███╗██║   ██║█████╗███████║███████╗█████╔╝ ███████║██████╔╝
  ██║   ██║██║   ██║╚════╝██╔══██║╚════██║██╔═██╗ ██╔══██║██╔══██╗
  ╚██████╔╝╚██████╔╝      ██║  ██║███████║██║  ██╗██║  ██║██║  ██║
   ╚═════╝  ╚═════╝       ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝
` + Reset + `
` + Magenta + `              A Laravel-inspired Go Framework` + Reset + `
` + Yellow + `              Version 1.0.0` + Reset + `
`
	fmt.Print(logo)
}

func printStep(message string) {
	fmt.Printf("%s▸%s %s\n", Green+Bold, Reset, message)
}

func printError(message string) {
	fmt.Printf("\n%s✗ Error:%s %s\n\n", Red+Bold, Reset, message)
}

func printSuccess(projectName string) {
	fmt.Printf("\n%s%s✓ Project created successfully!%s\n\n", Green, Bold, Reset)
	
	box := `
╔════════════════════════════════════════════════════════════╗
║                                                            ║
║  ` + Green + `✓` + Reset + ` Your go-askar project is ready!                      ║
║                                                            ║
║  ` + Cyan + `Next steps:` + Reset + `                                            ║
║                                                            ║
║    ` + Yellow + `cd ` + projectName + Reset + `                                           ║
║    ` + Yellow + `cp .env.example .env` + Reset + `                                ║
║    ` + Yellow + `make run` + Reset + `                                            ║
║                                                            ║
║  ` + Magenta + `Documentation:` + Reset + ` docs/                                   ║
║  ` + Magenta + `API:` + Reset + ` http://localhost:8080                            ║
║                                                            ║
╚════════════════════════════════════════════════════════════╝
`
	fmt.Println(box)
	fmt.Printf("%sHappy coding! 🎉%s\n\n", Bold, Reset)
}

func isValidProjectName(name string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", name)
	return matched && name != ""
}

func cloneFramework(projectName string) error {
	spinner := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	done := make(chan bool)
	
	go func() {
		i := 0
		for {
			select {
			case <-done:
				return
			default:
				fmt.Printf("\r  %s%s%s Cloning repository...", Cyan, spinner[i%len(spinner)], Reset)
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	cmd := exec.Command("git", "clone", "https://github.com/y-as-7/go-askar.git", projectName)
	err := cmd.Run()
	
	done <- true
	fmt.Print("\r  ✓ Repository cloned     \n")
	
	return err
}

func setupProject(projectName string) error {
	projectPath, _ := filepath.Abs(projectName)
	
	// Change to project directory
	if err := os.Chdir(projectPath); err != nil {
		return err
	}

	// Update go.mod
	printStep("📝 Configuring project...")
	time.Sleep(200 * time.Millisecond)
	
	if err := updateGoMod(projectName); err != nil {
		return err
	}

	// Update imports
	printStep("🔧 Updating imports...")
	time.Sleep(200 * time.Millisecond)
	
	if err := updateImports(projectName); err != nil {
		return err
	}

	// Create .env
	printStep("⚙️  Creating environment file...")
	time.Sleep(200 * time.Millisecond)
	
	if err := createEnvFile(); err != nil {
		return err
	}

	// Install dependencies
	printStep("📥 Installing dependencies...")
	time.Sleep(200 * time.Millisecond)
	
	cmd := exec.Command("go", "mod", "tidy")
	if err := cmd.Run(); err != nil {
		return err
	}

	// Remove .git directory
	os.RemoveAll(".git")

	// Initialize new git repo
	printStep("🔧 Initializing git repository...")
	time.Sleep(200 * time.Millisecond)
	
	exec.Command("git", "init").Run()
	exec.Command("git", "add", ".").Run()
	exec.Command("git", "commit", "-m", "Initial commit: "+projectName+" based on go-askar").Run()

	return nil
}

func updateGoMod(projectName string) error {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return err
	}

	updated := strings.ReplaceAll(string(content), "order-system", projectName)
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

			updated := strings.ReplaceAll(string(content), "order-system/", projectName+"/")
			return os.WriteFile(path, []byte(updated), 0644)
		}

		return nil
	})
}

func createEnvFile() error {
	if _, err := os.Stat(".env"); err == nil {
		return nil // .env already exists
	}

	// Read .env.example
	content, err := os.ReadFile(".env.example")
	if err != nil {
		return err
	}

	// Generate JWT secret
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
