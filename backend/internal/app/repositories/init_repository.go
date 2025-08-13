package repositories

import (
	"buddylink/internal/app/models"

	"gorm.io/gorm"
)

type InitRepository interface {
	Init() error
}

type initRepositoryImpl struct {
	db *gorm.DB
}

func (i initRepositoryImpl) Init() error {
	err1 := i.db.AutoMigrate(models.User{})
	if err1 != nil {
		return err1
	}
	err2 := i.db.AutoMigrate(models.SocialCard{})
	if err2 != nil {
		return err2
	}
	err3 := i.db.AutoMigrate(models.Comment{})
	if err3 != nil {
		return err3
	}
	err4 := i.db.AutoMigrate(models.Like{})
	if err4 != nil {
		return err4
	}
	return nil
}

func NewInitRepository(db *gorm.DB) InitRepository {
	return &initRepositoryImpl{
		db: db,
	}
}
