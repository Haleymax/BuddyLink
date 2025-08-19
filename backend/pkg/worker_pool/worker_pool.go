package workerpool

import "sync"

type WorkerPool struct {
	workerChan chan struct{}
	wg         sync.WaitGroup
}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		workerChan: make(chan struct{}, workerCount),
	}
}

func (wp *WorkerPool) Submit(task func()) {
	wp.wg.Add(1)
	wp.workerChan <- struct{}{}
	go func() {
		defer func() {
			<-wp.workerChan
			wp.wg.Done()
		}()

		task()
	}()
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

func (wp *WorkerPool) Close() {
	close(wp.workerChan)
}
