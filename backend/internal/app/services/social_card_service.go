package services

import (
	"buddylink/internal/app/models"
	"buddylink/internal/app/repositories"
	"fmt"
)

type SocialCardService interface {
	AddCard(card models.SocialCard) error
}

type socialCardService struct {
	SocialRepo repositories.SocialCardRepository
}

func NewSocialCardService(repo repositories.SocialCardRepository) SocialCardService {
	return &socialCardService{
		SocialRepo: repo,
	}
}

func (s *socialCardService) AddCard(card models.SocialCard) error {
	exists, err := s.SocialRepo.Exists(card)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("card with ID %d already exists", card.ID)
	}

	if err = s.SocialRepo.Insert(card); err != nil {
		return fmt.Errorf("failed to insert card: %w", err)
	}
	return nil
}
