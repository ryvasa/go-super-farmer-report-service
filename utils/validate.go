package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type errorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(data interface{}) []*errorResponse {
	var errors []*errorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &errorResponse{
				Field:   err.Field(),
				Message: getErrorMsg(err),
			})
		}
	}
	return errors
}

func getErrorMsg(fe validator.FieldError) string {
	msgMap := map[string]func(string) string{
		"required": func(param string) string { return "This field is required" },
		"min":      func(param string) string { return fmt.Sprintf("Minimum length is %s", param) },
		"max":      func(param string) string { return fmt.Sprintf("Maximum length is %s", param) },
		"gt":       func(param string) string { return fmt.Sprintf("The value must be greater than %s", param) },
	}
	msgFunc, ok := msgMap[fe.Tag()]
	if ok {
		return msgFunc(fe.Param())
	}
	return "Unknown error"
}
