package repomapper

import (
	"dirwatcher/internal/models"
	dbmodels "dirwatcher/repository/models"
)

func TaskRunSliceFromDBModel(taskRunSlice *dbmodels.TaskRunSlice) []models.TaskRun {
	taskRuns := make([]models.TaskRun, len(*taskRunSlice))
	for i, taskRun := range *taskRunSlice {
		taskRuns[i] = models.TaskRun{
			ID:           taskRun.ID,
			StartTime:    taskRun.StartTime.Time,
			EndTime:      taskRun.EndTime.Time,
			Runtime:      taskRun.TotalRuntime.String,
			Status:       taskRun.Status.String,
			FilesAdded:   taskRun.FilesAdded,
			FilesDeleted: taskRun.FilesDeleted,
			Occurrences:  taskRun.TotalOccurrences.Int,
			Directory:    taskRun.DirectoryPath.String,
			MagicString:  taskRun.MagicString.String,
		}
	}
	return taskRuns
}
