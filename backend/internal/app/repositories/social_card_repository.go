package repositories

import (
	"buddylink/internal/app/models"

	"gorm.io/gorm"
)

type SocialCardRepository interface {
	CreateTable() error
	Insert(entity models.SocialCard) error
	Update(oldData, entity models.SocialCard) error
	Delete(entity models.SocialCard) error
	DeleteByID(ID uint64) error
	DeleteByUserId(userID uint64) error
	FindByID(ID uint64) (models.SocialCard, error)
	FindByUserId(userID uint64) ([]models.SocialCard, error)
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

func (s SocialCardRepositoryImpl) Update(oldData, entity models.SocialCard) error {
	err := s.db.Model(&oldData).Updates(entity).Error
	return err
}

func (s SocialCardRepositoryImpl) Delete(entity models.SocialCard) error {
	return s.db.Delete(&entity).Error
}

func (s SocialCardRepositoryImpl) DeleteByID(ID uint64) error {
	return s.db.Where("id = ?", ID).Delete(&models.SocialCard{}).Error
}

func (s SocialCardRepositoryImpl) DeleteByUserId(userID uint64) error {
	return s.db.Where("user_id = ?", userID).Delete(&models.SocialCard{}).Error
}

func (s SocialCardRepositoryImpl) FindByID(ID uint64) (models.SocialCard, error) {
	var socialCard models.SocialCard
	err := s.db.Where("id = ?", ID).First(&socialCard).Error
	return socialCard, err
}

func (s SocialCardRepositoryImpl) FindByUserId(userID uint64) ([]models.SocialCard, error) {
	var socialCards []models.SocialCard
	err := s.db.Where("user_id = ?", userID).Find(&socialCards).Error
	if err != nil {
		return nil, err
	}
	return socialCards, nil
}

func (s SocialCardRepositoryImpl) FindAll() ([]models.SocialCard, error) {
	var socialCards []models.SocialCard
	err := s.db.Find(&socialCards).Error
	if err != nil {
		return nil, err
	}
	return socialCards, nil
}

func (s SocialCardRepositoryImpl) Exists(entity models.SocialCard) (bool, error) {
	var count int64
	err := s.db.Model(&models.SocialCard{}).Where("title = ?", entity.Title).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
