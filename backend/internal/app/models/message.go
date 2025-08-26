package models

type Message struct {
	BaseModel
	SenderID   uint64                 `gorm:"index" json:"sender_id"`
	ReceiverID uint64                 `gorm:"index" json:"receiver_id"`
	Type       string                 `gorm:"size:50;not null" json:"type"`
	Status     string                 `gorm:"size:50;default:'active'" json:"status"`
	Data       map[string]interface{} `gorm:"size:255;serializer:json" json:"data"`
}
