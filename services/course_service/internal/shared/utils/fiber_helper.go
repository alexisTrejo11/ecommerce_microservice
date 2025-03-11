package utils

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(v *validator.Validate, dto interface{}) (map[string]string, error) {
	if err := v.Struct(dto); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorsMap := make(map[string]string)
			for _, fieldErr := range validationErrors {
				errorsMap[fieldErr.Field()] = fieldErr.Tag()
			}

			return errorsMap, err
		}
		return nil, err
	}

	return nil, nil
}
