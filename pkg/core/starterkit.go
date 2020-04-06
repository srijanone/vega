package vega

import (
	"fmt"

	"github.com/otiai10/copy"
)

// StarterKit represents a starterkit repository.
type StarterKit struct {
	Name string
	Path string
}

type StarterKits []StarterKit

// DockerfileName: Default dockerfile name used in starterkits
const DockerfileName string = "Dockerfile"

func (sk *StarterKit) Create(dest string) error {
	// TODO: Merge gracefully in future
	fmt.Printf("Creating starterkit %s ... \n", sk.Name)
	err := copy.Copy(sk.Path, dest)
	return err
}
