package router_groups

import (
	"buddylink/internal/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupMessageRoutes(router *gin.RouterGroup, controllers controllers.MessageController, authMiddleware gin.HandlerFunc) {
	messageGroup := router.Group("/messages")
	{
		messageGroup.POST("/message", controllers.AddMessage)
		messageGroup.GET("/message/keys/:user_id", controllers.GetAllKeys)
	}
}
