package main

import (
	"fmt"
	"github.com/y-as-7/go-askar/library/dash-askar/core"
)

// Example demonstrating the new fluent form API
func main() {
	fmt.Println("DashAskar Fluent Forms Examples")
	fmt.Println("==============================")

	// Example 1: Basic text input with validation
	nameField := core.TextInput("name", "Full Name").
		Required().
		MinLength(2).
		MaxLength(50).
		Placeholder("Enter your full name").
		Help("Please provide your first and last name")

	fmt.Printf("Name Field: %+v\n\n", nameField.ToComponent())

	// Example 2: Email input with validation
	emailField := core.EmailInput("email", "Email Address").
		Required().
		MaxLength(255).
		Placeholder("user@example.com").
		Help("This will be your login email")

	fmt.Printf("Email Field: %+v\n\n", emailField.ToComponent())

	// Example 3: Password input
	passwordField := core.PasswordInput("password", "Password").
		Required().
		MinLength(8).
		MaxLength(128).
		Help("Must be at least 8 characters long")

	fmt.Printf("Password Field: %+v\n\n", passwordField.ToComponent())

	// Example 4: Numeric input with range validation
	ageField := core.IntegerInput("age", "Age").
		Required().
		MinValue(18).
		MaxValue(120).
		Placeholder("25").
		Help("Must be between 18 and 120 years old")

	fmt.Printf("Age Field: %+v\n\n", ageField.ToComponent())

	// Example 5: URL input
	websiteField := core.UrlInput("website", "Website").
		Placeholder("https://example.com").
		Help("Optional: Your personal or company website")

	fmt.Printf("Website Field: %+v\n\n", websiteField.ToComponent())

	// Example 6: Telephone input with custom regex
	phoneField := core.TelInput("phone", "Phone Number").
		Required().
		TelRegex(`^\+?[1-9]\d{1,14}$`).
		Placeholder("+1 234 567 8900").
		Help("Include country code")

	fmt.Printf("Phone Field: %+v\n\n", phoneField.ToComponent())

	// Example 7: Advanced text input with custom validation
	usernameField := core.TextInput("username", "Username").
		Required().
		MinLength(3).
		MaxLength(20).
		AlphaDash(). // Only letters, numbers, dashes, underscores
		Placeholder("john_doe123").
		Help("3-20 characters, letters, numbers, dash, underscore only")

	fmt.Printf("Username Field: %+v\n\n", usernameField.ToComponent())

	// Example 8: Numeric input with decimal support
	salaryField := core.NumericInput("salary", "Annual Salary").
		Required().
		MinValue(20000).
		MaxValue(1000000).
		Step(1000).
		Placeholder("50000").
		Help("Enter your annual salary in USD")

	fmt.Printf("Salary Field: %+v\n\n", salaryField.ToComponent())

	// Example 9: Validation testing
	fmt.Println("Validation Examples:")
	fmt.Println("===================")

	// Test email validation
	emailErrors := emailField.Validate("invalid-email")
	fmt.Printf("Email validation errors for 'invalid-email': %v\n", emailErrors)

	emailErrors = emailField.Validate("valid@example.com")
	fmt.Printf("Email validation errors for 'valid@example.com': %v\n", emailErrors)

	// Test name validation
	nameErrors := nameField.Validate("A")
	fmt.Printf("Name validation errors for 'A': %v\n", nameErrors)

	nameErrors = nameField.Validate("John Doe")
	fmt.Printf("Name validation errors for 'John Doe': %v\n", nameErrors)

	// Test age validation
	ageErrors := ageField.Validate("17")
	fmt.Printf("Age validation errors for '17': %v\n", ageErrors)

	ageErrors = ageField.Validate("25")
	fmt.Printf("Age validation errors for '25': %v\n", ageErrors)
}