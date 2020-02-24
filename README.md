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
	log.SetGlobalLogger(
        logger.NewLogger(logger.WithOutput(os.Stdout))
    )
	log.Info("test show info: ", "msg ", true, 45.65)
	log.Infof("test show infof: name: %s, age: %d", "sumo", 99)
	log.Infow("test show fields", map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	})
	// Output:
	// {"message":"test show info: msg true 45.65"}
	// {"message":"test show infof: name: sumo, age: 99"}
	// {"age":99,"alive":true,"message":"test show fields","name":"sumo"}
}
```

### Zerolog logger

```go
func ExampleWithOut() {
	log.SetGlobalLogger(
        zerolog_plugin.NewLogger(
            logger.WithOutput(os.Stdout),
            zerolog_plugin.WithTimeFormat("ddd"),
            zerolog_plugin.WithProductionMode(),
	    )
    )

	log.Info("testing: Info")
	log.Infof("testing: %s", "Infof")
	log.Infow("testing: Infow", logger.Fields{
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


### For Contributor

#### Prerequisites 

```bash
brew install hub
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
