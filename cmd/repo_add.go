package cmd

import (
	"fmt"
	"io"

	vega "github.com/srijanone/vega/pkg/core"

	"github.com/spf13/cobra"
)

type repoAddCmd struct {
	URL  string
	dst  string
	out  io.Writer
	home vega.Home
	name string
}

func newAddCmd(out io.Writer) *cobra.Command {
	rAddCmd := repoAddCmd{out: out}
	
	const addCmdDesc = "add starterkits repository"
	
	addCmd := &cobra.Command{
		Use:   "add [name] [url]",
		Short: addCmdDesc,
		Long:  addCmdDesc,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			rAddCmd.name = args[0]
			rAddCmd.URL = args[1]
			rAddCmd.home = vega.Home(homePath())
			rAddCmd.execute()
		},
	}
	
	return addCmd
}

func (rAddCmd *repoAddCmd) execute() {
	starterKitsRepo := vega.StarterKitRepo{
		Name: rAddCmd.name,
		URL:  rAddCmd.URL,
		Home: rAddCmd.home,
	}

	fmt.Fprintln(rAddCmd.out, "Adding new repo:", rAddCmd.URL)
	starterKitsRepo.Add()
}
