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
	go mq.messagePuller()

	for i := 0; i < numConsumers; i++ {
		mq.wg.Add(1)
		go mq.messageProcessor(handler, i)
	}
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
			if err != nil {
				log.Printf("failed to pull message: %v", err)
				continue
			}

			if len(result) < 2 {
				continue
			}

			var message TaskMessage
			err = json.Unmarshal([]byte(result[1]), &message)
			if err != nil {
				log.Printf("failed to unmarshal message: %v", err)
				continue
			}
			select {
			case mq.messageChan <- message:
				log.Printf("pulled message %s", message.ID)
			case <-mq.ctx.Done():
				return
			}
		}
	}
}

func (mq *MessageQueue) messageProcessor(handler MessageHandler, workerId int) {
	defer mq.wg.Done()

	for message := range mq.messageChan {
		mq.pool.Submit(func() {
			mq.handleMessage(message, handler, workerId)
		})
	}
	log.Printf("Message %d stopped", workerId)
}

func (mq *MessageQueue) handleMessage(message TaskMessage, handler MessageHandler, workerId int) {
	log.Printf("Worker %d processing message %s", workerId, message.ID)

	start := time.Now()
	if err := handler(message); err != nil {
		log.Printf("failed to handle message %s: %v", message.ID, err)
	} else {
		log.Printf("Worker %d completed message %s in %v", workerId, message.ID, time.Since(start))
	}
}

func (mq *MessageQueue) Close() {
	mq.cancelFunc()
	mq.wg.Wait()
	mq.pool.Wait()
	mq.pool.Close()
	_ = mq.client.Close()
	log.Panicln("Message queue closed")
}

func (mq *MessageQueue) Length() (int64, error) {
	return mq.client.LLen(mq.ctx, mq.queueName).Result()
}

func (mq *MessageQueue) ProduceBatch(messages []TaskMessage) error {
	pipe := mq.client.Pipeline()
	for i := range messages {
		messages[i].ID = fmt.Sprintf("%d", time.Now().UnixNano()+int64(i))
		messages[i].CreatedAt = time.Now()

		jsonData, err := json.Marshal(messages[i])
		if err != nil {
			return fmt.Errorf("failed to marshal message: %v", err)
		}
		pipe.LPush(mq.ctx, mq.queueName, jsonData)
	}
	_, err := pipe.Exec(mq.ctx)
	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}
	log.Panicln("Message produced successfully")
	return nil
}
