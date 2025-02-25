package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/jwt"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/response"
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
		return response.Unauthorized(c, "unauthorized", err.Error())
	}

	userId, err := uuid.Parse(claims.UserID)
	if err != nil {
		return response.BadRequest(c, "invalid user ID", err.Error())
	}

	secret, QRCodePath, err := u.MFAUseCase.SetupMFA(context.Background(), userId)
	if err != nil {
		return response.InternalServerError(c, "failed to enable MFA", err.Error())
	}

	return response.OK(c, "MFA enabled successfully", fiber.Map{
		"secret":  secret,
		"qr_code": QRCodePath,
	})
}

func (u *UserMfaHandler) DisableMfa(c *fiber.Ctx) error {
	claims, err := u.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return response.Unauthorized(c, "unauthorized", err.Error())
	}

	userId, err := uuid.Parse(claims.UserID)
	if err != nil {
		return response.BadRequest(c, "invalid user ID", err.Error())
	}

	if err := u.MFAUseCase.DisableMFA(context.Background(), userId, ""); err != nil {
		return response.InternalServerError(c, "failed to disable MFA", err.Error())
	}

	return response.OK(c, "MFA disabled successfully", nil)
}

func (u *UserMfaHandler) VerifyMfa(c *fiber.Ctx) error {
	claims, err := u.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return response.Unauthorized(c, "unauthorized", err.Error())
	}

	code := c.Query("code")
	if code == "" {
		return response.BadRequest(c, "code is required", nil)
	}

	userId, err := uuid.Parse(claims.UserID)
	if err != nil {
		return response.BadRequest(c, "invalid user ID", err.Error())
	}

	if err := u.MFAUseCase.DisableMFA(context.Background(), userId, code); err != nil {
		return response.BadRequest(c, "failed to disable MFA", err.Error())
	}

	return response.OK(c, "MFA successfully disabled", nil)
}

func (u *UserMfaHandler) GetMfa(c *fiber.Ctx) error {
	claims, err := u.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return response.Unauthorized(c, "unauthorized", err.Error())
	}

	userId, err := uuid.Parse(claims.UserID)
	if err != nil {
		return response.BadRequest(c, "invalid user ID", err.Error())
	}

	mfa, err := u.MFAUseCase.GetMFA(context.Background(), userId)
	if err != nil {
		return response.NotFound(c, "MFA not found", err.Error())
	}

	return response.OK(c, "MFA retrieved successfully", mfa)
}
