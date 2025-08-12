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
	//TODO implement me
	panic("implement me")
}

func (s SocialCardRepositoryImpl) Delete(entity models.SocialCard) error {
	//TODO implement me
	panic("implement me")
}

func (s SocialCardRepositoryImpl) Update(entity models.SocialCard) error {
	//TODO implement me
	panic("implement me")
}

func (s SocialCardRepositoryImpl) FindByID(ID uint) (models.SocialCard, error) {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}
