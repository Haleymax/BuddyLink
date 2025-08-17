package controllers

import (
	"buddylink/internal/app/models"
	"buddylink/internal/app/services"
	"net/http"
	"strconv"
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

	if socialCard.Title == "" || socialCard.Content == "" || socialCard.Location == nil || len(socialCard.Location) == 0 {
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

func (sc *SocialCardController) UpdateCard(c *gin.Context) {
	var socialCard models.SocialCard
	if err := c.ShouldBind(&socialCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	card_id := c.Param("card_id")
	if card_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "old data ID is required",
			"data":    nil,
		})
		return
	}

	card_id_uint, err := strconv.ParseUint(card_id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid card ID",
			"data":    nil,
		})
		return
	}

	if err := sc.SocialService.UpdateCard(card_id_uint, socialCard); err != nil {
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
	})
}

func (sc *SocialCardController) GetAllCard(c *gin.Context) {
	cards, err := sc.SocialService.FindAllCards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if cards == nil {
		cards = []map[string]interface{}{}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    cards,
	})
}

func (sc *SocialCardController) FindByUserId(c *gin.Context) {
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

	cards, err := sc.SocialService.FindByUserId(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if cards == nil {
		cards = []map[string]interface{}{}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    cards,
	})
}

func (sc *SocialCardController) FindByCardId(c *gin.Context) {
	cardIDStr := c.Param("card_id")
	if cardIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "card ID is required",
			"data":    nil,
		})
		return
	}

	cardID, err := strconv.ParseUint(cardIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid card ID",
			"data":    nil,
		})
		return
	}

	card, err := sc.SocialService.FindById(cardID)
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
		"message": "success",
		"data":    card,
	})
}

func (sc *SocialCardController) DeleteCard(c *gin.Context) {
	var card models.SocialCard
	if err := c.ShouldBind(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if err := sc.SocialService.Delete(card); err != nil {
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
	})
}

func (sc *SocialCardController) DeleteByCardId(c *gin.Context) {
	cardIDStr := c.Param("card_id")
	if cardIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "card ID is required",
			"data":    nil,
		})
		return
	}

	cardID, err := strconv.ParseUint(cardIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid card ID",
			"data":    nil,
		})
		return
	}

	if err := sc.SocialService.DeleteByID(cardID); err != nil {
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
	})
}

func (sc *SocialCardController) DeleteByUserId(c *gin.Context) {
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

	if err := sc.SocialService.DeleteByUserId(userID); err != nil {
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
	})
}
