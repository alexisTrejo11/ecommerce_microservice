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

// GetSessionByUserId retrieves all sessions of a user.
// @Summary Get user sessions
// @Description Retrieve all active sessions for a given user ID
// @Tags sessions
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.ApiResponse "User sessions retrieved successfully"
// @Failure 400 {object} response.ApiResponse "Invalid user ID"
// @Failure 500 {object} response.ApiResponse "Failed to retrieve user sessions"
// @Router /v1/api/sessions/{id} [get]
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

// DeleteSessionById deletes a specific user session.
// @Summary Delete user session
// @Description Delete a session based on session ID and user ID
// @Tags sessions
// @Accept json
// @Produce json
// @Param id path string true "Session ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} response.ApiResponse "Session successfully deleted"
// @Failure 400 {object} response.ApiResponse "Invalid session ID or user ID"
// @Failure 404 {object} response.ApiResponse "Session not found"
// @Router /v1/api/sessions/{id}/user/{user_id} [delete]
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
