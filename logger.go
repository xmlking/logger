// Package logger provides logger configuration
package logger

import (
	"context"
	"os"
	"sync"
	"time"
)

var (
	// Default logger
	DefaultLogger Logger = NewLogger()
)

// Logger is a generic logging interface
type Logger interface {
	// Init initializes options
	Init(options ...Option) error
	// The Logger options
	Options() Options
	// fields to be logged once
	WithFields(fields map[string]interface{}) Record
	// error to be logged once
	WithError(err error) Record
	// write args along with logger's default fields.
	Log(level Level, args ...interface{})
	// write formatted args along with logger's default fields.
	Logf(level Level, format string, args ...interface{})
	// String returns the name of logger
	String() string
}

/**
 * Default logger implementation
 */
type defaultLogger struct {
	sync.RWMutex
	opts Options
}

// Init(opts...) should only overwrite provided options
func (l *defaultLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.opts)
	}
	return nil
}

func (n *defaultLogger) Options() Options {
	return n.opts
}

func (l *defaultLogger) WithFields(fields map[string]interface{}) Record {
	if len(l.opts.Fields) > 0 {
		l.RLock()
		fields = MergeMaps(l.opts.Fields, fields)
		l.RUnlock()
	} else if len(fields) == 0 {
		fields = make(map[string]interface{}, 3)
	}
	return &defaultRecord{
		opts:   l.opts,
		level:  l.opts.Level,
		fields: fields,
	}
}

func (l *defaultLogger) WithError(err error) Record {
	l.RLock()
	fields := MergeMaps(l.opts.Fields)
	l.RUnlock()

	return &defaultRecord{
		opts:   l.opts,
		level:  l.opts.Level,
		fields: fields,
		err:    err,
	}
}

func (l *defaultLogger) Log(level Level, args ...interface{}) {
	if !l.opts.Level.Enabled(level) {
		return
	}

	r := defaultRecord{
		opts:   l.opts,
		level:  l.opts.Level,
		fields: l.opts.Fields,
	}
	r.Log(level, args...)

}

func (l *defaultLogger) Logf(level Level, format string, args ...interface{}) {
	if !l.opts.Level.Enabled(level) {
		return
	}

	r := defaultRecord{
		opts:   l.opts,
		level:  l.opts.Level,
		fields: l.opts.Fields,
	}
	r.Logf(level, format, args...)
}

func (l *defaultLogger) String() string {
	return "default"
}

// NewLogger builds a new logger based on options
func NewLogger(opts ...Option) Logger {
	// Default options
	options := Options{
		Level:      InfoLevel,
		TimeFormat: time.RFC3339,
		Fields:     make(map[string]interface{}),
		Out:        os.Stderr,
		Context:    context.Background(),
	}

	l := &defaultLogger{opts: options}
	_ = l.Init(opts...)
	return l
}

// MergeMaps will overwriting duplicate keys, you should handle that, if there is a need
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	var sum int
	for _, m := range maps {
		sum += len(m)
	}
	result := make(map[string]interface{}, sum)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
