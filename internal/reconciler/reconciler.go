package reconciler

import "github.com/charlie4284/github-runner-manager-server/internal/queue"

type Manager struct {
	queueManager    *queue.Manager
	desiredCapacity int
}

func New(queueManager *queue.Manager, desiredCapacity int) *Manager {
	return &Manager{
		queueManager:    queueManager,
		desiredCapacity: desiredCapacity,
	}
}

func (m *Manager) Reconcile() {
	go m.cleanup()
	go m.create()
}

// cleanup checks the state of the runners and deletes finished or errored runners
func (m *Manager) cleanup() {
	for {
		// check runners & see if job is terminated or
		// if the runner has encountered an error
		// shutdown and delete.
	}
}

// create checks the number of desired runners and reconciles to reach the number.
func (m *Manager) create() {
	for {
		// check current active runners
		// check queuesize
		// validate job (image, flavor) & push jobs to queue if necessary
	}
}
