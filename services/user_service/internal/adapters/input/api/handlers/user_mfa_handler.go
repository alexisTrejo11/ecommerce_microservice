package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserMfaHandler struct {
	MFAUseCase input.MFAUseCase
	jwtManager jwt.JWTManager
}

func NewUserMfaHandler(mfaUseCase input.MFAUseCase, jwtManager jwt.JWTManager) *UserMfaHandler {
	return &UserMfaHandler{
		MFAUseCase: mfaUseCase,
		jwtManager: jwtManager,
	}
}

func (u *UserMfaHandler) EnableMfa(c *fiber.Ctx) error {
	claims, err := u.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "unauthorized",
			"message": err.Error(),
		})
	}

	userId, _ := uuid.Parse(claims.UserID)
	secret, QRCodePath, err := u.MFAUseCase.SetupMFA(context.Background(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "internal server error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "MFA enabled",
		"secret":  secret,
		"qr_code": QRCodePath,
	})
}

func (u *UserMfaHandler) DisableMfa(c *fiber.Ctx) error {
	claims, err := u.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "unauthorized",
			"message": err.Error(),
		})
	}

	userId, _ := uuid.Parse(claims.UserID)
	err = u.MFAUseCase.DisableMFA(context.Background(), userId, "")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "internal server error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "MFA disabled",
	})
}

func (u *UserMfaHandler) VerifyMfa(c *fiber.Ctx) error {
	claims, err := u.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "unauthorized",
			"message": err.Error(),
		})
	}

	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad request",
			"message": "code is required",
		})
	}

	userId, _ := uuid.Parse(claims.UserID)
	err = u.MFAUseCase.DisableMFA(context.Background(), userId, code)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "can't disable MFA",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "MFA succesfully disabled",
	})

}

func (u *UserMfaHandler) GetMfa(c *fiber.Ctx) error {
	claims, err := u.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "unauthorized",
			"message": err.Error(),
		})
	}

	userId, _ := uuid.Parse(claims.UserID)
	mfa, err := u.MFAUseCase.GetMFA(context.Background(), userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(mfa)
}
