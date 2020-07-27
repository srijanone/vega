package cmd

import (
	"fmt"
	"io"
	"path/filepath"

	survey "github.com/AlecAivazis/survey/v2"
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
	path       string
}

func newCreateCmd(out io.Writer) *cobra.Command {
	cCmd := &createCmd{out: out}

	const createDesc = "create a project using starterkit"

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
	// cobra.MarkFlagRequired(flags, "starterkit")
	return createCmd
}

func (cCmd *createCmd) execute() error {
	// TODO: Check if starterkits files are already there or not properly
	dockerfileExists, err := common.Exists(filepath.Join(cCmd.dest, vega.DockerfileName))
	if err != nil {
		return fmt.Errorf("couldn't check if starterkit already exists: %v", err)
	}

	if dockerfileExists {
		fmt.Fprintln(cCmd.out, "starterkit already exists")
		return nil
	}

	repoPath := filepath.Join(cCmd.home.StarterKits(), cCmd.repo)
	starterkitRepo := vega.StarterKitRepo{
		Name: cCmd.repo,
		Path: repoPath,
	}

	fmt.Printf("checkthis this \n")

	var starterkit vega.StarterKit

	if cCmd.starterkit == "" {
		starterkits, err := starterkitRepo.StarterKitList()
		if err != nil {
			fmt.Fprintln(cCmd.out, "No starterkit found")
		}
		starterkit, err = askUserChoice(starterkits, repoPath)
		if err != nil {
			return fmt.Errorf("Bad choice")
		}
	} else {
		starterkits, err := starterkitRepo.Find(cCmd.starterkit)
		if err != nil || len(starterkits) == 0 {
			fmt.Fprintln(cCmd.out, "No starterkit found")
			return fmt.Errorf("No starterkit named %s found", cCmd.starterkit)
		}
		fmt.Fprintf(cCmd.out, "checkthis this %v", starterkits)
		if len(starterkits) == 1 {
			starterkit = starterkits[0]
			fmt.Println(starterkit)
		} else if len(starterkits) > 0 {
			fmt.Fprintf(cCmd.out, "multiple starterkit named %s found: %v", cCmd.starterkit, starterkits)
			starterkit, err = askUserChoice(starterkits, repoPath)
			if err != nil {
				return fmt.Errorf("Bad choice")
			}
		} else {
			return fmt.Errorf("no starterkit found with name %s in %s repo", cCmd.starterkit, cCmd.repo)
		}

	}

	err = starterkit.Create(cCmd.dest)
	if err != nil {
		fmt.Fprintln(cCmd.out, "Failed create starterkit")
		return err
	}
	fmt.Fprintln(cCmd.out, "Ready for development")
	return nil
}

func askUserChoice(starterkits []vega.StarterKit, repoPath string) (vega.StarterKit, error) {
	var starterkit vega.StarterKit
	prompt := &survey.Select{
		Message: "Select starterkit which you want to install:",
	}
	for _, starterkit := range starterkits {
		prompt.Options = append(prompt.Options, starterkit.Name)
	}
	var skName = ""
	err := survey.AskOne(prompt, &skName)
	if err != nil {
		return starterkit, err
	}

	starterkit = vega.StarterKit{
		Name: skName,
		Path: filepath.ToSlash(filepath.Join(repoPath, skName)),
	}

	return starterkit, nil

}
