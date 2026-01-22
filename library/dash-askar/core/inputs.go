package core

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// BaseInput represents the foundation for all form inputs
type BaseInput struct {
	Type         string                 `json:"type"`
	Name         string                 `json:"name"`
	Label        string                 `json:"label"`
	Placeholder  string                 `json:"placeholder,omitempty"`
	Required     bool                   `json:"required"`
	Value        interface{}            `json:"value,omitempty"`
	Rules        []string               `json:"rules,omitempty"`
	Attributes   map[string]interface{} `json:"attributes,omitempty"`
	Classes      []string               `json:"classes,omitempty"`
	IsReadOnly   bool                   `json:"readonly,omitempty"`
	IsDisabled   bool                   `json:"disabled,omitempty"`
	HelpText     string                 `json:"help_text,omitempty"`
	DefaultValue interface{}            `json:"default_value,omitempty"`
}

// TextInputField represents a text input field with fluent API
type TextInputField struct {
	*BaseInput
	minLength       *int
	maxLength       *int
	isEmail         bool
	isPassword      bool
	isUrl           bool
	isTel           bool
	isNumeric       bool
	isInteger       bool
	minValue        *float64
	maxValue        *float64
	step            *float64
	telRegex        string
	customType      string
	mask            string
	autocomplete    string
	autocapitalize  string
	inputMode       string
	pattern         string
	stripCharacters []string
	trim            bool
}

// NewTextInput creates a new TextInputField instance
func NewTextInput(name, label string) *TextInputField {
	return &TextInputField{
		BaseInput: &BaseInput{
			Type:       "text",
			Name:       name,
			Label:      label,
			Rules:      []string{},
			Attributes: make(map[string]interface{}),
			Classes:    []string{},
		},
		telRegex: `^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\.\/0-9]*$`,
		trim:     true,
	}
}

// Email sets the input as email type with validation
func (t *TextInputField) Email(condition ...bool) *TextInputField {
	cond := len(condition) == 0 || condition[0]
	if cond {
		t.isEmail = true
		t.Type = "email"
		t.addRule("email")
		t.setInputMode("email")
	}
	return t
}

// Password sets the input as password type
func (t *TextInputField) Password(condition ...bool) *TextInputField {
	cond := len(condition) == 0 || condition[0]
	if cond {
		t.isPassword = true
		t.Type = "password"
	}
	return t
}

// Url sets the input as url type with validation
func (t *TextInputField) Url(condition ...bool) *TextInputField {
	cond := len(condition) == 0 || condition[0]
	if cond {
		t.isUrl = true
		t.Type = "url"
		t.addRule("url")
	}
	return t
}

// Tel sets the input as telephone type with validation
func (t *TextInputField) Tel(condition ...bool) *TextInputField {
	cond := len(condition) == 0 || condition[0]
	if cond {
		t.isTel = true
		t.Type = "tel"
		t.setInputMode("tel")
		t.addRegexRule(t.telRegex)
	}
	return t
}

// TelRegex sets a custom regex for telephone validation
func (t *TextInputField) TelRegex(regex string) *TextInputField {
	t.telRegex = regex
	if t.isTel {
		t.addRegexRule(regex)
	}
	return t
}

// Numeric sets the input as numeric with validation
func (t *TextInputField) Numeric(condition ...bool) *TextInputField {
	cond := len(condition) == 0 || condition[0]
	if cond {
		t.isNumeric = true
		t.Type = "number"
		t.setInputMode("decimal")
		t.addRule("numeric")
		if t.step == nil {
			t.step = &[]float64{0.01}[0] // any
		}
	}
	return t
}

// Integer sets the input as integer with validation
func (t *TextInputField) Integer(condition ...bool) *TextInputField {
	cond := len(condition) == 0 || condition[0]
	if cond {
		t.isInteger = true
		t.isNumeric = true
		t.Type = "number"
		t.setInputMode("numeric")
		t.addRule("integer")
		t.step = &[]float64{1}[0]
	}
	return t
}

// MinLength sets minimum length validation
func (t *TextInputField) MinLength(length int) *TextInputField {
	t.minLength = &length
	t.addRule(fmt.Sprintf("min:%d", length))
	t.setAttribute("minlength", length)
	return t
}

// MaxLength sets maximum length validation
func (t *TextInputField) MaxLength(length int) *TextInputField {
	t.maxLength = &length
	t.addRule(fmt.Sprintf("max:%d", length))
	t.setAttribute("maxlength", length)
	return t
}

// Length sets both min and max length validation
func (t *TextInputField) Length(min, max int) *TextInputField {
	return t.MinLength(min).MaxLength(max)
}

// MinValue sets minimum value for numeric inputs
func (t *TextInputField) MinValue(value float64) *TextInputField {
	t.minValue = &value
	t.addRule(fmt.Sprintf("min:%g", value))
	t.setAttribute("min", value)
	return t
}

