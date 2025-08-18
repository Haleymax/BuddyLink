package cache_client

import (
	"time"
)

type CardMessageData struct {
	UserID uint64 `json:"user_id"`
	CardID uint64 `json:"card_id"`
}

type MessageQueue struct {
	ID        int               `json:"id"`
	Data      map[string]string `json:"data"`
	Timestamp time.Time         `json:"timestamp"`
}

type RedisQueueProducer struct {
	cache redisClientImpl
}

func NewRedisMessageQueue(redis *redisClientImpl) *RedisQueueProducer {
	return &RedisQueueProducer{
		cache: *redis,
	}
}
