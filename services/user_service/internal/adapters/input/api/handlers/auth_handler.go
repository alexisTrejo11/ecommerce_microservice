package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/api/dto"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUserCase input.AuthUseCase
	validator    *validator.Validate
}

func NewAuthHandler(authUserCase input.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUserCase: authUserCase,
		validator:    validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (ah *AuthHandler) Register(c *fiber.Ctx) error {
	var signupDTO dto.SignupDTO

	if err := c.BodyParser(&signupDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "can't parse request body",
		})
	}

	if err := ah.validator.Struct(&signupDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messsage": "data valdiation failed",
			"errors":   err.Error(),
		})
	}

	user, err := ah.authUserCase.Register(context.TODO(), signupDTO)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messsage": "can't create user",
			"errors":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"created": "user_created",
		"data":    user,
	})
}

func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	var loginDTO dto.LoginDTO

	if err := c.BodyParser(&loginDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "can't parse request body",
		})
	}

	if err := ah.validator.Struct(&loginDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messsage": "data valdiation failed",
			"errors":   err.Error(),
		})
	}

	tokenDetails, err := ah.authUserCase.Login(context.TODO(), loginDTO)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messsage": "invalid login credentials",
			"errors":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "login succesfully proccesed",
		"data":    tokenDetails,
	})
}
