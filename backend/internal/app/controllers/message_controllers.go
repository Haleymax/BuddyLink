package controllers

import (
	"buddylink/internal/app/services"
	messagepool "buddylink/pkg/message_pool"
	"net/http"
	"strconv"

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

func (ctrl *MessageController) GetAllKeys(c *gin.Context) {
	userIDStr := c.Param("user_id")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "user ID is required",
			"data":    nil,
		})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid user ID",
			"data":    nil,
		})
		return
	}
	keys, err := ctrl.messageService.GetAllKeys(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "keys retrieved",
		"data":    keys,
	})
}
