package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

func newHomeCmd(out io.Writer) *cobra.Command {
	const homeDesc = "print location of VEGA_HOME"

	homeCmd := &cobra.Command{
		Use:   "home",
		Short: homeDesc,
		Long:  homeDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(homePath())
		},
	}

	return homeCmd
}
