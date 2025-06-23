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
