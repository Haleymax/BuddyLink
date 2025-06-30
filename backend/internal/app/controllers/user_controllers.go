package controllers

import (
	"buddylink/internal/app/models"
	"buddylink/internal/app/services"
	"buddylink/pkg/cache_client"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func (uc *UserController) Register(c *gin.Context) {
	emil := c.PostForm("email")
	userName := c.PostForm("username")
	password := c.PostForm("password")
	verificationCode := c.PostForm("verification_code")
	if emil == "" || userName == "" || password == "" || verificationCode == "" {
		log.Println("email or username or password or verification_code is empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "email or username or password or verification_code is empty",
			"status":  false,
		})
		return
	}

	redis_client, err := cache_client.GetRedisClient()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  false,
		})
		return

	}

	verificationCode_check, err := uc.stmpService.VerifyCode(emil, verificationCode, redis_client)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	if !verificationCode_check {
		log.Println("verification code is not correct")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "verification code is not correct",
			"status":  false,
		})
		return
	}
	newUser := models.User{
		Email:        emil,
		Username:     userName,
		PasswordHash: password,
	}

	err = uc.userService.RegisterUser(newUser)
	if err != nil {
		log.Println("failed to register user", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"status":  false,
		})
		return
	}

	log.Println("success register user", newUser.Username)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success register user",
		"status":  true,
	})

}
