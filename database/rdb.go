package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RDB struct {
	client *redis.Client
}

// LRUExp - using an LRU cache with 1 hour expiry
const LRUExp = time.Hour

var (
	redisDatabase = &RDB{}
	ctx           = context.Background()
)

func InitDB() *RDB {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	// testing client connection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize redis DB. Err: %v", err))
	}
	fmt.Sprintf("Successfully initialized redis database - %s", pong)

	redisDatabase.client = client

	return redisDatabase
}

func StoreURL(shortened string, original string) {
	err := redisDatabase.client.Set(ctx, shortened, original, LRUExp).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to store URL. Err: %v", err))
	}
}

func RetrieveURL(shortened string) string {
	resp, err := redisDatabase.client.Get(ctx, shortened).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to retrieve URL. Err: %v", err))
	}
	return resp
}
