package handlers

import (
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/port/input"
	"github.com/gofiber/fiber/v2"
)

type UserReviewHandler struct {
	useCase input.ReviewUseCase
}

func NewUserReviewHandler(useCase input.ReviewUseCase) *UserReviewHandler {
	return &UserReviewHandler{
		useCase: useCase,
	}
}

func (h *UserReviewHandler) MyReviews(c *fiber.Ctx) error {

	return c.Status(204).JSON(fiber.Map{"course_id": "reviews"})
}

func (h *UserReviewHandler) CreateReview(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{"review": "review"})
}

func (h *UserReviewHandler) UpdatMyReview(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{"user_id": "reviews"})
}

func (h *UserReviewHandler) DeletMyReview(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{"course_id": "reviews"})
}
