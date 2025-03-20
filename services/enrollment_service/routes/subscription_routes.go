package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/controller"
	"github.com/gofiber/fiber/v2"
)

func SubscriptionRoutes(app *fiber.App, controller controller.SubscriptionController) {
	path := app.Group("/v1/api/subscriptions")

	path.Get("/my", controller.GetMySubscription)
	path.Patch("/cancel", controller.CancelMySubscription)

	path.Post("", controller.CreateSubscription)
	path.Patch(":user_id/type/:sub_type", controller.ChangeMySubscriptionType)
	path.Delete("/:subscription_id", controller.DeleteSubscription)
}
