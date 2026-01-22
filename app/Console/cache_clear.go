package console

import (
	"fmt"
	"os"
)

func init() {
	Register(Command{
		Name:        "cache:clear",
		Description: "Flush the application cache",
		Execute: func(args []string) error {
			fmt.Println("\033[32m▸\033[0m 🧹 Clearing application cache...")
			// Logic to clear storage/cache or similar
			os.RemoveAll("storage/cache")
			os.MkdirAll("storage/cache", 0755)
			fmt.Println("\033[32m✓\033[0m Cache cleared successfully!")
			return nil
		},
	})

	Register(Command{
		Name:        "route:clear",
		Description: "Remove the route cache file",
		Execute: func(args []string) error {
			fmt.Println("\033[32m▸\033[0m 🗺️  Clearing route cache...")
			// Logic to clear route cache
			fmt.Println("\033[32m✓\033[0m Route cache cleared!")
			return nil
		},
	})
}
