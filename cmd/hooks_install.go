package cmd

import (
	"fmt"
	"io"

	vega "github.com/srijanone/vega/pkg/core"

	"github.com/spf13/cobra"
)

type hooksInstallCmd struct {
	out  io.Writer
	home vega.Home
	path string
}

func newHooksInstallCmd(out io.Writer) *cobra.Command {
	hInstallCmd := hooksInstallCmd{out: out}

	const hooksInstallDesc = "install git hooks in existing git repository"

	installCmd := &cobra.Command{
		Use:   "install [project-path]",
		Short: hooksInstallDesc,
		Long:  hooksInstallDesc,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				// If no arguments are passed, then choose current directory as path
				hInstallCmd.path = "."
			} else {
				hInstallCmd.path = args[0]
			}
			hInstallCmd.home = vega.Home(homePath())
			hInstallCmd.execute()
		},
	}

	return installCmd
}

func (hInstallCmd *hooksInstallCmd) execute() {
	fmt.Fprintf(hInstallCmd.out, "Installing git hooks to %s\n", hInstallCmd.path)

	// Adding Git Hooks to Vega Home
	gitHooks := vega.GitHooks{
		Home: hInstallCmd.home,
		Out:  hInstallCmd.out,
	}

	// Installing Git Hooks to project
	gitHooks.Install(hInstallCmd.path)
}
