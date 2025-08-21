package services

import (
	messagepool "buddylink/pkg/message_pool"
)

type MessageService interface {
	SetMessage(message messagepool.TaskMessage) error
	GetMessage(key string) (messagepool.TaskMessage, error)
	GetAllKeys(userID uint64) ([]string, error)
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

func (m *messageServiceImpl) GetMessage(key string) (messagepool.TaskMessage, error) {
	return m.messagePool.GetMessage(key)
}

func (m *messageServiceImpl) GetAllKeys(userID uint64) ([]string, error) {
	keysChan := make(chan []string, 1)
	errChan := make(chan error, 1)

	m.messagePool.GetAllKeysAsync(userID, func(keys []string, err error) {
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
