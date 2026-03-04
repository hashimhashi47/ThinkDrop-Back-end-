package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

// -> connection and running redis

var Ctx = context.Background()

// func NewRedisClient() *redis.Client {
// 	// address := os.Getenv("REDIS_ADDRESS")
// 	// if address == "" {
// 	// 	log.Fatal("REDIS_ADDRESS not set")
// 	// }
// 	// rdb := redis.NewClient(&redis.Options{
// 	// 	Addr: address,
// 	// })

// 	// if err := rdb.Ping(Ctx).Err(); err != nil {
// 	// 	log.Fatalf("Redis connection failed: %v", err)
// 	// }

// 	// log.Println("Redis connected")
// 	// return rdb
// }

func NewRedisClient() *redis.Client {

	var rdb *redis.Client

	redisURL := os.Getenv("REDIS_URL")

	if redisURL != "" {
		opt, err := redis.ParseURL(redisURL)
		if err != nil {
			log.Fatalf("Invalid Redis URL: %v", err)
		}

		rdb = redis.NewClient(opt)

	} else {
		address := os.Getenv("REDIS_ADDRESS")
		if address == "" {
			log.Fatal("REDIS_ADDRESS not set")
		}

		rdb = redis.NewClient(&redis.Options{
			Addr: address,
		})
	}

	if err := rdb.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	log.Println("Redis connected ✅")

	return rdb
}
