package services

import (
	messagepool "buddylink/pkg/message_pool"
)

type MessageService interface {
	SetMessage(message messagepool.TaskMessage) error
	GetMessage(key string) (messagepool.TaskMessage, error)
}

type messageServiceImpl struct {
	messagePool *messagepool.MessagePool
}

func NewMessageService(messagePool *messagepool.MessagePool) MessageService {
	return &messageServiceImpl{
		messagePool: messagePool,
	}
}

func (m *messageServiceImpl) SetMessage(message messagepool.TaskMessage) error {
	m.messagePool.SetMessageAsync(message, nil)
	return nil
}

func (m *messageServiceImpl) GetMessage(key string) (messagepool.TaskMessage, error) {
	return m.messagePool.GetMessage(key)
}
