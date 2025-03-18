package main

import (
	"log"
	"os"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/usecase"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/config"
	rabbitmq "github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/message"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/port/input/v1/api/handlers"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/port/input/v1/api/routes"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/port/output/repository"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/pkg/jwt"
	logging "github.com/alexisTrejo11/ecommerce_microservice/rating-service/pkg/log"
	ratelimiter "github.com/alexisTrejo11/ecommerce_microservice/rating-service/pkg/rate_limiter"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// DB
	db := config.GORMConfig()
	config.InitRedis()

	// Messages
	rabbitConn, err := config.ConnectRabbitMQ()
	if err != nil {
		log.Fatalf("Can't Start Queue Messaging Queues %v", err)
	}
	rabbitClient, err := rabbitmq.NewRabbitMQClient(rabbitConn)
	if err != nil {
		log.Fatalf("Can't Start RabbitMQ Client %v", err)
	}

	// Logger
	logging.InitLogger()

	// Rate limiter
	rateLimiter := ratelimiter.NewRateLimiter(config.RedisClient, 20, 1*time.Minute)
	app.Use(rateLimiter.Limit)

	// JWT
	jwtManager, err := jwt.NewJWTManager()
	if err != nil {
		log.Fatalf("Can't Start Auth Manager %v", err)
	}

	// Application
	reviewRepository := repository.NewReviewRepositoryImpl(db)
	reviewUseCase := usecase.NewReviewUseCase(reviewRepository, rabbitClient)
	reviewHandler := handlers.NewReviewHandler(reviewUseCase)
	userReviewHandler := handlers.NewUserReviewHandler(reviewUseCase, *jwtManager)

	routes.ReviewRoutes(app, *reviewHandler)
	routes.UserReviewRoutes(app, *userReviewHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
