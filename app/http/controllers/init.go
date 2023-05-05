package controllers

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type controller struct {
	Db          *gorm.DB
	RedisClient *redis.Client
	Logger      *zap.Logger
}

func NewController(db *gorm.DB, redisClient *redis.Client, logger *zap.Logger) *controller {
	return &controller{
		Db:          db,
		RedisClient: redisClient,
		Logger:      logger,
	}
}
