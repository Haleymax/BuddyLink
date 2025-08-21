package messagepool

import (
	"buddylink/config"
	workerpool "buddylink/pkg/worker_pool"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type TaskMessage struct {
	ID        string                 `json:"id"`
	Sender    uint64                 `json:"sender"`
	Receiver  uint64                 `json:"receiver"`
	Type      string                 `json:"type"`
	Data      map[string]interface{} `json:"data"`
	CreatedAt time.Time              `json:"created_at"`
}

type MessagePool struct {
	client      *redis.Client
	poolName    string
	pool        *workerpool.WorkerPool
	wg          sync.WaitGroup
	ctx         context.Context
	cancelFunc  context.CancelFunc
	messageChan chan TaskMessage
}

func NewMessagePool(poolName string) (*MessagePool, error) {

	ctx, cancel := context.WithCancel(context.Background())
	cfg := config.GetConfig()
	url := cfg.Redis.Host + ":" + cfg.Redis.Port

	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: cfg.Redis.Password,
		DB:       cfg.Message.PoolDB,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		cancel()
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	pool := workerpool.NewWorkerPool(cfg.Message.WorkerCount)

	messageChan := make(chan TaskMessage)

	return &MessagePool{
		client:      client,
		poolName:    poolName,
		pool:        pool,
		wg:          sync.WaitGroup{},
		ctx:         ctx,
		cancelFunc:  cancel,
		messageChan: messageChan,
	}, nil
}

// SetMessageAsync 使用工作池异步设置消息
func (mp *MessagePool) SetMessageAsync(message TaskMessage, callback func(err error)) {
	mp.pool.Submit(func() {
		err := mp.SetMessageInternal(message)
		if callback != nil {
			callback(err)
		}
	})
}

func (mp *MessagePool) GetMessageAsync(user_id uint64, message_type string, message_id string, callback func(TaskMessage, error)) {
	mp.pool.Submit(func() {
		key := fmt.Sprintf("%s:%s:%d:%s", mp.poolName, message_type, user_id, message_id)
		msg, err := mp.GetMessageInternal(key)
		if callback != nil {
			callback(msg, err)
		}
	})
}

func (mp *MessagePool) GetAllKeysAsync(userId uint64, messageType string, callback func([]string, error)) {
	mp.pool.Submit(func() {
		keys, err := mp.GetAllKeys(userId, messageType)
		if callback != nil {
			callback(keys, err)
		}
	})
}

func (mp *MessagePool) GetAllKeys(userId uint64, messageType string) ([]string, error) {
	var message_ids []string

	pattern := fmt.Sprintf("%s:%s:%d:*", mp.poolName, messageType, userId)

	iter := mp.client.Scan(mp.ctx, 0, pattern, 0).Iterator()
	for iter.Next(mp.ctx) {
		key := iter.Val()
		keys := strings.Split(key, ":")
		message_ids = append(message_ids, keys[len(keys)-1])
	}

	if err := iter.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan keys: %v", err)
	}

	return message_ids, nil
}

func (mp *MessagePool) SetMessageInternal(message TaskMessage) error {
	message.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	message.CreatedAt = time.Now()

	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	key := fmt.Sprintf("%s:%s:%d:%s", mp.poolName, message.Type, message.Receiver, message.ID)

	err = mp.client.Set(mp.ctx, key, jsonData, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set message: %v", err)
	}
	return nil
}

func (mp *MessagePool) GetMessageInternal(key string) (TaskMessage, error) {
	data, err := mp.client.Get(mp.ctx, key).Bytes()
	if err != nil {
		return TaskMessage{}, err
	}
	var message TaskMessage
	if err = json.Unmarshal(data, &message); err != nil {
		return TaskMessage{}, err
	}
	return message, nil
}

// 同步操作
func (mp *MessagePool) SetMessage(message TaskMessage) error {
	return mp.SetMessageInternal(message)
}

func (mp *MessagePool) GetMessage(key string) (TaskMessage, error) {
	return mp.GetMessageInternal(key)
}

func (mp *MessagePool) Length() (int64, error) {
	return mp.client.DBSize(mp.ctx).Result()
}

func (mp *MessagePool) Wait() {
	mp.wg.Wait()
}

func (mp *MessagePool) Close() {
	mp.cancelFunc()
	mp.pool.Close()
	close(mp.messageChan)
}

func GetMessagePool() *MessagePool {
	cfg := config.GetConfig()
	pool, err := NewMessagePool(cfg.Message.PoolName)
	if err != nil {
		return nil
	}
	return pool
}
