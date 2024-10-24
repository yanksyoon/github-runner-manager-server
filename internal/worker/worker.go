package worker

import (
	"fmt"

	"github.com/charlie4284/github-runner-manager-server/internal/job"
	"github.com/charlie4284/github-runner-manager-server/internal/queue"
)

func Start(numThreads int, queueManager *queue.Manager) {
	for range numThreads {
		go startWork(queueManager)
	}
}

func startWork(queueManager *queue.Manager) {
	for {
		job, err := queueManager.Dequeue()
		if err != nil {
			// change to logger
			fmt.Println("Failed to dequeue job.")
			continue
		}
		createServer(job)
	}
}

func createServer(job *job.Job) error {
	// validate job (flavor, image) and create server.
	fmt.Println("Creating server...")
	return nil
}
