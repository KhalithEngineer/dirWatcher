package server

import "github.com/gin-gonic/gin"

func SetupRouter(handler RequestHandler) *gin.Engine {

	r := gin.Default()

	r.POST("/config", func(c *gin.Context) {
		handler.UpdateConfig(c)
	})
	r.GET("/getTaskRuns", func(c *gin.Context) {
		handler.GetTaskDetails(c)
	})
	r.POST("/getTaskRuns", func(c *gin.Context) {
		handler.GetTaskDetails(c)
	})
	r.POST("/stopMonitoringTask", func(c *gin.Context) {
		handler.StopMonitoringTask(c)
	})
	r.POST("/startMonitoringTask", func(c *gin.Context) {
		handler.StartMonitoringTask(c)
	})
	r.POST("/toggleMonitoringTask", func(c *gin.Context) {
		handler.ToggleMonitoringTask(c)
	})
	r.GET("/monitoringTaskStatus", func(c *gin.Context) {
		handler.MonitoringTaskStatus(c)
	})
	return r
}
