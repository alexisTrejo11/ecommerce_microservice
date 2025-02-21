package main

import (
	"log"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/api/handlers"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/api/routes"
	repository "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output"
	usecases "github.com/alexisTrejo11/ecommerce_microservice/internal/core/usercase"
	"github.com/alexisTrejo11/ecommerce_microservice/pkg/jwt"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := config.GORMConfig()
	config.InitRedis()

	jwtManager, err := jwt.NewJWTManager()
	if err != nil {
		log.Fatalf("Error initing JWTManager: %v", err)
	}

	// Repository
	userRepository := repository.NewUserRepository(db)
	addresRepository := repository.NewAddressRepository(db)
	sessionRepository := repository.NewSessionRepository(db)

	// UseCase
	tokenService := repository.NewTokenService(jwtManager)
	authUseCase := usecases.NewAuthUseCase(userRepository, tokenService, sessionRepository)
	addresUseCase := usecases.NewAddressUseCase(addresRepository)
	sessionUseCase := usecases.NewSessionUserCase(sessionRepository)

	// Handler
	authHandler := handlers.NewAuthHandler(authUseCase, *jwtManager)
	userAddresHandler := handlers.NewUserAddressHandler(addresUseCase, *jwtManager)
	sessionHandler := handlers.NewSessionHandler(sessionUseCase, *jwtManager)

	// Routes
	routes.AuthRoutes(app, authHandler)
	routes.UserAddressRoutes(app, userAddresHandler)
	routes.SessionRoutes(app, sessionHandler)

	log.Fatal(app.Listen(":3000"))
}
