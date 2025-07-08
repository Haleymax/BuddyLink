package routers

import (
	"buddylink/internal/app/controllers"
	"buddylink/internal/app/repositories"
	"buddylink/internal/app/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, db *gorm.DB) {

	userRepo := repositories.NewUsrRepo(db)

	userService := services.NewUserService(userRepo)
	stmpService := services.NewStmpService()

	userController := controllers.NewUserController(userService, stmpService)
	api := router.Group("/api/v1")
	index := api.Group("/index")
	{
		index.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello world",
			})
		})
	}

	init := api.Group("/init")
	{
		init.GET("/", userController.Create)
	}

	user := api.Group("/user")
	{
		user.POST("/add", userController.AddUser)
		user.DELETE("/delete", userController.DeleteUser)
		user.PUT("/update", userController.UpdateUser)
		user.POST("/send_captcha", userController.SendVerificationCode)
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}
}
