package zerolog

import (
	"fmt"
	// "errors"
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/xmlking/logger"
	"github.com/xmlking/logger/log"
)

func TestName(t *testing.T) {
	l := NewLogger()

	if l.String() != "zerolog" {
		t.Errorf("error: name expected 'zerolog' actual: %s", l.String())
	}

	t.Logf("testing logger name: %s", l.String())
}

func ExampleWithOut() {
	logger.DefaultLogger = NewLogger(
		logger.WithOutput(os.Stdout),
		logger.WithTimeFormat("ddd"),
		WithProductionMode(),
	)
	log.Info("testing: Info")
	log.Infof("testing: %s", "Infof")
	log.WithFields(map[string]interface{}{
		"sumo":  "demo",
		"human": true,
		"age":   99,
	}).Warn("testing: ", "Warn")
	// Output:
	// {"level":"info","time":"ddd","message":"testing: Info"}
	// {"level":"info","time":"ddd","message":"testing: Infof"}
	// {"level":"warn","age":99,"human":true,"sumo":"demo","time":"ddd","message":"testing: Warn"}
}

func ExampleWithGcp() {
	logger.DefaultLogger = NewLogger(
		logger.WithOutput(os.Stdout),
		WithGCPMode(),
		logger.WithTimeFormat("aaa"),
	)
	log.Info("testing: Info")
	log.Infof("testing: %s", "Infof")
	log.WithFields(map[string]interface{}{
		"sumo":  "demo",
		"human": true,
		"age":   99,
	}).Info("testing: Info with fields")
	logger.DefaultLogger.Init(ReportCaller())
	log.WithError(fmt.Errorf("Error %v: %w", "nested", errors.New("root error message"))).Error("TestWithGCPModeAndWithError")
	// Output:
	//{"severity":"Info","timestamp":"aaa","message":"testing: Info"}
	//{"severity":"Info","timestamp":"aaa","message":"testing: Infof"}
	//{"severity":"Info","age":99,"human":true,"sumo":"demo","timestamp":"aaa","message":"testing: Info with fields"}
	//{"severity":"Error","error":"Error nested: root error message","timestamp":"aaa","logging.googleapis.com/sourceLocation":{"file":"record.go","line":"17","function":"github.com/xmlking/logger/zerolog.(*zerologRecord).Log"},"message":"TestWithGCPModeAndWithError"}
}

func TestSetLevel(t *testing.T) {
	logger.DefaultLogger = NewLogger()

	log.SetLevel(logger.DebugLevel)
	log.Debugf("test show debug: %s", "debug msg")

	log.SetLevel(logger.InfoLevel)
	log.Debugf("test non-show debug: %s", "debug msg")
}

func TestWithReportCaller(t *testing.T) {
	logger.DefaultLogger = NewLogger(ReportCaller())

	log.Infof("testing: %s", "WithReportCaller")
}

func TestWithOutput(t *testing.T) {
	logger.DefaultLogger = NewLogger(logger.WithOutput(os.Stdout))

	log.Infof("testing: %s", "WithOutput")
}

func TestWithGCPMode(t *testing.T) {
	logger.DefaultLogger = NewLogger(WithGCPMode())

	log.Info("testing: TestWithGCPMode Info")
	log.Infof("testing: %s", "TestWithGCPMode Infof")
	log.WithFields(map[string]interface{}{
		"sumo":  "demo",
		"human": true,
		"age":   99,
	}).Info("testing: TestWithGCPMode Infow")

	logger.DefaultLogger.Init(ReportCaller())
	log.WithError(fmt.Errorf("Error %v: %w", "nested", errors.New("root error message"))).Info("TestWithGCPModeAndWithError")
	logger.DefaultLogger.Init(logger.WithTimeFormat(time.RFC3339Nano))
	log.Infof("testing: %s", "TestWithGCPMode")
	// reset `LevelFieldName` to make other tests pass.
	t.Cleanup(func() {
		NewLogger(WithProductionMode())
	})
}

func TestWithDevelopmentMode(t *testing.T) {
	logger.DefaultLogger = NewLogger(WithDevelopmentMode())

	log.Infof("testing: %s", "DevelopmentMode")
}

func TestSubLoggerWithMoreFields(t *testing.T) {
	logger.DefaultLogger = NewLogger(logger.WithFields(map[string]interface{}{
		"component": "AccountHandler",
	}))

	log.WithFields(map[string]interface{}{
		"name":  "demo",
		"human": true,
		"age":   77,
	}).Debug("testing: Infow with extra fields")
	log.Infof("testing: %s", "Infof with default fields")
	// Output:
	//{"level":"info","component":"AccountHandler","age":77,"human":true,"name":"demo","time":"2020-02-23T12:01:10-08:00","message":"testing: Infow with extra fields"}
	//{"level":"info","component":"AccountHandler","time":"2020-02-23T12:01:10-08:00","message":"testing: Infof with default fields"}
}

func TestWithError(t *testing.T) {
	logger.DefaultLogger = NewLogger()
	log.Error("TestWithError")
	log.Errorf("testing: %s", "TestWithError")
	log.WithError(fmt.Errorf("Error %v: %w", "nested", errors.New("root error message"))).Error("TestWithError")
}

func TestWithErrorAndDefaultFields(t *testing.T) {
	logger.DefaultLogger = NewLogger(logger.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}))
	log.Error("TestWithErrorAndDefaultFields")
	log.Errorf("testing: %s", "TestWithErrorAndDefaultFields")
	log.WithError(fmt.Errorf("Error %v: %w", "nested", errors.New("root error message"))).Error("TestWithErrorAndDefaultFields")
}

func TestWithHooks(t *testing.T) {
	simpleHook := zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
		e.Bool("has_level", level != zerolog.NoLevel)
		e.Str("test", "logged")
	})

	logger.DefaultLogger = NewLogger(WithHooks([]zerolog.Hook{simpleHook}))

	log.Infof("testing: %s", "WithHooks")
}
