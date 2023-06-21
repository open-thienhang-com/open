package command

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

var workerPool = make(chan chan Command)

// Worker contains the basic info to manage commands
type Worker struct {
	WorkerPool     chan chan Command
	JobChannel     chan Command
	CommandHandler CommandHandlerRegister
}

// Bus stores the command handler
type Bus struct {
	CommandHandler CommandHandlerRegister
	maxWorkers     int
}

// Start initialize a worker ready to receive jobs
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			job := <-w.JobChannel
			handler, err := w.CommandHandler.Get(job)
			if err != nil {
				continue
			}

			if !job.IsValid() {
				continue
			}

			if err = handler.Handle(job); err != nil {
				//TODO: log the error
				log.Error(err)
				log.Error(job)
			}
		}
	}()
}

// NewWorker initialize the values of worker and start it
func NewWorker(commandHandler CommandHandlerRegister) {
	w := Worker{
		WorkerPool:     workerPool,
		CommandHandler: commandHandler,
		JobChannel:     make(chan Command),
	}

	w.Start()
}

// HandleCommand ad a job to the queue
func (b *Bus) HandleCommand(cmd Command) {
	go func(c Command) {
		fmt.Println(cmd)
		workerJobQueue := <-workerPool
		workerJobQueue <- c
	}(cmd)
}

// NewBus return a bus with command handler register
func NewBus(register CommandHandlerRegister, maxWorkers int) *Bus {
	b := &Bus{
		CommandHandler: register,
		maxWorkers:     maxWorkers,
	}

	// start the bus
	b.Start()
	return b
}

// Start the bus
func (b *Bus) Start() {
	for i := 0; i < b.maxWorkers; i++ {
		NewWorker(b.CommandHandler)
	}
}
