package router_groups

import (
	"buddylink/internal/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupMessageRoutes(router *gin.RouterGroup, controllers controllers.MessageController, authMiddleware gin.HandlerFunc) {
	messageGroup := router.Group("/messages")
	messageGroup.Use(authMiddleware)
	{
		messageGroup.POST("", controllers.CreateMessage)
		messageGroup.GET("/:id", controllers.GetMessageByID)
		messageGroup.PUT("/:id", controllers.UpdateMessage)
		messageGroup.DELETE("/:id", controllers.DeleteMessage)
		messageGroup.GET("", controllers.GetAllMessages)
		messageGroup.GET("/search", controllers.GetMessagesByParams)
		messageGroup.GET("/sender/:sender_id", controllers.GetMessagesBySender)
		messageGroup.GET("/receiver/:receiver_id", controllers.GetMessagesByReceiver)
		messageGroup.GET("/type/:type", controllers.GetMessagesByType)
		messageGroup.GET("/mine", controllers.GetMessagesByIdAndType)
	}
}
