package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/charlie4284/github-runner-manager-server/internal/flavor"
	"github.com/charlie4284/github-runner-manager-server/internal/image"
	"github.com/charlie4284/github-runner-manager-server/internal/openstack"
	"github.com/charlie4284/github-runner-manager-server/internal/queue"
	"github.com/charlie4284/github-runner-manager-server/internal/reconciler"
	"github.com/charlie4284/github-runner-manager-server/internal/worker"
)

func main() {
	openstackProvider, openstackEndpointOpts, err := openstack.New()
	if err != nil {
		// log and quit
		panic(err)
	}
	flavorManager, err := flavor.New()
	if err != nil {
		panic(err)
	}
	imageManager, err := image.New()
	if err != nil {
		panic(err)
	}
	queueManager := queue.NewManager()

	terminationSignal := make(chan struct{})
	reconciler := reconciler.New(queueManager, 5)
	reconciler.Reconcile() // pass in termination channel
	worker.Start(1, queueManager)
	go func() {
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		close(terminationSignal)
	}()
	<-terminationSignal
}
