package queue

import (
	"errors"
	"fmt"
	"time"

	"github.com/charlie4284/github-runner-manager-server/internal/job"
	"github.com/enriquebris/goconcurrentqueue"
)

type QueueFullError struct {
	Capacity int // capacity of the queue in number of jobs.
	Current  int // current number of jobs in queue.
}

func (e *QueueFullError) Error() string {
	return fmt.Sprintf("queue size %d full, current %d.", e.Capacity, e.Current)
}

var ErrQueueEmpty = errors.New("queue is empty.")

type RetryJob struct {
	Job         *job.Job  // The job to retry
	Retries     int       // Number of times the job has been retried.
	LastAttempt time.Time // Last time when the job was tried.
}

// Manager instance manages the queue.
type Manager struct {
	errorQueue *goconcurrentqueue.FIFO // ErrorQueue handles jobs that have failed and should be retried.
	queue      *goconcurrentqueue.FIFO // Queue handles jobs that should be tried.
}

func NewManager() *Manager {
	errorQueue := goconcurrentqueue.NewFIFO()
	queue := goconcurrentqueue.NewFIFO()
	return &Manager{
		errorQueue: errorQueue,
		queue:      queue,
	}
}

func (m *Manager) Enqueue(job *job.Job) error {
	m.queue.Enqueue(job)
	return nil
}

func (m *Manager) Dequeue() (*job.Job, error) {
	jobElement, err := m.queue.DequeueOrWaitForNextElement()
	if err != nil {
		return nil, err
	}
	job := jobElement.(*job.Job)
	return job, nil
}
