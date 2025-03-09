package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

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
