package engine

import (
	"context"
	dbmodels "dirwatcher/repository/models"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/volatiletech/null/v8"
)

// MonitorDirectory is a task that will monitor the given directory for the magic string,
// files added, deleted and stores in the db
func (e *Engine) MonitorDirectory() {
	e.logger.Info("Running background task...")
	startTime := time.Now()

	dir, _, magicString := e.config.GetConfig()
	e.logger.Info("Directory: ", dir)
	e.logger.Info("Magic String: ", magicString)

	previousRun, err := e.repo.GetLatestTaskRun(context.Background())

	if err != nil {
		e.repo.InsertTaskRun(context.Background(),
			dbmodels.TaskRun{Status: null.NewString("Failed", true)})
		e.logger.Error("Failed to get latest task run: ", err)
	}
	// Scan the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		e.repo.InsertTaskRun(context.Background(),
			dbmodels.TaskRun{Status: null.NewString("Failed", true)})
		e.logger.Error("Failed to read directory: ", err)
		return
	}
	var totalOccurrences int
	var filesAdded []string
	var filesDeleted []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := filepath.Join(dir, file.Name())
		e.logger.Info("File path: ", filePath)

		content, err := os.ReadFile(filePath)
		if err != nil {
			e.logger.Error("Failed to read file: ", err)
			continue
		}
		occurrences := strings.Count(string(content), magicString)
		totalOccurrences += occurrences
		e.logger.Info("Occurrences: ", occurrences)

		filesAdded = append(filesAdded, file.Name())
	}

	filesAdded, filesDeleted = compareFileFromPreviousRun(previousRun.FilesAdded, filesAdded)

	endTime := time.Now()
	totalRuntime := endTime.Sub(startTime).Seconds()
	status := "Success"

	taskRun := dbmodels.TaskRun{
		StartTime:        null.TimeFrom(startTime),
		EndTime:          null.TimeFrom(endTime),
		TotalRuntime:     null.StringFrom(fmt.Sprint(totalRuntime)),
		FilesAdded:       filesAdded,
		FilesDeleted:     filesDeleted,
		MagicString:      null.StringFrom(magicString),
		DirectoryPath:    null.StringFrom(dir),
		Interval:         null.StringFrom(fmt.Sprint(e.config.Interval)),
		TotalOccurrences: null.IntFrom(totalOccurrences),
		Status:           null.StringFrom(status),
	}
	err = e.repo.InsertTaskRun(context.Background(), taskRun)
	if err != nil {
		e.logger.Error("Failed to save task run details to the database: ", err)
		e.repo.InsertTaskRun(context.Background(),
			dbmodels.TaskRun{Status: null.NewString("Failed", true)})
		return
	}

	e.logger.Info("Background task completed successfully")
}

func compareFileFromPreviousRun(prev, new []string) ([]string, []string) {

	// Find deleted and new files
	var deleted, added []string
	for _, file := range prev {
		found := false
		for _, newFile := range new {
			if file == newFile {
				found = true
				break
			}
		}
		if !found {
			deleted = append(deleted, file)
		}
	}
	for _, file := range new {
		found := false
		for _, prevFile := range prev {
			if file == prevFile {
				found = true
				break
			}
		}
		if !found {
			added = append(added, file)
		}
	}

	return added, deleted

}
