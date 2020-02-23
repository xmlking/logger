# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<a name="unreleased"></a>
## [Unreleased]


<a name="v0.1.0"></a>
## v0.1.0 - 2020-02-23
### Build
- **makefile:** polish makefile

### Docs
- **logger:** updated readme

### Feat
- **gcp:** support GCP Stackdriver logging
- **gcp:** support GCP Stackdriver logging
- **gorm:** adding gormlog support

### Fix
- **logrus:** properly adding default fields

### Improvement
- **logger:** removed SetLevel(Level) , Level() from interface
- **logger:** polish

### Perf
- **logger:** less memory allocation

### Refactor
- **logger:** setGlobalLevel -> SetLevel and moved to logger package
- **logger:** adding Options() to interface

### Style
- **format:** fix code format


[Unreleased]: https://github.com/xmlking/micro-starter-kit/compare/v0.1.0...HEAD
