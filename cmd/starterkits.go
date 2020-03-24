package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func newStarterKitCmd(out io.Writer) *cobra.Command {
	const starterkitsDesc = "Manage starter-kits used for initializeing projects"
	starterkitsCmd := &cobra.Command{
		Use:   "starterkits",
		Short: starterkitsDesc,
		Long:  starterkitsDesc,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				// os.Exit(0)
			}
		}}
	starterkitsCmd.AddCommand(newStarterKitListCmd(out))
	return starterkitsCmd
}
