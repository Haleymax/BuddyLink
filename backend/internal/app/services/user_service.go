package services

import (
	"buddylink/internal/app/models"
	"buddylink/internal/app/repositories"
	"buddylink/pkg/jwtutil"
	"buddylink/pkg/object_storage"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
)

type UserService interface {
	CreateUserTable() error
	AddUser(user models.User) error
	DeleteUser(user models.User) error
	UpdateUser(user models.User) error
	RegisterUser(user models.User) error
	LoginUser(user models.User) (string, error)
}

type UserServiceImpl struct {
	UserRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
	}
}

func (s *UserServiceImpl) CreateUserTable() error {
	return s.UserRepo.CreateTable()
}

func (u *UserServiceImpl) AddUser(user models.User) error {
	exis, err := u.UserRepo.Exist(user)
	if err != nil {
		return err
	}
	if exis {
		return fmt.Errorf("user already exists")
	}
	if err = u.UserRepo.Insert(user); err != nil {
		log.Println("failed to insert user", err.Error())
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(user models.User) error {
	exis, err := u.UserRepo.Exist(user)
	if err != nil {
		return err
	}
	if !exis {
		return fmt.Errorf("user not exists")
	}

	dbUser, err := u.UserRepo.FindByUuid(user.Uuid)
	if err != nil {
		return err
	}

	if err = u.UserRepo.Delete(dbUser); err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) UpdateUser(user models.User) error {
	dbUser, err := u.UserRepo.FindById(user.ID)
	if err != nil {
		return err
	}
	err = u.UserRepo.Update(dbUser, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) RegisterUser(user models.User) error {
	exis, err := u.UserRepo.Exist(user)
	if err != nil {
		log.Println("failed to check user existence", err.Error())
		return err
	}

	if exis {
		log.Println("user already exists")
		return fmt.Errorf("user already exists")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Uuid = uuid.New().String()
	user.Status = "active"
	user.Role = "user"

	minioBucket := "profile"
	defaultPngFile := "default.png"

	minio_client := object_storage.GetMinioClient()
	url, err := minio_client.GetUrl(minioBucket, defaultPngFile)
	if err != nil {
		return err
	}
	user.Avatar = url

	if err = u.UserRepo.Insert(user); err != nil {
		log.Println("failed to insert user", err.Error())
		return err
	}
	log.Println("user registered successfully")
	return nil
}

func (u *UserServiceImpl) LoginUser(user models.User) (string, error) {
	/*
		验证用户是否存在，用户验证成功后，返回一个jwt的token
	*/
	dbUser, err := u.UserRepo.FindByUuid(user.Uuid)
	if err != nil {
		log.Println("failed to check user existence", err.Error())
		return "", err
	}

	if dbUser.Email != user.Email {
		return "", fmt.Errorf("email not match")
	}

	if dbUser.PasswordHash != user.PasswordHash {
		return "", fmt.Errorf("password not match")
	}

	if dbUser.Status == "failed" {
		return "", fmt.Errorf("user status is failed")
	}

	jwtUtil := jwtutil.NewJWTUtil("secret")
	expirationTime := time.Now().Add(24 * time.Hour)
	token, err := jwtUtil.GenerateToken(dbUser.Uuid, dbUser.Username, expirationTime)
	if err != nil {
		log.Println("failed to generate token", err.Error())
		return "", err
	}
	log.Println("successful generated token")
	return token, nil
}
