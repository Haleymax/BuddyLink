package controllers

import "github.com/gin-gonic/gin"

type TestController struct {
}

func NewTestController() *TestController {
	return &TestController{}
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
