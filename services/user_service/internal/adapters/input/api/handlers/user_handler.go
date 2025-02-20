package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
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

	user, err := uh.userUseCase.GetUser(context.Background(), uuid.MustParse(userUUID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)

	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (uh *UserHandler) DeleteUserById(c *fiber.Ctx) error {
	userUUID := c.Query("id")

	if err := uh.userUseCase.DeleteUser(context.Background(), uuid.MustParse(userUUID)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"success": "user delted"})
}
