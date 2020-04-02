package vega

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	downloader "github.com/srijanone/vega/pkg/downloader"
)

// ErrStarterKitNotFoundInRepo is the error returned when a starterkit is not found in a starterkits
var ErrStarterKitNotFoundInRepo = errors.New("starterkit not found")

type StarterKitRepo struct {
	Name string
	Path string // local absolute path to repo
	Home Home
	URL  string
	Dir  string // starterkits directory name at source/remote
}

type StarterKitRepos []StarterKitRepo

// RepoList list of all the local Repositories
func RepoList(path string) ([]StarterKitRepo, error) {
	var repositories StarterKitRepos

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			repository := &StarterKitRepo{
				Name: file.Name(),
				Path: filepath.ToSlash(filepath.Join(path, file.Name())),
			}
			repositories = append(repositories, *repository)
		}
	}
	return repositories, nil
}

// StarterKitList Gets the list of all StarterKits for given repo.
func (repo *StarterKitRepo) StarterKitList() (StarterKits, error) {
	var starterkits StarterKits

	switch fileInfo, err := os.Stat(repo.Path); {
	case err != nil:
		if os.IsNotExist(err) {
			return nil, errors.New("No repo found, please type `vega init`")
		}
	case !fileInfo.IsDir():
		return nil, errors.New(fmt.Sprintf("%s is not a starterkit repo", repo.Path))
	}

	files, err := ioutil.ReadDir(repo.Path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			starterkit := &StarterKit{
				Name: file.Name(),
				Path: filepath.ToSlash(filepath.Join(repo.Path, file.Name())),
			}
			starterkits = append(starterkits, *starterkit)
		}
	}
	return starterkits, nil
}

// Find returns the starterkits matching with the given name
func (repo *StarterKitRepo) Find(name string) ([]StarterKit, error) {
	var starterkits StarterKits

	starterkitList, err := repo.StarterKitList()
	if err != nil {
		return nil, err
	}
	for _, starterkit := range starterkitList {
		// Trying to match exact name first
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

// Add adds repo and starterkits to vega
func (repo *StarterKitRepo) Add() {
	d := downloader.Downloader{}
	if repo.Dir == "" {
		repo.Dir = repo.Home.StarterKits()
	}
	sourceRepo := fmt.Sprintf("%s//%s", repo.URL, repo.Dir)
	fmt.Println("Downloading starterkits...")
	if repo.Path == "" {
		repo.Path = filepath.Join(repo.Home.StarterKits(), repo.Name)
	}
	d.Download(sourceRepo, repo.Path)
}

// Delete deletes the starterkit repo
func (repo StarterKitRepo) Delete() {
	if repo.Dir == "" {
		repo.Dir = repo.Home.StarterKits()
	}
	// TODO
}
