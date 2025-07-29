// Copyright (c) 2025 Michael D Henderson. All rights reserved.

package semver

import (
	"runtime/debug"
)

// Commit returns the VCS commit information for the current build.
// Returns the first 7 characters of the commit hash (matching Git's short format)
// with "-dirty" suffix if the working directory has uncommitted changes.
// Returns "dirty" if no commit hash is available but working directory is dirty.
// Returns empty string if no VCS information is available.
func Commit() string {
	var commitHash string
	var dirtyWorkingDirectory bool

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.modified":
				dirtyWorkingDirectory = setting.Value == "true"
			case "vcs.revision":
				if len(setting.Value) > 7 {
					commitHash = setting.Value[:7]
				} else {
					commitHash = setting.Value
				}
			}
		}
	}

	if commitHash != "" && dirtyWorkingDirectory {
		if dirtyWorkingDirectory {
			return commitHash + "-dirty"
		}
		return commitHash
	} else if dirtyWorkingDirectory {
		// edge case after "git init" and before the first commit
		return "dirty"
	}
	return ""
}
