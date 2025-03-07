package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/application/ports/input"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CartHandler struct {
	cartUseCase input.CartUseCase
}

func NewCartHandler(cartUseCase input.CartUseCase) *CartHandler {
	return &CartHandler{cartUseCase: cartUseCase}

}

func (h *CartHandler) InitCart(c *fiber.Ctx) error {
	userIdStr := c.Params("userId")
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid user id"})
	}

	err = h.cartUseCase.CreateCart(context.Background(), userId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON("Cart Successfully Init for New User")
}

func (h *CartHandler) GetCartByUserId(c *fiber.Ctx) error {
	userIdStr := c.Params("userId")
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid user id"})
	}

	cart, err := h.cartUseCase.GetCartByUserId(context.Background(), userId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(cart)
}

func (h *CartHandler) GetCartById(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid user id"})
	}

	cart, err := h.cartUseCase.GetCartById(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(cart)
}

func (h *CartHandler) DeleteCart(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid user id"})
	}

	err = h.cartUseCase.DeleteCart(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON("Cart Successfully Deleted")
}
