package repositories

import (
	"buddylink/internal/app/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	CreateTable() error
	Insert(user models.User) error
	Exist(user models.User) (bool, error)
	Delete(user models.User) error
	Update(oldData, NewDate models.User) error
	FindById(id uint64) (models.User, error)
	FindByUuid(uuid string) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByUsername(username string) ([]models.User, error)
	FindAll() ([]models.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUsrRepo(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) CreateTable() error {
	return r.db.AutoMigrate(&models.User{})
}

func (r *UserRepositoryImpl) Delete(user models.User) error {
	// 删除用户，目前根据用户的id去删除用户
	return r.db.Where("id = ?", user.ID).Delete(&models.User{}).Error
}

func (r *UserRepositoryImpl) Update(oldData, NewDate models.User) error {
	err := r.db.Model(&oldData).Updates(NewDate).Error
	return err
}

func (r *UserRepositoryImpl) FindById(id uint64) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) FindByUuid(uuid string) (models.User, error) {
	var user models.User
	err := r.db.Where("uuid = ?", uuid).First(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) FindByUsername(username string) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepositoryImpl) FindAll() ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepositoryImpl) Insert(user models.User) error {
	return r.db.Create(&user).Error
}
func (r *UserRepositoryImpl) Exist(user models.User) (bool, error) {
	var count int64

	// 性检查邮箱或UUID是否存在
	err := r.db.Model(&models.User{}).
		Where("email = ? OR uuid = ?", user.Email, user.Uuid).
		Count(&count).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}

	if count > 0 {
		log.Printf("user %s exists", user.Email)
		return true, nil
	}

	log.Printf("user %s does not exist", user.Email)
	return false, nil
}
