package main

import (
	"log"
	"os"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/config"
	certificateController "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/controller"
	certificateRepo "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/repository"
	certificateService "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/service"
	enrollmentController "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/controller"
	enrollmentRepo "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/repository"
	enrollmentService "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/service"
	progressController "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/controller"
	progressRepo "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/repository"
	progressService "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/suscription/controller"
	su_repository "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/suscription/repository"
	su_service "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/suscription/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/routes"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	"github.com/gofiber/fiber/v2"
)

type Enrollment struct {
}

func main() {
	// App
	app := fiber.New()

	// Config
	db := config.GORMConfig()

	// JWT
	jwtManager, err := jwt.NewJWTManager()
	if err != nil {
		log.Fatalf("Can't Start Auth Manager %v", err)
	}

	// Repository
	certificateRepository := certificateRepo.NewCertificateRepository(db)
	enrollmentRepository := enrollmentRepo.NewEnrollmentRepository(db)
	progressRepo := progressRepo.NewProgressRepository(db)
	subRepos := su_repository.NewSubscriptionRepository(db)

	// Service
	certificateService := certificateService.NewCertificateService(certificateRepository, enrollmentRepository)
	enrollmentService := enrollmentService.NewEnrollmentService(enrollmentRepository)
	progresssService := progressService.NewProgressService(progressRepo)
	subscriptionSeervice := su_service.NewSubscriptionService(subRepos)

	// Controller
	certificationController := certificateController.NewCertificateController(certificateService, *jwtManager)
	enrollmentCommandController := enrollmentController.NewEnrollmentComandController(enrollmentService)
	enrollmentQueryController := enrollmentController.NewEnrollmentQueryController(enrollmentService, *jwtManager)
	progressController := progressController.NewProgressController(progresssService, *jwtManager)
	subscriptionController := controller.NewSubscriptionController(subscriptionSeervice, *jwtManager)

	// routes
	routes.CerticationRoutes(app, *certificationController)
	routes.ProgressRoutes(app, *progressController)
	routes.EnrollmentsRoutes(app, *enrollmentCommandController, *enrollmentQueryController)
	routes.SubscriptionRoutes(app, *subscriptionController)

	// Run Server
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("Welcome To Enrollment Service")
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
