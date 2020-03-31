package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	vega "github.com/srijanone/vega/pkg/core"
)

type RepositoryListCmd struct {
	out  io.Writer
	home vega.Home
}

func newRepositoryList(out io.Writer) *cobra.Command {
	const listCmdDesc = "Display List of all the repositories available locally"
	repositoryListCmd := RepositoryListCmd{
		out: out,
	}
	listCmd := &cobra.Command{
		Use:   "list",
		Short: listCmdDesc,
		Long:  listCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			repositoryListCmd.execute()
		},
	}

	repositoryListCmd.home = vega.Home(homePath())

	return listCmd
}

func (cmd *RepositoryListCmd) execute() error {
	repositories, err := vega.Repositories(cmd.home.StarterKits())
	if err != nil {
		return err
	}
	fmt.Fprintln(cmd.out, "Availabe Repositories:")
	for _, repo := range repositories {
		fmt.Fprintln(cmd.out, repo.Name)
	}
	return nil
}
