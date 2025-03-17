package handlers

import (
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/port/input"
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

	return c.Status(200).JSON(fiber.Map{"review": "review"})
}

func (h *ReviewHandler) GetReviewByUserID(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{"user_id": "reviews"})
}

func (h *ReviewHandler) GetReviewByCourseID(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{"course_id": "reviews"})
}

func (h *ReviewHandler) DeleteReview(c *fiber.Ctx) error {

	return c.Status(204).JSON(fiber.Map{"course_id": "reviews"})
}
