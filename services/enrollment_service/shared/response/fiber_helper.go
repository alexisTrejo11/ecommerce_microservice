package response

import (
	"fmt"

	logging "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/logger"
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

func GetUUIDParam(c *fiber.Ctx, paramName, action string) (uuid.UUID, error) {
	idSTR := c.Params(paramName)
	if idSTR == "" {
		err := fmt.Errorf("%s is mandatory", paramName)
		logging.LogError(action, "Invalid  ID", map[string]interface{}{
			"error": err.Error(),
		})

		return uuid.Nil, err
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		err := fmt.Errorf("invalid %s", paramName)

		logging.LogError(action, "Invalid  ID", map[string]interface{}{
			"error": err.Error(),
		})

		return uuid.Nil, err
	}

	return id, nil
}
