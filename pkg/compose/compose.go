package compose

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

const (
	commandName  = "docker-compose"
	RequiredText = `
	docker-compose is not installed, which is required to run the application. 
	`
	InstallInstructions = `
		Install using:
		Linux : curl -L "https://github.com/docker/compose/releases/download/1.26.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
		Mac   : Install Docker Desktop

		And more info. can be found at: https://docs.docker.com/get-docker/
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

func Destroy(out io.Writer, arguments ...string) {
	arguments = append([]string{"down", "--volumes"}, arguments...)
	execute(out, arguments...)
}

func Run(out io.Writer, arguments ...string) {
	arguments = append([]string{"run", "--rm"}, arguments...)
	execute(out, arguments...)
}

func execute(out io.Writer, arguments ...string) error {
	if !IsInstalled() {
		fmt.Fprintf(out, RequiredText)
		fmt.Fprintf(out, InstallInstructions)
		return errors.New("docker-compose is not installed on system")
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
