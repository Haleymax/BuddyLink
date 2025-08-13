package controllers

import (
	"buddylink/internal/app/models"
	"buddylink/internal/app/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SocialCardController struct {
	SocialService services.SocialCardService
}

func NewSocialCardController(socialService services.SocialCardService) *SocialCardController {
	return &SocialCardController{
		SocialService: socialService,
	}
}

func (sc *SocialCardController) AddCard(c *gin.Context) {
	var socialCard models.SocialCard

	if err := c.ShouldBind(&socialCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if socialCard.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "userID is empty",
			"data":    nil,
		})
		return
	}

	if socialCard.Title == "" || socialCard.Content == "" || socialCard.Location == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "title, content or location is empty",
			"data":    nil,
		})
		return
	}

	socialCard.CreatedAt = time.Now()
	socialCard.UpdatedAt = time.Now()
	if err := sc.SocialService.AddCard(socialCard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    socialCard,
	})

}
