package vega

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"
)

// ErrStarterKitNotFoundInRepo is the error returned when a starterkit is not found in a starterkits
var ErrStarterKitNotFoundInRepo = errors.New("starterkit not found")

// StarterKit represents a starterkit repository.
type StarterKit struct {
	Name string
	Path string
}

type StarterKits struct {
}

// StarterKitFind finds a starterkits matching with the given name
func StarterKitFind(starterkitDir string, name string) ([]StarterKit, error) {
	starterkits := []StarterKit{}
	strterkitsList, err := StarterKitList(starterkitDir)
	if err != nil {
		return nil, err
	}
	for _, starterkit := range strterkitsList {
		if starterkit.Name == name {
			starterkits = nil
			starterkits = append(starterkits, starterkit)
			break
		}
		if strings.HasPrefix(starterkit.Name, name) {
			starterkits = append(starterkits, starterkit)
		}
	}
	return starterkits, nil
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

func StarterKitCreate(src string, dst string) error {
	err := copy.Copy(src, dst)
	return err
}
