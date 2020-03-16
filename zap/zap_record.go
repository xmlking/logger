package zap

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/xmlking/logger"
)

// zapRecord represents logger with Fields or Error
type zapRecord struct {
	*zap.Logger
	fields map[string]interface{}
}

func (r *zapRecord) Log(level logger.Level, args ...interface{}) {
	lvl := loggerToZapLevel(level)
	msg := fmt.Sprint(args...)
	if ce := r.Logger.Check(lvl, msg); ce != nil {
		zFields := make([]zap.Field, 0, len(r.fields))
		for k, v := range r.fields {
			zFields = append(zFields, zap.Any(k, v))
		}
		ce.Write(zFields...)
	}
}

func (r *zapRecord) Logf(level logger.Level, format string, args ...interface{}) {
	lvl := loggerToZapLevel(level)
	msg := fmt.Sprintf(format, args...)
	if ce := r.Logger.Check(lvl, msg); ce != nil {
		zFields := make([]zap.Field, 0, len(r.fields))
		for k, v := range r.fields {
			zFields = append(zFields, zap.Any(k, v))
		}
		ce.Write(zFields...)
	}
}

func (r *zapRecord) Trace(args ...interface{}) {
	r.Log(logger.TraceLevel, args...)
}
func (r *zapRecord) Tracef(format string, args ...interface{}) {
	r.Logf(logger.TraceLevel, format, args...)
}
func (r *zapRecord) Debug(args ...interface{}) {
	r.Log(logger.DebugLevel, args...)
}
func (r *zapRecord) Debugf(format string, args ...interface{}) {
	r.Logf(logger.DebugLevel, format, args...)
}
func (r *zapRecord) Info(args ...interface{}) {
	r.Log(logger.InfoLevel, args...)
}
func (r *zapRecord) Infof(format string, args ...interface{}) {
	r.Logf(logger.InfoLevel, format, args...)
}
func (r *zapRecord) Warn(args ...interface{}) {
	r.Log(logger.WarnLevel, args...)
}
func (r *zapRecord) Warnf(format string, args ...interface{}) {
	r.Logf(logger.WarnLevel, format, args...)
}
func (r *zapRecord) Error(args ...interface{}) {
	r.Log(logger.ErrorLevel, args...)
}
func (r *zapRecord) Errorf(format string, args ...interface{}) {
	r.Logf(logger.ErrorLevel, format, args...)
}
func (r *zapRecord) Panic(args ...interface{}) {
	r.Log(logger.PanicLevel, args...)
}
func (r *zapRecord) Panicf(format string, args ...interface{}) {
	r.Logf(logger.PanicLevel, format, args...)
}
func (r *zapRecord) Fatal(args ...interface{}) {
	r.Log(logger.FatalLevel, args...)
}
func (r *zapRecord) Fatalf(format string, args ...interface{}) {
	r.Logf(logger.FatalLevel, format, args...)
}
