package zerolog

import (
	"fmt"

	"github.com/rs/zerolog"

	"github.com/xmlking/logger"
)

type zerologRecord struct {
	*zerolog.Logger
	fields map[string]interface{}
	err error
}

func (r *zerologRecord) Log(level logger.Level, args ...interface{}) {
	r.Logger.WithLevel(loggerToZerologLevel(level)).Fields(r.fields).Err(r.err).Msg(fmt.Sprint(args...))
	// Should we use object pool to avoid allocation?
}

func (r *zerologRecord) Logf(level logger.Level, format string, args ...interface{}) {
	r.Logger.WithLevel(loggerToZerologLevel(level)).Fields(r.fields).Err(r.err).Msgf(format, args...)
}

func (r *zerologRecord) Trace(args ...interface{}) {
	r.Log(logger.TraceLevel, args...)
}
func (r *zerologRecord) Tracef(format string, args ...interface{}) {
	r.Logf(logger.TraceLevel, format, args...)
}
func (r *zerologRecord) Debug(args ...interface{}) {
	r.Log(logger.DebugLevel, args...)
}
func (r *zerologRecord) Debugf(format string, args ...interface{}) {
	r.Logf(logger.DebugLevel, format, args...)
}
func (r *zerologRecord) Info(args ...interface{}) {
	r.Log(logger.InfoLevel, args...)
}
func (r *zerologRecord) Infof(format string, args ...interface{}) {
	r.Logf(logger.InfoLevel, format, args...)
}
func (r *zerologRecord) Warn(args ...interface{}) {
	r.Log(logger.WarnLevel, args...)
}
func (r *zerologRecord) Warnf(format string, args ...interface{}) {
	r.Logf(logger.WarnLevel, format, args...)
}
func (r *zerologRecord) Error(args ...interface{}) {
	r.Log(logger.ErrorLevel, args...)
}
func (r *zerologRecord) Errorf(format string, args ...interface{}) {
	r.Logf(logger.ErrorLevel, format, args...)
}
func (r *zerologRecord) Panic(args ...interface{}) {
	r.Log(logger.PanicLevel, args...)
}
func (r *zerologRecord) Panicf(format string, args ...interface{}) {
	r.Logf(logger.PanicLevel, format, args...)
}
func (r *zerologRecord) Fatal(args ...interface{}) {
	r.Log(logger.FatalLevel, args...)
}
func (r *zerologRecord) Fatalf(format string, args ...interface{}) {
	r.Logf(logger.FatalLevel, format, args...)
}
