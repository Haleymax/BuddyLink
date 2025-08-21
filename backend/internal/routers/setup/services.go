package setup

import (
	"buddylink/internal/app/services"
	messagepool "buddylink/pkg/message_pool"
)

type Services struct {
	userService    services.UserService
	stmpService    services.StmpService
	initService    services.InitService
	socialService  services.SocialCardService
	messageService services.MessageService
}

func NewServices(repo *Repositories) *Services {

	mq := messagepool.GetMessagePool()

	userService := services.NewUserService(repo.userRepository)
	initService := services.NewInitService(repo.initRepository)
	stmpService := services.NewStmpService()
	socialService := services.NewSocialCardService(repo.socialCardRepository)
	messageService := services.NewMessageService(mq)

	return &Services{
		initService:    initService,
		userService:    userService,
		stmpService:    stmpService,
		socialService:  socialService,
		messageService: messageService,
	}
}
