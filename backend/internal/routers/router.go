package routers

import (
	"buddylink/internal/app/middleware"
	"buddylink/internal/routers/router_groups"
	"buddylink/internal/routers/setup"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, db *gorm.DB) {

	Repositories := setup.NewRepositories(db)

	Services := setup.NewServices(Repositories)

	Controllers := setup.NewControllers(Services)

	api := router.Group("/api/v1")

	router_groups.SetupInitRouter(api, Controllers.InitController)
	router_groups.SetupUserRouters(api, Controllers.UserController)
	router_groups.SetupTestRouters(api, Controllers.TestController, middleware.UserAuthMiddleware())
}
