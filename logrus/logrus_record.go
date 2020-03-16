package logrus

import (
	"github.com/sirupsen/logrus"
	"github.com/xmlking/logger"
)

type logrusRecord struct {
	*logrus.Entry
}

func (r *logrusRecord) Log(level logger.Level, args ...interface{}) {
	r.Entry.Log(loggerToLogrusLevel(level), args...)
}

func (r *logrusRecord) Logf(level logger.Level, format string, args ...interface{}) {
	r.Entry.Logf(loggerToLogrusLevel(level), format, args...)
}