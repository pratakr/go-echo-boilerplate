package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id            uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name          string    `gorm:"not null" json:"name"`
	Password      string    `gorm:"not null"`
	Email         string    `gorm:"uniqueIndex;not null" json:"email"`
	RememberToken string    `gorm:"column:remember_token"`
	IsActive      bool      `gorm:"default:1"`
	RoleId        uint      `gorm:"not null" json:"role_id"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
