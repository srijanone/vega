package vega

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	common "github.com/srijanone/vega/pkg/common"
	git "github.com/srijanone/vega/pkg/git"
)

const (
	scriptHeader = "#!/usr/bin/env bash"
)

type GitHooks struct {
	Home Home
	URL  string
	Dir  string // hooks directory name at source/remote
	Out  io.Writer
}

// Install installs Git Hooks to a git based project path
func (gitHook *GitHooks) Install(path string) {
	gitHooksPath := filepath.Join(path, ".git", "hooks")

	exist, err := common.Exists(gitHooksPath)
	if err != nil {
		fmt.Fprintf(gitHook.Out, "couldn't check if project is git based: %v\n", err)
		return
	}
	if !exist {
		fmt.Fprintf(gitHook.Out, "Project is not git based. in order to install git hooks, please run `git init`\n")
		return
	}

	// gitHook.createHook("pre-commit", gitHooksPath)

	fmt.Fprintf(gitHook.Out, "Setting up Git Hooks \n")
	os.Chdir(path) // change directory to project path if user is not in current directory
	args := []string{"secrets", "--install", "-f"}
	git.Execute(gitHook.Out, args...)
}
