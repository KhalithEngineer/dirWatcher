package main

import (
	"context"
	"dirwatcher/internal/engine"
	"dirwatcher/internal/server"
	"dirwatcher/repository"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	logger := setupLogger()
	// Load environment variables
	logger.Info("Loading envs...")
	if err != nil {
		logger.Error("envs loading failed")
		panic("Error loading .env file")
	}
	// Initialize database connection
	connectString := os.Getenv("connectString")
	logger.Info("Connecting to db...")
	db := repository.NewConnection(connectString)
	// Create a new repository with the database connection
	repo := repository.NewRepo(db, logger)

	ctx := context.Background()
	logger.Info("Loading config from db...")
	config, err := repo.GetConfig(ctx)

	if err != nil {
		logger.Error("error while getting config")
		panic("error while getting config")
	}

	if config == nil {
		// Load the default config if it is not present in the db
		logger.Info("Loading config with default values from env")
		interval := os.Getenv("defaultInterval")
		intervalInt, err := strconv.Atoi(interval)
		if err != nil {
			panic("Error converting interval to int")
		}
		config.SetConfig(os.Getenv("defaultDir"), intervalInt, os.Getenv("defaultMagicString"))
		repo.InsertConfig(ctx, config)
	}
	// Create a new engine with the repository
	logger.Info("Creating engine instance")
	engine := engine.NewEngine(repo, config, logger)
	// Start the engine
	logger.Info("starting engine...")
	engine.Start()
	// Create a new request handler with the repository and a new configuration
	handler := server.NewRequsthandler(repo, config, logger, engine)
	// Setup the router for the server
	r := server.SetupRouter(handler)
	// Start the server
	logger.Info("starting API server...")
	r.Run()
	logger.Info("Engine and server started")
}

func setupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	LogLevel, _ := strconv.Atoi(os.Getenv("LogLevel"))
	logger.SetLevel(logrus.Level(LogLevel))
	return logger
}
