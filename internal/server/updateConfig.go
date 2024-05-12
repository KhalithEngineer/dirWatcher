package server

import (
	"dirwatcher/internal/models"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) UpdateConfig(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.logger.Error("Failed to read request body: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	h.logger.Info("Request body read successfully")
	var config models.Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		h.logger.Error("Failed to unmarshal JSON: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	h.logger.Info("JSON unmarshalled successfully")
	if err := h.repo.UpdateConfig(ctx, &config); err != nil {
		h.logger.Error("Failed to update config in db: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot update config in db"})
	}
	h.logger.Info("Config updated in db successfully")
	h.config.SetConfig(config.Dir, config.Interval, config.MagicString)
	h.logger.Info("Config set successfully")
}
