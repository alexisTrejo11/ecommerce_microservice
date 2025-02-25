package main

import (
	"log"

	routes "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/http/v1"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/http/v1/handlers"
	repository "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/config"
	usecases "github.com/alexisTrejo11/ecommerce_microservice/internal/core/application"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/email"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/jwt"
	"github.com/alexisTrejo11/ecommerce_microservice/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := config.GORMConfig()
	emailConfig := config.GetEmailConfig()
	config.InitRedis()

	jwtManager, err := jwt.NewJWTManager()
	if err != nil {
		log.Fatalf("Error initing JWTManager: %v", err)
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

	log.Fatal(app.Listen(":3000"))
}
