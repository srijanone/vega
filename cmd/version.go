package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/srijanone/vega/pkg/version"
)

type versionCmd struct {
	out   io.Writer
	short bool
}

func newVersionCmd(out io.Writer) *cobra.Command {
	vCmd := &versionCmd{
		out: out,
	}

	const versionDesc = "print version"

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: versionDesc,
		Long:  versionDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version called")
			fmt.Println(version.New().FormatVersion(vCmd.short))
		},
	}

	flags := versionCmd.Flags()
	flags.BoolVarP(&vCmd.short, "short", "s", false, "shorten output version")

	return versionCmd
}
