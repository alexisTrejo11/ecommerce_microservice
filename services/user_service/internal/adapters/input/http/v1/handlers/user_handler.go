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

// GetUserById retrieves a user by ID.
// @Summary Get user by ID
// @Description Retrieve a user based on the provided ID
// @Tags users
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} response.ApiResponse "User retrieved successfully"
// @Failure 400 {object} response.ApiResponse "Invalid user ID"
// @Failure 404 {object} response.ApiResponse "User not found"
// @Router /v1/api/users [get]
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

// DeleteUserById deletes a user by ID.
// @Summary Delete user by ID
// @Description Delete a user based on the provided ID
// @Tags users
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} response.ApiResponse "User successfully deleted"
// @Failure 400 {object} response.ApiResponse "Invalid user ID"
// @Failure 404 {object} response.ApiResponse "User not found"
// @Router /v1/api/users [delete]
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
