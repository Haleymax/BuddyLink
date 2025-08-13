package setup

import (
	"buddylink/internal/app/controllers"
)

type Controllers struct {
	InitController       controllers.InitController
	UserController       controllers.UserController
	SocialCardController controllers.SocialCardController
	TestController       controllers.TestController
}

func NewControllers(services *Services) *Controllers {
	InitController := controllers.NewInitController(services.initService)
	userController := controllers.NewUserController(services.userService, services.stmpService)
	socialCardController := controllers.NewSocialCardController(services.socialService)
	testController := controllers.NewTestController()

	return &Controllers{
		UserController:       *userController,
		TestController:       *testController,
		InitController:       *InitController,
		SocialCardController: *socialCardController,
	}
}
