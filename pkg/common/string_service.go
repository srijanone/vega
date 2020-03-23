package common

import (
	"strings"

	generator "github.com/dustinkirkland/golang-petname"
)

func NormalizeApplicationName(name string) string {
	return strings.ToLower(name)
}

func GeneratePetName() string {
	return generator.Generate(2, "-")
}
