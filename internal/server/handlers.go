package server

import (
	"dirwatcher/internal/engine"
	"dirwatcher/internal/models"
	"dirwatcher/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RequestHandler interface {
	UpdateConfig(*gin.Context)
	GetTaskDetails(*gin.Context)
	ToggleMonitoringTask(c *gin.Context)
	MonitoringTaskStatus(c *gin.Context)
	StartMonitoringTask(c *gin.Context)
	StopMonitoringTask(c *gin.Context)
}

type Handlers struct {
	repo   repository.Repo
	config *models.Config
	logger *logrus.Logger
	engine *engine.Engine
}

func NewRequsthandler(repo repository.Repo, config *models.Config,
	logger *logrus.Logger, engine *engine.Engine) RequestHandler {
	return &Handlers{
		repo:   repo,
		config: config,
		logger: logger,
		engine: engine,
	}
}
