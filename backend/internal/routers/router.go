package routers

import (
	"buddylink/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, config config.Config, db *gorm.DB) {
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
