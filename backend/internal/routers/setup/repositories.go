package setup

import (
	"buddylink/internal/app/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
	userRepository repositories.UserRepository
	initRepository repositories.InitRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	userRepo := repositories.NewUsrRepo(db)
	initRepo := repositories.NewInitRepository(db)
	return &Repositories{
		userRepository: userRepo,
		initRepository: initRepo,
	}
}
