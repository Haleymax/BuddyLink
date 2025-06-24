package cache_client

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte, expiration time.Duration) error
}

type redisClientImpl struct {
	Client *redis.Client
}
