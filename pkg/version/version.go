// Package version provides build-time version information for the pentagi application.
package version

import (
	"fmt"
	"runtime"
)

var (
	// Version is the current version of pentagi.
	// This is typically set at build time via ldflags:
	// -ldflags "-X github.com/vxcontrol/pentagi/pkg/version.Version=1.0.0"
	Version = "dev"

	// GitCommit is the git commit hash at build time.
	GitCommit = "unknown"

	// GitBranch is the git branch at build time.
	GitBranch = "unknown"

	// BuildDate is the date when the binary was built.
	BuildDate = "unknown"
)

// Info holds all version-related information.
type Info struct {
	Version   string `json:"version"`
	GitCommit string `json:"git_commit"`
	GitBranch string `json:"git_branch"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
}

// Get returns the current version information.
func Get() Info {
	return Info{
		Version:   Version,
		GitCommit: GitCommit,
		GitBranch: GitBranch,
		BuildDate: BuildDate,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
}

// String returns a human-readable version string.
func (i Info) String() string {
	return fmt.Sprintf(
		"pentagi %s (commit: %s, branch: %s, built: %s, %s %s/%s)",
		i.Version,
		i.GitCommit,
		i.GitBranch,
		i.BuildDate,
		i.GoVersion,
		i.OS,
		i.Arch,
	)
}
