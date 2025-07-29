// Copyright (c) 2025 Michael D Henderson. All rights reserved.

// Package semver_test implements tests for the semver package.
package semver_test

import (
	"github.com/maloquacious/semver"
	"testing"
)

// Test for String method
func TestString(t *testing.T) {
	testCases := []struct {
		desc     string
		version  semver.Version
		expected string
	}{
		{
			desc:     "Major, minor and patch only",
			version:  semver.Version{Major: 1, Minor: 0, Patch: 0},
			expected: "1.0.0",
		},
		{
			desc:     "PreRelease only",
			version:  semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha"},
			expected: "1.0.0-alpha",
		},
		{
			desc:     "Build only",
			version:  semver.Version{Major: 1, Minor: 0, Patch: 0, Build: "0001"},
			expected: "1.0.0+0001",
		},
		{
			desc:     "PreRelease and Build",
			version:  semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "beta", Build: "0002"},
			expected: "1.0.0-beta+0002",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := tc.version.String()
			if actual != tc.expected {
				t.Errorf("Unexpected version string. expected: %v, actual: %v", tc.expected, actual)
			}
		})
	}
}

// Test for Equal method
func TestEqual(t *testing.T) {
	testCases := []struct {
		desc     string
		version1 semver.Version
		version2 semver.Version
		expected bool
	}{
		{
			desc:     "Same version",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0},
			expected: true,
		},
		{
			desc:     "Different patch",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 1},
			expected: false,
		},
		{
			desc:     "Different build",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0, Build: "0001"},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0, Build: "0002"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := tc.version1.Equal(tc.version2)
			if actual != tc.expected {
				t.Errorf("Unexpected comparison result for Equal method. expected: %v, actual: %v", tc.expected, actual)
			}
		})
	}
}

// Test for Less method
func TestLess(t *testing.T) {
	testCases := []struct {
		desc     string
		version1 semver.Version
		version2 semver.Version
		expected bool
	}{
		{
			desc:     "Same version",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0},
			expected: false,
		},
		{
			desc:     "version1 is less",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 1},
			expected: true,
		},
		{
			desc:     "version2 is less",
			version1: semver.Version{Major: 1, Minor: 1, Patch: 0},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 1},
			expected: false,
		},
		{
			desc:     "comparison with PreRelease versions",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha"},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "beta"},
			expected: true,
		},
		{
			desc:     "normal release > prerelease",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "rc.1"},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0},
			expected: true,
		},
		{
			desc:     "numeric vs numeric prerelease",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha.1"},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha.2"},
			expected: true,
		},
		{
			desc:     "numeric vs alphanumeric prerelease (numeric < alphanumeric)",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha.1"},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha.beta"},
			expected: true,
		},
		{
			desc:     "alphanumeric vs numeric prerelease (alphanumeric > numeric)",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha.beta"},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha.1"},
			expected: false,
		},
		{
			desc:     "different prerelease field lengths",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha"},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha.1"},
			expected: true,
		},
		{
			desc:     "build metadata ignored in comparison",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0, Build: "build1"},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0, Build: "build2"},
			expected: false,
		},
		{
			desc:     "complex prerelease comparison from SemVer spec",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha.beta"},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha.beta.1"},
			expected: true,
		},
		{
			desc:     "larger numeric identifiers",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "beta.2"},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "beta.11"},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := tc.version1.Less(tc.version2)
			if actual != tc.expected {
				t.Errorf("Unexpected comparison result. expected: %v, actual: %v", tc.expected, actual)
			}
		})
	}
}

// Test for Compare method
func TestCompare(t *testing.T) {
	testCases := []struct {
		desc     string
		version1 semver.Version
		version2 semver.Version
		expected int
	}{
		{
			desc:     "same precedence (equal versions)",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0},
			expected: 0,
		},
		{
			desc:     "same precedence (different build metadata)",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0, Build: "build1"},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0, Build: "build2"},
			expected: 0,
		},
		{
			desc:     "version1 less than version2",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 1},
			expected: -1,
		},
		{
			desc:     "version1 greater than version2",
			version1: semver.Version{Major: 1, Minor: 1, Patch: 0},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 1},
			expected: 1,
		},
		{
			desc:     "prerelease vs normal release",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha"},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0},
			expected: -1,
		},
		{
			desc:     "normal release vs prerelease",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha"},
			expected: 1,
		},
		{
			desc:     "numeric vs alphanumeric prerelease",
			version1: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha.1"},
			version2: semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha.beta"},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := tc.version1.Compare(tc.version2)
			if actual != tc.expected {
				t.Errorf("Unexpected comparison result. expected: %v, actual: %v", tc.expected, actual)
			}
		})
	}
}

// Test for Short method
func TestShort(t *testing.T) {
	testCases := []struct {
		desc     string
		version  semver.Version
		expected string
	}{
		{
			desc:     "Major, minor and patch only",
			version:  semver.Version{Major: 1, Minor: 0, Patch: 0},
			expected: "1.0.0",
		},
		{
			desc:     "PreRelease only",
			version:  semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha"},
			expected: "1.0.0-alpha",
		},
		{
			desc:     "Build only (stripped)",
			version:  semver.Version{Major: 1, Minor: 0, Patch: 0, Build: "build123"},
			expected: "1.0.0",
		},
		{
			desc:     "PreRelease and Build (build stripped)",
			version:  semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "beta", Build: "build123"},
			expected: "1.0.0-beta",
		},
		{
			desc:     "Complex prerelease, build stripped",
			version:  semver.Version{Major: 2, Minor: 1, Patch: 3, PreRelease: "rc.1", Build: "20230101.abc123"},
			expected: "2.1.3-rc.1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := tc.version.Short()
			if actual != tc.expected {
				t.Errorf("Unexpected version string. expected: %v, actual: %v", tc.expected, actual)
			}
		})
	}
}

// Test for Core method
func TestCore(t *testing.T) {
	testCases := []struct {
		desc     string
		version  semver.Version
		expected string
	}{
		{
			desc:     "Major, minor and patch only",
			version:  semver.Version{Major: 1, Minor: 0, Patch: 0},
			expected: "1.0.0",
		},
		{
			desc:     "PreRelease stripped",
			version:  semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha"},
			expected: "1.0.0",
		},
		{
			desc:     "Build stripped",
			version:  semver.Version{Major: 1, Minor: 0, Patch: 0, Build: "build123"},
			expected: "1.0.0",
		},
		{
			desc:     "Both PreRelease and Build stripped",
			version:  semver.Version{Major: 1, Minor: 0, Patch: 0, PreRelease: "beta", Build: "build123"},
			expected: "1.0.0",
		},
		{
			desc:     "Complex version numbers",
			version:  semver.Version{Major: 12, Minor: 34, Patch: 56, PreRelease: "rc.1", Build: "20230101.abc123"},
			expected: "12.34.56",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := tc.version.Core()
			if actual != tc.expected {
				t.Errorf("Unexpected version string. expected: %v, actual: %v", tc.expected, actual)
			}
		})
	}
}
