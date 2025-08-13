package setup

import (
	"buddylink/internal/app/services"
)

type Services struct {
	userService   services.UserService
	stmpService   services.StmpService
	initService   services.InitService
	socialService services.SocialCardService
}

func NewServices(repo *Repositories) *Services {
	userService := services.NewUserService(repo.userRepository)
	initService := services.NewInitService(repo.initRepository)
	stmpService := services.NewStmpService()
	socialService := services.NewSocialCardService(repo.socialCardRepository)
	return &Services{
		initService:   initService,
		userService:   userService,
		stmpService:   stmpService,
		socialService: socialService,
	}
}
