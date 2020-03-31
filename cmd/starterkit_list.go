package cmd

import (
	"fmt"
	"io"
	"path/filepath"

	vega "github.com/srijanone/vega/pkg/core"

	"github.com/spf13/cobra"
)

type starterkitListCmd struct {
	out  io.Writer
	home vega.Home
	repo string
}

func newStarterKitListCmd(out io.Writer) *cobra.Command {
	skListCmd := &starterkitListCmd{
		out: out,
	}

	const listCmdDesc = "List starterkits"

	listCmd := &cobra.Command{
		Use:   "list",
		Short: listCmdDesc,
		Long:  listCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return skListCmd.execute()
		},
	}
	flags := listCmd.Flags()
	flags.StringVarP(&skListCmd.repo, "repo", "r", "default", "name of the starterkit repo")

	skListCmd.home = vega.Home(homePath())

	return listCmd
}

func (cmd *starterkitListCmd) execute() error {
	path := filepath.Join(cmd.home.StarterKits(), cmd.repo)
	starterkitRepo := vega.StarterKitRepo{
		Name: cmd.repo,
		Path: path,
	}
	starterkits, err := starterkitRepo.List()
	if err != nil {
		return err
	}
	fmt.Fprintln(cmd.out, "Available starterkits:")
	for _, starterkit := range starterkits {
		fmt.Fprintf(cmd.out, "  %10s (%s)\n", starterkit.Name, starterkit.Path)
	}
	return nil
}
