// Package log provides a log interface
package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	// Default logger
	DefaultLogger Logger = NewLogger()
)

type Fields map[string]interface{}

// Logger is a generic logging interface
type Logger interface {
	// Init initializes options
	Init(options ...Option) error
	// log at given level with message, fmtArgs and context fields
	Log(level Level, template string, fmtArgs []interface{}, fields Fields)
	// log error at given level with message, fmtArgs and stack if enabled.
	Error(level Level, template string, fmtArgs []interface{}, err error)
	// String returns the name of logger
	String() string
}

type defaultLogger struct {
	level  Level
	fields Fields
	out    io.Writer
}

func (l *defaultLogger) Init(opts ...Option) error {
	options := &Options{
		Level:   InfoLevel,
		Fields:  make(map[string]interface{}),
		Out:     os.Stderr,
		Context: context.Background(),
	}

	for _, o := range opts {
		o(options)
	}

	l.level = options.Level
	l.fields = options.Fields
	l.out = options.Out

	return nil
}

func (l *defaultLogger) SetLevel(level Level) {
	l.level = level
}

func (l *defaultLogger) Level() Level {
	return l.level
}

func (l *defaultLogger) String() string {
	return "basic"
}

func (l *defaultLogger) Log(level Level, template string, fmtArgs []interface{}, fields Fields) {
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

	fields = mergeMaps(l.fields, fields)
	fields["message"] = msg

	enc := json.NewEncoder(l.out)

	if err := enc.Encode(fields); err != nil {
		log.Fatal(err)
	}
}

func (l *defaultLogger) Error(level Level, template string, fmtArgs []interface{}, err error) {
	if level < l.level {
		return
	}
	// Format with Sprint, Sprintf, or neither.
	msg := template
	if msg == "" && len(fmtArgs) > 0 {
		msg = fmt.Sprint(fmtArgs...)
	} else if msg != "" && len(fmtArgs) > 0 {
		msg = fmt.Sprintf(template, fmtArgs...)
	}

	fields := mergeMaps(l.fields, map[string]interface{}{
		"message": msg,
		"error":   err.Error(),
	})

	enc := json.NewEncoder(l.out)

	if err := enc.Encode(fields); err != nil {
		log.Fatal(err)
	}

}

// NewLogger builds a new logger based on options
func NewLogger(opts ...Option) Logger {
	l := &defaultLogger{}
	_ = l.Init(opts...)
	return l
}

// overwriting duplicate keys, you should handle that if there is a need
func mergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
