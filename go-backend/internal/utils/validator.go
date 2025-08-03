package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, formatValidationError(err))
		}
		return fmt.Errorf(strings.Join(errors, ", "))
	}
	return nil
}

func formatValidationError(err validator.FieldError) string {
	field := err.Field()
	tag := err.Tag()

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s is not valid format", field)
	case "max":
		return fmt.Sprintf("%s length max %s", field, err.Param())
	case "min":
		return fmt.Sprintf("%s must be at least %s", field, err.Param())
	default:
		return fmt.Sprintf("%s is not valid", field)
	}
}