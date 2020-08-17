package git_secrets

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/srijanone/vega/pkg/common"
	"github.com/srijanone/vega/pkg/git"
)

const (
	commandName  = "git-secrets"
	RequiredText = `
		git-secrets is not installed, which is required to run the application. 
	`
	InstallInstructions = `
		Install using: curl -fsSL https://raw.githubusercontent.com/srijanone/vega/develop/scripts/install_git_secrets.sh| bash
	`
)

func IsInstalled() bool {
	_, err := exec.LookPath(commandName)
	return err == nil
}

func Configure(out io.Writer) {
	templateDir := filepath.Join(common.DefaultHome(), ".git-templates", "git-secrets")
	// This is a very rudimentary check, it checks if host, port, password etc in the database
	// array in settings.php(drupal) is written in plain text. In case these are written in plain
	// text the developer might write them in "", or '' and in case these are externalise typically
	// developers would use https://www.php.net/manual/en/function.getenv.php or some other function.
	drupalSecretRegex := "(\"|')?(host|port|password|username)(\"|')?\\s*(=>)\\s*(\"|')+(.*)(\"|')+\\s*"

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
