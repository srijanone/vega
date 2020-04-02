package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	vega "github.com/srijanone/vega/pkg/core"
)

type repositoryListCmd struct {
	out  io.Writer
	home vega.Home
}

func newRepositoryCmdList(out io.Writer) *cobra.Command {
	rListCmd := repositoryListCmd{out: out}

	const listCmdDesc = "list all the starterkit repositories available locally"

	listCmd := &cobra.Command{
		Use:   "list",
		Short: listCmdDesc,
		Long:  listCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			rListCmd.execute()
		},
	}

	rListCmd.home = vega.Home(homePath())

	return listCmd
}

func (cmd *repositoryListCmd) execute() error {
	repositories, err := vega.RepoList(cmd.home.StarterKits())
	if err != nil {
		return err
	}
	fmt.Fprintln(cmd.out, "Available Repositories:")
	for _, repo := range repositories {
		fmt.Fprintf(cmd.out, "  %20s (%s)\n", repo.Name, repo.Path)
	}
	return nil
}
