package caches

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
	"os"
)

func Connect() *redis.Client{
	redisAddr := fmt.Sprintf("%s:%s",os.Getenv("REDIS_HOST"),os.Getenv("REDIS_PORT"))
	log.Info(fmt.Sprintf("Redis:%s",redisAddr))
	return redis.NewClient(&redis.Options{
		Addr: redisAddr   ,
	})
}