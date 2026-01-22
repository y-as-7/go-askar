package console

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func init() {
	Register(Command{
		Name:        "key:generate",
		Description: "Set the application key (JWT_SECRET)",
		Execute:     handleKeyGenerate,
	})
}

func handleKeyGenerate(args []string) error {
	fmt.Println("\033[32m▸\033[0m 🔑 Generating secure application key...")

	// Generate 32 bytes of random data for the secret
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return err
	}
	secret := hex.EncodeToString(bytes)

	// Check for .env file
	envPath := ".env"
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		return fmt.Errorf(".env file not found. Please create one from .env.example first")
	}

	// Read .env file
	content, err := os.ReadFile(envPath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	found := false
	for i, line := range lines {
		if strings.HasPrefix(line, "JWT_SECRET=") {
			lines[i] = "JWT_SECRET=" + secret
			found = true
			break
		}
	}

	if !found {
		lines = append(lines, "JWT_SECRET="+secret)
	}

	// Write back to .env
	err = os.WriteFile(envPath, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("\033[32m✓\033[0m Application key set successfully: %s\n", secret)
	return nil
}
