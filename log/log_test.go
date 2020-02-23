package log_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/xmlking/logger"
	"github.com/xmlking/logger/log"
)

func TestName(t *testing.T) {
	l := logger.DefaultLogger

	if l.String() != "default" {
		t.Errorf("error: name expected 'default' actual: %s", l.String())
	}

	t.Logf("testing logger name: %s", l.String())
}

func TestSetLevel(t *testing.T) {
	logger.SetLevel(logger.DebugLevel)
	log.Debugf("test show debug: %s", "debug msg")

	logger.SetLevel(logger.InfoLevel)
	log.Debugf("test non-show debug: %s", "debug msg")
}

func TestSubLogger(t *testing.T) {
	log.Infof("Logging with Default Options: %v", logger.DefaultLogger.Options())
	subLogger := logger.NewLogger(logger.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}))
	subLogger.Log(logger.WarnLevel, "Logging with subLogger Options: %v", []interface{}{subLogger.Options()}, nil)
	log.Warnf("Logging with Default Options: %v", logger.DefaultLogger.Options())
}

func TestWithFields(t *testing.T) {
	logger.DefaultLogger = logger.NewLogger(logger.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}))
	log.Info("test with default fields")
	log.Infow("test with extra fields", map[string]interface{}{"weight": 3.14159265359, "name": "demo"})
	log.Infow("test with duplicate fields", map[string]interface{}{"name": "sumo1"})
}

func TestWithError(t *testing.T) {
	log.Error("TestWithError")
	log.Errorf("testing: %s", "TestWithError")
	log.Errorw("TestWithError", fmt.Errorf("Error %v: %w", "nested", errors.New("root error message")))
}

func TestWithErrorAndDefaultFields(t *testing.T) {
	logger.DefaultLogger = logger.NewLogger(logger.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}))
	log.Error("TestWithErrorAndDefaultFields")
	log.Errorf("testing: %s", "TestWithErrorAndDefaultFields")
	log.Errorw("TestWithErrorAndDefaultFields", fmt.Errorf("Error %v: %w", "nested", errors.New("root error message")))
}

func ExampleLog() {
	logger.DefaultLogger = logger.NewLogger(logger.WithOutput(os.Stdout))
	log.Info("test show info: ", "msg ", true, 45.65)
	log.Infof("test show infof: name: %s, age: %d", "sumo", 99)
	log.Infow("test show fields", map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	})
	// Output:
	//{"level":"info","message":"test show info: msg true 45.65"}
	//{"level":"info","message":"test show infof: name: sumo, age: 99"}
	//{"age":99,"alive":true,"level":"info","message":"test show fields","name":"sumo"}
}
