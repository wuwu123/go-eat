package common

import (
	"github.com/go-redis/redis"
)

func GetRedis() *redis.Client  {
	client := redis.NewClient(&redis.Options{
		Addr:     "47.100.82.54:6379",
		Password: "ky1024",
		DB:       1,
	})
	return client
}