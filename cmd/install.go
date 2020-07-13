package cmd

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/cobra"
	survey "github.com/AlecAivazis/survey/v2"

	vega "github.com/srijanone/vega/pkg/core"
)

type installCmd struct {
	out  io.Writer
	home vega.Home
	repo string // StarterKit Repo
	path string // Project path where starter kit is going to install
}

func newInstallCmd(out io.Writer) *cobra.Command {
	iCmd := installCmd{out: out}

	const installDesc = "install vega to existing repository"

	installCobraCmd := &cobra.Command{
		Use:   "install",
		Short: installDesc,
		Long:  installDesc,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				// If no arguments are passed, then choose current directory as path
				iCmd.path = "."
			} else {
				iCmd.path = args[0]
			}
			iCmd.home = vega.Home(homePath())

			iCmd.execute()
		},
	}

	flags := installCobraCmd.Flags()
	flags.StringVarP(&iCmd.repo, "repo", "r", "default", "name of the starterkit repo")

	return installCobraCmd
}

func (iCmd *installCmd) execute() {
	fmt.Fprintf(iCmd.out, "Installing vega to %s\n", iCmd.path)

	repoPath := filepath.Join(iCmd.home.StarterKits(), iCmd.repo)
	starterkitRepo := vega.StarterKitRepo{
		Name: iCmd.repo,
		Path: repoPath,
	}
	starterkits, err := starterkitRepo.StarterKitList()
	if err != nil {
		fmt.Fprintln(iCmd.out, "No starterkit found")
	}

	starterkit := ""
	prompt := &survey.Select{
		Message: "Select starterkit which you want to install:",
	}
	for _, starterkit := range starterkits {
		prompt.Options = append(prompt.Options, starterkit.Name)
	}
	survey.AskOne(prompt, &starterkit)
	fmt.Fprintf(iCmd.out,"You have selected " + starterkit + "\n")

	sk := vega.StarterKit{
		Name: starterkit,
		Path: filepath.ToSlash(filepath.Join(repoPath, starterkit)),
	}

	err = sk.Install(iCmd.path)
	if err != nil {
		fmt.Fprintf(iCmd.out, "Error in Installing : %v\n", err)
	} else {
		fmt.Fprintln(iCmd.out, "Starterkit installed successfully")
	}
}

