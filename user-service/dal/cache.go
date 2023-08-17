package dal

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var redis_client *redis.Client

func init() {
	redis_client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	cmdResult := redis_client.Ping(context.Background())
	if cmdResult.Err() != nil {
		panic(cmdResult.Err())
	}
}
