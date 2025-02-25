package handlers

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/http/v1/dto"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/jwt"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/response"
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
		return response.Unauthorized(c, "unauthorized", err.Error())
	}

	userId, _ := uuid.Parse(claims.UserID)

	addresses, err := uah.addressUseCase.GetUserAddresses(context.Background(), userId)
	if err != nil {
		return response.InternalServerError(c, "internal server error", err.Error())
	}

	return response.OK(c, "Addresses retrieved successfully", addresses)
}

func (uah *UserAddressHandler) AddAddress(c *fiber.Ctx) error {
	claims, err := uah.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return response.Unauthorized(c, "unauthorized", err.Error())
	}

	var addressDTO dto.AddressInsertDTO

	if err := c.BodyParser(&addressDTO); err != nil {
		return response.BadRequest(c, "can't parse request body", nil)
	}

	if err := uah.validator.Struct(&addressDTO); err != nil {
		return response.BadRequest(c, "data validation failed", err.Error())
	}

	userId, _ := uuid.Parse(claims.UserID)
	addressDTO.UserID = userId
	err = uah.addressUseCase.AddAddress(context.Background(), &addressDTO)
	if err != nil {
		return response.BadRequest(c, "invalid data", err.Error())
	}

	return response.Created(c, "Address successfully created", nil)
}

func (uah *UserAddressHandler) UpdateMyAddress(c *fiber.Ctx) error {
	claims, err := uah.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return response.Unauthorized(c, "unauthorized", err.Error())
	}

	var addressDTO dto.AddressInsertDTO

	if err := c.BodyParser(&addressDTO); err != nil {
		return response.BadRequest(c, "can't parse request body", nil)
	}

	if err := uah.validator.Struct(&addressDTO); err != nil {
		return response.BadRequest(c, "data validation failed", err.Error())
	}

	addressID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "invalid address ID", nil)
	}

	userId, err := uuid.Parse(claims.UserID)
	if err != nil {
		return response.BadRequest(c, "invalid user ID", err.Error())
	}

	addressDTO.UserID = userId
	err = uah.addressUseCase.UpdateAddress(context.Background(), uint(addressID), &addressDTO)
	if err != nil {
		return response.BadRequest(c, "invalid data", err.Error())
	}

	return response.OK(c, "Address successfully updated", nil)
}

func (uah *UserAddressHandler) DeleteAddress(c *fiber.Ctx) error {
	claims, err := uah.jwtManager.ExtractAndValidateToken(c)
	if err != nil {
		return response.Unauthorized(c, "unauthorized", err.Error())
	}

	addressID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "invalid address ID", nil)
	}

	userId, _ := uuid.Parse(claims.UserID)
	err = uah.addressUseCase.DeleteAddress(context.Background(), uint(addressID), userId)
	if err != nil {
		return response.BadRequest(c, "invalid data", err.Error())
	}

	return response.OK(c, "Address successfully deleted", nil)
}
