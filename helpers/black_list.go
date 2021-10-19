package helpers

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:6379",
		Password: "",
		DB:       0,
	})
}

func Save(key string, value string) error {
	ctx := context.Background()
	status := rdb.Set(ctx, key, value, 15 * time.Minute)

	_, err := status.Result()
	return err
}

func Get(key string) (string, error){
	ctx := context.Background()
	status := rdb.Get(ctx, key)	
	result, err := status.Result()	
	return result, err	
}