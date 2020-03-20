// Package `log` provides default logger's public API
package log

import (
	"github.com/xmlking/logger"
)

func WithFields(fields map[string]interface{}) logger.Record {
	return logger.DefaultLogger.WithFields(fields)
}
func WithError(err error) logger.Record {
	return logger.DefaultLogger.WithError(err)
}

// Set DefaultLogger Level
func SetLevel(lvl logger.Level) {
	if err := logger.DefaultLogger.Init(logger.WithLevel(lvl)); err != nil {
		print(err)
	}
}

// Get DefaultLogger name
func String() string {
	return logger.DefaultLogger.String()
}

// Sugar methods
func Trace(args ...interface{}) {
	logger.DefaultLogger.Log(logger.TraceLevel, args...)
}
func Tracef(format string, args ...interface{}) {
	logger.DefaultLogger.Logf(logger.TraceLevel, format, args...)
}
func Debug(args ...interface{}) {
	logger.DefaultLogger.Log(logger.DebugLevel, args...)
}
func Debugf(format string, args ...interface{}) {
	logger.DefaultLogger.Logf(logger.DebugLevel, format, args...)
}
func Info(args ...interface{}) {
	logger.DefaultLogger.Log(logger.InfoLevel, args...)
}
func Infof(format string, args ...interface{}) {
	logger.DefaultLogger.Logf(logger.InfoLevel, format, args...)
}
func Warn(args ...interface{}) {
	logger.DefaultLogger.Log(logger.WarnLevel, args...)
}
func Warnf(format string, args ...interface{}) {
	logger.DefaultLogger.Logf(logger.WarnLevel, format, args...)
}
func Error(args ...interface{}) {
	logger.DefaultLogger.Log(logger.ErrorLevel, args...)
}
func Errorf(format string, args ...interface{}) {
	logger.DefaultLogger.Logf(logger.ErrorLevel, format, args...)
}
func Panic(args ...interface{}) {
	logger.DefaultLogger.Log(logger.PanicLevel, args...)
}
func Panicf(format string, args ...interface{}) {
	logger.DefaultLogger.Logf(logger.PanicLevel, format, args...)
}
func Fatal(args ...interface{}) {
	logger.DefaultLogger.Log(logger.FatalLevel, args...)
}
func Fatalf(format string, args ...interface{}) {
	logger.DefaultLogger.Logf(logger.FatalLevel, format, args...)
}
