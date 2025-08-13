package router_groups

import (
	"buddylink/internal/app/controllers"
	"buddylink/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func SetupUserRouters(api *gin.RouterGroup, controller controllers.UserController) {
	user := api.Group("/user")
	{
		user.POST("/add", controller.AddUser)
		user.DELETE("/delete", controller.DeleteUser)
		user.PUT("/update", controller.UpdateUser)
		user.POST("/send_captcha", controller.SendVerificationCode)
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
		user.GET("/get_user", middleware.UserAuthMiddleware(), controller.GetUser)
	}
}
