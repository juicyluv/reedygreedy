package database

import (
	"github.com/juicyluv/reedygreedy/internal/reedygreedy/infrastructure/configuration"
	"github.com/juicyluv/rgdb/pkg/rgdb"
	"github.com/juicyluv/rgutils/pkg/logger"
	"strconv"
)

func New(logger *logger.Logger, cfg configuration.Interface) (*rgdb.Client, error) {
	client, err := rgdb.New(logger, &rgdb.Config{
		ConnString: cfg.GetDatabaseConnString(),
		Host:       cfg.GetDatabaseHost(),
		User:       cfg.GetDatabaseUser(),
		Password:   cfg.GetDatabasePassword(),
		Port:       strconv.Itoa(cfg.GetDatabasePort()),
		Database:   cfg.GetDatabaseName(),
		SSLMode:    cfg.GetDatabaseSSLMode(),
		MaxConns:   cfg.GetDatabaseMaxConnectionCount(),
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}
