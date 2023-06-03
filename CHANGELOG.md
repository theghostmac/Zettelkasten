# Changelog

All notable changes to the Zettelkasten project will be documented in this file.

## [Unreleased]

### Added
- User authentication feature
- Ability to create and edit zettels
- Search functionality
- Tagging system for zettels
- Markdown support for zettel content

### Changed
- Improved user interface design
- Refactored codebase for better performance
- Enhanced error handling and validation

### Fixed
- Resolved issue with zettel deletion
- Fixed search result pagination bug
- Addressed security vulnerabilities

## [1.0.0] - 3rd June, 2023

### Added
- Optionally access Zettelkasten using any address of your choice by adding the `-addr` flag:
```go
go run ./cmd/http -addr=":9999"
```
To learn more, run `go run ./cmd/http -help`.
- Initial release of Zettelkasten
- Basic CRUD functionality for zettels
- User registration and login
- Simple layout and navigation

