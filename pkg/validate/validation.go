package validate

import (
	"fmt"
	"regexp"
	"strings"

	validator "github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var v = validator.New()

func init() {
	RegisterCustomValidation("kenyan_phone", func(fl validator.FieldLevel) bool {
		phone := fl.Field().String()

		matched, _ := regexp.MatchString(`^(\+254|0)\d{9}$`, phone)

		return matched
	})
}

func ValidateStruct(s interface{}) []FieldError {
	err := v.Struct(s)

	if err == nil {
		return nil
	}

	var errors []FieldError

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, FieldError{
			Field:   toSnakeCase(e.Field()),
			Message: buildMessage(e),
		})
	}

	return errors
}

func RegisterCustomValidation(tag string, fn validator.Func) error {
	return v.RegisterValidation(tag, fn)
}

// buildMessage creates human-readable error messages per validation tag
func buildMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", e.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", e.Field(), e.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", e.Field(), e.Param())
	case "kenyan_phone":
		return fmt.Sprintf("%s must be a valid Kenyan phone number (e.g. +254XXXXXXXXX or 0XXXXXXXXX)", e.Field())
	default:
		return fmt.Sprintf("%s failed on %s validation", e.Field(), e.Tag())
	}
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
