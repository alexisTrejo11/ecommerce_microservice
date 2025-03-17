package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/port/input/v1/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func ReviewRoutes(app *fiber.App, handler handlers.ReviewHandler) {
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Rating Service")
	})

	path := app.Group("v1/api/reviews")

	path.Get("/:id", handler.GetReviewByID)
	path.Get("user/:user_id", handler.GetReviewByUserID)
	path.Get("course/:course_id", handler.GetReviewByCourseID)
	path.Delete("/:id", handler.DeleteReview)
}

func UserReviewRoutes(app *fiber.App, handler handlers.UserReviewHandler) {
	path := app.Group("v1/api/users/reviews")

	path.Get("/:id", handler.MyReviews)
	path.Post("", handler.CreateReview)
	path.Put("/:id", handler.UpdatMyReview)
	path.Delete("/:id", handler.DeletMyReview)
}
