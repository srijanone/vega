package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	tilt "github.com/srijanone/vega/pkg/tilt"
)

var watch, noBrowser bool
var port string

func newUpCmd(out io.Writer) *cobra.Command {
	const upDesc = "start the application"

	upCmd := &cobra.Command{
		Use:   "up",
		Short: upDesc,
		Long:  upDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(out, "Running the application")
			upArgs := []string{"--port", port}
			if noBrowser {
				upArgs = append(upArgs, "--no-browser")
			}
			if watch == false {
				upArgs = append(upArgs, "--watch", "false")
			}
			tilt.Up(out, upArgs...)
		},
	}

	flags := upCmd.Flags()
	flags.BoolVar(&noBrowser, "no-browser", false, "If true, Web UI will not open on startup")
	flags.BoolVar(&watch, "watch", true, "If true, services will be automatically rebuilt and redeployed when files change")
	flags.StringVar(&port, "port", "9090", "Port for the Logging HTTP server")
	return upCmd
}
