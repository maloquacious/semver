package semver

// Current returns the version of this semver package itself.
// The build metadata is automatically populated with VCS commit information
// using the Commit() function.
func Current() Version {
	return Version{
		Major:      0,
		Minor:      3,
		Patch:      0,
		PreRelease: "",
		Build:      Commit(),
	}
}
