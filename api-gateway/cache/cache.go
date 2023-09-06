package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RdsClient *redis.Client

func init() {
	RdsClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	cmdResult := RdsClient.Ping(context.Background())
	if cmdResult.Err() != nil {
		panic(cmdResult.Err())
	}
}
