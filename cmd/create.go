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
	appName        string
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
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("create called")
			if len(args) > 0 {
				cCmd.dest = args[0]
			}
			return cCmd.execute()
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			// transforming name as Docker doesn't allow upper case names for repositories
			//common.normalizeApplicationName(cCmd.appName)
		},
	}

	cCmd.home = vega.Home(homePath())

	flags := createCmd.Flags()
	flags.StringVarP(&cCmd.appName, "app", "a", "", "name of the app. (default: randomly generated name)")
	flags.StringVarP(&cCmd.starterKit, "starterkit", "s", "", "name of the Vega starterkit to scaffold the app")
	// TODO: make starterkit required
	return createCmd
}

// Temp. Struct, will be replaced by pkg's struct
type starterKitStruct struct {
	dockerfileName string
}

func (sk *starterKitStruct) Find(dir string, name string) ([]string, error) {
	return []string{}, nil
}

func (sk *starterKitStruct) CreateFrom(dest string, source string, name string) error {
	return nil
}

var starterKit starterKitStruct = starterKitStruct{"Dockerfile"}

func (cCmd *createCmd) execute() error {
	if cCmd.appName == "" {
		cCmd.appName = common.GeneratePetName()
	}
	// TODO: Check if starterkits files are already there or not
	dockerfileExists, err := common.Exists(filepath.Join(cCmd.dest, starterKit.dockerfileName))
	if err != nil {
		return fmt.Errorf("there was an error checking if %s exists: %v", starterKit.dockerfileName, err)
	}
	if dockerfileExists {
		fmt.Fprintln(cCmd.out, "starterkit already exists")
		return nil
	}

	starterKits, err := starterKit.Find(cCmd.home.StarterKits(), cCmd.starterKit)
	if err != nil {
		return err
	}
	if len(starterKits) == 1 {
		starterKitSrc := starterKits[0]
		fmt.Fprintln(cCmd.out, "Found starterkit")
		if err = starterKit.CreateFrom(cCmd.dest, starterKitSrc, cCmd.appName); err != nil {
			return err
		}
	} else if len(starterKits) > 0 {
		return fmt.Errorf("Multiple starterkits named %s found: %v", cCmd.starterKit, starterKits)
	} else {
		return fmt.Errorf("No starterkit found with name %s", cCmd.starterKit)
	}

	fmt.Fprintln(cCmd.out, "Ready for development")
	return nil
}
