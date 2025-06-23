# Semver

This is a simple Go package for handling semantic versioning (SemVer).

## Features

- Representation of semantic versions with major, minor, and patch version numbers, as well as optional pre-release and build metadata.
- Comparison of versions with `Less` method, according to the rules described in the [Semver Spec](https://semver.org/).
- Equality check with `Equal` method.

## Usage

Here's an example of how to use the `Version` struct:

```go
v1 := semver.Version{ Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha", Build: "001", }
fmt.Println(v1.String()) // Output: "1.0.0-alpha+001"
v2 := semver.Version{ Major: 1, Minor: 0, Patch: 1, }
if v1.Less(v2) { fmt.Println("v1 is less than v2") }
if !v1.Equal(v2) { fmt.Println("v1 is not equal to v2") }
```

## Testing

You can run the unit tests included in the project with the following command:

```shell
go test
```

## Contributing

If you would like to contribute, please fork the repository and use a feature branch. All contributions are welcome!

## License

This project is open source, under the [MIT License](LICENSE).# Semver
