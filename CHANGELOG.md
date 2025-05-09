# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], [markdownlint],
and this project adheres to [Semantic Versioning].

## [0.2.2] - 2025-04-21

### Changed in 0.2.2

- Update dependencies

## [0.2.1] - 2024-04-19

### Changed in 0.2.1

- Update dependencies
  - github.com/stretchr/testify v1.9.0
  - google.golang.org/grpc v1.63.2

## [0.2.0] - 2023-12-29

### Changed in 0.2.0

- Renamed module to `github.com/senzing-garage/go-grpcing`
- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - google.golang.org/grpc v1.60.0

## [0.1.3] - 2023-10-15

### Changed in 0.1.3

- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - google.golang.org/grpc v1.58.3

## [0.1.2] - 2023-08-04

### Changed in 0.1.2

- Update dependencies
  - google.golang.org/grpc v1.57.0

## [0.1.1] - 2023-06-16

- Refactored to `template-go`

### Changed in 0.1.1

- Update dependencies
  - github.com/stretchr/testify v1.8.4
  - google.golang.org/grpc v1.56.0

## [0.1.0] - 2023-05-24

### Added to 0.1.0

- Initial functionality
- Support for grpc.Dial(target)
- Support for the following gRPC client `grpc.DialOption`:
  - `grpc.WithTransportCredentials(insecure.NewCredentials()))`
- Support for the following gRPC server `grpc.ServerOption`.
  - (none)

[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[markdownlint]: https://dlaa.me/markdownlint/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html
