package common

import (
	generator "github.com/dustinkirkland/golang-petname"
)

func NormalizeApplicationName(name string) string {
	return name
}

func GeneratePetName () string {
	return generator.Generate(2, "-")
}