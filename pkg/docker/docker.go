package docker

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

const (
	commandName  = "docker"
	RequiredText = `
	docker is not installed, which is required to run the application. 
	`
	InstallInstructions = `
		Instructions to install can be found at: https://docs.docker.com/get-docker/
	`
)

func IsInstalled() bool {
	_, err := exec.LookPath(commandName)
	return err == nil
}

func DeleteImagesByLabel(out io.Writer, label string) {
	dockerCmd := `docker image rm $(docker image ls -q --filter "label=` + label + `") 2> /dev/null`
	arguments := append([]string{"-c"}, dockerCmd)
	execute(out, arguments...)
}

func execute(out io.Writer, arguments ...string) error {
	if !IsInstalled() {
		fmt.Fprintf(out, RequiredText)
		fmt.Fprintf(out, InstallInstructions)
		return errors.New("docker is not installed on system")
	}
	// TODO: Windows compatibility
	command := exec.Command("bash", arguments...)
	command.Stdout = out
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}
