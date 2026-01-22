package console

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	Register(Command{
		Name:        "run",
		Description: "Start the application server",
		Execute:     handleServe,
	})
	
	// Alias 'serve' to 'run'
	Register(Command{
		Name:        "serve",
		Description: "Alias for run",
		Execute:     handleServe,
	})
}

func handleServe(args []string) error {
	watch := false
	for _, arg := range args {
		if arg == "--watch" {
			watch = true
			break
		}
	}

	if watch {
		fmt.Println("\033[32m▸\033[0m 👀 Starting server with hot reload...")
		if _, err := exec.LookPath("air"); err != nil {
			return fmt.Errorf("hot reload requires 'air'. Please install it with: go install github.com/air-verse/air@latest")
		}

		cmd := exec.Command("air")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	fmt.Println("\033[32m▸\033[0m 🚀 Starting server...")
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
