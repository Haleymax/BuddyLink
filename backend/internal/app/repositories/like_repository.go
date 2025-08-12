package repositories

import "buddylink/internal/app/models"

type LikeRepository interface {
	Create(like *models.Like) error
	Insert(like *models.Like) error
	Update(like *models.Like) error
	Delete(like *models.Like) error
	FindByID(like *models.Like) (*models.Like, error)
	FindByCardId(CardId uint) (*models.Like, error)
	FindAll() ([]*models.Like, error)
}
