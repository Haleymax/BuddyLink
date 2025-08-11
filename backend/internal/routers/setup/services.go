package setup

import (
	"buddylink/internal/app/services"
)

type Services struct {
	userService services.UserService
	stmpService services.StmpService
}

func NewServices(repo *Repositories) *Services {
	userService := services.NewUserService(repo.userRepository)
	stmpService := services.NewStmpService()
	return &Services{
		userService: userService,
		stmpService: stmpService,
	}
}
