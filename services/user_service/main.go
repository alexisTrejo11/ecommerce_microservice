package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/alexisTrejo11/ecommerce_microservice/docs"
	routes "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/http/v1"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/http/v1/handlers"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/http/v1/middleware"
	repository "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/config"
	usecases "github.com/alexisTrejo11/ecommerce_microservice/internal/core/application"
	port "github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output/logger"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/email"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/jwt"
	logging "github.com/alexisTrejo11/ecommerce_microservice/internal/shared/logger"
	"github.com/alexisTrejo11/ecommerce_microservice/pkg/rabbitmq"
	ratelimiter "github.com/alexisTrejo11/ecommerce_microservice/pkg/rate_limiter"
	swagger "github.com/arsmn/fiber-swagger/v2"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a global context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize Logger
	log := logging.InitLogger()
	defer log.Flush()

	log.Info(ctx, "Starting application",
		port.Field{Key: "environment", Value: os.Getenv("ENVIRONMENT")},
	)

	// Initialize database before starting the server
	db := config.GORMConfig(log)

	// Email Configuration
	emailConfig := config.GetEmailConfig()

	// Initialize Redis
	config.InitRedis()

	// Create Fiber instance
	app := fiber.New()

	// Middleware
	app.Use(middleware.LoggerMiddleware(log))
	app.Use(middleware.RequestIDMiddleware())

	// Rate Limiter (50 Requests per minute)
	rateLimiter := ratelimiter.NewRateLimiter(config.RedisClient, 50, 1*time.Minute)
	app.Use(rateLimiter.Limit)

	// Initialize JWT Manager
	jwtManager, err := jwt.NewJWTManager()
	if err != nil {
		log.Fatal(ctx, "Error initializing JWTManager",
			port.Field{Key: "error", Value: err.Error()},
		)
	}

	// Repository
	userRepository := repository.NewUserRepository(db)
	addresRepository := repository.NewAddressRepository(db)
	sessionRepository := repository.NewSessionRepository(db)
	mfaRepository := repository.NewMFARepository(db)

	// Mail Client
	mailClient := email.NewMailClient(emailConfig)

	// UseCase
	tokenService := repository.NewTokenService(jwtManager)
	authUseCase := usecases.NewAuthUseCase(userRepository, tokenService, sessionRepository, mfaRepository)
	addresUseCase := usecases.NewAddressUseCase(addresRepository)
	sessionUseCase := usecases.NewSessionUserCase(sessionRepository)
	mfaUseCase := usecases.NewMFAUseCase(mfaRepository, tokenService)
	emailUseCase := usecases.NewEmailUseCase(mailClient, userRepository, tokenService)

	// Handler
	authHandler := handlers.NewAuthHandler(authUseCase, *jwtManager, emailUseCase)
	userAddresHandler := handlers.NewUserAddressHandler(addresUseCase, *jwtManager)
	sessionHandler := handlers.NewSessionHandler(sessionUseCase, *jwtManager)
	mfaHandler := handlers.NewUserMfaHandler(mfaUseCase, *jwtManager)

	// RabbitMQ
	emailConsumer := rabbitmq.NewEmailConsumer(emailUseCase)
	go emailConsumer.ConsumeEmail()

	// Routes
	routes.AuthRoutes(app, authHandler)
	routes.UserAddressRoutes(app, userAddresHandler)
	routes.SessionRoutes(app, sessionHandler)
	routes.UserMFARoutes(app, mfaHandler)

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Channel to capture termination signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		app_port := os.Getenv("APP_PORT")
		log.Info(ctx, "Starting HTTP server",
			port.Field{Key: "port", Value: app_port},
		)

		if err := app.Listen(":" + app_port); err != nil {
			log.Fatal(ctx, "Error starting server",
				port.Field{Key: "error", Value: err.Error()},
			)
		}
	}()

	// Wait for termination signal
	<-sigChan
	log.Info(ctx, "Shutdown signal received, closing server...")

	// Properly shutdown the server
	if err := app.Shutdown(); err != nil {
		log.Error(ctx, "Error shutting down server",
			port.Field{Key: "error", Value: err.Error()},
		)
	}

	log.Info(ctx, "Server shut down successfully.")
}