// MaxValue sets maximum value for numeric inputs
func (t *TextInputField) MaxValue(value float64) *TextInputField {
	t.maxValue = &value
	t.addRule(fmt.Sprintf("max:%g", value))
	t.setAttribute("max", value)
	return t
}

// Step sets the step value for numeric inputs
func (t *TextInputField) Step(value float64) *TextInputField {
	t.step = &value
	t.setAttribute("step", value)
	return t
}

// Required makes the field required
func (t *TextInputField) Required(condition ...bool) *TextInputField {
	cond := len(condition) == 0 || condition[0]
	t.BaseInput.Required = cond
	if cond {
		t.addRule("required")
	}
	return t
}

// Placeholder sets the placeholder text
func (t *TextInputField) Placeholder(text string) *TextInputField {
	t.BaseInput.Placeholder = text
	return t
}

// Default sets the default value
func (t *TextInputField) Default(value interface{}) *TextInputField {
	t.BaseInput.DefaultValue = value
	t.BaseInput.Value = value
	return t
}

// ReadOnly makes the field read-only
func (t *TextInputField) ReadOnly(condition ...bool) *TextInputField {
	cond := len(condition) == 0 || condition[0]
	t.BaseInput.IsReadOnly = cond
	t.setAttribute("readonly", cond)
	return t
}

// Disabled makes the field disabled
func (t *TextInputField) Disabled(condition ...bool) *TextInputField {
	cond := len(condition) == 0 || condition[0]
	t.BaseInput.IsDisabled = cond
	t.setAttribute("disabled", cond)
	return t
}

// Help sets help text for the field
func (t *TextInputField) Help(text string) *TextInputField {
	t.BaseInput.HelpText = text
	return t
}

// Autocomplete sets the autocomplete attribute
func (t *TextInputField) Autocomplete(value string) *TextInputField {
	t.autocomplete = value
	t.setAttribute("autocomplete", value)
	return t
}

// Autocapitalize sets the autocapitalize attribute
func (t *TextInputField) Autocapitalize(value string) *TextInputField {
	t.autocapitalize = value
	t.setAttribute("autocapitalize", value)
	return t
}

// Pattern sets a regex pattern for validation
func (t *TextInputField) Pattern(pattern string) *TextInputField {
	t.pattern = pattern
	t.setAttribute("pattern", pattern)
	t.addRegexRule(pattern)
	return t
}

// Mask sets an input mask
func (t *TextInputField) Mask(mask string) *TextInputField {
	t.mask = mask
	t.setAttribute("data-mask", mask)
	return t
}

// StripCharacters sets characters to strip from input
func (t *TextInputField) StripCharacters(chars []string) *TextInputField {
	t.stripCharacters = chars
	return t
}

// Trim enables/disables trimming whitespace
func (t *TextInputField) Trim(condition ...bool) *TextInputField {
	cond := len(condition) == 0 || condition[0]
	t.trim = cond
	return t
}

// Alpha validates that the field contains only alphabetic characters
func (t *TextInputField) Alpha() *TextInputField {
	t.addRule("alpha")
	return t
}

// AlphaNum validates that the field contains only alphanumeric characters
func (t *TextInputField) AlphaNum() *TextInputField {
	t.addRule("alpha_num")
	return t
}

// AlphaDash validates that the field contains only alphanumeric characters, dashes, and underscores
func (t *TextInputField) AlphaDash() *TextInputField {
	t.addRule("alpha_dash")
	return t
}

// Regex adds a custom regex validation rule
func (t *TextInputField) Regex(pattern string) *TextInputField {
	t.addRegexRule(pattern)
	return t
}

// Rule adds a custom validation rule
func (t *TextInputField) Rule(rule string) *TextInputField {
	t.addRule(rule)
	return t
}

// Class adds CSS classes
func (t *TextInputField) Class(classes ...string) *TextInputField {
	t.BaseInput.Classes = append(t.BaseInput.Classes, classes...)
	return t
}

// Attribute sets a custom HTML attribute
func (t *TextInputField) Attribute(key string, value interface{}) *TextInputField {
	t.setAttribute(key, value)
	return t
}

// Helper methods

func (t *TextInputField) addRule(rule string) {
	// Check if rule already exists
	for _, existingRule := range t.BaseInput.Rules {
		if existingRule == rule {
			return
		}
	}
	t.BaseInput.Rules = append(t.BaseInput.Rules, rule)
}

func (t *TextInputField) addRegexRule(pattern string) {
	rule := fmt.Sprintf("regex:%s", pattern)
	t.addRule(rule)
}

func (t *TextInputField) setAttribute(key string, value interface{}) {
	if t.BaseInput.Attributes == nil {
		t.BaseInput.Attributes = make(map[string]interface{})
	}
	t.BaseInput.Attributes[key] = value
}

func (t *TextInputField) setInputMode(mode string) {
	t.inputMode = mode
	t.setAttribute("inputmode", mode)
}

