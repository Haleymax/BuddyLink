package messagepool

import (
	"buddylink/config"
	workerpool "buddylink/pkg/worker_pool"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type TaskMessage struct {
	ID        string                 `json:"id"`
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

func (mp *MessagePool) GetMessageAsync(key string, callback func(TaskMessage, error)) {
	mp.pool.Submit(func() {
		msg, err := mp.GetMessageInternal(key)
		if callback != nil {
			callback(msg, err)
		}
	})
}

func (mp *MessagePool) SetMessageInternal(message TaskMessage) error {
	message.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	message.CreatedAt = time.Now()

	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	err = mp.client.Set(mp.ctx, mp.poolName+":"+message.ID, jsonData, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set message: %v", err)
	}
	return nil
}

func (mp *MessagePool) GetMessageInternal(key string) (TaskMessage, error) {
	data, err := mp.client.Get(mp.ctx, mp.poolName+":"+key).Bytes()
	if err != nil {
		return TaskMessage{}, err
	}
	var message TaskMessage
	if err = json.Unmarshal(data, &message); err != nil {
		return TaskMessage{}, err
	}
	return message, nil
}
