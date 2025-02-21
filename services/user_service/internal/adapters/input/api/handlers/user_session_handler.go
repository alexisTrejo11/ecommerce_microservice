package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/pkg/jwt"
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
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	sessions, err := ush.sessionUseCase.GetUserSessions(context.Background(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(sessions)
}

func (ush *SessionHandler) DeleteSessionById(c *fiber.Ctx) error {
	idSTR := c.Params("id")
	if idSTR == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "session id not provided"})
	}
	sessonId, _ := uuid.Parse(idSTR)

	userIdSTR := c.Params("user_id")
	userId, err := uuid.Parse(userIdSTR)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "user id not valid"})
	}

	err = ush.sessionUseCase.DeleteSession(context.Background(), sessonId, userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"succession": "session successfully deleted"})
}
