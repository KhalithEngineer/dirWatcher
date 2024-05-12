package engine

import (
	"dirwatcher/internal/models"
	"dirwatcher/repository"
	"time"

	"github.com/sirupsen/logrus"
)

type Engine struct {
	repo     repository.Repo
	config   *models.Config
	logger   *logrus.Logger
	stopChan chan struct{}
	Running  bool
}

// Create a new instance of Engine
func NewEngine(repo repository.Repo, config *models.Config, logger *logrus.Logger) *Engine {
	return &Engine{
		repo:   repo,
		config: config,
		logger: logger,
	}
}

// Run starts the engine and runs the background task at regular intervals
func (e *Engine) Run() {

	// Create a new ticker with the configured interval
	e.logger.Info("Creating a scheduler for monitoring directory")
	ticker := time.NewTicker(time.Duration(e.config.Interval) * time.Second)
	// Start a new goroutine
	go func() {
		defer ticker.Stop()
		// Run the background task each time the ticker ticks
		for {
			select {
			case <-ticker.C:
				e.MonitorDirectory()
			case <-e.stopChan:
				e.logger.Info("Background task stopped")
				return
			}
		}

	}()
}

func (e *Engine) Start() {
	if !e.Running {
		e.stopChan = make(chan struct{})
		e.Running = true
		go e.Run()
	}
}

func (e *Engine) Stop() {
	if e.Running {
		e.Running = false
		close(e.stopChan)
	}
}
