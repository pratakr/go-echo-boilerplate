package response

import "time"

type UserResponse struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name            string    `gorm:"column:name;not null" json:"name"`
	Email           string    `gorm:"column:email;not null" json:"email"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at" json:"email_verified_at"`
	Password        string    `gorm:"column:password;not null" json:"password"`
	RememberToken   string    `gorm:"column:remember_token" json:"remember_token"`
	IsActive        int32     `gorm:"column:is_active" json:"is_active"`
	RoleID          int32     `gorm:"column:role_id" json:"role_id"`
	CreatedAt       time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}
