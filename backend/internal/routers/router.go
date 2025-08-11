package routers

import (
	"buddylink/internal/app/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, db *gorm.DB) {

	Repositories := NewRepositories(db)

	Services := NewServices(Repositories)

	Controllers := NewControllers(Services)

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
		init.GET("/", Controllers.UserController.Create)
	}

	user := api.Group("/user")
	{
		user.POST("/add", Controllers.UserController.AddUser)
		user.DELETE("/delete", Controllers.UserController.DeleteUser)
		user.PUT("/update", Controllers.UserController.UpdateUser)
		user.POST("/send_captcha", Controllers.UserController.SendVerificationCode)
		user.POST("/register", Controllers.UserController.Register)
		user.POST("/login", Controllers.UserController.Login)
	}

	test := api.Group("test")
	test.Use(middleware.UserAuthMiddleware())
	{
		test.GET("/data", Controllers.TestController.GetTestData)
	}
}
