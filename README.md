# Semver

This is a simple Go package for handling semantic versioning (SemVer).

## Features

- Representation of semantic versions with major, minor, and patch version numbers, as well as optional pre-release and build metadata.
- Comparison of versions with `Less` method, according to the rules described in the [Semver Spec](https://semver.org/).
- Equality check with `Equal` method.
- Automatic VCS commit information extraction with `Commit()` function for build metadata.
- Package version introspection with `Current()` function.

## Usage

Here's an example of how to use the `Version` struct:

```go
v1 := semver.Version{ Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha", Build: "001", }
fmt.Println(v1.String()) // Output: "1.0.0-alpha+001"
v2 := semver.Version{ Major: 1, Minor: 0, Patch: 1, }
if v1.Less(v2) { fmt.Println("v1 is less than v2") }
if !v1.Equal(v2) { fmt.Println("v1 is not equal to v2") }
```

### Using Build Metadata with VCS Information

The `Commit()` function automatically extracts VCS commit information to populate build metadata:

```go
// Automatically populate build metadata with commit info
version := semver.Version{
    Major: 1,
    Minor: 2,
    Patch: 3,
    PreRelease: "beta",
    Build: semver.Commit(), // Uses Git commit hash from build info
}

fmt.Println(version.String())
// Examples of output:
// "1.2.3-beta+abc1234"        (clean build from commit abc1234)
// "1.2.3-beta+abc1234-dirty"  (build with uncommitted changes)
// "1.2.3-beta+*-dirty"        (dirty build, no commit hash available)
```

This is particularly useful for:
- **CI/CD pipelines** - Track which commit produced each build
- **Deployment tracking** - Link deployed artifacts to source code
- **Debugging** - Identify the exact code version causing issues

### Package Version Information

You can get the version of the semver package itself:

```go
fmt.Printf("Using semver package version: %s\n", semver.Current().String())
// Output: "Using semver package version: 0.2.0+abc1234"
```

### Sorting Versions

The `Compare()` method enables easy sorting of version slices:

```go
versions := []semver.Version{
    {Major: 1, Minor: 0, Patch: 0, PreRelease: "beta"},
    {Major: 1, Minor: 0, Patch: 0},
    {Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha"},
    {Major: 0, Minor: 9, Patch: 0},
}

sort.Slice(versions, func(i, j int) bool {
    return versions[i].Compare(versions[j]) < 0
})

// Result: [0.9.0, 1.0.0-alpha, 1.0.0-beta, 1.0.0]
```

### String Formatting

Different string representations for various use cases:

```go
version := semver.Version{
    Major: 1, Minor: 2, Patch: 3, 
    PreRelease: "beta.1", 
    Build: "20250129.abc123",
}

fmt.Println(version.String()) // "1.2.3-beta.1+20250129.abc123" (full version)
fmt.Println(version.Short())  // "1.2.3-beta.1" (without build metadata)
fmt.Println(version.Core())   // "1.2.3" (core version only)
```

## Testing

You can run the unit tests included in the project with the following command:

```shell
go test
```

## Contributing

If you would like to contribute, please fork the repository and use a feature branch. All contributions are welcome!

## License

This project is open source, under the [MIT License](LICENSE).
