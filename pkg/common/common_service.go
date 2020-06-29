package common

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

// Exists checks if a file or folder exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// EnsureDir ensures a directory is present, if it doesn't then creates it
func EnsureDir(dir string) error {
	if info, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("Could not create %s: %s", dir, err)
		}
	} else if !info.IsDir() {
		// CHECK: Why not errors.New()?
		return fmt.Errorf("%s must be a directory", dir)
	}

	return nil
}

// DefaultHome returns User's Home Directory
func DefaultHome() string {
	homeEnvPath := os.Getenv("HOME")
	if homeEnvPath == "" && runtime.GOOS == "windows" {
		homeEnvPath = os.Getenv("USERPROFILE")
	}
	return homeEnvPath
}

// CopyFile copies a file from source to destinationn
func CopyFile(sourceFile string, destFile string) error {
	from, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("couldn't open source file %v: %v", sourceFile, err)
	}
	defer from.Close()

	to, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("couldn't create/open destination file %v: %v", destFile, err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		return fmt.Errorf("couldn't copy file: %v", err)
	}
	return nil
}
