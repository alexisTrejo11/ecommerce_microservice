package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/http/v1/dto"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/jwt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authUserCase input.AuthUseCase
	emailUseCase input.EmailUseCase
	validator    *validator.Validate
	jwtManager   jwt.JWTManager
}

func NewAuthHandler(authUserCase input.AuthUseCase, jwtManager jwt.JWTManager, emailUseCase input.EmailUseCase) *AuthHandler {
	return &AuthHandler{
		authUserCase: authUserCase,
		validator:    validator.New(validator.WithRequiredStructEnabled()),
		jwtManager:   jwtManager,
		emailUseCase: emailUseCase,
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

	_, _, err := ah.authUserCase.Register(context.TODO(), signupDTO)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messsage": "can't create user",
			"errors":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"succes":  "singup succesfully proccesed",
		"message": "An email will be sended to activate your account",
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

func (ah *AuthHandler) Logout(c *fiber.Ctx) error {
	claims, err := ah.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "unauthorized",
			"message": err.Error(),
		})
	}

	refresh_token := c.Params("refresh_token")
	if refresh_token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "refresh_token is empty",
		})
	}

	userId, _ := uuid.Parse(claims.UserID)
	err = ah.authUserCase.Logout(context.TODO(), refresh_token, userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messsage": "can't logout token",
			"errors":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "logout succesfully proccesed",
	})
}

func (ah *AuthHandler) LogoutAll(c *fiber.Ctx) error {
	claims, err := ah.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "unauthorized",
			"message": err.Error(),
		})
	}

	userId, _ := uuid.Parse(claims.UserID)
	err = ah.authUserCase.LogoutAll(context.TODO(), userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messsage": "can't delete all sesions",
			"errors":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "all logout succesfully proccesed",
	})
}

func (ah *AuthHandler) RefreshAccesToken(c *fiber.Ctx) error {
	refresh_token := c.Params("refresh_token")
	if refresh_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "acces_token is empty",
		})
	}

	tokenDetails, err := ah.authUserCase.RefreshTokens(context.TODO(), refresh_token, "", "")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messsage": "can't create access token",
			"errors":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": tokenDetails,
	})
}

func (ah *AuthHandler) ActivateAccount(c *fiber.Ctx) error {
	token := c.Params("token")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "token is empty",
		})
	}

	err := ah.authUserCase.ActivateAccount(context.TODO(), token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messsage": "can't activate account",
			"errors":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Account successfully activated",
	})
}

func (ah *AuthHandler) ResendCode(c *fiber.Ctx) error {
	code_type := c.Params("code_type")
	if code_type == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "code_type is empty",
		})
	}

	userIdSTR := c.Params("user_id")
	if userIdSTR == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "token is empty",
		})
	}
	userId, err := uuid.Parse(userIdSTR)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "user_id is invalid",
		})
	}

	err = ah.authUserCase.ResendCode(context.TODO(), code_type, userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messsage": "can't resend code",
			"errors":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Code successfully resended",
	})
}

func (ah *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	token := c.Params("token")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "token is empty",
		})
	}

	newPassword := c.Params("new_password")
	if newPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "new Password is empty",
		})
	}

	err := ah.authUserCase.ResetPassword(context.TODO(), token, newPassword)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messsage": "can't reset password",
			"errors":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Code successfully resended",
	})
}
