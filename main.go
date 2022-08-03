package main

import (
	"github.com/juicyluv/reedygreedy/internal/reedygreedy/app"
	"github.com/juicyluv/reedygreedy/internal/reedygreedy/interfaces/cli"
)

func main() {
	apl, err := app.New()

	if err != nil {
		panic(err)
	}

	router := cli.NewRouter(apl)

	if err := router.Execute(); err != nil {
		apl.GetLogger().Err(err).Msg("cannot execute command")
	}
}
