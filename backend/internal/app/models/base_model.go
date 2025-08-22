package models

import "time"

type BaseModel struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp(3);default:CURRENT_TIMESTAMP(3);not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp(3);default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);not null" json:"updated_at"`
}
