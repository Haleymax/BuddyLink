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
	userIDStr := c.Query("user_id")
	messageType := c.Query("type")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "user ID is required",
			"data":    nil,
		})
		return
	}
	if messageType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "type is required",
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
	keys, err := ctrl.messageService.GetAllKeys(userID, messageType)
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

func (ctrl *MessageController) GetMessage(c *gin.Context) {
	userIDStr := c.Query("user_id")
	messageType := c.Query("type")
	messageId := c.Query("message_id")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "user ID is required",
			"data":    nil,
		})
		return
	}
	if messageType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "type is required",
			"data":    nil,
		})
		return
	}
	if messageId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "message ID is required",
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
	message, err := ctrl.messageService.GetMessage(userID, messageType, messageId)
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
		"message": "message retrieved",
		"data":    message,
	})
}

func (ctrl *MessageController) GetAllMessages(c *gin.Context) {
	userIDStr := c.Query("user_id")
	messageType := c.Query("type")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "user ID is required",
			"data":    nil,
		})
		return
	}
	if messageType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "type is required",
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
	messages, err := ctrl.messageService.GetAllMessage(userID, messageType)
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
		"message": "messages retrieved",
		"data":    messages,
	})
}
