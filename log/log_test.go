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
	log.SetLevel(logger.DebugLevel)
	log.Debugf("test show debug: %s", "debug msg")

	log.SetLevel(logger.InfoLevel)
	log.Debugf("test non-show debug: %s", "debug msg")
}

func TestSubLogger(t *testing.T) {
	log.Infof("Logging with Default Options: %v", logger.DefaultLogger.Options())
	subLogger := logger.NewLogger(logger.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}))
	subLogger.Logf(logger.WarnLevel, "Logging with subLogger Options: %v", []interface{}{subLogger.Options()}, nil)
	log.Warnf("Logging with Default Options: %v", logger.DefaultLogger.Options())
}

func TestWithFields(t *testing.T) {
	logger.DefaultLogger = logger.NewLogger(logger.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}))
	log.Info("test with default fields")
	log.WithFields(map[string]interface{}{"weight": 3.14159265359, "name": "demo"}).Info("test with extra fields")
	log.WithFields(map[string]interface{}{"name": "sumo1"}).Info("test with duplicate fields")
}

func TestWithError(t *testing.T) {
	log.Error("TestWithError")
	log.Errorf("testing: %s", "TestWithError")
	log.WithError(errors.New("error message")).Errorf("TestWithError: %s", "root error message")
}

func TestWithErrorAndDefaultFields(t *testing.T) {
	logger.DefaultLogger = logger.NewLogger(logger.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}))
	log.Error("TestWithErrorAndDefaultFields")
	log.Errorf("testing: %s", "TestWithErrorAndDefaultFields")
	log.WithError(fmt.Errorf("Error %v: %w", "nested", errors.New("root error message"))).Error("TestWithErrorAndDefaultFields")
}

func ExampleLog() {
	logger.DefaultLogger = logger.NewLogger(logger.WithOutput(os.Stdout), logger.WithTimeFormat("ddd"))
	log.Info("test show info")
	log.Infof("test show infof: name: %s, age: %d", "sumo", 99)
	log.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}).Info("test show fields")
	// Output:
	//{"level":"info","message":"test show info: msg true 45.65","time":"ddd"}
	//{"level":"info","message":"test show infof: name: sumo, age: 99","time":"ddd"}
	//{"age":99,"alive":true,"level":"info","message":"test show fields","name":"sumo","time":"ddd"}
}
