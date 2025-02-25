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

// MyAddresses retrieves the addresses of the authenticated user.
// @Summary Get user addresses
// @Description Retrieves a list of addresses associated with the authenticated user
// @Tags Addresses
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponse "Addresses retrieved successfully"
// @Failure 401 {object} response.ApiResponse "Unauthorized"
// @Router /v1/api/users/addresses [get]
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

// AddAddress adds a new address for the authenticated user.
// @Summary Add user address
// @Description Adds a new address for the authenticated user
// @Tags Addresses
// @Accept json
// @Produce json
// @Param address body dto.AddressInsertDTO true "Address details"
// @Success 201 {object} response.ApiResponse "Address successfully created"
// @Failure 400 {object} response.ApiResponse "Bad request"
// @Failure 401 {object} response.ApiResponse "Unauthorized"
// @Router /v1/api/users/addresses [post]
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

// UpdateMyAddress updates an existing address of the authenticated user.
// @Summary Update user address
// @Description Updates an existing address for the authenticated user
// @Tags Addresses
// @Accept json
// @Produce json
// @Param id path int true "Address ID"
// @Param address body dto.AddressInsertDTO true "Updated address details"
// @Success 200 {object} response.ApiResponse "Address successfully updated"
// @Failure 400 {object} response.ApiResponse "Bad request"
// @Failure 401 {object} response.ApiResponse "Unauthorized"
// @Router /v1/api/users/addresses/{id} [put]
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

// DeleteAddress deletes an existing address of the authenticated user.
// @Summary Delete user address
// @Description Deletes an existing address for the authenticated user
// @Tags Addresses
// @Accept json
// @Produce json
// @Param id path int true "Address ID"
// @Success 200 {object} response.ApiResponse "Address successfully deleted"
// @Failure 400 {object} response.ApiResponse "Bad request"
// @Failure 401 {object} response.ApiResponse "Unauthorized"
// @Router /v1/api/users/addresses/{id} [delete]
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
