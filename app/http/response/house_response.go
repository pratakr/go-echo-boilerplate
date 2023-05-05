package response

import "time"

type HouseResponse struct {
	Id          uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Code        string    `gorm:"uniqueIndex;not null" json:"code"`
	Name        string    `gorm:"not null" json:"name"`
	PlantTotal  uint      `gorm:"not null" json:"plant_total"`
	Temperature float32   `gorm:"not null" json:"temperature"`
	Humidity    float32   `gorm:"not null" json:"humidity"`
	Light       string    `gorm:"not null" json:"light"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
