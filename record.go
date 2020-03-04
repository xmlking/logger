package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Record interface {
	Log(level Level, message string)
	Logf(level Level, template string, fmtArgs ...interface{})
	// Sugar methods
	Trace(message string)
	Tracef(template string, fmtArgs ...interface{})
	Debug(message string)
	Debugf(template string, fmtArgs ...interface{})
	Info(message string)
	Infof(template string, fmtArgs ...interface{})
	Warn(message string)
	Warnf(template string, fmtArgs ...interface{})
	Error(message string)
	Errorf(template string, fmtArgs ...interface{})
	Panic(message string)
	Panicf(template string, fmtArgs ...interface{})
	Fatal(message string)
	Fatalf(template string, fmtArgs ...interface{})
}

/**
 * Default record implementation
 */
type defaultRecord struct {
	opts   Options
	level  Level
	fields map[string]interface{}
	err    error
}

func (l *defaultRecord) Log(level Level, message string) {
	if !l.level.Enabled(level) {
		return
	}
	l.fields["time"] = time.Now().Format(l.opts.TimeFormat)
	l.fields["level"] = level.String()
	l.fields["message"] = message

	enc := json.NewEncoder(l.opts.Out)

	if err := enc.Encode(l.fields); err != nil {
		log.Fatal(err)
	}
}

func (l *defaultRecord) Logf(level Level, template string, fmtArgs ...interface{}) {
	if !l.level.Enabled(level) {
		return
	}
	// Format with Sprint, Sprintf, or neither.
	msg := template
	if msg == "" && len(fmtArgs) > 0 {
		msg = fmt.Sprint(fmtArgs...)
	} else if msg != "" && len(fmtArgs) > 0 {
		msg = fmt.Sprintf(template, fmtArgs...)
	}

	l.fields["time"] = time.Now().Format(l.opts.TimeFormat)
	l.fields["level"] = level.String()
	l.fields["message"] = msg

	enc := json.NewEncoder(l.opts.Out)

	if err := enc.Encode(l.fields); err != nil {
		log.Fatal(err)
	}
}

// Sugar methods
func (l *defaultRecord) Trace(message string) {
	l.Log(TraceLevel, message)
}
func (l *defaultRecord) Tracef(template string, fmtArgs ...interface{}) {
	l.Logf(TraceLevel, template, fmtArgs)
}
func (l *defaultRecord) Debug(message string) {
	l.Log(DebugLevel, message)
}
func (l *defaultRecord) Debugf(template string, fmtArgs ...interface{}) {
	l.Logf(DebugLevel, template, fmtArgs)
}
func (l *defaultRecord) Info(message string) {
	l.Log(InfoLevel, message)
}
func (l *defaultRecord) Infof(template string, fmtArgs ...interface{}) {
	l.Logf(InfoLevel, template, fmtArgs)
}
func (l *defaultRecord) Warn(message string) {
	l.Log(WarnLevel, message)
}
func (l *defaultRecord) Warnf(template string, fmtArgs ...interface{}) {
	l.Logf(WarnLevel, template, fmtArgs)
}
func (l *defaultRecord) Error(message string) {
	l.Log(ErrorLevel, message)
}
func (l *defaultRecord) Errorf(template string, fmtArgs ...interface{}) {
	l.Logf(ErrorLevel, template, fmtArgs)
}
func (l *defaultRecord) Panic(message string) {
	l.Log(PanicLevel, message)
}
func (l *defaultRecord) Panicf(template string, fmtArgs ...interface{}) {
	l.Logf(PanicLevel, template, fmtArgs)
}
func (l *defaultRecord) Fatal(message string) {
	l.Log(FatalLevel, message)
}
func (l *defaultRecord) Fatalf(template string, fmtArgs ...interface{}) {
	l.Logf(FatalLevel, template, fmtArgs...)
}
