package cmd

import (
	"fmt"
	"io"

	vega "github.com/srijanone/vega/pkg/core"

	"github.com/spf13/cobra"
)

type repoAddCmd struct {
	out  io.Writer
	home vega.Home
	name string
	URL  string
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

	fmt.Fprintf(rAddCmd.out, "Adding new repo %s from %s", rAddCmd.name, rAddCmd.URL)
	starterKitsRepo.Add()
}
