package console

import (
	"fmt"
	"strings"
	"unicode"
)

func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

func toPascalCase(s string) string {
	words := strings.Split(s, "_")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	return strings.Join(words, "")
}

func pluralize(word string) string {
	// Simple pluralization rules
	if strings.HasSuffix(word, "y") && len(word) > 1 {
		// Check if the letter before 'y' is a consonant
		beforeY := rune(word[len(word)-2])
		if !isVowel(beforeY) {
			return word[:len(word)-1] + "ies"
		}
	}

	if strings.HasSuffix(word, "s") || strings.HasSuffix(word, "x") ||
		strings.HasSuffix(word, "ch") || strings.HasSuffix(word, "sh") {
		return word + "es"
	}

	return word + "s"
}

func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, r)
}

func validateModelName(name string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	// Check if first character is a letter
	firstChar := rune(name[0])
	if !unicode.IsLetter(firstChar) {
		return fmt.Errorf("name must start with a letter")
	}

	// Check if all characters are alphanumeric
	for _, char := range name {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return fmt.Errorf("name must contain only letters and numbers")
		}
	}

	return nil
}
