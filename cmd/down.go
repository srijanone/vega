package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	tilt "github.com/srijanone/vega/pkg/tilt"
)

func newDownCmd(out io.Writer) *cobra.Command {
	const downDesc = "stop the application"

	downCmd := &cobra.Command{
		Use:   "down",
		Short: downDesc,
		Long:  downDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(out, "Stopping the application")
			tilt.Down(out, args...)
		},
	}

	return downCmd
}
