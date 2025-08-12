
type Comment struct {
	BaseModel
	CardID	 uint64    `gorm:"index" json:"card_id"`
	UserID	 uint64    `gorm:"index" json:"user_id"`
	Content  string    `gorm:"type:text;not null" json:"content"`
	ParentID uint64    *uint 
	ReplyCount uint64    `gorm:"default:0" json:"reply_count"`

	// Relationships
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Card   SocialCard `gorm:"foreignKey:CardID" json:"card"`
}