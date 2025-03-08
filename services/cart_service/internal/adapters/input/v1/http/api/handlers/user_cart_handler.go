package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/core/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/shared/dtos"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserCartHandler struct {
	cartUseCase input.CartUseCase
}

func NewUserCartHandler(cartUseCase input.CartUseCase) *UserCartHandler {
	return &UserCartHandler{cartUseCase: cartUseCase}

}

// TODO Import or Whatever to get User Id from JWT
func (h *UserCartHandler) GetMyCart(c *fiber.Ctx) error {
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

func (h *UserCartHandler) AddItems(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userId, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid user id"})
	}

	var insertDTO []dtos.CartItemInserDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid items ids"})
	}

	err = h.cartUseCase.AddItems(context.Background(), userId, insertDTO)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON("Items removed")
}

func (h *UserCartHandler) RemoveItems(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userId, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid user id"})
	}

	var itemsIds []uuid.UUID
	if err := c.BodyParser(&itemsIds); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid items ids"})
	}

	err = h.cartUseCase.RemoveItems(context.Background(), userId, itemsIds)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON("Items removed")
}

func (h *UserCartHandler) Buy(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userId, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid user id"})
	}

	// Allow empty
	var exludeItemsIds []*uuid.UUID
	if err := c.BodyParser(&exludeItemsIds); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid items ids"})
	}

	err = h.cartUseCase.Buy(context.Background(), userId, exludeItemsIds)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON("Cart Operation for Buy Completed")
}

// Implement this
func (h *UserCartHandler) BuyProduct(c *fiber.Ctx) error {
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
