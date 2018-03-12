package main

import (
	"github.com/go-redis/redis"
)

func initRedis() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})
	if _, err := redisClient.Ping().Result(); err != nil {
		return err
	}
	return nil
}
