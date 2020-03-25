package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func newStarterKitCmd(out io.Writer) *cobra.Command {
	const starterkitsDesc = "Manage starter-kits used for initializeing projects"
	starterkitsCmd := &cobra.Command{
		Use:   "starterkit",
		Short: starterkitsDesc,
		Long:  starterkitsDesc,
	}
	starterkitsCmd.AddCommand(newStarterKitListCmd(out))
	return starterkitsCmd
}
