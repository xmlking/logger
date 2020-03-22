# Logger

Logger provides a simple facade over most popular logging systems for __GoLang__, allowing you to log in your application without vendor lock-in.
Think SLF4J for GoLang.

[![Total alerts](https://img.shields.io/lgtm/alerts/g/xmlking/logger.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/xmlking/logger/alerts/)
[![Language grade: Go](https://img.shields.io/lgtm/grade/go/g/xmlking/logger.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/xmlking/logger/context:go)

## Usage

Import dependencies. Use latest version.

```go
import (
	github.com/xmlking/logger v0.1.4
    // required: your choice of logger plugins
	github.com/xmlking/logger/zerolog v0.1.4
    //optional: gormlog
	github.com/xmlking/logger/gormlog v0.1.4

)
```

### Default logger

```go
func ExampleLog() {
	logger.DefaultLogger = logger.NewLogger(logger.WithOutput(os.Stdout), logger.WithTimeFormat("ddd"))

	log.Info("test show info: ", "msg ", true, 45.65)
	log.Infof("test show infof: name: %s, age: %d", "sumo", 99)
	log.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}).Info("test show fields")
	log.WithError(errors.New("error message")).Errorf("Testing: %s", "WithError")
	// Output:
	//{"level":"info","message":"test show info: msg true 45.65","time":"ddd"}
	//{"level":"info","message":"test show infof: name: sumo, age: 99","time":"ddd"}
	//{"age":99,"alive":true,"level":"info","message":"test show fields","name":"sumo","time":"ddd"}
	//{"error":"error message","level":"error","message":"Testing: WithError","time":"ddd"}
}
```

### Zerolog logger

```go
func ExampleWithFields() {
	logger.DefaultLogger = zerolog.NewLogger(
		logger.WithOutput(os.Stdout),
		logger.WithTimeFormat("ddd"),
		zerolog.WithProductionMode(),
	)

	log.Info("testing: Info")
	log.Infof("testing: %s", "Infof")
	log.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}).Info("testing: with fields")
	// Output:
	// {"level":"info","time":"ddd","message":"testing: Info"}
	// {"level":"info","time":"ddd","message":"testing: Infof"}
	// {"level":"info","age":99,"human":true,"sumo":"demo","time":"ddd","message":"testing: with fields"}
}
```


### For Contributors

#### Prerequisites

```bash
brew install hub
# goup checks if there are any updates for imports in your module.
GO111MODULE=on go get github.com/rvflash/goup
# for static check/linter
GO111MODULE=off go get github.com/golangci/golangci-lint/cmd/golangci-lint
```

#### Test

```bash
make download
make test
make bench
# Benchmark specific test
cd zerolog
go test -run=^$ -bench=^BenchmarkInfoLog ./...
go test -run=^$ -bench=^BenchmarkInfoLog -benchtime 15s -count 2 -cpu 1,2,4 ./...
# Results:
# zerolog: 326443       3866 ns/op
# wrapper: 324487       4280 ns/op
```

#### Release

```bash
make download
git add .
# Start release on develop branch
git flow release start v0.1.0
# on release branch
git-chglog -c .github/chglog/config.yml -o CHANGELOG.md --next-tag v0.1.0
# update `github.com/xmlking/logger` version in each `go.mod` file.
# commit all changes.
# finish release on release branch
git flow release finish
# on master branch, (gpoat = git push origin --all && git push origin --tags)
gpoat
# add git tags for sub-modules
make release TAG=v0.1.0
```

### Reference
- [How Zap Package is Optimized](https://medium.com/a-journey-with-go/go-how-zap-package-is-optimized-dbf72ef48f2d)
- [zapdriver for Stackdriver logging](https://github.com/blendle/zapdriver)