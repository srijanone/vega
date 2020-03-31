package vega

import (
	"path/filepath"
)

// This builds paths relative to a Vega Home directory.
type Home string

func (home Home) String() string {
	return string(home)
}

func (home Home) Path(elements ...string) string {
	path := []string{home.String()}
	path = append(path, elements...)
	return filepath.Join(path...)
}

func (home Home) StarterKits() string {
	return home.Path("starterkits")
}

func (home Home) Repos() string {
	return home.Path("repos")
}

func (home Home) Logs() string {
	return home.Path("logs")
}
