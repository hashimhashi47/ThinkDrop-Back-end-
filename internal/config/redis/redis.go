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
	Adress := os.Getenv("REDIS_ADDRESS")
	rdb := redis.NewClient(&redis.Options{
		Addr: Adress,
	})

	if err := rdb.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	log.Println("Redis connected")
	return rdb
}
