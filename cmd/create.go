package cmd

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/cobra"
	common "github.com/srijanone/vega/pkg/common"
	vega "github.com/srijanone/vega/pkg/core"
)

type createCmd struct {
	out        io.Writer
	home       vega.Home
	starterkit string
	dest       string
	repo       string
}

func newCreateCmd(out io.Writer) *cobra.Command {
	cCmd := &createCmd{out: out}

	const createDesc = "create starterkit"

	createCmd := &cobra.Command{
		Use:   "create [path]",
		Short: createDesc,
		Long:  createDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				cCmd.dest = args[0]
			} else {
				// If no path is given, then create in current directory
				cCmd.dest = "."
			}
			return cCmd.execute()
		},
	}

	cCmd.home = vega.Home(homePath())

	flags := createCmd.Flags()
	flags.StringVarP(&cCmd.starterkit, "starterkit", "s", "", "name of the vega starterkit to scaffold the app")
	flags.StringVarP(&cCmd.repo, "repo", "r", "default", "name of the starterkit repo")
	cobra.MarkFlagRequired(flags, "starterkit")
	return createCmd
}

func (cCmd *createCmd) execute() error {
	// TODO: Check if starterkits files are already there or not properly
	dockerfileExists, err := common.Exists(filepath.Join(cCmd.dest, vega.DockerfileName))
	if err != nil {
		return fmt.Errorf("there was an error checking if %s exists: %v", vega.DockerfileName, err)
	}

	if dockerfileExists {
		fmt.Fprintln(cCmd.out, "starterkit already exists")
		return nil
	}

	path := filepath.Join(cCmd.home.StarterKits(), cCmd.repo)
	starterkitRepo := vega.StarterKitRepo{
		Name: cCmd.repo,
		Path: path,
	}

	starterkits, err := starterkitRepo.Find(cCmd.starterkit)
	if err != nil {
		return err
	}

	if len(starterkits) == 1 {
		starterkit := starterkits[0]
		starterkit.Create(cCmd.dest)
	} else if len(starterkits) > 0 {
		// TODO: display proper list of matching kits
		return fmt.Errorf("Multiple starterkit named %s found: %v", cCmd.starterkit, starterkits)
	} else {
		return fmt.Errorf("No starterkit found with name %s in %s repo", cCmd.starterkit, cCmd.repo)
	}

	fmt.Fprintln(cCmd.out, "Ready for development")
	return nil
}
