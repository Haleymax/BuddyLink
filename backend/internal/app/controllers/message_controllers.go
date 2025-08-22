package controllers

import (
	"buddylink/internal/app/models"
	"buddylink/internal/app/repositories"
	"buddylink/internal/app/services"
	"log"
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

func (ctrl *MessageController) CreateMessage(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	if err := ctrl.messageService.CreateMessage(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to create message: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Message created successfully",
		"data":    message,
	})
}

func (ctrl *MessageController) GetMessageByID(c *gin.Context) {
	messageIDStr := c.Param("id")
	messageID, err := strconv.ParseUint(messageIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid message ID",
			"data":    nil,
		})
		return
	}

	message, err := ctrl.messageService.GetMessageByID(messageID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Message not found: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Message retrieved successfully",
		"data":    message,
	})
}

func (ctrl *MessageController) GetMessagesBySender(c *gin.Context) {
	senderIDStr := c.Param("sender_id")
	senderID, err := strconv.ParseUint(senderIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid sender ID",
			"data":    nil,
		})
		return
	}

	messages, err := ctrl.messageService.GetMessagesBySender(senderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to retrieve messages: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Messages retrieved successfully",
		"data":    messages,
	})
}

func (ctrl *MessageController) GetMessagesByReceiver(c *gin.Context) {
	receiverIDStr := c.Param("receiver_id")
	receiverID, err := strconv.ParseUint(receiverIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid receiver ID",
			"data":    nil,
		})
		return
	}

	messages, err := ctrl.messageService.GetMessagesByReceiver(receiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to retrieve messages: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Messages retrieved successfully",
		"data":    messages,
	})
}

func (ctrl *MessageController) GetMessagesByType(c *gin.Context) {
	messageType := c.Param("type")
	if messageType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Message type is required",
			"data":    nil,
		})
		return
	}

	messages, err := ctrl.messageService.GetMessagesByType(messageType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to retrieve messages: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Messages retrieved successfully",
		"data":    messages,
	})
}

func (ctrl *MessageController) GetMessagesByParams(c *gin.Context) {
	var params repositories.FindParam

	if senderIDStr := c.Query("sender_id"); senderIDStr != "" {
		if senderID, err := strconv.ParseUint(senderIDStr, 10, 64); err == nil {
			params.SenderID = senderID
		}
	}

	if receiverIDStr := c.Query("receiver_id"); receiverIDStr != "" {
		if receiverID, err := strconv.ParseUint(receiverIDStr, 10, 64); err == nil {
			params.ReceiverID = receiverID
		}
	}

	if messageType := c.Query("type"); messageType != "" {
		params.Type = messageType
	}

	messages, err := ctrl.messageService.GetMessagesByParams(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to retrieve messages: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Messages retrieved successfully",
		"data":    messages,
	})
}

func (ctrl *MessageController) UpdateMessage(c *gin.Context) {
	messageIDStr := c.Param("id")
	messageID, err := strconv.ParseUint(messageIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid message ID",
			"data":    nil,
		})
		return
	}

	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
		return
	}

	message.ID = messageID
	log.Println(message)
	if err := ctrl.messageService.UpdateMessage(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to update message: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Message updated successfully",
		"data":    message,
	})
}

func (ctrl *MessageController) DeleteMessage(c *gin.Context) {
	messageIDStr := c.Param("id")
	messageID, err := strconv.ParseUint(messageIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid message ID",
			"data":    nil,
		})
		return
	}

	if err := ctrl.messageService.DeleteMessage(messageID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to delete message: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Message deleted successfully",
		"data":    nil,
	})
}

func (ctrl *MessageController) GetAllMessages(c *gin.Context) {
	messages, err := ctrl.messageService.GetAllMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to retrieve messages: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "All messages retrieved successfully",
		"data":    messages,
	})
}
