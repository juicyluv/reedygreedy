package reedygreedy

import (
	"github.com/juicyluv/reedygreedy/internal/reedygreedy/infrastructure/configuration"
	"github.com/juicyluv/rgutils/pkg/logger"
	"github.com/kardianos/service"
)

type App interface {
	service.Interface

	Service() service.Service

	GetLogger() *logger.Logger

	SetLogger(logger *logger.Logger)

	GetConfiguration() configuration.Interface

	SetConfiguration(cfg configuration.Interface)
}
