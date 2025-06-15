package controllers

import (
	"buddylink/internal/app/models"
	"buddylink/internal/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Create(c *gin.Context) {
	err := uc.userService.CreateUserTable()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"status":  true,
	})
}

func (uc *UserController) AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	if user.Username == "" || user.Email == "" || user.PasswordHash == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "username or email or password is empty",
			"status":  false,
		})
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err := uc.userService.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success add user",
		"status":  true,
	})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	err := uc.userService.DeleteUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success delete user",
		"status":  true,
	})
}
