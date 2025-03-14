package utils

import (
	"errors"
	"fmt"
	"strconv"

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

type Page struct {
	PageNumber int
	PageSize   int
}

func GetPageData(c *fiber.Ctx) (*Page, error) {
	size := c.Query("page_size", "10")
	number := c.Query("page_number", "1")

	pageSize, err := strconv.Atoi(size)
	if err != nil || pageSize <= 0 {
		return nil, errors.New("invalid page size")
	}

	pageNumber, err := strconv.Atoi(number)
	if err != nil || pageNumber <= 0 {
		return nil, errors.New("invalid page number")
	}

	return &Page{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}, nil
}
