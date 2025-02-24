package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"

	models "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GORMConfig() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create the Data Source Name (DSN) for the MySQL connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	var db *gorm.DB
	var err error

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		log.Printf("Attempt %d: Failed to connect to the database: %v\n", i+1, err)
		if i == maxRetries-1 {
			log.Fatal("Maximum retry attempts reached. Failed to connect to the database:", err)
		}

		time.Sleep(5 * time.Second)
	}

	if err := db.AutoMigrate(
		&models.UserModel{},
		&models.RoleModel{},
		&models.PasswordResetModel{},
		&models.AddressModel{},
		&models.MFAModel{},
		&models.SessionModel{},
	); err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	if err := seedRoles(db); err != nil {
		log.Fatal("Failed to seed roles:", err)
	}

	log.Println("Database connected successfully!")
	return db
}

func seedRoles(db *gorm.DB) error {
	defaultRoles := []models.RoleModel{
		{Name: "admin", Description: "Administrator role"},
		{Name: "common_user", Description: "Regular user role"},
		{Name: "premuim_user", Description: "Common user role"},
		{Name: "guest", Description: "Guest role"},
	}

	for _, role := range defaultRoles {
		var existingRole models.RoleModel
		if err := db.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := db.Create(&role).Error; err != nil {
					return fmt.Errorf("failed to create role %s: %w", role.Name, err)
				}
			} else {
				return fmt.Errorf("failed to query role %s: %w", role.Name, err)
			}
		}
	}

	return nil
}

var RedisClient *redis.Client

func InitRedis() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic("Can not connect to Redis: " + err.Error())
	}

	fmt.Println("Succesfully connected to Redis")
}

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
	FromEmail    string
}

func GetEmailConfig() EmailConfig {
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatalf("Error converting SMTP_PORT to int: %v", err)
	}

	return EmailConfig{
		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPPort:     port,
		SMTPUser:     os.Getenv("SMTP_USER"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
		FromEmail:    os.Getenv("FROM_EMAIL"),
	}
}
