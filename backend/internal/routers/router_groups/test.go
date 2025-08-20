package router_groups

import (
	"buddylink/internal/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupTestRouters(api *gin.RouterGroup, controller controllers.TestController, authMiddleware gin.HandlerFunc) {
	test := api.Group("test")
	{
		test.GET("/data", authMiddleware, controller.GetTestData)
		test.POST("/producer", controller.Producer)
		test.GET("/consumer", controller.Consumer)
	}
}
