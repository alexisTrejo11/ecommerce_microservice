package config

import (
	"fmt"
	"log"
	"os"
	"time"

	certificate "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/model"
	enrollment "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/model"
	progress "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/model"
	subscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GORMConfig() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

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
		&enrollment.Enrollment{},
		&progress.CompletedLesson{},
		&certificate.Certificate{},
		&subscription.Subscription{},
	); err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	log.Println("Database connected successfully!")
	return db
}
