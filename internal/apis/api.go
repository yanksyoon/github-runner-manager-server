package api

import (
	"github.com/charlie4284/github-runner-manager-server/internal/flavor"
	"github.com/charlie4284/github-runner-manager-server/internal/image"
	"github.com/charlie4284/github-runner-manager-server/internal/queue"
	"github.com/labstack/echo/v4"
)

func New(queueManager *queue.Manager, flavorManager *flavor.Manager, imageManager *image.Manager) *echo.Echo {
	e := echo.New()
	registerQueueRoutes(e, queueManager, flavorManager, imageManager)
	return e
}
