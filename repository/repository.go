package repository

import (
	"context"
	"dirwatcher/internal/models"
	dbmodels "dirwatcher/repository/models"
)

type Repo interface {
	InsertConfig(ctx context.Context, config *models.Config) error
	UpdateConfig(ctx context.Context, config *models.Config) error
	GetConfig(ctx context.Context) (*models.Config, error)
	InsertTaskRun(ctx context.Context, taskRun dbmodels.TaskRun) error
	GetLatestTaskRun(ctx context.Context) (*dbmodels.TaskRun, error)
	GetAllTaskRun(ctx context.Context) ([]models.TaskRun, error)
}
