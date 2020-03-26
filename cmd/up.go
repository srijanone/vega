package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func newUpCmd(out io.Writer) *cobra.Command {
	const upDesc = "Runs the application in the image"

	upCmd := &cobra.Command{
		Use:   "up",
		Short: upDesc,
		Long:  upDesc,
		Run: func(cmd *cobra.Command, args []string) {
			command := exec.Command("/usr/local/bin/tilt", "up")
			command.Stdout = out
			command.Stderr = os.Stderr
			fmt.Fprintln(out, "Running the application")
			err := command.Run()
			if err != nil {
				fmt.Fprintln(out, "Error in Running", err)
			}
		},
	}

	return upCmd
}
