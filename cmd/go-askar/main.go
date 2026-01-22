package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	_ "github.com/y-as-7/go-askar/app/Console" // Import to register commands
	console "github.com/y-as-7/go-askar/app/Console"
)

const Version = "0.1.0"

// ANSI Color Codes
const (
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
	ColorBold   = "\033[1m"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h" {
		displayLogo()
		displayUsage()
		return
	}

	command := os.Args[1]
	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	// Run console command
	if err := console.Run(command, args); err != nil {
		fmt.Printf("%s✗ Error: %s%s\n\n", ColorYellow, err.Error(), ColorReset)
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

              %sThe Expressive Go Framework%s
              Version %s%s%s ✨
`
	fmt.Printf(logo, ColorBold, ColorReset, ColorCyan, Version, ColorReset)
}

func displayUsage() {
	fmt.Printf("\n%sUsage:%s\n", ColorGreen, ColorReset)
	fmt.Printf("  go askar <command> [options]\n")

	commands := console.GetCommands()

	// Categorize commands
	categories := map[string][]string{
		"System":       {},
		"Scaffolding": {},
		"Database":    {},
		"Other":       {},
	}

	for name := range commands {
		if strings.HasPrefix(name, "make:") {
			categories["Scaffolding"] = append(categories["Scaffolding"], name)
		} else if strings.HasPrefix(name, "migrate:") || name == "migrate" {
			categories["Database"] = append(categories["Database"], name)
		} else if name == "run" || name == "serve" {
			categories["System"] = append(categories["System"], name)
		} else {
			categories["Other"] = append(categories["Other"], name)
		}
	}

	// Sort categories order
	catOrder := []string{"System", "Scaffolding", "Database", "Other"}

	// Get max command name length for alignment
	maxLen := 0
	for name := range commands {
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}

	fmt.Printf("\n%sAvailable commands:%s\n", ColorGreen, ColorReset)

	for _, catName := range catOrder {
		cmds := categories[catName]
		if len(cmds) == 0 {
			continue
		}
		sort.Strings(cmds)

		fmt.Printf("\n  %s%s%s\n", ColorYellow, catName, ColorReset)
		for _, name := range cmds {
			cmd := commands[name]
			padding := strings.Repeat(" ", maxLen-len(name)+4)
			fmt.Printf("    %s%s%s%s%s\n", ColorCyan, name, ColorReset, padding, cmd.Description)
		}
	}

	fmt.Printf("\n%sExample:%s\n", ColorGreen, ColorReset)
	fmt.Printf("  go askar run --watch\n")
	fmt.Printf("  go askar make:model Product\n\n")
}
