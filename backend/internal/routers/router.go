package routers

import (
	"buddylink/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, config config.Config) {
	api := router.Group("/api/v1")
	index := api.Group("/index")
	{
		index.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello world",
			})
		})
	}
}
