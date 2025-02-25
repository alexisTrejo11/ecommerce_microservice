package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/jwt"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SessionHandler struct {
	sessionUseCase input.SessionUseCase
	jwtManager     jwt.JWTManager
}

func NewSessionHandler(sessionUseCase input.SessionUseCase, jwtManager jwt.JWTManager) *SessionHandler {
	return &SessionHandler{
		sessionUseCase: sessionUseCase,
		jwtManager:     jwtManager,
	}
}

func (ush *SessionHandler) GetSessionByUserId(c *fiber.Ctx) error {
	userIdSTR := c.Params("id")
	userId, err := uuid.Parse(userIdSTR)
	if err != nil {
		return response.BadRequest(c, "invalid user ID", err.Error())
	}

	sessions, err := ush.sessionUseCase.GetUserSessions(context.Background(), userId)
	if err != nil {
		return response.InternalServerError(c, "failed to retrieve user sessions", err.Error())
	}

	return response.OK(c, "User sessions retrieved successfully", sessions)
}

func (ush *SessionHandler) DeleteSessionById(c *fiber.Ctx) error {
	idSTR := c.Params("id")
	if idSTR == "" {
		return response.BadRequest(c, "session ID not provided", nil)
	}

	sessionId, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "invalid session ID", err.Error())
	}

	userIdSTR := c.Params("user_id")
	userId, err := uuid.Parse(userIdSTR)
	if err != nil {
		return response.BadRequest(c, "invalid user ID", err.Error())
	}

	err = ush.sessionUseCase.DeleteSession(context.Background(), sessionId, userId)
	if err != nil {
		return response.NotFound(c, "session not found", err.Error())
	}

	return response.OK(c, "Session successfully deleted", nil)
}
