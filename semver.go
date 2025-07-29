// Copyright (c) 2025 Michael D Henderson. All rights reserved.

// Package semver implements semantic versioning per the specification at https://semver.org/.
// It provides types and methods for creating, parsing, and comparing semantic version numbers.
package semver

import (
	"fmt"
	"strconv"
	"strings"
)

// Version represents a semantic version with major, minor, and patch version numbers,
// plus optional pre-release and build metadata components.
//
// Example versions:
//   - 1.0.0 (basic version)
//   - 1.0.0-alpha (with pre-release)
//   - 1.0.0+20130313144700 (with build metadata)
//   - 1.0.0-beta+exp.sha.5114f85 (with both)
type Version struct {
	Major      int    // Major version number (breaking changes)
	Minor      int    // Minor version number (new features, backward compatible)
	Patch      int    // Patch version number (bug fixes, backward compatible)
	PreRelease string // Pre-release identifier (alpha, beta, rc.1, etc.)
	Build      string // Build metadata (+20130313144700, +exp.sha.5114f85, etc.)
}

// String implements the fmt.Stringer interface and returns the semantic version
// string formatted according to https://semver.org/ rules.
//
// Format: MAJOR.MINOR.PATCH[-PRERELEASE][+BUILD]
//
// Examples:
//   - Version{1, 0, 0, "", ""} returns "1.0.0"
//   - Version{1, 0, 0, "alpha", ""} returns "1.0.0-alpha"
//   - Version{1, 0, 0, "", "20130313144700"} returns "1.0.0+20130313144700"
//   - Version{1, 0, 0, "beta", "exp.sha.5114f85"} returns "1.0.0-beta+exp.sha.5114f85"
func (v Version) String() string {
	hasPreRelease, hasBuild := v.PreRelease != "", v.Build != ""
	if hasPreRelease && hasBuild {
		return fmt.Sprintf("%d.%d.%d-%s+%s", v.Major, v.Minor, v.Patch, v.PreRelease, v.Build)
	} else if hasPreRelease && !hasBuild {
		return fmt.Sprintf("%d.%d.%d-%s", v.Major, v.Minor, v.Patch, v.PreRelease)
	} else if !hasPreRelease && hasBuild {
		return fmt.Sprintf("%d.%d.%d+%s", v.Major, v.Minor, v.Patch, v.Build)
	}
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Short returns the semantic version string without build metadata.
// This is useful when you want the version with pre-release information
// but don't need the build metadata for display purposes.
//
// Examples:
//   - Version{1, 0, 0, "", ""} returns "1.0.0"
//   - Version{1, 0, 0, "alpha", ""} returns "1.0.0-alpha"
//   - Version{1, 0, 0, "", "build123"} returns "1.0.0"
//   - Version{1, 0, 0, "beta", "build123"} returns "1.0.0-beta"
func (v Version) Short() string {
	if v.PreRelease != "" {
		return fmt.Sprintf("%d.%d.%d-%s", v.Major, v.Minor, v.Patch, v.PreRelease)
	}
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Core returns only the core semantic version numbers (major.minor.patch).
// This strips both pre-release and build metadata, returning just the
// base version number.
//
// Examples:
//   - Version{1, 0, 0, "", ""} returns "1.0.0"
//   - Version{1, 0, 0, "alpha", ""} returns "1.0.0"
//   - Version{1, 0, 0, "", "build123"} returns "1.0.0"
//   - Version{1, 0, 0, "beta", "build123"} returns "1.0.0"
func (v Version) Core() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// ByVersion implements sort.Interface for []Version based on semantic version precedence.
// This allows sorting slices of versions using the standard sort package.
//
// Example usage:
//   versions := []Version{{1, 0, 0, "beta", ""}, {1, 0, 0, "", ""}, {1, 0, 0, "alpha", ""}}
//   sort.Sort(ByVersion(versions))
//   // Result: versions are now sorted as [1.0.0-alpha, 1.0.0-beta, 1.0.0]
type ByVersion []Version

// Len returns the number of versions in the slice.
func (v ByVersion) Len() int {
	return len(v)
}

// Less reports whether the version at index i should sort before the version at index j.
// It uses the Compare method to determine semantic version precedence.
func (v ByVersion) Less(i, j int) bool {
	return v[i].Compare(v[j]) < 0
}

// Swap swaps the versions at indices i and j.
func (v ByVersion) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

// Equal returns true if this version is identical to v2 in all components.
// This includes major, minor, patch, pre-release, and build metadata.
// Note: Two versions that differ only in build metadata are considered different
// by Equal, even though they have the same precedence per SemVer spec.
// Unlike Less(), Equal cannot use Compare() because build metadata is significant
// for equality but ignored for precedence comparison.
func (v Version) Equal(v2 Version) bool {
	return v.Major == v2.Major && v.Minor == v2.Minor && v.Patch == v2.Patch && v.PreRelease == v2.PreRelease && v.Build == v2.Build
}

// Compare returns an integer comparing two versions according to semantic versioning precedence
// defined in https://semver.org/#spec-item-11.
// The result will be 0 if v == v2 (same precedence), -1 if v < v2, or +1 if v > v2.
// Build metadata is ignored in precedence comparison, consistent with the SemVer specification.
//
// Comparison precedence (from lowest to highest):
//  1. Major version number
//  2. Minor version number  
//  3. Patch version number
//  4. Pre-release version (when major.minor.patch are equal)
//     - Normal version (no pre-release) has higher precedence than pre-release
//     - Pre-release identifiers are compared lexically in ASCII sort order
//     - Numeric identifiers are compared numerically (1 < 2 < 10)
//     - Numeric identifiers have lower precedence than non-numeric (1 < alpha)
//     - Larger set of pre-release fields has higher precedence (1.0.0-alpha < 1.0.0-alpha.1)
//
// Examples: 1.0.0-alpha < 1.0.0-alpha.1 < 1.0.0-alpha.beta < 1.0.0-beta < 1.0.0-beta.2 < 1.0.0-beta.11 < 1.0.0-rc.1 < 1.0.0
//
// This method is useful for sorting and integrates well with Go's sort package:
//   sort.Slice(versions, func(i, j int) bool {
//       return versions[i].Compare(versions[j]) < 0
//   })
func (v Version) Compare(v2 Version) int {
	// Compare major version number
	if v.Major < v2.Major {
		return -1
	} else if v.Major > v2.Major {
		return 1
	}
	// Major is equal, compare minor
	if v.Minor < v2.Minor {
		return -1
	} else if v.Minor > v2.Minor {
		return 1
	}
	// Major and minor are equal, compare patch
	if v.Patch < v2.Patch {
		return -1
	} else if v.Patch > v2.Patch {
		return 1
	}
	// Major, minor, patch are equal, compare pre-release.
	// Per SemVer spec: normal release > prerelease version
	if v.PreRelease == "" && v2.PreRelease != "" {
		return 1
	}
	if v.PreRelease != "" && v2.PreRelease == "" {
		return -1
	}
	// Both have prerelease, compare them
	fields1 := strings.Split(v.PreRelease, ".")
	fields2 := strings.Split(v2.PreRelease, ".")
	for i := 0; i < len(fields1) && i < len(fields2); i++ {
		n1, err1 := strconv.Atoi(fields1[i])
		n2, err2 := strconv.Atoi(fields2[i])
		if err1 == nil && err2 == nil { // both fields are int
			if n1 < n2 {
				return -1
			} else if n1 > n2 {
				return 1
			}
			// n1 and n2 are equal, so compare the next field
		} else if err1 == nil { // only field1 is an int
			return -1 // numeric identifiers have lower precedence than non-numeric
		} else if err2 == nil { // only field2 is an int
			return 1  // numeric identifiers have lower precedence than non-numeric
		} else if fields1[i] < fields2[i] { // compare as text
			return -1
		} else if fields1[i] > fields2[i] { // compare as text
			return 1
		}
	}
	// All common fields are equal, compare field lengths
	if len(fields1) < len(fields2) {
		return -1
	} else if len(fields1) > len(fields2) {
		return 1
	}
	// Versions have same precedence (build metadata is ignored)
	return 0
}

// IsZero returns true if the version is 0.0.0.
func (v Version) IsZero() bool {
	return v == Version{}
}

// Less returns true if this version has lower precedence than v2 according to
// semantic versioning rules. This is a convenience method that calls Compare(v2) < 0.
// See Compare() for detailed precedence rules and examples.
func (v Version) Less(v2 Version) bool {
	return v.Compare(v2) < 0
}
