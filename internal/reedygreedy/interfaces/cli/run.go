package cli

import "github.com/spf13/cobra"

func (r *router) run() *cobra.Command {
	cmd := &cobra.Command{}

	cmd.Use = `run`
	cmd.Short = `Start the application`

	cmd.Run = func(cmd *cobra.Command, args []string) {
		err := r.app.Start(r.app.Service())

		if err != nil {
			cmd.PrintErrln(err)
			return
		}
	}

	return cmd
}
