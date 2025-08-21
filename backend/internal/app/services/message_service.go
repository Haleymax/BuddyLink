package services

import (
	messagepool "buddylink/pkg/message_pool"
)

type MessageService interface {
	SetMessage(message messagepool.TaskMessage) error
	GetMessage(userId uint64, messageType string, messageId string) (messagepool.TaskMessage, error)
	GetAllKeys(userID uint64, messageType string) ([]string, error)
	GetAllMessage(userID uint64, messageType string) ([]messagepool.TaskMessage, error)
	RemoveMessage(userId uint64, messageType string, messageId string) error
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
	errChan := make(chan error, 1)
	m.messagePool.SetMessageAsync(message, func(err error) {
		errChan <- err
	})
	err := <-errChan
	return err
}

func (m *messageServiceImpl) GetMessage(userId uint64, messageType string, messageId string) (messagepool.TaskMessage, error) {
	var message messagepool.TaskMessage
	var err error
	errChan := make(chan error, 1)
	m.messagePool.GetMessageAsync(userId, messageType, messageId, func(msg messagepool.TaskMessage, err error) {
		if err != nil {
			errChan <- err
		} else {
			message = msg
			errChan <- nil
		}
	})
	err = <-errChan
	return message, err
}

func (m *messageServiceImpl) GetAllKeys(userID uint64, messageType string) ([]string, error) {
	keysChan := make(chan []string, 1)
	errChan := make(chan error, 1)

	m.messagePool.GetAllKeysAsync(userID, messageType, func(keys []string, err error) {
		if err != nil {
			errChan <- err
		} else {
			keysChan <- keys
		}
	})

	select {
	case keys := <-keysChan:
		return keys, nil
	case err := <-errChan:
		return nil, err
	}
}

func (m *messageServiceImpl) GetAllMessage(userID uint64, messageType string) ([]messagepool.TaskMessage, error) {
	keys, err := m.GetAllKeys(userID, messageType)
	if err != nil {
		return nil, err
	}

	var messages []messagepool.TaskMessage
	for _, key := range keys {
		message, err := m.GetMessage(userID, messageType, key)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func (m *messageServiceImpl) RemoveMessage(userId uint64, messageType string, messageId string) error {
	errChan := make(chan error, 1)
	m.messagePool.RemoveMessage(userId, messageType, messageId, func(result int64, err error) {
		if err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	})
	return <-errChan
}
