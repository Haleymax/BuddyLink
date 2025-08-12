package router_groups

import (
	"buddylink/internal/app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupInitRouter(api *gin.RouterGroup, controller controllers.InitController) {
	init := api.Group("init")
	{
		init.GET("/init", controller.Init)
		init.GET("/index", controller.Index)
	}
}
