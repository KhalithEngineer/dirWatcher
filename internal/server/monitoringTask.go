package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) ToggleMonitoringTask(c *gin.Context) {
	if h.engine.Running {
		h.engine.Stop()
		c.JSON(http.StatusOK, gin.H{"message": "Task stopped successfully"})
	} else {
		h.engine.Start()
		c.JSON(http.StatusOK, gin.H{"message": "Task started successfully"})
	}

}

func (h *Handlers) StartMonitoringTask(c *gin.Context) {
	h.logger.Info("Starting monitoring task...")
	if !h.engine.Running {
		h.engine.Start()
		c.JSON(http.StatusOK, gin.H{"message": "Task started successfully"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task already running"})
}

func (h *Handlers) StopMonitoringTask(c *gin.Context) {
	h.logger.Info("Stopping monitoring task...")
	if h.engine.Running {
		h.engine.Stop()
		c.JSON(http.StatusOK, gin.H{"message": "Task stopped successfully"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task stopped already"})
}

func (h *Handlers) MonitoringTaskStatus(c *gin.Context) {
	h.logger.Info("Stopping monitoring task...")
	if h.engine.Running {
		c.JSON(http.StatusOK, gin.H{"message": "Task is running"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Task is not running"})
	}
}
