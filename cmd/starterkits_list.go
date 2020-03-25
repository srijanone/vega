package cmd

import (
	"fmt"
	"io"

	vega "github.com/srijanone/vega/pkg/core"

	"github.com/spf13/cobra"
)

type starterkitListCmd struct {
	out  io.Writer
	home vega.Home
}

func newStarterKitListCmd(out io.Writer) *cobra.Command {
	const listCmdDesc = "List starter-kits"
	list := &starterkitListCmd{out: out}
	listCmd := &cobra.Command{
		Use:   "list",
		Short: listCmdDesc,
		Long:  listCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			list.execute()
		},
	}
	list.home = vega.Home(homePath())
	return listCmd
}

func (cmd *starterkitListCmd) execute() error {
	starterkits, err := vega.StarterKitList(cmd.home.StarterKits())
	if err != nil {
		return err
	}
	fmt.Fprintln(cmd.out, "Available Starter-kits:")
	for _, starterkit := range starterkits {
		fmt.Fprintf(cmd.out, "  %s\n", starterkit.Name)
	}
	return nil
}
