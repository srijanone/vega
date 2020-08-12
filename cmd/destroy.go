package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/srijanone/vega/pkg/compose"
)

func newDestroyCmd(out io.Writer) *cobra.Command {
	const destroyDesc = "Delete application and persisted data"

	destroyCmd := &cobra.Command{
		Use:   "destroy",
		Short: destroyDesc,
		Long:  destroyDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(out, "Deleting the application")
			compose.Destroy(out, args...)
		},
	}

	return destroyCmd
}
