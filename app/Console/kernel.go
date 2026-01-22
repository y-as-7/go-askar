package console

import (
	"fmt"
)

// Command interface that all Artisan commands must implement
type Command struct {
	Name        string
	Description string
	Execute     func(args []string) error
}

var commands = make(map[string]Command)

// Register a new command
func Register(cmd Command) {
	commands[cmd.Name] = cmd
}

// Run executes a command from the registry
func Run(name string, args []string) error {
	cmd, ok := commands[name]
	if !ok {
		return fmt.Errorf("unknown command: %s", name)
	}
	return cmd.Execute(args)
}

// GetCommands returns all registered commands
func GetCommands() map[string]Command {
	return commands
}

// Reset the registry (mainly for testing or re-init)
func Reset() {
	commands = make(map[string]Command)
}
