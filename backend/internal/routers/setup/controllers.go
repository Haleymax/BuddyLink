package setup

import (
	"buddylink/internal/app/controllers"
)

type Controllers struct {
	UserController controllers.UserController
	TestController controllers.TestController
}

func NewControllers(services *Services) *Controllers {
	userController := controllers.NewUserController(services.userService, services.stmpService)
	testController := controllers.NewTestController()

	return &Controllers{
		UserController: *userController,
		TestController: *testController,
	}
}