// GetType returns the appropriate input type
func (t *TextInputField) GetType() string {
	if t.customType != "" {
		return t.customType
	}
	if t.isEmail {
		return "email"
	}
	if t.isPassword {
		return "password"
	}
	if t.isNumeric || t.isInteger {
		return "number"
	}
	if t.isTel {
		return "tel"
	}
	if t.isUrl {
		return "url"
	}
	return "text"
}

// Validate validates the input value against the rules
func (t *TextInputField) Validate(value interface{}) []string {
	var errors []string
	strValue := ""

	if value != nil {
		strValue = fmt.Sprintf("%v", value)
	}

	// Process value
	if t.trim {
		strValue = strings.TrimSpace(strValue)
	}

	if t.stripCharacters != nil {
		for _, char := range t.stripCharacters {
			strValue = strings.ReplaceAll(strValue, char, "")
		}
	}

	// Validate rules
	for _, rule := range t.BaseInput.Rules {
		if err := t.validateRule(rule, strValue); err != "" {
			errors = append(errors, err)
		}
	}

	return errors
}

func (t *TextInputField) validateRule(rule, value string) string {
	parts := strings.SplitN(rule, ":", 2)
	ruleName := parts[0]
	ruleValue := ""
	if len(parts) > 1 {
		ruleValue = parts[1]
	}

	switch ruleName {
	case "required":
		if value == "" {
			return fmt.Sprintf("%s is required", t.BaseInput.Label)
		}
	case "email":
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		if matched, _ := regexp.MatchString(emailRegex, value); !matched && value != "" {
			return fmt.Sprintf("%s must be a valid email address", t.BaseInput.Label)
		}
	case "url":
		urlRegex := `^https?://[^\s/$.?#].[^\s]*$`
		if matched, _ := regexp.MatchString(urlRegex, value); !matched && value != "" {
			return fmt.Sprintf("%s must be a valid URL", t.BaseInput.Label)
		}
	case "numeric":
		if _, err := strconv.ParseFloat(value, 64); err != nil && value != "" {
			return fmt.Sprintf("%s must be a number", t.BaseInput.Label)
		}
	case "integer":
		if _, err := strconv.Atoi(value); err != nil && value != "" {
			return fmt.Sprintf("%s must be an integer", t.BaseInput.Label)
		}
	case "alpha":
		alphaRegex := `^[a-zA-Z]+$`
		if matched, _ := regexp.MatchString(alphaRegex, value); !matched && value != "" {
			return fmt.Sprintf("%s must contain only letters", t.BaseInput.Label)
		}
	case "alpha_num":
		alphaNumRegex := `^[a-zA-Z0-9]+$`
		if matched, _ := regexp.MatchString(alphaNumRegex, value); !matched && value != "" {
			return fmt.Sprintf("%s must contain only letters and numbers", t.BaseInput.Label)
		}
	case "alpha_dash":
		alphaDashRegex := `^[a-zA-Z0-9_-]+$`
		if matched, _ := regexp.MatchString(alphaDashRegex, value); !matched && value != "" {
			return fmt.Sprintf("%s must contain only letters, numbers, dashes, and underscores", t.BaseInput.Label)
		}
	case "min":
		minVal, _ := strconv.Atoi(ruleValue)
		if len(value) < minVal && value != "" {
			return fmt.Sprintf("%s must be at least %d characters", t.BaseInput.Label, minVal)
		}
	case "max":
		maxVal, _ := strconv.Atoi(ruleValue)
		if len(value) > maxVal {
			return fmt.Sprintf("%s must not exceed %d characters", t.BaseInput.Label, maxVal)
		}
	case "regex":
		if matched, _ := regexp.MatchString(ruleValue, value); !matched && value != "" {
			return fmt.Sprintf("%s format is invalid", t.BaseInput.Label)
		}
	}

	return ""
}

// ToComponent converts TextInput to the original Component format for compatibility
func (t *TextInputField) ToComponent() Component {
	return Component{
		Type:        t.GetType(),
		Name:        t.BaseInput.Name,
		Label:       t.BaseInput.Label,
		Placeholder: t.BaseInput.Placeholder,
		Required:    t.BaseInput.Required,
		Value:       t.BaseInput.Value,
	}
}

// Convenience functions for creating common input types

// TextInput creates a new text input (factory function)
func TextInput(name, label string) *TextInputField {
	return NewTextInput(name, label)
}

// EmailInput creates a new email input (factory function)
func EmailInput(name, label string) *TextInputField {
	return NewTextInput(name, label).Email()
}

// PasswordInput creates a new password input
func PasswordInput(name, label string) *TextInputField {
	return NewTextInput(name, label).Password()
}

// UrlInput creates a new URL input
func UrlInput(name, label string) *TextInputField {
	return NewTextInput(name, label).Url()
}

// TelInput creates a new telephone input
func TelInput(name, label string) *TextInputField {
	return NewTextInput(name, label).Tel()
}

// NumericInput creates a new numeric input
func NumericInput(name, label string) *TextInputField {
	return NewTextInput(name, label).Numeric()
}

// IntegerInput creates a new integer input
func IntegerInput(name, label string) *TextInputField {
	return NewTextInput(name, label).Integer()
}