# Logger

Logger provides a simple facade over most popular logging systems for __GoLang__, allowing you to log in your application without vendor lock-in.
Think SLF4J for GoLang.


## Usage

Import dependencies. Use latest version.

```go
import (
	github.com/xmlking/logger v0.1.2
    // required: your choice of logger plugins
	github.com/xmlking/logger/zerolog v0.1.2
    //optional: gormlog
	github.com/xmlking/logger/gormlog v0.1.2

)
```

### Default logger

```go
func ExampleLog() {
	logger.DefaultLogger = logger.NewLogger(logger.WithOutput(os.Stdout), logger.WithTimeFormat("ddd"))

	log.Info("test show info: ", "msg ", true, 45.65)
	log.Infof("test show infof: name: %s, age: %d", "sumo", 99)
	log.Infow("test show fields", map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	})
	// Output:
	// {"level":"info","time":"ddd","message":"test show info: msg true 45.65"}
	// {"level":"info","time":"ddd","message":"test show infof: name: sumo, age: 99"}
	// {"level":"info","age":99,"alive":true,"name":"sumo","time":"ddd","message":"test show fields"}
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
	log.Infow("testing: Infow", map[string]interface{}{
		"sumo":  "demo",
		"human": true,
		"age":   99,
	})
	// Output:
	// {"level":"info","time":"ddd","message":"testing: Info"}
	// {"level":"info","time":"ddd","message":"testing: Infof"}
	// {"level":"info","age":99,"human":true,"sumo":"demo","time":"ddd","message":"testing: Infow"}
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
```

#### Release

```bash
make download
git add .
# Start release on develop branch
git flow release start v0.1.0
# on release branch
git-chglog -c .github/chglog/config.yml -o CHANGELOG.md --next-tag v0.1.0
# finish release on release branch
git flow release finish v0.1.0
# on master branch, (gpoat = git push origin --all && git push origin --tags)
gpoat
# add git tags for sub-modules
make release TAG=v0.1.0
```
