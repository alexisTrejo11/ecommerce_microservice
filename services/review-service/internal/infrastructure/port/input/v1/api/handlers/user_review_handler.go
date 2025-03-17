package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/port/input"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/shared/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserReviewHandler struct {
	useCase   input.ReviewUseCase
	validator *validator.Validate
}

func NewUserReviewHandler(useCase input.ReviewUseCase) *UserReviewHandler {
	return &UserReviewHandler{
		useCase:   useCase,
		validator: validator.New(),
	}
}

// TODO: Fetch from JWT
func (h *UserReviewHandler) MyReviews(c *fiber.Ctx) error {
	reviewID, err := response.GetUUIDParam(c, "id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_course_id")
	}

	review, err := h.useCase.GetReviewById(context.Background(), reviewID)
	if err != nil {
		return response.NotFound(c, "Review Not Found", "COURSE_NOT_FOUND")
	}

	return response.OK(c, "Review Succesfully Retrieved", review)
}

func (h *UserReviewHandler) CreateReview(c *fiber.Ctx) error {
	var insertDTO dtos.ReviewInsertDTO

	if err := c.BodyParser(&insertDTO); err != nil {
		return response.BadRequest(c, err.Error(), "INVALID_BODY_REQUEST")
	}

	_, err := response.ValidateStruct(h.validator, &insertDTO)
	if err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	reviewCreate, err := h.useCase.CreateReview(context.Background(), insertDTO)
	if err != nil {
		return response.BadRequest(c, err.Error(), "INVALID_INPUT")
	}

	return response.Created(c, "Review Sucessfully Created", reviewCreate)
}

// Check User Auth
func (h *UserReviewHandler) UpdatMyReview(c *fiber.Ctx) error {
	reviewID, err := response.GetUUIDParam(c, "id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_course_id")
	}

	var insertDTO dtos.ReviewInsertDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		return response.BadRequest(c, err.Error(), "INVALID_BODY_REQUEST")
	}

	_, err = response.ValidateStruct(h.validator, &insertDTO)
	if err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	reviewUpdate, err := h.useCase.UpdateReview(context.Background(), reviewID, insertDTO)
	if err != nil {
		return response.BadRequest(c, err.Error(), "INVALID_INPUT")
	}

	return response.Created(c, "Review Sucessfully Updated", reviewUpdate)
}

// Check User Auth
func (h *UserReviewHandler) DeletMyReview(c *fiber.Ctx) error {
	reviewID, err := response.GetUUIDParam(c, "review_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_review_id")
	}

	err = h.useCase.DeleteReview(context.Background(), reviewID)
	if err != nil {
		return response.NotFound(c, "Review Not Found", "COURSE_NOT_FOUND")
	}

	return response.OK(c, "User Review Succesfully Deleted", nil)
}
