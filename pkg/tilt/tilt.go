package tilt

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

const (
	commandName  = "tilt"
	RequiredText = `
		Tilt is not installed, which is required to run the application. 
	`
	InstallInstructions = `
		Install using: curl -fsSL https://raw.githubusercontent.com/windmilleng/tilt/master/scripts/install.sh | bash
		And more info. can be found at: https://docs.tilt.dev/install.html
	`
)

func IsInstalled() bool {
	_, err := exec.LookPath(commandName)
	return err == nil
}

func Up(out io.Writer, arguments ...string) {
	arguments = append([]string{"up"}, arguments...)
	execute(out, arguments...)
}

func Down(out io.Writer, arguments ...string) {
	arguments = append([]string{"down"}, arguments...)
	execute(out, arguments...)
}

func execute(out io.Writer, arguments ...string) {
	command := exec.Command(commandName, arguments...)
	command.Stdout = out
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Fprintln(out, "Error in Executing", err)
	}
}
