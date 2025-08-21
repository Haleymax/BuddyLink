package controllers

import (
	"buddylink/internal/app/services"
	messagepool "buddylink/pkg/message_pool"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	messageService services.MessageService
}

func NewMessageController(messageService services.MessageService) *MessageController {
	return &MessageController{
		messageService: messageService,
	}
}

func (ctrl *MessageController) AddMessage(c *gin.Context) {
	var message messagepool.TaskMessage
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if err := ctrl.messageService.SetMessage(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "message added",
		"data":    nil,
	})
}
