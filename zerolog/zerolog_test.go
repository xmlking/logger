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
		WithTimeFormat("ddd"),
		WithProductionMode(),
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

func TestSetLevel(t *testing.T) {
	logger.DefaultLogger = NewLogger()

	logger.SetLevel(logger.DebugLevel)
	log.Debugf("test show debug: %s", "debug msg")

	logger.SetLevel(logger.InfoLevel)
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

func TestWithDevelopmentMode(t *testing.T) {
	logger.DefaultLogger = NewLogger(WithDevelopmentMode(), WithTimeFormat(time.Kitchen))

	log.Infof("testing: %s", "DevelopmentMode")
}

func TestSubLoggerWithMoreFields(t *testing.T) {
	logger.DefaultLogger = NewLogger(logger.WithFields(map[string]interface{}{
		"component": "AccountHandler",
	}))

	log.Infow("testing: Infow with extra fields", map[string]interface{}{
		"name":  "demo",
		"human": true,
		"age":   77,
	})
	log.Infof("testing: %s", "Infof with default fields")
	// Output:
	//{"level":"info","component":"AccountHandler","age":77,"human":true,"name":"demo","time":"2020-02-23T12:01:10-08:00","message":"testing: Infow with extra fields"}
	//{"level":"info","component":"AccountHandler","time":"2020-02-23T12:01:10-08:00","message":"testing: Infof with default fields"}
}

func TestWithError(t *testing.T) {
	logger.DefaultLogger = NewLogger()
	log.Error("TestWithError")
	log.Errorf("testing: %s", "TestWithError")
	log.Errorw("TestWithError", fmt.Errorf("Error %v: %w", "nested", errors.New("root error message")))
}

func TestWithErrorAndDefaultFields(t *testing.T) {
	logger.DefaultLogger = NewLogger(logger.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}))
	log.Error("TestWithErrorAndDefaultFields")
	log.Errorf("testing: %s", "TestWithErrorAndDefaultFields")
	log.Errorw("TestWithErrorAndDefaultFields", fmt.Errorf("Error %v: %w", "nested", errors.New("root error message")))
}

func TestWithHooks(t *testing.T) {
	simpleHook := zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
		e.Bool("has_level", level != zerolog.NoLevel)
		e.Str("test", "logged")
	})

	logger.DefaultLogger = NewLogger(WithHooks([]zerolog.Hook{simpleHook}))

	log.Infof("testing: %s", "WithHooks")
}
