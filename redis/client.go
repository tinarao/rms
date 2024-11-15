package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var Client *redis.Client

func Init(password string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: password,
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("failed \"ping\" command: %s", err)
	}

	Client = rdb
}
