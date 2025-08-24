package worker

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Task interface {
	Run() (any, error)
	GetName() string
}

type WorkerPool struct {
	workers map[int]*Worker
	// signals to know which workers are idle
	signal chan int
}

func NewWorkerPool(number uint) *WorkerPool {
	pool := &WorkerPool{
		workers: make(map[int]*Worker),
		signal:  make(chan int, number),
	}
	for i := range number {
		id := int(i)
		pool.signal <- id
		worker := Worker{
			id:     strconv.Itoa(id),
			signal: pool.signal,
			Logger: log.New(os.Stdout, fmt.Sprintf("[WORKER %d]", i), log.LstdFlags),
		}
		pool.workers[id] = &worker
	}
	return pool
}

func (wp *WorkerPool) Submit(tasks []Task) {
	for _, t := range tasks {
		id := <-wp.signal
		worker, ok := wp.workers[id]
		if !ok {
			panic(fmt.Sprintf("Worker not found: %d\n", id))
		}
		go func() {
			worker.Execute(t)
		}()
	}
}
