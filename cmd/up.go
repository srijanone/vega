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
			if tilt.IsInstalled() {
				fmt.Fprintln(out, "Running the application")
				args := []string{"--port", port}
				if noBrowser {
					args = append(args, "--no-browser")
				}
				if watch == false {
					args = append(args, "--watch", "false")
				}
				tilt.Up(out, args...)
			} else {
				fmt.Fprintf(out, tilt.RequiredText)
				fmt.Fprintf(out, tilt.InstallInstructions)
			}
		},
	}

	flags := upCmd.Flags()
	flags.BoolVar(&noBrowser, "no-browser", false, "If true, Web UI will not open on startup")
	flags.BoolVar(&watch, "watch", true, "If true, services will be automatically rebuilt and redeployed when files change")
	flags.StringVar(&port, "port", "9090", "Port for the Logging HTTP server")
	return upCmd
}
