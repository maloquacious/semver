// Copyright (c) 2025 Michael D Henderson. All rights reserved.

package semver

import (
	"runtime/debug"
)

// Commit returns the VCS commit information for the current build.
// Returns the first 7 characters of the commit hash (matching Git's short format)
// with "-dirty" suffix if the working directory has uncommitted changes.
// Returns "*-dirty" if no commit hash is available but working directory is dirty.
// Returns empty string if no VCS information is available.
func Commit() string {
	var commitHash, dirtyWorkingDirectory string

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.modified":
				dirtyWorkingDirectory = setting.Value
			case "vcs.revision":
				if len(setting.Value) > 7 {
					commitHash = setting.Value[:7]
				} else {
					commitHash = setting.Value
				}
			}
		}
	}

	if commitHash != "" && dirtyWorkingDirectory != "" {
		return commitHash + "-dirty"
	} else if commitHash != "" {
		return commitHash
	} else if dirtyWorkingDirectory != "" {
		return "*-dirty"
	}
	return ""
}
