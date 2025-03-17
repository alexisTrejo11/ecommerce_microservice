package ratelimiter

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type RateLimiter struct {
	redisClient *redis.Client
	rate        int
	window      time.Duration
}

func NewRateLimiter(redisClient *redis.Client, rate int, window time.Duration) *RateLimiter {
	if redisClient == nil {
		log.Fatal("redisClient cannot be nil")
	}
	return &RateLimiter{
		redisClient: redisClient,
		rate:        rate,
		window:      window,
	}
}

func (rl *RateLimiter) Limit(c *fiber.Ctx) error {
	ctx := context.Background()
	ip := c.IP()
	now := time.Now().Unix()

	windowSeconds := int64(rl.window / time.Second)
	key := ip + ":" + strconv.FormatInt(now/windowSeconds, 10)

	count, err := rl.redisClient.Incr(ctx, key).Result()
	if err != nil {
		log.Printf("Error incrementing in Redis for %s: %v", key, err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}

	if count == 1 {
		err = rl.redisClient.Expire(ctx, key, rl.window).Err()
		if err != nil {
			log.Printf("Error setting expiration for %s: %v", key, err)
		}
	}

	if count > int64(rl.rate) {
		log.Printf("Rate limit exceeded for %s", ip)
		return c.Status(fiber.StatusTooManyRequests).SendString("Too many requests")
	}

	return c.Next()
}
