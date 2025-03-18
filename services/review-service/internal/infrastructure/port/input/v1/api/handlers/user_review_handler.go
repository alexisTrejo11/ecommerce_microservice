package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/port/input"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/pkg/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/pkg/jwt"
	logging "github.com/alexisTrejo11/ecommerce_microservice/rating-service/pkg/log"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserReviewHandler struct {
	useCase    input.ReviewUseCase
	validator  *validator.Validate
	jwtManager jwt.JWTManager
}

func NewUserReviewHandler(useCase input.ReviewUseCase, jwtManager jwt.JWTManager) *UserReviewHandler {
	return &UserReviewHandler{
		useCase:    useCase,
		validator:  validator.New(),
		jwtManager: jwtManager,
	}
}

// TODO: Fetch from JWT
func (h *UserReviewHandler) MyReviews(c *fiber.Ctx) error {
	userID, err := h.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.Unauthorized(c, "unauthorized", err.Error())
	}

	logging.LogIncomingRequest(c, "my_reviews")

	reviews, err := h.useCase.GetReviewsByUserId(context.Background(), userID)
	if err != nil {
		return response.NotFound(c, "Review Not Found", "COURSE_NOT_FOUND")
	}

	logging.LogSuccess("my_reviews", "User Review Succesfully Retrieved", map[string]interface{}{
		"userID": userID,
	})

	return response.OK(c, "UserReview Succesfully Retrieved", reviews)
}

func (h *UserReviewHandler) CreateReview(c *fiber.Ctx) error {
	userID, err := h.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.Unauthorized(c, "unauthorized", err.Error())
	}

	logging.LogIncomingRequest(c, "create_review")

	var insertDTO dtos.ReviewInsertDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		return response.BadRequest(c, err.Error(), "invalid_body_request")
	}

	errorMap, err := response.ValidateStruct(h.validator, &insertDTO)
	if err != nil {
		logging.LogError("create_review", "can't parse body request", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, errorMap, "invalid_data")
	}

	reviewCreate, err := h.useCase.CreateReview(context.Background(), userID, insertDTO)
	if err != nil {
		return response.BadRequest(c, err.Error(), "INVALID_INPUT")
	}

	logging.LogSuccess("create_review", "Review Sucessfully Created", map[string]interface{}{
		"review_id": reviewCreate.ID,
	})

	return response.Created(c, "Review Sucessfully Created", reviewCreate)
}

// Check User Auth
func (h *UserReviewHandler) UpdatMyReview(c *fiber.Ctx) error {
	userID, err := h.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.Unauthorized(c, "unauthorized", err.Error())
	}

	logging.LogIncomingRequest(c, "update_review")

	reviewID, err := response.GetUUIDParam(c, "id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_course_id")
	}

	var insertDTO dtos.ReviewInsertDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		logging.LogError("update_review", "can't parse body request", map[string]interface{}{
			"error": err.Error(),
		})

		return response.BadRequest(c, err.Error(), "invalid_body_request")
	}

	errorMap, err := response.ValidateStruct(h.validator, &insertDTO)
	if err != nil {
		return response.BadRequest(c, errorMap, "invalid_data")
	}

	reviewUpdate, err := h.useCase.UpdateReview(context.Background(), reviewID, userID, insertDTO)
	if err != nil {
		logging.LogError("update_review", "can't parse body request", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "INVALID_INPUT")
	}

	logging.LogSuccess("update_my_review", "Review Sucessfully Updated", map[string]interface{}{
		"review_id": reviewID,
	})

	return response.Created(c, "Review Sucessfully Updated", reviewUpdate)
}

// Check User Auth
func (h *UserReviewHandler) DeletMyReview(c *fiber.Ctx) error {
	userID, err := h.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.Unauthorized(c, "unauthorized", err.Error())
	}

	logging.LogIncomingRequest(c, "delete_review")

	reviewID, err := response.GetUUIDParam(c, "id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_review_id")
	}

	err = h.useCase.DeleteReview(context.Background(), userID, reviewID)
	if err != nil {
		return response.NotFound(c, "Review Not Found", "COURSE_NOT_FOUND")
	}

	logging.LogSuccess("delete_my_review", "User Review Succesfully Delete", map[string]interface{}{
		"review_id": reviewID,
	})

	return response.OK(c, "User Review Succesfully Deleted", nil)
}
