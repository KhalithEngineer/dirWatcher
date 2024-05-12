package repomapper

import (
	"dirwatcher/internal/models"
	dbmodels "dirwatcher/repository/models"

	"github.com/volatiletech/null/v8"
)

func ConfigToDBModel(config *models.Config) *dbmodels.Config {
	return &dbmodels.Config{
		Directory:    null.StringFrom(config.Dir),
		TimeInterval: null.IntFrom(config.Interval),
		MagicString:  null.StringFrom(config.MagicString),
	}
}

func ConfigFromDBModel(config *dbmodels.Config) *models.Config {
	return &models.Config{
		Dir:         config.Directory.String,
		Interval:    config.TimeInterval.Int,
		MagicString: config.MagicString.String,
	}
}
