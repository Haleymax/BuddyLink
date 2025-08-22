package services

import (
	"buddylink/internal/app/models"
	"buddylink/internal/app/repositories"
	"fmt"
)

type MessageService interface {
	CreateMessage(message models.Message) error
	GetMessageByID(messageID uint64) (models.Message, error)
	GetMessagesBySender(senderID uint64) ([]models.Message, error)
	GetMessagesByReceiver(receiverID uint64) ([]models.Message, error)
	GetMessagesByType(messageType string) ([]models.Message, error)
	GetMessagesByParams(params repositories.FindParam) ([]models.Message, error)
	UpdateMessage(message models.Message) error
	DeleteMessage(messageID uint64) error
	GetAllMessages() ([]models.Message, error)
}

type messageServiceImpl struct {
	messageRepo repositories.MessageRepository
}

func NewMessageService(messageRepo repositories.MessageRepository) MessageService {
	return &messageServiceImpl{
		messageRepo: messageRepo,
	}
}

func (m *messageServiceImpl) CreateMessage(message models.Message) error {
	return m.messageRepo.Insert(message)
}

func (m *messageServiceImpl) GetMessageByID(messageID uint64) (models.Message, error) {
	return m.messageRepo.FindByID(messageID)
}

func (m *messageServiceImpl) GetMessagesBySender(senderID uint64) ([]models.Message, error) {
	return m.messageRepo.FindAllBysenderID(senderID)
}

func (m *messageServiceImpl) GetMessagesByReceiver(receiverID uint64) ([]models.Message, error) {
	return m.messageRepo.FindAllByReceiverID(receiverID)
}

func (m *messageServiceImpl) GetMessagesByType(messageType string) ([]models.Message, error) {
	return m.messageRepo.FindAllByType(messageType)
}

func (m *messageServiceImpl) GetMessagesByParams(params repositories.FindParam) ([]models.Message, error) {
	return m.messageRepo.FindByParams(params)
}

func (m *messageServiceImpl) UpdateMessage(message models.Message) error {
	exists, err := m.messageRepo.Exist(message.ID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("message with ID %d does not exist", message.ID)
	}
	return m.messageRepo.Update(message)
}

func (m *messageServiceImpl) DeleteMessage(messageID uint64) error {
	exists, err := m.messageRepo.Exist(messageID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("message with ID %d does not exist", messageID)
	}
	return m.messageRepo.Delete(messageID)
}

func (m *messageServiceImpl) GetAllMessages() ([]models.Message, error) {
	return m.messageRepo.FindAll()
}
