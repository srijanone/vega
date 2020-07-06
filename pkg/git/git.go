package git

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

const (
	commandName  = "git"
	RequiredText = `
		Git is not installed, which is required to install Git Hooks. 
	`
	InstallInstructions = `
		Install using: https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
	`
)

func IsInstalled() bool {
	_, err := exec.LookPath(commandName)
	return err == nil
}

func Execute(out io.Writer, arguments ...string) error {
	if !IsInstalled() {
		fmt.Fprintf(out, RequiredText)
		fmt.Fprintf(out, InstallInstructions)
		return errors.New("git is not installed on system")
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
