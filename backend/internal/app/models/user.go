package models

import "time"

type User struct {
	BaseModel
	Uuid         string    `gorm:"unique;size:255;not null" json:"uuid"`
	Email        string    `gorm:"unique;size:255;not null" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"password"`
	Username     string    `gorm:"size:255;not null" json:"username"`
	Avatar       string    `gorm:"size:1024;" json:"avatar"`
	Role         string    `gorm:"size:255;not null;default:'user'" json:"role"`
	Status       string    `gorm:"size:255;default:'active'" json:"-"`
	LastLoginAt  time.Time `gorm:"type:timestamp(3);default:CURRENT_TIMESTAMP(3)" json:"-"`
	LastLoginIp  string    `gorm:"size:255;" json:"-"`
}
