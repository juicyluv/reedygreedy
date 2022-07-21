package reedygreedy

import (
	"github.com/juicyluv/reedygreedy/internal/reedygreedy/infrastructure/configuration"
	"github.com/kardianos/service"
	"go.uber.org/zap"
)

type App interface {
	service.Interface

	Service() service.Service

	GetLogger() *zap.Logger

	SetLogger(logger *zap.Logger)

	GetConfiguration() configuration.Interface

	SetConfiguration(cfg configuration.Interface)
}
