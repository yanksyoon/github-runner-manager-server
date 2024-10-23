package api

import (
	"net/http"

	"github.com/charlie4284/github-runner-manager-server/internal/flavor"
	"github.com/charlie4284/github-runner-manager-server/internal/image"
	"github.com/charlie4284/github-runner-manager-server/internal/job"
	"github.com/charlie4284/github-runner-manager-server/internal/queue"
	"github.com/labstack/echo/v4"
)

func registerQueueRoutes(e *echo.Echo, queueManager *queue.Manager, flavorManager *flavor.Manager, imageManager *image.Manager) {
	api := queueApi{
		queueManager:  queueManager,
		flavorManager: flavorManager,
		imageManager:  imageManager,
	}
	group := e.Group("/jobs")
	group.POST("/", api.EnqueueJobs)
}

type queueApi struct {
	queueManager  *queue.Manager
	flavorManager *flavor.Manager
	imageManager  *image.Manager
}

type JobDTO struct {
	Flavor string `json:"flavor"`
	Image  string `json:"image"`
}

func (h *queueApi) EnqueueJobs(c echo.Context) error {
	jobDto := new(JobDTO)
	if err := c.Bind(jobDto); err != nil {
		return err
	}
	flavor, err := h.flavorManager.FindFlavor(jobDto.Flavor)
	if err != nil {
		return c.String(http.StatusNotFound, "Matching flavor not found.")
	}
	image, err := h.imageManager.FindImage(jobDto.Image)
	if err != nil {
		return c.String(http.StatusNotFound, "Matching image not found.")
	}
	if err := h.queueManager.Enqueue(&job.Job{
		Flavor: flavor,
		Image:  image,
	}); err != nil {
		c.Logger().Error("Failed to queue job, %s", err)
		return c.String(http.StatusInternalServerError, "Internal server error")
	}
	return c.NoContent(http.StatusOK)
}
