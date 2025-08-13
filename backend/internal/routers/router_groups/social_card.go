package router_groups

import (
	"buddylink/internal/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupSocialCardRouters(api *gin.RouterGroup, controller controllers.SocialCardController, authMiddleware gin.HandlerFunc) {
	socialCard := api.Group("social_card")
	socialCard.Use(authMiddleware)
	{
		socialCard.POST("/card", controller.AddCard)
		socialCard.PUT("/card/:card_id", controller.UpdateCard)
		socialCard.GET("/card/:card_id", controller.FindByCardId)
		socialCard.GET("/card/user/:user_id", controller.FindByUserId)
	}
}
