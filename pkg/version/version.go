package version

import "fmt"

var (
	SemVer    = "v0.0.1"
	GitCommit = ""
	BuildTime = ""
)

func New() *Version {
	return &Version{
		Version:   getVersion(),
		GitCommit: GitCommit,
	}
}

type Version struct {
	Version   string
	GitCommit string
}

func (v *Version) String() string {
	return v.Version
}

func (v *Version) FormatVersion(short bool) string {
	if short {
		return fmt.Sprintf("%s", v.Version)
	}
	return fmt.Sprintf("%#v", v)
}

func getVersion() string {
	if BuildTime == "" {
		return SemVer
	}
	return fmt.Sprintf("%s+%s", SemVer, BuildTime)
}
