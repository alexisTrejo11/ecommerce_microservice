package main

import (
	"log"
	"os"

	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/usecase"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/config"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/port/input/v1/api/handlers"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/port/input/v1/api/routes"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/port/output/repository"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Config
	db := config.GORMConfig()
	config.InitRedis()

	// Application
	reviewRepository := repository.NewReviewRepositoryImpl(db)
	reviewUseCase := usecase.NewReviewUseCase(reviewRepository)
	reviewHandler := handlers.NewReviewHandler(reviewUseCase)

	routes.ReviewRoutes(app, *reviewHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
