package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/input/v1/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func ModulesRoutes(app *fiber.App, moduleHanlders handlers.ModuleHandler) {
	path := app.Group("v1/api/modules")
	path.Get("/:id", moduleHanlders.GetModuleById)
	path.Get("/course/:course_id", moduleHanlders.GetModulesByCourseId)
	path.Post("", moduleHanlders.CreateModule)
	path.Put("/:id", moduleHanlders.UpdateModule)
	path.Delete("/:id", moduleHanlders.DeleteModule)
}
