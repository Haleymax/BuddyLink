package services

import (
	"buddylink/internal/app/models"
	"buddylink/internal/app/repositories"
	"fmt"
)

type SocialCardService interface {
	AddCard(card models.SocialCard) error
	UpdateCard(oldDataID uint64, newData models.SocialCard) error
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
	// add card and check if it already exists (Card title is cannot be repeated)

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

func (s *socialCardService) UpdateCard(oldData_id uint64, newData models.SocialCard) error {
	// update card
	oldData, err := s.SocialRepo.FindByID(oldData_id)
	if err != nil {
		return fmt.Errorf("failed to find old card: %w", err)
	}

	if err := s.SocialRepo.Update(oldData, newData); err != nil {
		return fmt.Errorf("failed to update card: %w", err)
	}
	return nil
}
