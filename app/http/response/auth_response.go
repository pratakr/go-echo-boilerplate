package response

type UserResponse struct {
	Id     uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name   string `gorm:"uniqueIndex;not null" json:"name"`
	Email  string `gorm:"uniqueIndex;not null" json:"email"`
	RoleId uint   `gorm:"not null" json:"role_id"`
	Token  string `json:"token"`
}

type ProfileResponse struct {
	Id     uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name   string `gorm:"uniqueIndex;not null" json:"name"`
	Email  string `gorm:"uniqueIndex;not null" json:"email"`
	RoleId uint   `gorm:"not null" json:"role_id"`
}
