package setup

import (
	"buddylink/internal/app/controllers"
)

type Controllers struct {
	UserController controllers.UserController
	TestController controllers.TestController
	InitController controllers.InitController
}

func NewControllers(services *Services) *Controllers {
	userController := controllers.NewUserController(services.userService, services.stmpService)
	InitController := controllers.NewInitController(services.initService)
	testController := controllers.NewTestController()

	return &Controllers{
		UserController: *userController,
		TestController: *testController,
		InitController: *InitController,
	}
}
