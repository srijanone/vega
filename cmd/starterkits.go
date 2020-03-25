package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func newStarterKitCmd(out io.Writer) *cobra.Command {
	const starterkitDesc = "Manage starterkit used for initializeing project"
	starterkitCmd := &cobra.Command{
		Use:   "starterkit",
		Short: starterkitDesc,
		Long:  starterkitDesc,
	}
	starterkitCmd.AddCommand(newStarterKitListCmd(out))
	return starterkitCmd
}
