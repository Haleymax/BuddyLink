package cache_client

import (
	"buddylink/config"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte, expiration time.Duration) error
	Del(key string) error
	Close() error
}

type redisClientImpl struct {
	Client *redis.Client
}

func NewRedisClient(cfg config.RedisConfig) (RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Host,
		Password:     cfg.Password,
		DB:           cfg.Db,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		PoolSize:     5,
		MinIdleConns: 5,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &redisClientImpl{
		client,
	}, nil
}

func (r redisClientImpl) Get(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	val, err := r.Client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil // 键如果不存子，不判为error
		}
		return nil, err
	}
	return val, nil
}

func (r redisClientImpl) Set(key string, data []byte, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return r.Client.Set(ctx, key, data, expiration).Err()
}

func (r redisClientImpl) Del(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return r.Client.Del(ctx, key).Err()
}

func (r redisClientImpl) Close() error {
	return r.Client.Close()
}
