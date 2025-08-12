type SocialCard struct {
	BaseModel
	UserID      uint64    `gorm:"index" json:"user_id"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	Images	  string    `gorm:"type:text" json:"images"`
	LikesCount uint64	`gorm:"default:0" json:"likes_count"`
	ViewsCount uint64	`gorm:"default:0" json:"views_count"`
	CommentsCount uint64	`gorm:"default:0" json:"comments_count"`
	IsPrivate  bool      `gorm:"default:false" json:"is_private"`
	Location   string    `gorm:"size:255" json:"location"`
	Date 	 time.Time `gorm:"type:timestamp(3);default:CURRENT_TIMESTAMP(3)" json:"date"`
	Tags	   string    `gorm:"type:text" json:"tags"`

	// Relationships
	User       User      `gorm:"foreignKey:UserID" json:"user"`
	Comments   []Comment `gorm:"foreignKey:CardID" json:"comments"`
	Likes      []Like	`gorm:"foreignKey:CardID" json:"likes"`
}
