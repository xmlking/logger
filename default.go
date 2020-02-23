package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type defaultLogger struct {
	opts Options
}

// Init(opts...) should only overwrite provided options
func (l *defaultLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.opts)
	}
	return nil
}

func (l *defaultLogger) String() string {
	return "default"
}

func (l *defaultLogger) Log(level Level, template string, fmtArgs []interface{}, fields map[string]interface{}) {
	if !l.opts.Level.Enabled(level) {
		return
	}
	// Format with Sprint, Sprintf, or neither.
	msg := template
	if msg == "" && len(fmtArgs) > 0 {
		msg = fmt.Sprint(fmtArgs...)
	} else if msg != "" && len(fmtArgs) > 0 {
		msg = fmt.Sprintf(template, fmtArgs...)
	}

	if len(l.opts.Fields) > 0 {
		fields = MergeMaps(l.opts.Fields, fields)
	} else if len(fields) == 0 {
		fields = make(map[string]interface{}, 2)
	}

	fields["level"] = level.String()
	fields["message"] = msg

	enc := json.NewEncoder(l.opts.Out)

	if err := enc.Encode(fields); err != nil {
		log.Fatal(err)
	}
}

func (l *defaultLogger) Error(level Level, template string, fmtArgs []interface{}, err error) {
	if level < l.opts.Level {
		return
	}
	// Format with Sprint, Sprintf, or neither.
	msg := template
	if msg == "" && len(fmtArgs) > 0 {
		msg = fmt.Sprint(fmtArgs...)
	} else if msg != "" && len(fmtArgs) > 0 {
		msg = fmt.Sprintf(template, fmtArgs...)
	}

	fields := map[string]interface{}{
		"level":   level.String(),
		"message": msg,
		"error":   err.Error(),
	}

	if len(l.opts.Fields) > 0 {
		fields = MergeMaps(l.opts.Fields, fields)
	}

	enc := json.NewEncoder(l.opts.Out)

	if err := enc.Encode(fields); err != nil {
		log.Fatal(err)
	}

}

func (n *defaultLogger) Options() Options {
	return n.opts
}

// NewLogger builds a new logger based on options
func NewLogger(opts ...Option) Logger {
	// Default options
	options := Options{
		Level:   InfoLevel,
		Fields:  make(map[string]interface{}),
		Out:     os.Stderr,
		Context: context.Background(),
	}

	l := &defaultLogger{opts: options}
	_ = l.Init(opts...)
	return l
}
