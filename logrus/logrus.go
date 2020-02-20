package logrus

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/xmlking/logger"
)

type logrusLogger struct {
	*logrus.Logger
}

func (l *logrusLogger) Init(opts ...logger.Option) error {
	// Default options
	options := &Options{
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

	for _, o := range opts {
		o(&options.Options)
	}

	if formatter, ok := options.Context.Value(formatterKey{}).(logrus.Formatter); ok {
		options.Formatter = formatter
	}
	if hs, ok := options.Context.Value(hooksKey{}).(logrus.LevelHooks); ok {
		options.Hooks = hs
	}
	if caller, ok := options.Context.Value(reportCallerKey{}).(bool); ok && caller {
		options.ReportCaller = caller
	}
	if exitFunction, ok := options.Context.Value(exitKey{}).(func(int)); ok {
		options.ExitFunc = exitFunction
	}

	l.Logger = &logrus.Logger{
		Out:          options.Out,
		Formatter:    options.Formatter,
		Hooks:        options.Hooks,
		Level:        loggerToLogrusLevel(options.Level),
		ExitFunc:     options.ExitFunc,
		ReportCaller: options.ReportCaller,
	}

	return nil
}

func (l *logrusLogger) SetLevel(level logger.Level) {
	l.Logger.SetLevel(loggerToLogrusLevel(level))
}

func (l *logrusLogger) Level() logger.Level {
	return logrusToLoggerLevel(l.Logger.Level)
}

func (l *logrusLogger) String() string {
	return "logrus"
}

func (l *logrusLogger) Log(level logger.Level, template string, fmtArgs []interface{}, fields logger.Fields) {
	var fld map[string]interface{} = fields

	// Format with Sprint, Sprintf, or neither.
	msg := template
	if msg == "" && len(fmtArgs) > 0 {
		msg = fmt.Sprint(fmtArgs...)
	} else if msg != "" && len(fmtArgs) > 0 {
		msg = fmt.Sprintf(template, fmtArgs...)
	}

	l.Logger.WithFields(fld).Log(loggerToLogrusLevel(level), msg)
}
func (l *logrusLogger) Error(level logger.Level, template string, fmtArgs []interface{}, err error) {

	// Format with Sprint, Sprintf, or neither.
	msg := template
	if msg == "" && len(fmtArgs) > 0 {
		msg = fmt.Sprint(fmtArgs...)
	} else if msg != "" && len(fmtArgs) > 0 {
		msg = fmt.Sprintf(template, fmtArgs...)
	}

	l.Logger.WithError(err).Log(loggerToLogrusLevel(level), msg)
}

// New builds a new logger based on options
func NewLogger(opts ...logger.Option) logger.Logger {
	l := &logrusLogger{}
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

func logrusToLoggerLevel(level logrus.Level) logger.Level {
	switch level {
	case logrus.TraceLevel:
		return logger.TraceLevel
	case logrus.DebugLevel:
		return logger.DebugLevel
	case logrus.InfoLevel:
		return logger.InfoLevel
	case logrus.WarnLevel:
		return logger.WarnLevel
	case logrus.ErrorLevel:
		return logger.ErrorLevel
	case logrus.PanicLevel:
		return logger.PanicLevel
	case logrus.FatalLevel:
		return logger.FatalLevel
	default:
		return logger.InfoLevel
	}
}
