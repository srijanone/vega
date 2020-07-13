package vega

import (
	"fmt"
	"path/filepath"

	"github.com/otiai10/copy"

	common "github.com/srijanone/vega/pkg/common"
)

// StarterKit represents a starterkit repository.
type StarterKit struct {
	Name string
	Path string
}

type StarterKits []StarterKit

// DockerfileName: Default dockerfile name used in starterkits
const DockerfileName string = "Dockerfile"

// Create creates a new project at dest from given starter kit
func (sk *StarterKit) Create(dest string) error {
	// TODO: Merge gracefully in future
	fmt.Printf("Creating starterkit %s ... \n", sk.Name)
	err := copy.Copy(sk.Path, dest)
	return err
}

// Install installs a starterkit at existing project
func (sk *StarterKit) Install(dest string) error {
	// TODO: This filesToCopy is dependent of type of starterkit
	// files for Drupal starterkit
	filesToCopy := []string{
		// Docker Related
		"Dockerfile",
		".dockerignore",
		"docker-compose.yml",

		// Tilt Related
		"Tiltfile",

		// Others
		".env",
	}

	for _, file := range filesToCopy {
		srcFile := filepath.ToSlash(filepath.Join(sk.Path, file))
		destFile := filepath.ToSlash(filepath.Join(dest, file))
		err := common.CopyFile(srcFile, destFile)
		if err != nil {
			return err
		}
	}

	// Directory To Copy
	dirToCopy := []string{
		// Debug Related
		".vscode",
	}

	for _, dir := range dirToCopy {
		srcDir := filepath.ToSlash(filepath.Join(sk.Path, dir))
		destDir := filepath.ToSlash(filepath.Join(dest, dir))
		err := copy.Copy(srcDir, destDir)
		if err != nil {
			return err
		}
	}

	return nil
}
