package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("no_space", noSpace)

	return &Validator{
		validate: validate,
	}
}

func (v *Validator) Validate(data any) map[string]string {
	err := v.validate.Struct(data)
	if err == nil {
		return nil
	}

	validationErrs := make(map[string]string, 0)
	for _, v := range err.(validator.ValidationErrors) {
		var e error
		switch v.Tag() {
		case "required":
			e = fmt.Errorf("Field '%s' cannot be empty", v.Field())
		case "email":
			e = fmt.Errorf("Field '%s' must be a valid email address", v.Field())
		case "len":
			e = fmt.Errorf("Field '%s' must be exactly %v characters long", v.Field(), v.Param())
		case "min":
			e = fmt.Errorf("Field '%s' must at least '%v' characters long", v.Field(), v.Param())
		case "max":
			e = fmt.Errorf("Field '%s' must not exceed '%v' characters long", v.Field(), v.Param())
		case "no_space":
			e = fmt.Errorf("Field '%s' must not contains space character", v.Field())
		default:
			e = fmt.Errorf("Field '%s' must satisfy '%s' '%v' criteria", v.Field(), v.Tag(), v.Param())
		}

		validationErrs[strings.ToLower(v.Field())] = e.Error()
	}

	return validationErrs
}

// INFO: Custom Validation
func noSpace(fl validator.FieldLevel) bool {
	return !strings.ContainsRune(fl.Field().String(), ' ')
}
