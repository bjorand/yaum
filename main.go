package main

import (
	"os"
	"strings"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (
	listenAddr    = "127.0.0.1:8080"
	redisAddr     = "127.0.0.1:6379"
	redisPassword = ""
	redisDB       = 0
	redisClient   *redis.Client
)

func main() {

	if os.Getenv("LISTEN_ADDR") != "" {
		listenAddr = os.Getenv("LISTEN_ADDR")
	}
	if os.Getenv("REDIS_ADDR") != "" {
		redisAddr = os.Getenv("REDIS_ADDR")
	}
	switch level := strings.ToLower(os.Getenv("LOG_LEVEL")); level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	if err := initRedis(); err != nil {
		log.Fatal(err)
	}
	r := ginEngine()
	r.Run(listenAddr)
}
