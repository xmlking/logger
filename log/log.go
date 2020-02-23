package log

import (
	"github.com/xmlking/logger"
)

func Trace(args ...interface{}) {
	logger.Log(logger.TraceLevel, "", args, nil)
}
func Tracef(template string, args ...interface{}) {
	logger.Log(logger.TraceLevel, template, args, nil)
}
func Tracew(msg string, fields logger.Fields) {
	logger.Log(logger.TraceLevel, msg, nil, fields)
}

func Debug(args ...interface{}) {
	logger.Log(logger.DebugLevel, "", args, nil)
}
func Debugf(template string, args ...interface{}) {
	logger.Log(logger.DebugLevel, template, args, nil)
}
func Debugw(msg string, fields logger.Fields) {
	logger.Log(logger.DebugLevel, msg, nil, fields)
}

func Info(args ...interface{}) {
	logger.Log(logger.InfoLevel, "", args, nil)
}
func Infof(template string, args ...interface{}) {
	logger.Log(logger.InfoLevel, template, args, nil)
}
func Infow(msg string, fields logger.Fields) {
	logger.Log(logger.InfoLevel, msg, nil, fields)
}

func Warn(args ...interface{}) {
	logger.Log(logger.WarnLevel, "", args, nil)
}
func Warnf(template string, args ...interface{}) {
	logger.Log(logger.WarnLevel, template, args, nil)
}
func Warnw(msg string, fields logger.Fields) {
	logger.Log(logger.WarnLevel, msg, nil, fields)
}

func Error(args ...interface{}) {
	logger.Log(logger.ErrorLevel, "", args, nil)
}
func Errorf(template string, args ...interface{}) {
	logger.Log(logger.ErrorLevel, template, args, nil)
}
func Errorw(msg string, err error) {
	logger.Error(logger.ErrorLevel, msg, nil, err)
}

func Panic(args ...interface{}) {
	logger.Log(logger.PanicLevel, "", args, nil)
}
func Panicf(template string, args ...interface{}) {
	logger.Log(logger.PanicLevel, template, args, nil)
}
func Panicw(msg string, fields logger.Fields) {
	logger.Log(logger.PanicLevel, msg, nil, fields)
}

func Fatal(args ...interface{}) {
	logger.Log(logger.FatalLevel, "", args, nil)
}
func Fatalf(template string, args ...interface{}) {
	logger.Log(logger.FatalLevel, template, args, nil)
}
func Fatalw(msg string, fields logger.Fields) {
	logger.Log(logger.FatalLevel, msg, nil, fields)
}
