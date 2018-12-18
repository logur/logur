# Change Log


All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).


## [Unreleased]

### Added

- Contextual logger (instead of `Logger.WithFields`)
- Field parameter to log functions

### Changed

- Replace log func variadic arguments with a single message argument
- Check if level is enabled (to prevent unwanted context conversions) when the underlying logger supports it

### Removed

- Remove format functions from `Logger` interface
- Remove ln functions from `Logger` interface
- Simple log adapter (implementing format and ln functions)
- `Logger.WithFields` method (use field parameter of log functions instead)


## [0.5.0] - 2018-12-17

### Added

- [Watermill](https://watermill.io) compatible logger

### Changed

- Dropped the custom `Fields` type from the `Logger` interface (replaced with `map[string]interface{}`)


## [0.4.0] - 2018-12-11

### Added

- Benchmarks
- [github.com/rs/zerolog](https://github.com/rs/zerolog) adapter
- [github.com/go-kit/kit](https://github.com/go-kit/kit) adapter


## [0.3.0] - 2018-12-11

### Added

- [github.com/goph/emperror](https://github.com/goph/emperror) compatible error handler
- Uber Zap adapter

### Changed

- Removed *Level* suffix from level constants


## [0.2.0] - 2018-12-10

### Added

- [github.com/InVisionApp/go-logger](https://github.com/InVisionApp/go-logger) integration
- `simplelogadapter` to make logger library integration easier
- [github.com/hashicorp/go-hclog](https://github.com/hashicorp/go-hclog) adapter

### Changed

- Renamed `logrusshim` to `logrusadapter`

## 0.1.0 - 2018-12-09

- Initial release


[Unreleased]: https://github.com/goph/logur/compare/v0.5.0...HEAD
[0.5.0]: https://github.com/goph/logur/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/goph/logur/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/goph/logur/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/goph/logur/compare/0.1.0...v0.2.0
