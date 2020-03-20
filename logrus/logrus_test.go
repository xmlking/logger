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
	log.WithFields(map[string]interface{}{
		"sumo":  "demo",
		"human": true,
		"age":   99,
	}).Info("testing: Info with fields")
}

func TestJSON(t *testing.T) {
	logger.DefaultLogger = NewLogger(WithJSONFormatter(&logrus.JSONFormatter{}))

	log.Infof("test logf: %s", "name")
}

func TestSetLevel(t *testing.T) {
	logger.DefaultLogger = NewLogger()
	log.SetLevel(logger.DebugLevel)
	log.Debugf("test show debug: %s", "debug msg")

	log.SetLevel(logger.InfoLevel)
	log.Debugf("test non-show debug: %s", "debug msg")

	t.Cleanup(func() {
		log.SetLevel(logger.InfoLevel)
	})
}

func TestWithReportCaller(t *testing.T) {
	logger.DefaultLogger = NewLogger(ReportCaller())

	log.Infof("testing: %s", "WithReportCaller")
}

func TestWithError(t *testing.T) {
	logger.DefaultLogger = NewLogger()
	log.Error("testing: Error")
	log.Errorf("testing: %s", "Errorf")
	log.WithError(fmt.Errorf("Error %v: %w", "nested", errors.New("root error message"))).Error("TestWithError")
}

func TestWithDefaultFields(t *testing.T) {
	logger.DefaultLogger = NewLogger(WithJSONFormatter(&logrus.JSONFormatter{}),
		logger.WithFields(map[string]interface{}{
			"component": "AccountHandler",
		}))

	log.WithFields(map[string]interface{}{
		"name":  "demo",
		"human": true,
		"age":   77,
	}).Info("testing: Info with extra fields")
	log.Infof("testing: %s", "Infof with default fields")
	// Output:
	//{"age":77,"component":"AccountHandler","human":true,"level":"info","msg":"testing: Info with extra fields","name":"demo","time":"2020-03-15T15:07:51-07:00"}
	//{"component":"AccountHandler","level":"info","msg":"testing: Infof with default fields","time":"2020-03-15T15:07:51-07:00"}
}

func TestWithErrorAndDefaultFields(t *testing.T) {
	logger.DefaultLogger = NewLogger(logger.WithFields(map[string]interface{}{
		"name":  "sumo",
		"age":   99,
		"alive": true,
	}))
	log.Errorf("testing: %s", "ErrorfWithDefaultFields")
	log.WithError(fmt.Errorf("Error %v: %w", "nested", errors.New("root error message"))).Error("testing: Error with default fields")
	log.WithError(errors.New("root error message")).Errorf("testing: %s", "Errorf with default fields")
	// Output:
	//time="2020-03-15T15:11:50-07:00" level=error msg="testing: ErrorfWithDefaultFields" age=99 alive=true name=sumo
	//time="2020-03-15T15:11:50-07:00" level=error msg="testing: Error with default fields" age=99 alive=true error="Error nested: root error message" name=sumo
	//time="2020-03-15T15:11:50-07:00" level=error msg="testing: Errorf with default fields" age=99 alive=true error="root error message" name=sumo
}
