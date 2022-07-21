package app

import (
	"fmt"
	"github.com/juicyluv/reedygreedy/internal/reedygreedy"
	"github.com/juicyluv/reedygreedy/internal/reedygreedy/infrastructure/configuration"
	"github.com/juicyluv/reedygreedy/internal/reedygreedy/infrastructure/logs"
	"github.com/kardianos/service"
	"go.uber.org/zap"
)

type app struct {
	service service.Service
	logger  *zap.Logger
	config  configuration.Interface
}

func New() (reedygreedy.App, error) {
	logger, err := logs.NewLogger()

	if err != nil {
		return nil, fmt.Errorf("init logger: %v", err)
	}

	app := &app{logger: logger}

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

func (a *app) GetLogger() *zap.Logger {
	return a.logger
}

func (a *app) SetLogger(logger *zap.Logger) {
	a.logger = logger
}

func (a *app) GetConfiguration() configuration.Interface {
	return a.config
}

func (a *app) SetConfiguration(cfg configuration.Interface) {
	a.config = cfg
}
