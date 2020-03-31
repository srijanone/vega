package cmd

import (
	"fmt"
	"io"

	vega "github.com/srijanone/vega/pkg/core"
	"github.com/srijanone/vega/pkg/downloader"

	"github.com/spf13/cobra"
)

type starterkitAddCmd struct {
	URL  string
	out  io.Writer
	home vega.Home
}

func newAddCmd(out io.Writer) *cobra.Command {
	const addCmdDesc = "Add new starterkits repository"
	skAddCmd := starterkitAddCmd{out: out}

	addCmd := &cobra.Command{
		Use:   "add [address]",
		Short: addCmdDesc,
		Args:  cobra.ExactArgs(1),
		Long:  addCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			skAddCmd.URL = args[0]
			skAddCmd.home = vega.Home(homePath())
			skAddCmd.add()
		},
	}
	return addCmd
}

func (cmd *starterkitAddCmd) add() error {
	//TODO: Merge multiple repo
	//TODO: Allow local starter kit repo
	fmt.Println("Adding new repo:", cmd.URL)
	d := downloader.Downloader{}
	sourceRepo := fmt.Sprintf("%s//%s", cmd.URL, "starterkits")
	fmt.Fprintln(cmd.out, "Downloading starterkits...")
	d.Download(sourceRepo, cmd.home.Caches())
	return nil
}
