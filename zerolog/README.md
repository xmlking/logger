# Zerolog

[Zerolog](https://github.com/rs/zerolog) logger implementation

## Usage

### Production

```go
import (
	"github.com/xmlking/logger"
	"github.com/xmlking/logger/log"
  "github.com/xmlking/logger/zerolog"
)

func ExampleLog() {
	logger.DefaultLogger = zerolog.NewLogger(
		logger.WithOutput(os.Stdout),
		logger.WithTimeFormat("ddd"),
		zerolog.WithProductionMode(),
	)
	log.Info("test show info: ", "msg ", true, 45.65)
	log.Infof("test show infof: name: %s, age: %d", "sumo", 99)
	log.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}).Info("test show fields")
	// Output:
	// {"level":"info","time":"ddd","message":"test show info: msg true 45.65"}
	// {"level":"info","time":"ddd","message":"test show infof: name: sumo, age: 99"}
	// {"level":"info","age":99,"alive":true,"name":"sumo","time":"ddd","message":"test show fields"}
}
```

### GCP

set log level name to `Severity` for __GCP__ `StackDriver`

```go
import (
	"os"
	"testing"

	"github.com/xmlking/logger"
	"github.com/xmlking/logger/log"
	"github.com/xmlking/logger/zerolog"
)

func ExampleWithGcp() {
	logger.DefaultLogger = zerolog.NewLogger(logger.WithOutput(os.Stdout), logger.WithTimeFormat("aaa"), zerolog.WithGCPMode())

	log.Info("testing: Info")
	log.Infof("testing: %s", "Infof")
	log.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}).Info("testing: with fields")
	logger.DefaultLogger.Init(ReportCaller())
	log.WithError(fmt.Errorf("Error %v: %w", "nested", errors.New("root error message"))).Error("TestWithGCPModeAndWithError")
	// Output:
	//{"severity":"Info","timestamp":"aaa","message":"testing: Info"}
	//{"severity":"Info","timestamp":"aaa","message":"testing: Infof"}
	//{"severity":"Info","age":99,"human":true,"sumo":"demo","timestamp":"aaa","message":"testing: with fields"}
	//{"severity":"Error","error":"Error nested: root error message","timestamp":"aaa","logging.googleapis.com/sourceLocation":{"file":"zerolog.go","line":"170","function":"github.com/xmlking/logger/zerolog.(*zeroLogger).Error"},"message":"TestWithGCPModeAndWithError"}
}

```

### Reference
- https://github.com/arquivei/foundationkit/blob/master/log/stackdriver/stackdriver.go
- https://github.com/yfuruyama/crzerolog