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
	out            io.Writer
	starterKit     string
	home           vega.Home
	dest           string
	repositoryName string
}

func newCreateCmd(out io.Writer) *cobra.Command {
	cCmd := &createCmd{
		out: out,
	}

	const createDesc = "create starterkit"

	createCmd := &cobra.Command{
		Use:   "create [path]",
		Short: createDesc,
		Long:  createDesc,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				cCmd.dest = args[0]
			}
			return cCmd.execute()
		},
	}

	cCmd.home = vega.Home(homePath())

	flags := createCmd.Flags()
	flags.StringVarP(&cCmd.starterKit, "starterkit", "s", "", "name of the Vega starterkit to scaffold the app")
	cobra.MarkFlagRequired(flags, "starterkit")
	return createCmd
}

//DockerfileName : Default dockerfile name used in starterkits
const DockerfileName string = "Dockerfile"

func (cCmd *createCmd) execute() error {
	// TODO: Check if starterkits files are already there or not
	dockerfileExists, err := common.Exists(filepath.Join(cCmd.dest, DockerfileName))
	if err != nil {
		return fmt.Errorf("there was an error checking if %s exists: %v", DockerfileName, err)
	}

	if dockerfileExists {
		fmt.Fprintln(cCmd.out, "starterkit already exists")
		return nil
	}

	starterKits, err := vega.StarterKitFind(cCmd.home.StarterKits(), cCmd.starterKit)
	if err != nil {
		return err
	}
	if len(starterKits) == 1 {
		starterKitSrc := starterKits[0].Path
		fmt.Fprintln(cCmd.out, "Found starterkit")
		vega.StarterKitCreate(starterKitSrc, cCmd.dest)
	} else if len(starterKits) > 0 {
		// TODO: display proper list of matching kits
		return fmt.Errorf("Multiple starterkits named %s found: %v", cCmd.starterKit, starterKits)
	} else {
		return fmt.Errorf("No starterkit found with name %s", cCmd.starterKit)
	}

	fmt.Fprintln(cCmd.out, "Ready for development")
	return nil
}
