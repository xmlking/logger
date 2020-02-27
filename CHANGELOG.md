# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<a name="unreleased"></a>
## [Unreleased]


<a name="v0.1.3"></a>
## [v0.1.3] - 2020-02-26
### Chore
- **golang:** updated golang to 1.4

### Docs
- **readme:** updated example code
- **readme:** updated Usage

### Perf
- **benchmark:** adding native vs. logger wrapper Benchmark tests for Zerolog
- **results:** adding test results

### Test
- **zerolog:** using new t.Cleanup() from GoLang 1.14


<a name="zerolog/v0.1.2"></a>
## [zerolog/v0.1.2] - 2020-02-23

<a name="logrus/v0.1.2"></a>
## [logrus/v0.1.2] - 2020-02-23

<a name="zap/v0.1.2"></a>
## [zap/v0.1.2] - 2020-02-23

<a name="gormlog/v0.1.2"></a>
## [gormlog/v0.1.2] - 2020-02-23

<a name="v0.1.2"></a>
## [v0.1.2] - 2020-02-23
### Chore
- **changelog:** update changelog

### Docs
- **readme:** update Release docs

### Fix
- **release:** adding proper tags for nested modules


<a name="zerolog/v0.1.1"></a>
## [zerolog/v0.1.1] - 2020-02-23

<a name="logrus/v0.1.1"></a>
## [logrus/v0.1.1] - 2020-02-23

<a name="zap/v0.1.1"></a>
## [zap/v0.1.1] - 2020-02-23

<a name="gormlog/v0.1.1"></a>
## [gormlog/v0.1.1] - 2020-02-23

<a name="v0.1.1"></a>
## [v0.1.1] - 2020-02-23
### Chore
- **changelog:** update changelog

### Docs
- **readme:** updated Contributor section

### Feat
- **makefile:** adding release task

### Improvement
- **gormlog:** now using defaultLogger for testing
- **logger:** adding WithTimeFormat(...) option for root options and defaultLogger


<a name="v0.1.0"></a>
## v0.1.0 - 2020-02-23
### Build
- **makefile:** polish makefile

### Chore
- **changelog:** adding changelog

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


[Unreleased]: https://github.com/xmlking/micro-starter-kit/compare/v0.1.3...HEAD
[v0.1.3]: https://github.com/xmlking/micro-starter-kit/compare/zerolog/v0.1.2...v0.1.3
[zerolog/v0.1.2]: https://github.com/xmlking/micro-starter-kit/compare/logrus/v0.1.2...zerolog/v0.1.2
[logrus/v0.1.2]: https://github.com/xmlking/micro-starter-kit/compare/zap/v0.1.2...logrus/v0.1.2
[zap/v0.1.2]: https://github.com/xmlking/micro-starter-kit/compare/gormlog/v0.1.2...zap/v0.1.2
[gormlog/v0.1.2]: https://github.com/xmlking/micro-starter-kit/compare/v0.1.2...gormlog/v0.1.2
[v0.1.2]: https://github.com/xmlking/micro-starter-kit/compare/zerolog/v0.1.1...v0.1.2
[zerolog/v0.1.1]: https://github.com/xmlking/micro-starter-kit/compare/logrus/v0.1.1...zerolog/v0.1.1
[logrus/v0.1.1]: https://github.com/xmlking/micro-starter-kit/compare/zap/v0.1.1...logrus/v0.1.1
[zap/v0.1.1]: https://github.com/xmlking/micro-starter-kit/compare/gormlog/v0.1.1...zap/v0.1.1
[gormlog/v0.1.1]: https://github.com/xmlking/micro-starter-kit/compare/v0.1.1...gormlog/v0.1.1
[v0.1.1]: https://github.com/xmlking/micro-starter-kit/compare/v0.1.0...v0.1.1
