package main

import (
	"log"
	"os"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/config"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/adapters/input/v1/http/api/handlers"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/adapters/input/v1/http/api/routes"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/adapters/output/repository"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/core/application/usecases"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/pkg/facadeService"
	ratelimiter "github.com/alexisTrejo11/ecommerce_microservice/cart-service/pkg/rate_limiter"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// ROUTER
	app := fiber.New()

	// db
	gormDB := config.GORMConfig()

	//CONFIG
	//redis
	config.InitRedis()

	// rate limiter (50 Requests per minute)
	rateLimiter := ratelimiter.NewRateLimiter(config.RedisClient, 50, 1*time.Minute)
	app.Use(rateLimiter.Limit)

	// APP
	// repository
	cartRepository := repository.NewCartRepository(gormDB)

	// usecases
	productService := facadeService.NewProductFacadeService()
	cartUseCase := usecases.NewCartUseCase(cartRepository, productService)
	cartHandler := handlers.NewCartHandler(cartUseCase)

	// handlers
	userCartHandler := handlers.NewUserCartHandler(cartUseCase)

	// routes
	routes.CartRoutes(app, *cartHandler)
	routes.UserCartRoutes(app, *userCartHandler)

	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Cart Service")
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
