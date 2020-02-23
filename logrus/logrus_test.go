package logrus

import (
	"fmt"
	"os"
	"testing"

	"github.com/pkg/errors"

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
	log.Infow("testing: Infow", map[string]interface{}{
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

func TestWithDefaultFields(t *testing.T) {
	logger.DefaultLogger = NewLogger(WithJSONFormatter(&logrus.JSONFormatter{}),
		logger.WithFields(map[string]interface{}{
			"component": "AccountHandler",
		}))

	log.Infow("testing: Infow with extra fields", map[string]interface{}{
		"name":  "demo",
		"human": true,
		"age":   77,
	})
	log.Infof("testing: %s", "Infof with default fields")
	// Output:
	//{"age":77,"component":"AccountHandler","human":true,"level":"info","msg":"testing: Infow with extra fields","name":"demo","time":"2020-02-23T11:56:51-08:00"}
	//{"component":"AccountHandler","level":"info","msg":"testing: Infof with default fields","time":"2020-02-23T11:56:51-08:00"}
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
