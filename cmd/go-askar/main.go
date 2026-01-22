package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	_ "github.com/y-as-7/go-askar/app/Console" // Import to register commands
	console "github.com/y-as-7/go-askar/app/Console"
)

const Version = "1.4.0"

func main() {
	displayLogo()

	if len(os.Args) < 2 {
		displayUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	// Handle special commands
	if command == "create/project" || command == "create" && len(args) > 0 && args[0] == "project" {
		handleCreateProject(args)
		return
	}

	// Run console command
	if err := console.Run(command, args); err != nil {
		fmt.Printf("✗ Error: %s\n\n", err.Error())
		displayUsage()
		os.Exit(1)
	}
}

func displayLogo() {
	logo := `

   █████╗ ███████╗██╗  ██╗ █████╗ ██████╗ 
  ██╔══██╗██╔════╝██║ ██╔╝██╔══██╗██╔══██╗
  ███████║███████╗█████╔╝ ███████║██████╔╝
  ██╔══██║╚════██║██╔═██╗ ██╔══██║██╔══██╗
  ██║  ██║███████║██║  ██╗██║  ██║██║  ██║
  ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝

              A Laravel-inspired Go Framework
              Version ` + Version + ` ✨
`
	fmt.Println(logo)
}

func displayUsage() {
	fmt.Println("\nUsage:")
	fmt.Println("  go askar <command> [options]")
	fmt.Println("\nAvailable commands:")

	commands := console.GetCommands()
	
	// Get max command name length for alignment
	maxLen := len("create/project") // Start with create/project length
	for name := range commands {
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}

	// Display commands
	for name, cmd := range commands {
		padding := strings.Repeat(" ", maxLen-len(name)+2)
		fmt.Printf("  %s%s%s\n", name, padding, cmd.Description)
	}

	// Add create/project to the list
	createPadding := strings.Repeat(" ", maxLen-len("create/project")+2)
	fmt.Printf("  %s%s%s\n", "create/project", createPadding, "Create a new askar project")

	fmt.Println("\nExample:")
	fmt.Println("  go askar run --watch")
	fmt.Println("  go askar make:model Product")
}

func handleCreateProject(args []string) {
	// Remove "project" from args if it's there
	projectArgs := args
	if len(args) > 0 && args[0] == "project" {
		projectArgs = args[1:]
	}

	if len(projectArgs) == 0 {
		fmt.Println("✗ Error: project name is required")
		fmt.Println("\nUsage:")
		fmt.Println("  go askar create/project <project-name>")
		os.Exit(1)
	}

	projectName := projectArgs[0]
	
	fmt.Printf("🚀 Creating new askar project: %s\n\n", projectName)
	fmt.Println("📦 Cloning template...")
	
	// Clone the repository
	if err := runCommand("git", "clone", "https://github.com/y-as-7/go-askar.git", projectName); err != nil {
		fmt.Printf("✗ Failed to clone repository: %v\n", err)
		os.Exit(1)
	}

	// Change to project directory and run setup script
	if err := os.Chdir(projectName); err != nil {
		fmt.Printf("✗ Failed to enter project directory: %v\n", err)
		os.Exit(1)
	}

	// Run the install script
	fmt.Println("\n📝 Setting up project...")
	if err := runCommand("bash", "install.sh", projectName); err != nil {
		fmt.Printf("✗ Failed to setup project: %v\n", err)
		os.Exit(1)
	}
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
