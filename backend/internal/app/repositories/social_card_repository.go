package repositories

import (
	"buddylink/internal/app/models"

	"gorm.io/gorm"
)

type SocialCardRepository interface {
	CreateTable() error
	Insert(entity models.SocialCard) error
	Delete(entity models.SocialCard) error
	Update(entity models.SocialCard) error
	FindByID(ID uint) (models.SocialCard, error)
	FindByUserId(userID uint) ([]models.SocialCard, error)
	FindAll() ([]models.SocialCard, error)
	Exists(entity models.SocialCard) (bool, error)
}

type SocialCardRepositoryImpl struct {
	db *gorm.DB
}

func NewSocialCardRepository(db *gorm.DB) SocialCardRepository {
	return &SocialCardRepositoryImpl{
		db: db,
	}
}

func (s SocialCardRepositoryImpl) CreateTable() error {
	return s.db.AutoMigrate(&models.SocialCard{})
}

func (s SocialCardRepositoryImpl) Insert(entity models.SocialCard) error {
	return s.db.Create(&entity).Error
}

func (s SocialCardRepositoryImpl) Delete(entity models.SocialCard) error {
	return s.db.Delete(&entity).Error
}

func (s SocialCardRepositoryImpl) Update(entity models.SocialCard) error {
	return s.db.Save(&entity).Error
}

func (s SocialCardRepositoryImpl) FindByID(ID uint) (models.SocialCard, error) {
	var socialCard models.SocialCard
	err := s.db.Where("id = ?", ID).First(&socialCard).Error
	return socialCard, err
}

func (s SocialCardRepositoryImpl) FindByUserId(userID uint) ([]models.SocialCard, error) {
	//TODO implement me
	panic("implement me")
}

func (s SocialCardRepositoryImpl) FindAll() ([]models.SocialCard, error) {
	//TODO implement me
	panic("implement me")
}

func (s SocialCardRepositoryImpl) Exists(entity models.SocialCard) (bool, error) {
	var count int64
	err := s.db.Model(&models.SocialCard{}).Where("title = ?", entity.Title).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
