package controllers

import (
	"buddylink/internal/app/models"
	"buddylink/internal/app/services"
	"buddylink/pkg/cache_client"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type UserController struct {
	userService services.UserService
	stmpService services.StmpService
}

func NewUserController(userService services.UserService, stmpService services.StmpService) *UserController {
	return &UserController{
		userService: userService,
		stmpService: stmpService,
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

func (uc *UserController) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  false,
		})
		return
	}
	err := uc.userService.UpdateUser(user)
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
		"message": "success update user",
		"status":  true,
	})
}

func (uc *UserController) SendVerificationCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "email is empty",
			"status":  false,
		})
	}
	redisClient, err := cache_client.GetRedisClient()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	go func() {
		err := uc.stmpService.SendVerification(email, redisClient)
		if err != nil {
			log.Println(err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "send verification code, if not received, please try again later.",
		"status":  true,
	})
}
