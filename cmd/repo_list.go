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

func newRepositoryList(out io.Writer) *cobra.Command {
	const listCmdDesc = "list all the starterkit repositories available locally"
	repoListCmd := repositoryListCmd{
		out: out,
	}
	listCmd := &cobra.Command{
		Use:   "list",
		Short: listCmdDesc,
		Long:  listCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			repoListCmd.execute()
		},
	}

	repoListCmd.home = vega.Home(homePath())

	return listCmd
}

func (cmd *repositoryListCmd) execute() error {
	repositories, err := vega.RepoList(cmd.home.StarterKits())
	if err != nil {
		return err
	}
	fmt.Fprintln(cmd.out, "Availabe Repositories:")
	for _, repo := range repositories {
		fmt.Fprintln(cmd.out, repo.Name)
	}
	return nil
}
