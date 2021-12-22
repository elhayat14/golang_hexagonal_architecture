package validator

import (
	"fmt"
	goValidator "github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *goValidator.Validate
}

func (v *Validator) Validate(schema interface{}) error {
	return v.Validator.Struct(schema)
}

type ErrorInfo struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type Error struct {
	Field []ErrorInfo `json:"field"`
}

func BuildErrorBodyRequestValidator(err error) Error {
	var errors Error
	for _, err := range err.(goValidator.ValidationErrors) {
		message := ErrorString(err.Tag(), err.Field(), err.Param())
		errorInfo := ErrorInfo{Name: err.Field(), Message: message}
		errors.Field = append(errors.Field, errorInfo)
	}
	return errors
}
func ErrorString(tag string, field string, param string) string {
	msgError := ""
	switch tag {
	case "required":
		msgError = fmt.Sprintf("Value is required")
	case "email":
		msgError = fmt.Sprintf("%s is not valid email", field)
	case "min":
		msgError = fmt.Sprintf("%s must be greater than %s", field, param)
	case "max":
		msgError = fmt.Sprintf("%s must be lower than %s", field, param)
	case "numeric":
		msgError = fmt.Sprintf("Param has illegal value")
	}
	return msgError
}
