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

func TestOptions(t *testing.T) {
	log.Infof("Default Options: %v", logger.DefaultLogger.Options())

	subLogger := logger.NewLogger(logger.WithFields(logger.Fields{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}))
	log.Infof("Modified Options: %v", subLogger.Options())
	subLogger.Log(logger.WarnLevel, "Logging with subLogger: Default Options: %v", []interface{}{logger.DefaultLogger.Options()}, nil)
}

func TestWithFields(t *testing.T) {
	logger.DefaultLogger = logger.NewLogger(logger.WithFields(logger.Fields{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}))
	log.Info("test with fields")
	log.Infow("test with fields", map[string]interface{}{"weight": 3.14159265359, "name": "demo"})
	log.Infow("testing replace", logger.Fields{"name": "sumo1"})
}

func TestWithError(t *testing.T) {
	logger.DefaultLogger = logger.NewLogger(logger.WithFields(logger.Fields{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}))
	log.Error("test with fields")
	log.Errorw("test with fields", fmt.Errorf("Error %v: %w", "nested", errors.New("root error message")))
	log.Infof("testing: %s", "TestWithError")
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
