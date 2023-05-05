package service

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	Db     *gorm.DB
	Logger *zap.Logger
}

func NewService(db *gorm.DB) *Service {
	return &Service{Db: db}
}
