package controllers

import (
	"buddylink/internal/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InitController struct {
	initService services.InitService
}

func NewInitController(initService services.InitService) *InitController {
	return &InitController{
		initService: initService,
	}
}

func (ctrl *InitController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"code":    http.StatusOK,
		"message": "ok",
		"data":    nil,
	})
}

func (ctrl *InitController) Init(c *gin.Context) {
	err := ctrl.initService.Init()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    nil,
	})
}
