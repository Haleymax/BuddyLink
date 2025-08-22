package repositories

import (
	"buddylink/internal/app/models"
	"time"

	"gorm.io/gorm"
)

type FindParam struct {
	SenderID   uint64
	ReceiverID uint64
	Type       string
}

type MessageRepository interface {
	Insert(message models.Message) error
	Delete(message_id uint64) error
	Update(message models.Message) error
	FindByID(message_id uint64) (models.Message, error)
	FindAll() ([]models.Message, error)
	FindAllBysenderID(sender_id uint64) ([]models.Message, error)
	FindAllByReceiverID(receiver_id uint64) ([]models.Message, error)
	FindAllByType(message_type string) ([]models.Message, error)
	FindByParams(params FindParam) ([]models.Message, error)
	Exist(message_id uint64) (bool, error)
}

type MessageRepositoryImpl struct {
	db *gorm.DB
}

// Delete implements MessageRepository.
func (m *MessageRepositoryImpl) Delete(message_id uint64) error {
	return m.db.Delete(&models.Message{}, message_id).Error
}

// FindAll implements MessageRepository.
func (m *MessageRepositoryImpl) FindAll() ([]models.Message, error) {
	var messages []models.Message
	err := m.db.Find(&messages).Error
	return messages, err
}

// FindAllByReceiverID implements MessageRepository.
func (m *MessageRepositoryImpl) FindAllByReceiverID(receiver_id uint64) ([]models.Message, error) {
	var messages []models.Message
	err := m.db.Where("receiver_id = ?", receiver_id).Find(&messages).Error
	return messages, err
}

// FindAllByType implements MessageRepository.
func (m *MessageRepositoryImpl) FindAllByType(message_type string) ([]models.Message, error) {
	var messages []models.Message
	err := m.db.Where("type = ?", message_type).Find(&messages).Error
	return messages, err
}

// FindAllBysenderID implements MessageRepository.
func (m *MessageRepositoryImpl) FindAllBysenderID(sender_id uint64) ([]models.Message, error) {
	var messages []models.Message
	err := m.db.Where("sender_id = ?", sender_id).Find(&messages).Error
	return messages, err
}

// FindByID implements MessageRepository.
func (m *MessageRepositoryImpl) FindByID(message_id uint64) (models.Message, error) {
	var message models.Message
	err := m.db.First(&message, message_id).Error
	return message, err
}

// FindByParams implements MessageRepository.
func (m *MessageRepositoryImpl) FindByParams(params FindParam) ([]models.Message, error) {
	var messages []models.Message
	query := m.db

	if params.SenderID != 0 {
		query = query.Where("sender_id = ?", params.SenderID)
	}

	if params.ReceiverID != 0 {
		query = query.Where("receiver_id = ?", params.ReceiverID)
	}

	if params.Type != "" {
		query = query.Where("type = ?", params.Type)
	}

	err := query.Find(&messages).Error
	return messages, err
}

// Insert implements MessageRepository.
func (m *MessageRepositoryImpl) Insert(message models.Message) error {
	// 确保时间字段正确设置
	now := time.Now()
	message.CreatedAt = now
	message.UpdatedAt = now
	return m.db.Create(&message).Error
}

// Update implements MessageRepository.
func (m *MessageRepositoryImpl) Update(message models.Message) error {
	message.UpdatedAt = time.Now()
	return m.db.Model(&message).Select("sender_id", "receiver_id", "type", "data", "updated_at").Updates(message).Error
}

func (m *MessageRepositoryImpl) Exist(message_id uint64) (bool, error) {
	var count int64
	err := m.db.Model(&models.Message{}).Where("id = ?", message_id).Count(&count).Error
	return count > 0, err
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &MessageRepositoryImpl{
		db: db,
	}
}
