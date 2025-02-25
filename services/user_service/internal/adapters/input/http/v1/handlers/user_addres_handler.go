package handlers

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/http/v1/dto"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/jwt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserAddressHandler struct {
	addressUseCase input.AddressUseCase
	validator      *validator.Validate
	jwtManager     jwt.JWTManager
}

func NewUserAddressHandler(addressUseCase input.AddressUseCase, jwtManager jwt.JWTManager) *UserAddressHandler {
	return &UserAddressHandler{
		addressUseCase: addressUseCase,
		jwtManager:     jwtManager,
		validator:      validator.New(),
	}
}

func (uah *UserAddressHandler) MyAddresses(c *fiber.Ctx) error {
	claims, err := uah.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "unauthorized",
			"message": err.Error(),
		})
	}

	userId, _ := uuid.Parse(claims.UserID)

	addresses, err := uah.addressUseCase.GetUserAddresses(context.Background(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "internal server error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(addresses)
}

func (uah *UserAddressHandler) AddAddress(c *fiber.Ctx) error {
	claims, err := uah.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "unauthorized",
			"message": err.Error(),
		})
	}

	var addressDTO dto.AddressInsertDTO

	if err := c.BodyParser(&addressDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "can't parse request body",
		})
	}

	if err := uah.validator.Struct(&addressDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "data validation failed",
			"errors":  err.Error(),
		})
	}

	userId, _ := uuid.Parse(claims.UserID)
	addressDTO.UserID = userId
	err = uah.addressUseCase.AddAddress(context.Background(), &addressDTO)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid data",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Address successfully created",
	})
}

func (uah *UserAddressHandler) UpdateMyAddress(c *fiber.Ctx) error {
	claims, err := uah.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "unauthorized",
			"message": err.Error(),
		})
	}

	var addressDTO dto.AddressInsertDTO

	if err := c.BodyParser(&addressDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "can't parse request body",
		})
	}

	if err := uah.validator.Struct(&addressDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "data validation failed",
			"errors":  err.Error(),
		})
	}

	addressID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid address ID",
		})
	}

	userId, err := uuid.Parse(claims.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	addressDTO.UserID = userId
	err = uah.addressUseCase.UpdateAddress(context.Background(), uint(addressID), &addressDTO)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid data",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Address successfully updated",
	})
}

func (uah *UserAddressHandler) DeleteAddress(c *fiber.Ctx) error {
	claims, err := uah.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "unauthorized",
			"message": err.Error(),
		})
	}

	addressID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid address ID",
		})
	}

	userId, _ := uuid.Parse(claims.UserID)
	err = uah.addressUseCase.DeleteAddress(context.Background(), uint(addressID), userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid data",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Address successfully deleted",
	})
}
