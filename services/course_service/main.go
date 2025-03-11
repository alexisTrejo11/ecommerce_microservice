package main

import (
	"log"
	"os"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/config"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/input/v1/api/handlers"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/input/v1/api/routes"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/repository"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/usecase"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Router
	app := fiber.New()

	// Config
	db := config.GORMConfig()
	config.InitRedis()

	// Repository
	courseRepository := repository.NewCourseRepository(*db)
	lessonRepository := repository.NewLessonRepository(*db)
	moduleRepository := repository.NewModuleRepository(*db)
	resourceRepository := repository.NewResourceRepository(*db)

	// Use Case
	courseUseCase := usecase.NewCourseUseCase(courseRepository)
	moduleUseCase := usecase.NewModuleUseCase(moduleRepository, courseRepository)
	lessonUseCase := usecase.NewLessonUseCase(lessonRepository, moduleRepository)
	resourceUseCase := usecase.NewResourceUseCase(resourceRepository, lessonRepository)

	// Handler
	courseHandler := handlers.NewCourseHandler(courseUseCase)
	lessonHandler := handlers.NewLessonHandler(lessonUseCase)
	moduleHandler := handlers.NewModuleHandler(moduleUseCase)
	resourceHandler := handlers.NewResourceHandler(resourceUseCase)

	// Routes
	routes.CourseRoutes(app, *courseHandler)
	routes.LessonRoutes(app, *lessonHandler)
	routes.ModulesRoutes(app, *moduleHandler)
	routes.ResourceRoutes(app, *resourceHandler)

	// Run Server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
