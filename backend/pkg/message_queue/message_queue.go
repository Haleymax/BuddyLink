package messagequeue

import (
	"buddylink/config"
	workerpool "buddylink/pkg/worker_pool"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type TaskMessage struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageQueue struct {
	client      *redis.Client
	queueName   string
	pool        *workerpool.WorkerPool
	wg          sync.WaitGroup
	ctx         context.Context
	cancelFunc  context.CancelFunc
	messageChan chan TaskMessage
}

func NewMessageQueue(queueName string, workerCount int) (*MessageQueue, error) {
	ctx, cancel := context.WithCancel(context.Background())

	cfg := config.GetConfig()

	url := cfg.Redis.Host + ":" + cfg.Redis.Port

	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: cfg.Redis.Password,
		DB:       6,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		cancel()
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	pool := workerpool.NewWorkerPool(workerCount)

	return &MessageQueue{
		client:      client,
		queueName:   queueName,
		pool:        pool,
		wg:          sync.WaitGroup{},
		ctx:         ctx,
		cancelFunc:  cancel,
		messageChan: make(chan TaskMessage),
	}, nil
}

func (mq *MessageQueue) Produce(message TaskMessage) error {
	message.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	message.CreatedAt = time.Now()

	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	err = mq.client.LPush(mq.ctx, mq.queueName, jsonData).Err()
	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	log.Panicln("Message produced successfully")
	return nil
}

type MessageHandler func(message TaskMessage) error

func (mq *MessageQueue) StartConsumer(handler MessageHandler, numConsumers int) {
	mq.wg.Add(1)
	go mq.
}


func (mq *MessageQueue) messagePuller() {
	defer mq.wg.Done()

	for {
		select {
		case <-mq.ctx.Done():
		    log.Println("Message puller stopping...")
			close(mq.messageChan)
			return
		default:
			result, err := mq.client.BRPop(mq.ctx, 1*time.Second, mq.queueName).Result()
		}
	}
}