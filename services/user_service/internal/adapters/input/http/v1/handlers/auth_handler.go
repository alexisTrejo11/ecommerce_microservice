package handlers

import (
	"context"
	"log"

	_ "github.com/alexisTrejo11/ecommerce_microservice/docs"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/http/v1/dto"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/jwt"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/response"
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

// Register handles user registration by validating input, creating a user, and sending a verification email.
// @Summary Register a new user
// @Description Registers a new user and sends an email with a verification token
// @Tags Auth
// @Accept json
// @Produce json
// @Param signup body dto.SignupDTO true "User registration data"
// @Success 201 {object} response.ApiResponse "Signup successful message"
// @Failure 400 {object} response.ApiResponse "Bad request"
// @Router /v1/api/register [post]
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

	// Return the success response
	return response.Created(c, "Signup successfully processed", map[string]string{
		"message": "An email will be sent to activate your account",
	})
}

// Login handles user login by validating credentials and issuing a JWT token.
// @Summary User login
// @Description Authenticates a user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body dto.LoginDTO true "Login credentials"
// @Success 200 {object} response.ApiResponse "Login successful"
// @Failure 400 {object} response.ApiResponse "Bad request"
// @Failure 401 {object} response.ApiResponse "Invalid credentials"
// @Router /v1/api/login [post]
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

// Logout handles user logout and invalidates the session.
// @Summary Logout user
// @Description Logs the user out and invalidates the session
// @Tags Auth
// @Accept json
// @Produce json
// @Param refresh_token path string true "Refresh token"
// @Success 200 {object} response.ApiResponse "Logout successful"
// @Failure 400 {object} response.ApiResponse "Bad request"
// @Failure 401 {object} response.ApiResponse "Unauthorized"
// @Router /v1/api/logout/{refresh_token} [post]
func (ah *AuthHandler) Logout(c *fiber.Ctx) error {
	claims, err := ah.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return response.Unauthorized(c, "Unauthorized", err.Error())
	}

	refreshToken := c.Params("refresh_token")
	if refreshToken == "" {
		return response.BadRequest(c, "Refresh token is empty", nil)
	}

	userID, _ := uuid.Parse(claims.UserID)

	err = ah.authUserCase.Logout(context.TODO(), refreshToken, userID)
	if err != nil {
		return response.BadRequest(c, "Unable to logout", err.Error())
	}

	return response.OK(c, "Logout successfully processed", nil)
}

// LogoutAll handles logging out all sessions for a user.
// @Summary Logout all sessions
// @Description Logs out all active sessions for the user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponse "All sessions logged out"
// @Failure 400 {object} response.ApiResponse "Bad request"
// @Failure 401 {object} response.ApiResponse "Unauthorized"
// @Router /v1/api/logout-all [post]
func (ah *AuthHandler) LogoutAll(c *fiber.Ctx) error {
	claims, err := ah.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return response.Unauthorized(c, "Unauthorized", err.Error())
	}

	userID, _ := uuid.Parse(claims.UserID)

	err = ah.authUserCase.LogoutAll(context.TODO(), userID)
	if err != nil {
		return response.BadRequest(c, "Unable to delete all sessions", err.Error())
	}

	return response.OK(c, "All sessions successfully logged out", nil)
}

// RefreshAccessToken handles the refresh of an access token using the provided refresh token.
// @Summary Refresh access token
// @Description Refreshes the access token using a valid refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param refresh_token path string true "Refresh token"
// @Success 201 {object} response.ApiResponse "New access token details"
// @Failure 400 {object} response.ApiResponse "Bad request"
// @Failure 401 {object} response.ApiResponse "Unauthorized"
// @Router /v1/api/refresh-acces-token/{refresh_token} [get]
func (ah *AuthHandler) RefreshAccessToken(c *fiber.Ctx) error {
	refreshToken := c.Params("refresh_token")
	if refreshToken == "" {
		return response.Unauthorized(c, "Access token is empty", nil)
	}

	tokenDetails, err := ah.authUserCase.RefreshTokens(context.TODO(), refreshToken, "", "")
	if err != nil {
		return response.BadRequest(c, "Unable to create access token", err.Error())
	}

	return response.Created(c, "Access token successfully refreshed", tokenDetails)
}

// ActivateAccount handles the activation of a user account using a verification token.
// @Summary Activate user account
// @Description Activates the user account using a valid verification token
// @Tags Auth
// @Accept json
// @Produce json
// @Param token path string true "Verification token"
// @Success 200 {object} response.ApiResponse "Account activated successfully"
// @Failure 400 {object} response.ApiResponse "Bad request"
// @Failure 401 {object} response.ApiResponse "Unauthorized"
// @Router /v1/api/activate-account/{token} [post]
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

// ResendCode handles resending a verification or password reset code.
// @Summary Resend verification or reset code
// @Description Resends a verification or reset code to the user
// @Tags Auth
// @Accept json
// @Produce json
// @Param code_type path string true "Type of code to resend (e.g., verification, reset)"
// @Param user_id path string true "User ID"
// @Success 200 {object} response.ApiResponse "Code resent successfully"
// @Failure 400 {object} response.ApiResponse "Bad request"
// @Failure 401 {object} response.ApiResponse "Unauthorized"
// @Router /v1/api/resend-code/{code_type} [post]
func (ah *AuthHandler) ResendCode(c *fiber.Ctx) error {
	codeType := c.Params("code_type")
	if codeType == "" {
		return response.BadRequest(c, "Code type is empty", nil)
	}

	userIDStr := c.Params("user_id")
	if userIDStr == "" {
		return response.BadRequest(c, "User ID is empty", nil)
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return response.BadRequest(c, "User ID is invalid", err.Error())
	}

	err = ah.authUserCase.ResendCode(context.TODO(), codeType, userID)
	if err != nil {
		return response.BadRequest(c, "Unable to resend code", err.Error())
	}

	return response.OK(c, "Code successfully resent", nil)
}

// ResetPassword handles resetting a user's password using a reset token.
// @Summary Reset user password
// @Description Resets the user's password using a valid reset token
// @Tags Auth
// @Accept json
// @Produce json
// @Param token path string true "Password reset token"
// @Param new_password path string true "New password"
// @Success 200 {object} response.ApiResponse "Password reset successfully"
// @Failure 400 {object} response.ApiResponse "Bad request"
// @Failure 401 {object} response.ApiResponse "Unauthorized"
// @Router /v1/api/reset-password/{token} [post]
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
