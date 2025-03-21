package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	if redisHost == "" || redisPort == "" || redisPassword == "" {
		log.Fatal("Missing Redis configuration in environment variables")
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
		DB:       0,
	})

	ctx := context.Background()
	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		RedisClient.Close()
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	fmt.Println("Successfully connected to Redis")
}
