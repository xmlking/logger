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

func ExampleWithOut() {
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
	log.Infow("testing: Infow", map[string]interface{}{
		"sumo":  "demo",
		"human": true,
		"age":   99,
	})
	logger.DefaultLogger.Init(ReportCaller())
	log.Errorw("TestWithGCPModeAndWithError", fmt.Errorf("Error %v: %w", "nested", errors.New("root error message")))
	// Output:
	//{"severity":"Info","timestamp":"aaa","message":"testing: Info"}
	//{"severity":"Info","timestamp":"aaa","message":"testing: Infof"}
	//{"severity":"Info","age":99,"human":true,"sumo":"demo","timestamp":"aaa","message":"testing: Infow"}
	//{"severity":"Error","error":"Error nested: root error message","timestamp":"aaa","logging.googleapis.com/sourceLocation":{"file":"zerolog.go","line":"170","function":"github.com/xmlking/logger/zerolog.(*zeroLogger).Error"},"message":"TestWithGCPModeAndWithError"}
}

```

### Reference
- https://github.com/arquivei/foundationkit/blob/master/log/stackdriver/stackdriver.go
- https://github.com/yfuruyama/crzerolog