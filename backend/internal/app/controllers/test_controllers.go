package controllers

import (
	messagequeue "buddylink/pkg/message_queue"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type MessageData struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type TestController struct {
	mq  *messagequeue.MessageQueue
	err error
}

func NewTestController() *TestController {
	mq, err := messagequeue.NewMessageQueue("test_queue", 5)
	if err != nil {
		return nil
	}
	return &TestController{
		mq:  mq,
		err: err,
	}
}

// GetTestData 返回测试数据
func (tc *TestController) GetTestData(c *gin.Context) {

	email, emailExists := c.Get("email")
	if !emailExists {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "email not found in context",
			"status":  false,
		})
		return
	}
	id, idExists := c.Get("id")
	if !idExists {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "id not found in context",
			"status":  false,
		})
		return
	}
	username, nameExists := c.Get("username")
	if !nameExists {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "username not found in context",
			"status":  false,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"status":  true,
		"data": gin.H{
			"email":    email,
			"id":       id,
			"username": username,
		},
	})

}

func (tc *TestController) Producer(c *gin.Context) {

	var messageData MessageData
	if err := c.ShouldBindJSON(&messageData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	message := messagequeue.TaskMessage{
		Type: messageData.Type,
		Data: messageData.Data,
	}

	err := tc.mq.Produce(message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "failed to produce message",
			"status":  false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "producer",
		"status":  true,
	})
}

func (tc *TestController) Consumer(c *gin.Context) {
	if tc.mq == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "message queue not initialized",
			"status":  false,
		})
		return
	}

	queueLength, err := tc.mq.Length()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "failed to get queue length: " + err.Error(),
			"status":  false,
		})
		return
	}

	var processedMessages []map[string]interface{}
	var errors []string

	handler := func(message messagequeue.TaskMessage) error {
		var data map[string]interface{}
		if message.Data != nil {
			data = message.Data
		} else {
			data = make(map[string]interface{})
		}

		processedMessage := map[string]interface{}{
			"id":           message.ID,
			"type":         message.Type,
			"data":         data,
			"created_at":   message.CreatedAt,
			"processed_at": time.Now(),
		}

		processedMessages = append(processedMessages, processedMessage)
		return nil
	}

	if queueLength > 0 {
		tc.mq.StartConsumer(handler, 1)
	}

	response := gin.H{
		"code":    200,
		"message": "consumer processed successfully",
		"status":  true,
		"data": gin.H{
			"queue_length":       queueLength,
			"processed_count":    len(processedMessages),
			"processed_messages": processedMessages,
		},
	}

	if len(errors) > 0 {
		response["errors"] = errors
	}

	c.JSON(http.StatusOK, response)
}
