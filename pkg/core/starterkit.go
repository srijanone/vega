package vega

import (
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
	err := copy.Copy(sk.Path, dest)
	return err
}
