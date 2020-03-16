package logrus

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/xmlking/logger"
)

type logrusLogger struct {
	*logrus.Logger
	opts Options
}

func (l *logrusLogger) Init(opts ...logger.Option) error {
	for _, o := range opts {
		o(&l.opts.Options)
	}

	if formatter, ok := l.opts.Context.Value(formatterKey{}).(logrus.Formatter); ok {
		l.opts.Formatter = formatter
	}
	if hs, ok := l.opts.Context.Value(hooksKey{}).(logrus.LevelHooks); ok {
		l.opts.Hooks = hs
	}
	if caller, ok := l.opts.Context.Value(reportCallerKey{}).(bool); ok && caller {
		l.opts.ReportCaller = caller
	}
	if exitFunction, ok := l.opts.Context.Value(exitKey{}).(func(int)); ok {
		l.opts.ExitFunc = exitFunction
	}

	l.Logger = &logrus.Logger{
		Out:          l.opts.Out,
		Formatter:    l.opts.Formatter,
		Hooks:        l.opts.Hooks,
		Level:        loggerToLogrusLevel(l.opts.Level),
		ExitFunc:     l.opts.ExitFunc,
		ReportCaller: l.opts.ReportCaller,
	}

	return nil
}

func (l *logrusLogger) Options() logger.Options {
	// FIXME: How to return full opts?
	return l.opts.Options
}

func (l *logrusLogger) WithFields(fields map[string]interface{}) logger.Record {
	// Adding seed fields if exist
	if l.opts.Fields != nil {
		return &logrusRecord{l.Logger.WithFields(l.opts.Fields).WithFields(fields)}
	} else {
		return &logrusRecord{l.Logger.WithFields(fields)}
	}

}

func (l *logrusLogger) WithError(err error) logger.Record {
	// Adding seed fields if exist
	if l.opts.Fields != nil {
		return &logrusRecord{l.Logger.WithFields(l.opts.Fields).WithError(err)}
	} else {
		return &logrusRecord{l.Logger.WithError(err)}
	}
}

func (l *logrusLogger) Log(level logger.Level, args ...interface{}) {
	// Adding seed fields if exist
	if l.opts.Fields != nil {
		l.Logger.WithFields(l.opts.Fields).Log(loggerToLogrusLevel(level), args...)
	} else {
		l.Logger.Log(loggerToLogrusLevel(level), args...)
	}

}

func (l *logrusLogger) Logf(level logger.Level, format string, args ...interface{}) {
	// Adding seed fields if exist
	if l.opts.Fields != nil {
		l.Logger.WithFields(l.opts.Fields).Logf(loggerToLogrusLevel(level), format, args...)
	} else {
		l.Logger.Logf(loggerToLogrusLevel(level), format, args...)
	}
}

func (l *logrusLogger) String() string {
	return "logrus"
}

// New builds a new logger based on options
func NewLogger(opts ...logger.Option) logger.Logger {
	// Default options
	options := Options{
		Options: logger.Options{
			Level:   logger.InfoLevel,
			Fields:  make(map[string]interface{}),
			Out:     os.Stderr,
			Context: context.Background(),
		},
		Formatter:    new(logrus.TextFormatter),
		Hooks:        make(logrus.LevelHooks),
		ReportCaller: false,
		ExitFunc:     os.Exit,
	}
	l := &logrusLogger{opts: options}
	_ = l.Init(opts...)
	return l
}

func loggerToLogrusLevel(level logger.Level) logrus.Level {
	switch level {
	case logger.TraceLevel:
		return logrus.TraceLevel
	case logger.DebugLevel:
		return logrus.DebugLevel
	case logger.InfoLevel:
		return logrus.InfoLevel
	case logger.WarnLevel:
		return logrus.WarnLevel
	case logger.ErrorLevel:
		return logrus.ErrorLevel
	case logger.PanicLevel:
		return logrus.PanicLevel
	case logger.FatalLevel:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
