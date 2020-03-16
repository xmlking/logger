package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// An Record is the final or intermediate logging entry. It contains all
// the fields passed with WithField{,s}. It's finally logged when Trace, Debug,
// Info, Warn, Error, Fatal or Panic is called on it. These objects can be
// reused and passed around as much as you wish to avoid field duplication.
type Record interface {
	Log(level Level, args ...interface{})
	Logf(level Level, format string, args ...interface{})
	// Sugar methods
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
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

func (l *defaultRecord) Log(level Level, args ...interface{}) {
	if !l.level.Enabled(level) {
		return
	}
	l.fields["time"] = time.Now().Format(l.opts.TimeFormat)
	l.fields["level"] = level.String()
	l.fields["message"] = fmt.Sprint(args...)

	enc := json.NewEncoder(l.opts.Out)

	if err := enc.Encode(l.fields); err != nil {
		log.Fatal(err)
	}
}

func (l *defaultRecord) Logf(level Level, format string, args ...interface{}) {
	if !l.level.Enabled(level) {
		return
	}

	l.fields["time"] = time.Now().Format(l.opts.TimeFormat)
	l.fields["level"] = level.String()
	l.fields["message"] = fmt.Sprintf(format, args...)

	enc := json.NewEncoder(l.opts.Out)

	if err := enc.Encode(l.fields); err != nil {
		log.Fatal(err)
	}
}

// Sugar methods
func (l *defaultRecord) Trace(args ...interface{}) {
	l.Log(TraceLevel, args...)
}
func (l *defaultRecord) Tracef(format string, args ...interface{}) {
	l.Logf(TraceLevel, format, args...)
}
func (l *defaultRecord) Debug(args ...interface{}) {
	l.Log(DebugLevel, args...)
}
func (l *defaultRecord) Debugf(format string, args ...interface{}) {
	l.Logf(DebugLevel, format, args...)
}
func (l *defaultRecord) Info(args ...interface{}) {
	l.Log(InfoLevel, args...)
}
func (l *defaultRecord) Infof(format string, args ...interface{}) {
	l.Logf(InfoLevel, format, args...)
}
func (l *defaultRecord) Warn(args ...interface{}) {
	l.Log(WarnLevel, args...)
}
func (l *defaultRecord) Warnf(format string, args ...interface{}) {
	l.Logf(WarnLevel, format, args...)
}
func (l *defaultRecord) Error(args ...interface{}) {
	l.Log(ErrorLevel, args...)
}
func (l *defaultRecord) Errorf(format string, args ...interface{}) {
	l.Logf(ErrorLevel, format, args...)
}
func (l *defaultRecord) Panic(args ...interface{}) {
	l.Log(PanicLevel, args...)
}
func (l *defaultRecord) Panicf(format string, args ...interface{}) {
	l.Logf(PanicLevel, format, args...)
}
func (l *defaultRecord) Fatal(args ...interface{}) {
	l.Log(FatalLevel, args...)
}
func (l *defaultRecord) Fatalf(format string, args ...interface{}) {
	l.Logf(FatalLevel, format, args...)
}
