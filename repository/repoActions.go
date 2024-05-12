package repository

import (
	"context"
	"database/sql"
	"dirwatcher/internal/models"
	dbmodels "dirwatcher/repository/models"
	repoMapper "dirwatcher/repository/repoMapper"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repository struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewRepo(db *sql.DB, logger *logrus.Logger) Repo {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) UpdateConfig(ctx context.Context, config *models.Config) error {

	dbConfig := repoMapper.ConfigToDBModel(config)
	_, err := dbConfig.Update(ctx, r.db, boil.Infer())
	if err != nil {
		r.logger.Error("Failed to update config: ", err)
		return fmt.Errorf("failed to update config: %v", err)
	}
	r.logger.Info("Config updated successfully")
	return nil
}

func (r *repository) InsertConfig(ctx context.Context, config *models.Config) error {

	dbConfig := repoMapper.ConfigToDBModel(config)
	err := dbConfig.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		r.logger.Error("Failed to insert config: ", err)
		return fmt.Errorf("failed to update config: %v", err)
	}
	r.logger.Info("Config inserted successfully")
	return nil
}

func (r *repository) GetConfig(ctx context.Context) (*models.Config, error) {

	configDB, err := dbmodels.Configs(qm.Limit(1)).One(ctx, r.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.Config{}, nil
		}
		r.logger.Error("Failed to fetch config: ", err)
		return nil, fmt.Errorf("failed to fetch config: %v", err)
	}
	config := repoMapper.ConfigFromDBModel(configDB)
	r.logger.Info("Config fetched successfully")
	return config, nil
}

func (r *repository) GetLatestTaskRun(ctx context.Context) (*dbmodels.TaskRun, error) {

	taskRun, err := dbmodels.TaskRuns(qm.OrderBy("start_time DESC"), qm.Limit(1)).One(ctx, r.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return &dbmodels.TaskRun{}, nil
		}
		r.logger.Error("Failed to fetch latest task run: ", err)
		return nil, fmt.Errorf("failed to fetch latest task run: %v", err)
	}
	r.logger.Info("Latest task run fetched successfully")
	return taskRun, nil

}

func (r *repository) GetAllTaskRun(ctx context.Context) ([]models.TaskRun, error) {

	taskRuns, err := dbmodels.TaskRuns(qm.OrderBy("start_time DESC")).All(ctx, r.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return []models.TaskRun{}, nil
		}
		r.logger.Error("Failed to fetch latest task run: ", err)
		return nil, fmt.Errorf("failed to fetch latest task run: %v", err)
	}

	return repoMapper.TaskRunSliceFromDBModel(&taskRuns), nil

}

func (r *repository) InsertTaskRun(ctx context.Context, taskRun dbmodels.TaskRun) error {
	if err := taskRun.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.logger.Error("Failed to insert task run: ", err)
		return fmt.Errorf("failed to update task run: %v", err)
	}
	r.logger.Info("Task run inserted successfully")
	return nil
}
