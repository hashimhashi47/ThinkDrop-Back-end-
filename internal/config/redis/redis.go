package redis

import (
	"context"
	"log"
	"os"
	"github.com/redis/go-redis/v9"
)

// -> connection and running redis

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
	address := os.Getenv("REDIS_ADDRESS")
	if address == "" {
		log.Fatal("REDIS_ADDRESS not set")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: address,
	})

	if err := rdb.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	log.Println("Redis connected")
	return rdb
}
