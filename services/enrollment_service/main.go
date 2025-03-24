package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/config"
	_ "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/docs"
	certificateController "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/controller"
	certificateRepo "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/repository"
	certificateService "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/service"
	enrollmentController "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/controller"
	enrollmentRepo "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/repository"
	enrollmentService "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/service"
	progressController "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/controller"
	progressRepo "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/repository"
	progressService "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/controller"
	su_repository "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/repository"
	su_service "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/routes"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	logging "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/logger"
	rabbitmq "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/messaging"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/middleware"
	ratelimiter "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/rate_limiter"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

type Enrollment struct {
}

func main() {
	// App
	app := fiber.New()

	// Config
	db := config.GORMConfig()
	config.InitRedis()

	// Mongo
	mongoDBName := os.Getenv("MONGO_DB_NAME")
	if mongoDBName == "" {
		mongoDBName = "ENROLLMENT_DB"
	}
	mongoClient := config.InitMongoClient()
	mongoCollections, err := config.CreateCollections(mongoClient, mongoDBName)
	if err != nil {
		log.Fatalf("Can't Start Mongo Collection %v", err)
	}

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Log
	logging.InitLogger()

	// Rate Limiter
	rateLimiter := ratelimiter.NewRateLimiter(config.RedisClient, 50, 1*time.Minute)
	app.Use(rateLimiter.Limit)

	// JWT
	jwtManager, err := jwt.NewJWTManager()
	if err != nil {
		log.Fatalf("Can't Start Auth Manager %v", err)
	}

	// RabbitMQ
	rabbitConn, err := config.ConnectRabbitMQ()
	if err != nil {
		log.Fatalf("Can't Connect to rabbitMQ: %v", err)
	}

	client, err := config.NewRabbitMQClient(rabbitConn)
	if err != nil {
		log.Fatalf("Can't setup rabbitMQ: %v", err)
	}

	// Repository
	certificateRepository := certificateRepo.NewCertificateRepository(db)
	enrollmentRepository := enrollmentRepo.NewEnrollmentRepository(db)
	progressRepository := progressRepo.NewProgressRepository(db)
	courseRepository := progressRepo.NewMongoDBCourseRepository(mongoClient, mongoDBName, *mongoCollections)
	subRepos := su_repository.NewSubscriptionRepository(db)

	// Service
	certificateService := certificateService.NewCertificateService(certificateRepository, enrollmentRepository)
	enrollmentService := enrollmentService.NewEnrollmentService(enrollmentRepository)
	progressService := progressService.NewProgressService(progressRepository, enrollmentRepository, courseRepository)
	subscriptionService := su_service.NewSubscriptionService(subRepos)

	// Controller
	certificationController := certificateController.NewCertificateController(certificateService, jwtManager)
	enrollmentCommandController := enrollmentController.NewEnrollmentComandController(enrollmentService, certificateService, progressService)
	enrollmentQueryController := enrollmentController.NewEnrollmentQueryController(enrollmentService, jwtManager)
	progressController := progressController.NewProgressController(progressService, jwtManager)
	subscriptionController := controller.NewSubscriptionController(subscriptionService, jwtManager)

	// Receiver
	courseReceiver := rabbitmq.NewCourseQueueReceiver(client, "course_queue", time.Second*5, courseRepository)
	moduleReceiver := rabbitmq.NewModuleQueueReceiver(client, "module_queue", 5, courseRepository)
	lessonReceiver := rabbitmq.NewLessonQueueReceiver(client, "lesson_queue", 5, courseRepository)

	go courseReceiver.ReceiveCourse(context.Background())
	go moduleReceiver.ReceiveModule(context.Background())
	go lessonReceiver.ReceiveLesson(context.Background())

	// routes
	// Public Routes
	routes.UserCerticationRoutes(app, *certificationController)
	routes.CerticationRoutes(app, *certificationController)
	routes.SubscriptionRoutes(app, *subscriptionController)
	routes.EnrollmentsRoutes(app, *enrollmentCommandController, *enrollmentQueryController)

	// Auth Routes
	app.Use(middleware.JWTAuthMiddleware(jwtManager))
	routes.UserSubscriptionRoutes(app, *subscriptionController)
	routes.UserEnrollmentsRoutes(app, *enrollmentCommandController, *enrollmentQueryController)
	routes.ProgressRoutes(app, *progressController)

	// Checker to Expire Notification
	go subscriptionService.StartSubscriptionChecker(5 * time.Minute)

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
