package console

import (
	"fmt"
)

func init() {
	Register(Command{
		Name:        "mytest",
		Description: "Command description",
		Execute:     handleMyTestCommand,
	})
}

func handleMyTestCommand(args []string) error {
	fmt.Println("Hello from mytest command!")
	return nil
}
