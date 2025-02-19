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

	jwtManager, err := jwt.NewJWTManager()
	if err != nil {
		log.Fatalf("Error initing JWTManager: %v", err)
	}

	userRepository := repository.NewUserRepository(db)
	token_service := repository.NewTokenService(jwtManager)
	authUseCase := usecases.NewAuthUseCase(userRepository, token_service)
	authHandler := handlers.NewAuthHandler(authUseCase)

	routes.AuthRoutes(app, authHandler)

	log.Fatal(app.Listen(":3000"))
}
