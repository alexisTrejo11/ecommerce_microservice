package response

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func GetUUIDParam(c *fiber.Ctx, paramName string) (uuid.UUID, error) {
	idSTR := c.Params(paramName)
	if idSTR == "" {
		return uuid.Nil, fmt.Errorf("%s is mandatory", paramName)
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid %s", paramName)
	}

	return id, nil
}
