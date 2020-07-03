package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	common "github.com/srijanone/vega/pkg/common"
	vega "github.com/srijanone/vega/pkg/core"
)

const (
	starterKitsRepoName = "git@github.com:srijanone/vega.git"
	starterKitsDirName  = "starterkits"
	gitHooksRepoName    = "git@github.com:srijanone/vega.git"
	gitHooksDirName     = "hooks"
)

type initCmd struct {
	in     io.Reader
	out    io.Writer
	home   vega.Home
	dryRun bool
}

func newInitCmd(out io.Writer, in io.Reader) *cobra.Command {
	init := &initCmd{
		in:  in,
		out: out,
	}

	const initDesc = `sets up local configuration in $VEGA_HOME with default starterkits`

	initCmd := &cobra.Command{
		Use:   "init",
		Short: initDesc,
		Long:  initDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				return errors.New("Command does not accept arguments")
			}
			init.home = vega.Home(homePath())
			return init.execute()
		},
	}

	return initCmd
}

func (iCmd *initCmd) execute() error {

	if !iCmd.dryRun {
		if err := iCmd.setupVegaHome(); err != nil {
			return err
		}
	}

	fmt.Fprintln(iCmd.out, "$VEGA_HOME has been initialized at", vegaHome)
	return nil
}

func (iCmd *initCmd) setupVegaHome() error {
	directories := []string{
		iCmd.home.String(),
		iCmd.home.StarterKits(),
		iCmd.home.GitHooks(),
		iCmd.home.Logs(),
	}

	// Ensuring that required directory exists or not
	for _, path := range directories {
		if err := common.EnsureDir(path); err != nil {
			return err
		}
		fmt.Fprintln(iCmd.out, "Initializing", path)
	}

	// Adding default starter kits to Vega Home
	defaultStarterKit := vega.StarterKitRepo{
		Name: "default",
		Home: iCmd.home,
		URL:  starterKitsRepoName,
		Dir:  starterKitsDirName,
	}
	defaultStarterKit.Add()

	// Adding Git Hooks to Vega Home
	gitHooks := vega.GitHooks{
		Home: iCmd.home,
		URL:  gitHooksRepoName,
		Dir:  gitHooksDirName,
		Out:  iCmd.out,
	}
	gitHooks.Add()

	// Installing Git Hooks as Global hooks
	gitHooks.InstallGlobally()
	return nil
}
