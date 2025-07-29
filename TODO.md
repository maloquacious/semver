# TODO: Missing API Helpers and Improvements

## ✅ Recently Completed
- [x] **SemVer Precedence Bug Fix** - Fixed numeric vs alphanumeric ordering in prerelease comparison
- [x] **Compare() Method** - Efficient primary comparison method with full SemVer compliance
- [x] **Less() Optimization** - Refactored to use Compare() for better performance  
- [x] **Comprehensive Test Coverage** - Added 8 additional test cases covering edge cases
- [x] **Documentation Enhancement** - Complete method documentation with examples
- [x] **IsZero() Method** - Basic version 0.0.0 detection
- [x] **Architecture Improvement** - Single source of truth for comparison logic
- [x] **String Formatting** - Added Short() and Core() methods for flexible version display

## High Priority - API Completeness

### Parse Functions
- [ ] `Parse(s string) (Version, error)` - Parse version string with validation
- [ ] `MustParse(s string) Version` - Parse with panic on error
- [ ] Validation rules:
  - [ ] Major/Minor/Patch must be non-negative integers
  - [ ] No leading zeros except for "0" itself
  - [ ] Pre-release identifiers: alphanumeric + hyphen only, no empty identifiers
  - [ ] Build metadata: alphanumeric + hyphen only, no empty identifiers

### Comparison Helpers
- [x] `Compare(v2 Version) int` - Returns -1, 0, 1 (for sorting compatibility) ✅
- [ ] `Validate() error` - Validate current version struct

### Version Manipulation
- [ ] `NextMajor() Version` - Increment major, reset minor/patch to 0
- [ ] `NextMinor() Version` - Increment minor, reset patch to 0  
- [ ] `NextPatch() Version` - Increment patch
- [ ] `WithPreRelease(pre string) Version` - Set pre-release identifier
- [ ] `WithBuild(build string) Version` - Set build metadata
- [ ] `StripPreRelease() Version` - Remove pre-release identifier
- [ ] `StripBuild() Version` - Remove build metadata

### Query Methods
- [ ] `IsPreRelease() bool` - Check if version has pre-release identifier
- [ ] `IsStable() bool` - Check if version is stable (no pre-release)
- [x] `IsZero() bool` - Check if version is 0.0.0 ✅
- [ ] `Major() int`, `Minor() int`, `Patch() int` - Getter methods
- [ ] `PreRelease() string`, `Build() string` - Getter methods

## Medium Priority - Serialization & Compatibility

### Text Marshaling
- [ ] `MarshalText() ([]byte, error)` - For JSON/XML/flags compatibility
- [ ] `UnmarshalText(text []byte) error` - For JSON/XML/flags compatibility

### Range Operations
- [ ] `InRange(constraint string) bool` - Check if version satisfies constraint
- [ ] `Satisfies(constraint string) bool` - Alias for InRange
- [ ] Constraint syntax: `^1.2.3`, `~1.2.3`, `>=1.0.0 <2.0.0`

### Sorting Integration
- [ ] `ByVersion` type for sorting slices of versions
- [x] Example usage in documentation ✅ (Compare() examples in README)

## Low Priority - Enhanced Features

### Collection Operations
- [ ] `Latest(versions []Version) Version` - Find highest version
- [ ] `Sort(versions []Version)` - Sort versions in place
- [ ] `Filter(versions []Version, constraint string) []Version`

### String Formatting
- [x] `String() string` variants:
  - [x] `Short() string` - Without build metadata ✅
  - [x] `Core() string` - Only major.minor.patch ✅
  - [x] `Format(template string) string` - Not needed: users can access public fields directly ✅

### Advanced Validation
- [ ] `IsValidIdentifier(s string) bool` - Check pre-release/build identifier
- [ ] `NormalizePreRelease(s string) string` - Clean up pre-release format
- [ ] Support for version ranges and constraints

## Documentation & Examples

### README Updates
- [ ] Add Parse/MustParse examples
- [ ] Add constraint checking examples  
- [x] Add sorting examples ✅ (Compare() sorting example added)
- [ ] Add JSON marshaling examples

### Go Doc Examples
- [x] Add testable examples for all major functions ✅ (comprehensive documentation added)
- [ ] Benchmark tests for performance-critical operations

## Breaking Changes (v2.0+)

### Immutable Design
- [ ] Consider making Version fields private
- [ ] Force construction through Parse/New functions
- [ ] Ensure all operations return new Version instances

### API Consistency  
- [ ] Standardize error messages and types
- [ ] Consider context.Context support for Parse operations
- [ ] Consistent naming patterns across all methods
