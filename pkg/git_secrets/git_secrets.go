package git_secrets

import (
	"errors"
	"fmt"
	"github.com/srijanone/vega/pkg/common"
	"github.com/srijanone/vega/pkg/git"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	commandName  = "git-secrets"
	RequiredText = `
		git-secrets is not installed, which is required to run the application. 
	`
	InstallInstructions = `
		Install using: curl -fsSL https://raw.githubusercontent.com/srijanone/vega/develop/scripts/install.sh| bash
	`
)

func IsInstalled() bool {
	_, err := exec.LookPath(commandName)
	return err == nil
}

// @TODO:
func Configure(out io.Writer) {
	templateDir := filepath.Join(common.DefaultHome(), ".git-templates", "git-secrets")
	drupalSecretRegex := "(\"|')?(host|port|password|username)(\"|')?\\s*(:|=>|=)\\s*(\"|')?(\".*\")(\"|')?\\s*"

	fmt.Print("Adding common AWS patterns to the git config...\n")
	execute(out, "--register-aws", "--global")

	fmt.Printf("Adding hooks to all local repositories...\n")
	execute(out, "--install", "-f", templateDir)
	args := []string{"config", "--global", "init.templateDir", templateDir}
	git.Execute(out, args...)

	fmt.Printf("Registering Drupal secrets patters...\n")
	execute(out, "--add", "--global", drupalSecretRegex)
}

func execute(out io.Writer, arguments ...string) error {
	if !IsInstalled() {
		fmt.Fprintf(out, RequiredText)
		fmt.Fprintf(out, InstallInstructions)
		return errors.New("git-secrets is not installed on system")
	}

	command := exec.Command(commandName, arguments...)
	command.Stdout = out
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}
