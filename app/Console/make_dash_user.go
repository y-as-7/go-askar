package console

import (
	"fmt"
	"strings"

	"github.com/y-as-7/go-askar/app/Models"
	"github.com/y-as-7/go-askar/config"
)

func init() {
	Register(Command{
		Name:        "make:dash-user",
		Description: "Create a new DashAskar admin user",
		Execute:     handleMakeDashUser,
	})
}

func handleMakeDashUser(args []string) error {
	var name, email, password string

	fmt.Print("Enter User Name: ")
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("name is required")
	}

	fmt.Print("Enter Email Address: ")
	fmt.Scanln(&email)
	email = strings.TrimSpace(email)
	if email == "" {
		return fmt.Errorf("email is required")
	}

	fmt.Print("Enter Password: ")
	fmt.Scanln(&password)
	password = strings.TrimSpace(password)
	if password == "" {
		return fmt.Errorf("password is required")
	}

	// Ensure database is connected
	if config.DB == nil {
		config.ConnectDatabase()
	}

	user := Models.User{
		Name:     name,
		Email:    email,
		Role:     "admin",
		IsActive: true,
	}

	if err := user.HashPassword(password); err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	fmt.Printf("\n✅ Admin user created successfully!\n")
	fmt.Printf("   Email: %s\n", email)
	fmt.Printf("   Role:  admin\n\n")

	return nil
}
