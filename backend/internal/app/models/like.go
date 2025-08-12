package models

type Like struct {
	BaseModel
	CardID uint64 `gorm:"index" json:"card_id"`
	UserID uint64 `gorm:"index" json:"user_id"`

	// Relationships
	User User       `gorm:"foreignKey:UserID" json:"user"`
	Card SocialCard `gorm:"foreignKey:CardID" json:"card"`
}
