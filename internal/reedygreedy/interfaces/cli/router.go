package cli

import (
	"github.com/juicyluv/reedygreedy/internal/reedygreedy"
	"github.com/spf13/cobra"
)

type router struct{ app reedygreedy.App }

func NewRouter(app reedygreedy.App) *cobra.Command {
	r := &router{app: app}

	cmd := &cobra.Command{}

	cmd.AddCommand(
		r.run(),
	)

	return cmd
}
