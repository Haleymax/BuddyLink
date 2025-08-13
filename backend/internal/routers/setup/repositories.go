package setup

import (
	"buddylink/internal/app/repositories"

	"gorm.io/gorm"
)

type Repositories struct {
	initRepository       repositories.InitRepository
	userRepository       repositories.UserRepository
	socialCardRepository repositories.SocialCardRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	userRepo := repositories.NewUsrRepo(db)
	initRepo := repositories.NewInitRepository(db)
	socialCardRepository := repositories.NewSocialCardRepository(db)
	return &Repositories{
		userRepository:       userRepo,
		initRepository:       initRepo,
		socialCardRepository: socialCardRepository,
	}
}
