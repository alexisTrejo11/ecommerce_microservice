package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	models "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GORMConfig initializes and returns a GORM database connection.
// It reads database configuration from environment variables and retries the connection if necessary.
func GORMConfig() *gorm.DB {
	// Read database configuration from environment variables
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
