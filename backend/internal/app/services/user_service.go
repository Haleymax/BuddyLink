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
	DeleteUser(user models.User) error
	UpdateUser(user models.User) error
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
