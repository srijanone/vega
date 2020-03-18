package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

func newInitCmd(out io.Writer) *cobra.Command {
	const initDesc = `This command sets up local configuration in $VEGA_HOME (default ~/.vega/) with default starter-kits`

	initCmd := &cobra.Command{
		Use:   "init",
		Short: initDesc,
		Long:  initDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("init called")
		},
	}

	return initCmd
}
