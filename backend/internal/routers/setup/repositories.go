package setup

import (
	"buddylink/internal/app/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
	userRepository repositories.UserRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	userRepo := repositories.NewUsrRepo(db)
	return &Repositories{
		userRepository: userRepo,
	}
}
