package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/srijanone/vega/pkg/compose"
)

func newDestroyCmd(out io.Writer) *cobra.Command {
	const downDesc = "Delete application and persisted data"

	downCmd := &cobra.Command{
		Use:   "down",
		Short: downDesc,
		Long:  downDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(out, "Deleting the application")
			compose.Destroy(out, args...)
		},
	}

	return downCmd
}
