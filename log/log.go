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
func Trace(message string) {
	logger.DefaultLogger.Log(logger.TraceLevel, message)
}
func Tracef(template string, fmtArgs ...interface{}) {
	logger.DefaultLogger.Logf(logger.TraceLevel, template, fmtArgs)
}
func Debug(message string) {
	logger.DefaultLogger.Log(logger.DebugLevel, message)
}
func Debugf(template string, fmtArgs ...interface{}) {
	logger.DefaultLogger.Logf(logger.DebugLevel, template, fmtArgs)
}
func Info(message string) {
	logger.DefaultLogger.Log(logger.InfoLevel, message)
}
func Infof(template string, fmtArgs ...interface{}) {
	logger.DefaultLogger.Logf(logger.InfoLevel, template, fmtArgs)
}
func Warn(message string) {
	logger.DefaultLogger.Log(logger.WarnLevel, message)
}
func Warnf(template string, fmtArgs ...interface{}) {
	logger.DefaultLogger.Logf(logger.WarnLevel, template, fmtArgs)
}
func Error(message string) {
	logger.DefaultLogger.Log(logger.ErrorLevel, message)
}
func Errorf(template string, fmtArgs ...interface{}) {
	logger.DefaultLogger.Logf(logger.ErrorLevel, template, fmtArgs)
}
func Panic(message string) {
	logger.DefaultLogger.Log(logger.PanicLevel, message)
}
func Panicf(template string, fmtArgs ...interface{}) {
	logger.DefaultLogger.Logf(logger.PanicLevel, template, fmtArgs)
}
func Fatal(message string) {
	logger.DefaultLogger.Log(logger.FatalLevel, message)
}
func Fatalf(template string, fmtArgs ...interface{}) {
	logger.DefaultLogger.Logf(logger.FatalLevel, template, fmtArgs)
}
