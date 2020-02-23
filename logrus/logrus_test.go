package logrus

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/xmlking/logger"
	"github.com/xmlking/logger/log"
)

func TestName(t *testing.T) {
	l := NewLogger()

	if l.String() != "logrus" {
		t.Errorf("error: name expected 'logrus' actual: %s", l.String())
	}

	t.Logf("testing logger name: %s", l.String())
}

func TestWithFields(t *testing.T) {
	logger.DefaultLogger = NewLogger(logger.WithOutput(os.Stdout))

	log.Info("testing: Info")
	log.Infof("testing: %s", "Infof")
	log.Infow("testing: Infow", logger.Fields{
		"sumo":  "demo",
		"human": true,
		"age":   99,
	})
}

func TestJSON(t *testing.T) {
	logger.DefaultLogger = NewLogger(WithJSONFormatter(&logrus.JSONFormatter{}))

	log.Infof("test logf: %s", "name")
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
