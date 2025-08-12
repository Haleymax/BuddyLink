package services

import "buddylink/internal/app/repositories"

type InitService interface {
	Init() error
}

type InitServiceImpl struct {
	InitRepo repositories.InitRepository
}

func NewInitService(initRepo repositories.InitRepository) InitService {
	return &InitServiceImpl{
		InitRepo: initRepo,
	}
}

func (s *InitServiceImpl) Init() error {
	return s.InitRepo.Init()
}
