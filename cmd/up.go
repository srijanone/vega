package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	compose "github.com/srijanone/vega/pkg/compose"
	detector "github.com/srijanone/vega/pkg/detector"
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
			upArgs := []string{"--port", port}
			if noBrowser {
				upArgs = append(upArgs, "--no-browser")
			}
			if watch == false {
				upArgs = append(upArgs, "--watch", "false")
			}
			if detector.IsDrupal() {
				fmt.Fprintln(out, "Building the application")
				compose.Run(out, "cli")
			}
			fmt.Fprintln(out, "Running the application")
			tilt.Up(out, upArgs...)
		},
	}
	flags := upCmd.Flags()
	flags.BoolVar(&noBrowser, "no-browser", false, "If true, Web UI will not open on startup")
	flags.BoolVar(&watch, "watch", true, "If true, services will be automatically rebuilt and redeployed when files change")
	flags.StringVar(&port, "port", "9090", "Port for the Logging HTTP server")
	return upCmd
}
