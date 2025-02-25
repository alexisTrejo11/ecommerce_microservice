package handlers

import (
	"context"
	"log"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/http/v1/dto"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/jwt"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/response" // Import the new package
	"github.com/alexisTrejo11/ecommerce_microservice/pkg/rabbitmq"
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
		return response.BadRequest(c, "Unable to parse request body", err.Error())
	}

	if err := ah.validator.Struct(&signupDTO); err != nil {
		return response.BadRequest(c, "Data validation failed", err.Error())
	}

	user, verificationToken, err := ah.authUserCase.Register(context.TODO(), signupDTO)
	if err != nil {
		return response.BadRequest(c, "Unable to create user", err.Error())
	}

	conn, err := rabbitmq.ConnectRabbitMQ()
	if err != nil {
		log.Println("Error connecting to RabbitMQ:", err)
	} else {
		defer conn.Close()

		message := rabbitmq.EmailMessage{
			UserID:            user.ID,
			VerificationToken: verificationToken,
		}
		err = rabbitmq.PublishMessage(context.Background(), conn, "email_queue", message)
		if err != nil {
			log.Println("Error publishing message:", err)
		}
	}

	return response.Created(c, "Signup successfully processed", map[string]string{
		"message": "An email will be sent to activate your account",
	})
}

func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	var loginDTO dto.LoginDTO

	if err := c.BodyParser(&loginDTO); err != nil {
		return response.BadRequest(c, "Unable to parse request body", err.Error())
	}

	if err := ah.validator.Struct(&loginDTO); err != nil {
		return response.BadRequest(c, "Data validation failed", err.Error())
	}

	tokenDetails, err := ah.authUserCase.Login(context.TODO(), loginDTO)
	if err != nil {
		return response.Unauthorized(c, "Invalid login credentials", err.Error())
	}

	return response.OK(c, "Login successfully processed", tokenDetails)
}

func (ah *AuthHandler) Logout(c *fiber.Ctx) error {
	claims, err := ah.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return response.Unauthorized(c, "Unauthorized", err.Error())
	}

	refresh_token := c.Params("refresh_token")
	if refresh_token == "" {
		return response.BadRequest(c, "Refresh token is empty", nil)
	}

	userId, _ := uuid.Parse(claims.UserID)
	err = ah.authUserCase.Logout(context.TODO(), refresh_token, userId)
	if err != nil {
		return response.BadRequest(c, "Unable to logout", err.Error())
	}

	return response.OK(c, "Logout successfully processed", nil)
}

func (ah *AuthHandler) LogoutAll(c *fiber.Ctx) error {
	claims, err := ah.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return response.Unauthorized(c, "Unauthorized", err.Error())
	}

	userId, _ := uuid.Parse(claims.UserID)
	err = ah.authUserCase.LogoutAll(context.TODO(), userId)
	if err != nil {
		return response.BadRequest(c, "Unable to delete all sessions", err.Error())
	}

	return response.OK(c, "All sessions successfully logged out", nil)
}

func (ah *AuthHandler) RefreshAccesToken(c *fiber.Ctx) error {
	refresh_token := c.Params("refresh_token")
	if refresh_token == "" {
		return response.Unauthorized(c, "Access token is empty", nil)
	}

	tokenDetails, err := ah.authUserCase.RefreshTokens(context.TODO(), refresh_token, "", "")
	if err != nil {
		return response.BadRequest(c, "Unable to create access token", err.Error())
	}

	return response.Created(c, "Access token successfully refreshed", tokenDetails)
}

func (ah *AuthHandler) ActivateAccount(c *fiber.Ctx) error {
	token := c.Params("token")
	if token == "" {
		return response.BadRequest(c, "Token is empty", nil)
	}

	err := ah.authUserCase.ActivateAccount(context.TODO(), token)
	if err != nil {
		return response.BadRequest(c, "Unable to activate account", err.Error())
	}

	return response.OK(c, "Account successfully activated", nil)
}

func (ah *AuthHandler) ResendCode(c *fiber.Ctx) error {
	code_type := c.Params("code_type")
	if code_type == "" {
		return response.BadRequest(c, "Code type is empty", nil)
	}

	userIdSTR := c.Params("user_id")
	if userIdSTR == "" {
		return response.BadRequest(c, "User ID is empty", nil)
	}

	userId, err := uuid.Parse(userIdSTR)
	if err != nil {
		return response.BadRequest(c, "User ID is invalid", err.Error())
	}

	err = ah.authUserCase.ResendCode(context.TODO(), code_type, userId)
	if err != nil {
		return response.BadRequest(c, "Unable to resend code", err.Error())
	}

	return response.OK(c, "Code successfully resent", nil)
}

func (ah *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	token := c.Params("token")
	if token == "" {
		return response.BadRequest(c, "Token is empty", nil)
	}

	newPassword := c.Params("new_password")
	if newPassword == "" {
		return response.BadRequest(c, "New password is empty", nil)
	}

	err := ah.authUserCase.ResetPassword(context.TODO(), token, newPassword)
	if err != nil {
		return response.BadRequest(c, "Unable to reset password", err.Error())
	}

	return response.OK(c, "Password successfully reset", nil)
}
