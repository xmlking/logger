package zap

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"

	"github.com/xmlking/logger"
	"github.com/xmlking/logger/log"
)

func TestName(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}

	if l.String() != "zap" {
		t.Errorf("name is error %s", l.String())
	}

	t.Logf("test logger name: %s", l.String())
}

func TestLogf(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger = l
	log.Infof("test logf: %s", "name")
}

func TestSetLevel(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger = l

	log.SetLevel(logger.DebugLevel)
	log.Debugf("test show debug: %s", "debug msg")

	log.SetLevel(logger.InfoLevel)
	log.Debugf("test non-show debug: %s", "debug msg")
}

func TestError(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger = l

	err2 := errors.Wrap(errors.New("error message"), "from error")
	err3 := fmt.Errorf("Error %v: %w", "nested", errors.New("root error message"))

	log.Error("testing: Error")
	log.Errorf("testing: %s", "Errorf")
	log.WithError(err2).Error("testing: WithError")
	log.WithError(err3).Error("testing: WithError")
}

func TestErrorWithDefaultFields(t *testing.T) {
	l, err := NewLogger(logger.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}))
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger = l

	log.Error("testing: Error with default fields")
	log.Errorf("testing: %s", "Errorf with default fields")
	log.WithError(fmt.Errorf("Error %v: %w", "nested", errors.New("root error message"))).Error("testing: WithError with default fields")
}

func TestFields(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger = l

	log.WithFields(map[string]interface{}{
		"sumo":  "demo",
		"human": true,
		"age":   99,
	}).Info("testing: Fields")
}

func TestSubLoggerWithFields(t *testing.T) {
	l, err := NewLogger(logger.WithFields(map[string]interface{}{
		"category": "test",
		"alive":    true,
	}))
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger = l

	log.WithFields(map[string]interface{}{
		"name":  "demo",
		"human": true,
		"age":   77,
	}).Info("testing: SubLogger WithFields")
	log.Warnf("testing: SubLogger %s", "Warn")
	// Output:
	// {"level":"info","ts":1582075193.56922,"caller":"zap/zap.go:87","msg":"testing: WithFields","category":"test","alive":true,"name":"demo","human":true,"age":77}
}

func TestWithNamespace(t *testing.T) {
	l, err := NewLogger(WithNamespace("micro"))
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger = l

	log.WithFields(map[string]interface{}{
		"name":  "demo",
		"human": true,
		"age":   77,
	}).Info("testing: WithFields")
	// Output:
	// {"level":"info","ts":1582075193.569254,"caller":"zap/zap.go:87","msg":"testing: WithFields","micro":{"name":"demo","human":true,"age":77}}
}
