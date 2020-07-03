package vega

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	common "github.com/srijanone/vega/pkg/common"
	downloader "github.com/srijanone/vega/pkg/downloader"
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

// Add downloads git hooks to vega home
func (gitHook *GitHooks) Add() {
	d := downloader.Downloader{}
	if gitHook.Dir == "" {
		gitHook.Dir = Home("").GitHooks()
	}
	sourceRepo := fmt.Sprintf("%s//%s", gitHook.URL, gitHook.Dir)
	fmt.Println("Downloading git hooks...")
	d.Download(sourceRepo, gitHook.Home.GitHooks())
}

// InstallGlobally installs Git Hooks as Global Git Hooks
func (gitHook *GitHooks) InstallGlobally() {
	globalHooksDir := filepath.Join(common.DefaultHome(), ".git", "hooks")

	fmt.Fprintf(gitHook.Out, "Creating Global Hooks Directory\n")
	if err := common.EnsureDir(globalHooksDir); err != nil {
		fmt.Fprintf(gitHook.Out, "Error in global hook directory: %v\n", err)
	}

	gitHook.createHook("pre-commit", globalHooksDir)

	fmt.Fprintf(gitHook.Out, "Setting Global Git Hooks: %v\n", globalHooksDir)
	args := []string{"config", "--global", "core.hooksPath", globalHooksDir}
	git.Execute(gitHook.Out, args...)
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

	gitHook.createHook("pre-commit", gitHooksPath)

	fmt.Fprintf(gitHook.Out, "Setting Up Local Git Hooks \n")
	os.Chdir(path) // change directory to project path if user is not in current directory
	args := []string{"config", "core.hooksPath", ".git/hooks"}
	git.Execute(gitHook.Out, args...)
}

func (gitHook *GitHooks) createHook(hookName string, path string) {
	fmt.Fprintf(gitHook.Out, "Installing %v hooks\n", hookName)
	var shellScripts []string

	preCommitHooksDir := filepath.Join(gitHook.Home.GitHooks(), "generic", hookName)
	preCommitScriptBody := scriptHeader + "\n"

	shellScripts = common.ListFiles(preCommitHooksDir)
	for _, shellScript := range shellScripts {
		fmt.Fprintf(gitHook.Out, "Adding hook: %v\n", shellScript)
		preCommitScriptBody = preCommitScriptBody + "\n" + shellScript
	}

	err := ioutil.WriteFile(filepath.Join(path, hookName), []byte(preCommitScriptBody), 0755)
	if err != nil {
		fmt.Fprintf(gitHook.Out, "couldn't create %v hook: %v\n", hookName, err)
	}
}