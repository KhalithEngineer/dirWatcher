package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTaskDetails fetches the task run from db and send it as response
func (h *Handlers) GetTaskDetails(ctx *gin.Context) {

	taskRuns, err := h.repo.GetAllTaskRun(ctx)

	if err != nil {
		h.logger.Error("Failed to get all task runs: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	h.logger.Info("All task runs fetched successfully")
	ctx.JSON(http.StatusFound, taskRuns)

}
