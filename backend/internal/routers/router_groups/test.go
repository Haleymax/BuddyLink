package router_groups

import (
	"buddylink/internal/app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTestRouters(api *gin.RouterGroup, controller controllers.TestController, authMiddleware gin.HandlerFunc) {
	test := api.Group("test")
	test.Use(authMiddleware)
	{
		test.GET("/data", controller.GetTestData)
	}
}
