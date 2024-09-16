package main

import (
	"os"

	"github.com/Arturchick1/person-test/internal/config"
	"github.com/Arturchick1/person-test/internal/handlers"
	"github.com/Arturchick1/person-test/internal/logic"
	"github.com/Arturchick1/person-test/internal/repository"
	"github.com/Arturchick1/person-test/pkg/database"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.New()

	logger := setupLogger(cfg.Env)

	logger.Info("server is starting")

	storage, err := database.New(cfg.StoragePath)
	if err != nil {
		logger.Error(err)
		return
	}
	defer storage.DB.Close()

	personRepository := repository.New(storage.DB)
	personLogic := logic.New(personRepository)
	personHandler := handlers.New(personLogic, logger)

	e := echo.New()

	e.GET("/person/:id", personHandler.GetOne)
	e.GET("/person", personHandler.Get)
	e.GET("/person?limit&offset", personHandler.Get)
	e.PUT("/person/:id", personHandler.Update)
	e.POST("/person", personHandler.Create)
	e.DELETE("/person/:id", personHandler.Delete)

	err = e.Start(cfg.Address)
	if err != nil {
		logger.Error(err)
		return
	}
}

func setupLogger(env string) *log.Logger {
	var (
		logFormatter log.Formatter
		logLevel     log.Level
	)

	switch env {
	case envLocal:
		logFormatter = new(log.TextFormatter)
		logLevel = log.DebugLevel
	case envDev:
		logFormatter = new(log.JSONFormatter)
		logLevel = log.DebugLevel
	case envProd:
		logFormatter = new(log.JSONFormatter)
		logLevel = log.InfoLevel
	}

	logger := &log.Logger{
		Out:       os.Stdout,
		Formatter: logFormatter,
		Level:     logLevel,
	}

	return logger
}
