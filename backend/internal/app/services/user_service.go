package services

import (
	"buddylink/internal/app/models"
	"buddylink/internal/app/repositories"
	"fmt"
	"log"
)

type UserService interface {
	CreateUserTable() error
	AddUser(user models.User) error
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
