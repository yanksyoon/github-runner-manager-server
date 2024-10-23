package reconciler

import "github.com/charlie4284/github-runner-manager-server/internal/queue"

type Manager struct {
	queueManager    *queue.Manager
	desiredCapacity int
}

func New(queueManager *queue.Manager) *Manager {
	return &Manager{
		queueManager: queueManager,
	}
}

func (m *Manager) Reconcile() {
	go m.cleanup()
	go m.queueRunners()
}

func (m *Manager) cleanup() {
	for {

	}
}

func (m *Manager) queueRunners() {
	for {

	}
}
