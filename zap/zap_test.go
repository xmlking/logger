package zap

import (
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

	log.SetGlobalLevel(logger.DebugLevel)
	log.Debugf("test show debug: %s", "debug msg")

	log.SetGlobalLevel(logger.InfoLevel)
	log.Debugf("test non-show debug: %s", "debug msg")
}

func TestError(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger = l

	err2 := errors.Wrap(errors.New("error message"), "from error")
	log.Error("test Error")
	log.Errorw("test Errorw", err2)
}

func TestFields(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger = l

	log.Infow("testing: Fields", logger.Fields{
		"sumo":  "demo",
		"human": true,
		"age":   99,
	})
}

func TestSubLoggerWithFields(t *testing.T) {
	l, err := NewLogger(logger.WithFields(logger.Fields{
		"category": "test",
		"alive":    true,
	}))
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger = l

	log.Infow("testing: WithFields", logger.Fields{
		"name":  "demo",
		"human": true,
		"age":   77,
	})
	// Output:
	// {"level":"info","ts":1582075193.56922,"caller":"zap/zap.go:87","msg":"testing: WithFields","category":"test","alive":true,"name":"demo","human":true,"age":77}
}

func TestWithNamespace(t *testing.T) {
	l, err := NewLogger(WithNamespace("micro"))
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger = l

	log.Infow("testing: WithFields", logger.Fields{
		"name":  "demo",
		"human": true,
		"age":   77,
	})
	// Output:
	// {"level":"info","ts":1582075193.569254,"caller":"zap/zap.go:87","msg":"testing: WithFields","micro":{"name":"demo","human":true,"age":77}}
}
