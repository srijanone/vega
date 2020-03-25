package vega

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ErrStarterKitNotFoundInRepo is the error returned when a starterkit is not found in a starterkits
var ErrStarterKitNotFoundInRepo = errors.New("starterkit not found")

// StarterKit represents a starterkit repository.
type StarterKit struct {
	Name string
	Path string
}

// StarterKitFind finds a starterkits with the given name  and returns path
func StarterKitFind(starterkitDir string, name string) (string, error) {
	if _, err := os.Stat(starterkitDir); os.IsNotExist(err) {
		return "", fmt.Errorf("starterkit dir %s not found", starterkitDir)
	}

	targetDir := filepath.Join(starterkitDir, name)
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		return "", ErrStarterKitNotFoundInRepo
	}

	return targetDir, nil
}

// StarterKitList returns a list of all Starter-Kits.
func StarterKitList(starterkitDir string) ([]StarterKit, error) {
	switch fi, err := os.Stat(starterkitDir); {
	case err != nil:
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("starterkit directory %s not found", starterkitDir)
		}
	case !fi.IsDir():
		return nil, fmt.Errorf("%s is not a directory", starterkitDir)
	}
	var starterkits []StarterKit
	files, err := ioutil.ReadDir(starterkitDir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			starterkit := &StarterKit{}
			starterkit.Name = file.Name()
			starterkit.Path = filepath.ToSlash(filepath.Join(starterkitDir, file.Name()))
			starterkits = append(starterkits, *starterkit)
		}
	}
	return starterkits, nil
}
