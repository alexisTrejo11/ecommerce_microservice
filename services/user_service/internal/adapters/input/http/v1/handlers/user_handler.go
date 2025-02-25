package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	userUseCase input.UserUseCase
}

func NewUserHandler(userUseCase input.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (uh *UserHandler) GetUserById(c *fiber.Ctx) error {
	userUUID := c.Query("id")
	if userUUID == "" {
		return response.BadRequest(c, "user ID not provided", nil)
	}

	userID, err := uuid.Parse(userUUID)
	if err != nil {
		return response.BadRequest(c, "invalid user ID", err.Error())
	}

	user, err := uh.userUseCase.GetUser(context.Background(), userID)
	if err != nil {
		return response.NotFound(c, "user not found", err.Error())
	}

	return response.OK(c, "User retrieved successfully", user)
}

func (uh *UserHandler) DeleteUserById(c *fiber.Ctx) error {
	userUUID := c.Query("id")
	if userUUID == "" {
		return response.BadRequest(c, "user ID not provided", nil)
	}

	userID, err := uuid.Parse(userUUID)
	if err != nil {
		return response.BadRequest(c, "invalid user ID", err.Error())
	}

	if err := uh.userUseCase.DeleteUser(context.Background(), userID); err != nil {
		return response.NotFound(c, "user not found", err.Error())
	}

	return response.OK(c, "User successfully deleted", nil)
}
