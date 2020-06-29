package vega

import (
	"fmt"
	"io"
	"path/filepath"

	downloader "github.com/srijanone/vega/pkg/downloader"
	common "github.com/srijanone/vega/pkg/common"
	git "github.com/srijanone/vega/pkg/git"
)

type GitHooks struct {
	Path string // local absolute path to repo
	Home Home
	URL  string
	Dir  string // hooks directory name at source/remote
	Out io.Writer
}

// Add downloads git hooks to vega home
func (gitHook *GitHooks) Add() {
	d := downloader.Downloader{}
	if gitHook.Dir == "" {
		gitHook.Dir = Home("").GitHooks()
	}
	sourceRepo := fmt.Sprintf("%s//%s", gitHook.URL, gitHook.Dir)
	fmt.Println("Downloading git hooks...")
	if gitHook.Path == "" {
		gitHook.Path = gitHook.Home.GitHooks()
	}
	d.Download(sourceRepo, gitHook.Path)
}

// Install installs Git Hooks as Global Git Hooks
func (gitHook *GitHooks) Install() {
	// TODO: Use code instead of bash script to accommodate windows OS

	globalHooksDir := filepath.Join(common.DefaultHome(), ".git", "hooks")

	fmt.Fprintf(gitHook.Out, "Creating Global Hooks Directory\n")
	if err := common.EnsureDir(globalHooksDir); err != nil {
		fmt.Fprintf(gitHook.Out, "Error in global hook directory: %v\n", err)
	}

	args := []string{"config", "--global", "core.hooksPath", globalHooksDir}

	fmt.Fprintf(gitHook.Out, "Setting Global Git Hooks: %v\n", globalHooksDir)
	git.Execute(gitHook.Out, args...)

	sourceDir := filepath.Join(gitHook.Path, "generic", "pre-commit")

	fmt.Fprintf(gitHook.Out, "Installing pre-commit hooks\n")
	// TODO: Make it generic to install all hooks
	err := common.CopyFile(filepath.Join(sourceDir, "check-aws-credentials.sh"), filepath.Join(globalHooksDir, "pre-commit"))
	if err != nil {
		fmt.Fprintf(gitHook.Out, "Error: %v\n", err)
	}
}