package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/port/input"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/shared/response"
	"github.com/gofiber/fiber/v2"
)

type ReviewHandler struct {
	useCase input.ReviewUseCase
}

func NewReviewHandler(useCase input.ReviewUseCase) *ReviewHandler {
	return &ReviewHandler{
		useCase: useCase,
	}
}

func (h *ReviewHandler) GetReviewByID(c *fiber.Ctx) error {
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

func (h *ReviewHandler) GetReviewByUserID(c *fiber.Ctx) error {
	userID, err := response.GetUUIDParam(c, "user_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_review_id")
	}

	review, err := h.useCase.GetReviewsByUserId(context.Background(), userID)
	if err != nil {
		return response.NotFound(c, "Review Not Found", "COURSE_NOT_FOUND")
	}

	return response.OK(c, "User Review Succesfully Retrieved", review)
}

func (h *ReviewHandler) GetReviewByCourseID(c *fiber.Ctx) error {
	courseID, err := response.GetUUIDParam(c, "course_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_course_id")
	}

	reviews, err := h.useCase.GetReviewsByCourseId(context.Background(), courseID)
	if err != nil {
		return response.NotFound(c, err.Error(), "COURSE_NOT_FOUND")
	}

	return response.OK(c, "Course Reviews Succesfully Retrieved", reviews)
}

func (h *ReviewHandler) DeleteReview(c *fiber.Ctx) error {
	reviewID, err := response.GetUUIDParam(c, "id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_review_id")
	}

	err = h.useCase.DeleteReview(context.Background(), reviewID)
	if err != nil {
		return response.NotFound(c, "Review Not Found", "COURSE_NOT_FOUND")
	}

	return response.OK(c, "User Review Succesfully Deleted", nil)
}
