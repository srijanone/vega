package cmd

import (
	"fmt"
	"io"

	vega "github.com/srijanone/vega/pkg/core"

	"github.com/spf13/cobra"
)

type starterkitAddCmd struct {
	URL  string
	dst  string
	out  io.Writer
	home vega.Home
	name string
}

func newAddCmd(out io.Writer) *cobra.Command {
	const addCmdDesc = "Add new starterkits repository"
	skAddCmd := starterkitAddCmd{out: out}

	addCmd := &cobra.Command{
		Use:   "add [name] [url]",
		Short: addCmdDesc,
		Args:  cobra.ExactArgs(2),
		Long:  addCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			skAddCmd.name = args[0]
			skAddCmd.URL = args[1]
			skAddCmd.home = vega.Home(homePath())
			skAddCmd.add()
		},
	}
	return addCmd
}

func (cmd *starterkitAddCmd) add() {
	starterKitsRepo := vega.StarterKitRepo{
		Name: cmd.name,
		URL:  cmd.URL,
		Home: cmd.home,
	}

	fmt.Fprintln(cmd.out, "Adding new repo:", cmd.URL)
	starterKitsRepo.Add()
}
