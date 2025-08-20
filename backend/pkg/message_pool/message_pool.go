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

func (mp *MessagePool) Produce(message TaskMessage) error {
	message.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	message.CreatedAt = time.Now()

	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	err = mp.client.Set(mp.ctx, mp.poolName+":"+message.ID, jsonData, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to store message in Redis: %v", err)
	}

	return nil
}
