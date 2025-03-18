package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/port/input"
	logging "github.com/alexisTrejo11/ecommerce_microservice/rating-service/pkg/log"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/pkg/response"
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
	logging.LogIncomingRequest(c, "get_review_by_id")

	reviewID, err := response.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("get_review_by_id", "Invalid review ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid_review_id")
	}

	review, err := h.useCase.GetReviewById(context.Background(), reviewID)
	if err != nil {
		return response.NotFound(c, "Review Not Found", "COURSE_NOT_FOUND")
	}

	logging.LogSuccess("get_review_by_id", "Review Succesfully Retrieved", map[string]interface{}{
		"review_id": reviewID,
	})

	return response.OK(c, "Review Succesfully Retrieved", review)
}

func (h *ReviewHandler) GetReviewByUserID(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "get_review_by_user_id")

	userID, err := response.GetUUIDParam(c, "user_id")
	if err != nil {
		logging.LogError("get_review_by_user_id", "Invalid user ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	review, err := h.useCase.GetReviewsByUserId(context.Background(), userID)
	if err != nil {
		return response.NotFound(c, "Review Not Found", "COURSE_NOT_FOUND")
	}

	logging.LogSuccess("get_review_by_user_id", "User Review Succesfully Retrieved", map[string]interface{}{
		"user_id": userID,
	})

	return response.OK(c, "User Review Succesfully Retrieved", review)
}

func (h *ReviewHandler) GetReviewByCourseID(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "get_review_by_course_id")

	courseID, err := response.GetUUIDParam(c, "course_id")
	if err != nil {
		logging.LogError("get_review_by_course_id", "Invalid course ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid_course_id")
	}

	reviews, err := h.useCase.GetReviewsByCourseId(context.Background(), courseID)
	if err != nil {
		return response.NotFound(c, err.Error(), "COURSE_NOT_FOUND")
	}

	logging.LogSuccess("get_review_by_course_id", "Course Reviews Succesfully Retrieved", map[string]interface{}{
		"course_id": courseID,
	})

	return response.OK(c, "Course Reviews Succesfully Retrieved", reviews)
}

func (h *ReviewHandler) DeleteReview(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "delete_review")

	reviewID, err := response.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("delete_review", "Invalid review ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid_review_id")
	}

	userID, err := response.GetUUIDParam(c, "user_id")
	if err != nil {
		logging.LogError("delete_review", "Invalid review ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid_review_id")
	}

	err = h.useCase.DeleteReview(context.Background(), userID, reviewID)
	if err != nil {
		return response.NotFound(c, "Review Not Found", "COURSE_NOT_FOUND")
	}

	logging.LogSuccess("delete_review", "User Review Succesfully Deleted", map[string]interface{}{
		"review_id": reviewID,
	})

	return response.OK(c, "User Review Succesfully Deleted", nil)
}
