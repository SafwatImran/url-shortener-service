package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var Ctx = context.Background()

var client *redis.Client

func GetClient(dbNo int) *redis.Client {
	if client == nil {
		rdb := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("DB_ADDR"),
			Password: os.Getenv("DB_PASS"),
			DB:       dbNo,
		})

		client = rdb
	}
	return client
}
