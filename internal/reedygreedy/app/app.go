package app

import (
	"fmt"
	"github.com/juicyluv/reedygreedy/internal/reedygreedy"
	"github.com/juicyluv/reedygreedy/internal/reedygreedy/infrastructure/configuration"
	"github.com/juicyluv/rgutils/pkg/logger"
	"github.com/kardianos/service"
)

type app struct {
	service service.Service
	logger  *logger.Logger
	config  configuration.Interface
}

func New() (reedygreedy.App, error) {
	// TODO: set the logger from config
	l := logger.New(&logger.Config{
		LogToConsole:     true,
		EncodeLogsAsJson: true,
		LogToFile:        true,
		Directory:        "logs",
		Filename:         "logs.log",
		MaxSize:          10,
		MaxBackups:       1,
		MaxAge:           10,
	})

	app := &app{logger: l}

	srv, err := service.New(app, &service.Config{Name: "ReedyGreedy"})

	if err != nil {
		return nil, fmt.Errorf("init service: %v", err)
	}

	app.service = srv

	return app, nil
}

func (a *app) Service() service.Service {
	return a.service
}

func (a *app) GetLogger() *logger.Logger {
	return a.logger
}

func (a *app) SetLogger(logger *logger.Logger) {
	a.logger = logger
}

func (a *app) GetConfiguration() configuration.Interface {
	return a.config
}

func (a *app) SetConfiguration(cfg configuration.Interface) {
	a.config = cfg
}
