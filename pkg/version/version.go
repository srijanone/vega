package version

import "fmt"

type Version struct {
	SemVer       string `json:"semver"`
	GitCommit    string `json:"git-commit"`
	GitTreeState string `json:"git-tree-state"`
}

func (v *Version) String() string {
	return v.SemVer
}

func (v *Version) FormatVersion(short bool) string {
	if short {
		return fmt.Sprintf("%s+g%s", v.SemVer, v.GitCommit[:7])
	}
	return fmt.Sprintf("%#v", v)
}

var (
	Release       = "canary"
	BuildMetadata = ""
	GitCommit     = ""
	GitTreeState  = ""
)

func getVersion() string {
	if BuildMetadata == "" {
		return Release
	}
	return Release + "+" + BuildMetadata
}

func New() *Version {
	return &Version{
		SemVer:       getVersion(),
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
	}
}
