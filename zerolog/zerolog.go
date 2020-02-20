package zerolog

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/xmlking/logger"
)

type Mode uint8

const (
	Production Mode = iota
	Development
)

type zeroLogger struct {
	zLog     zerolog.Logger
	exitFunc func(int)
}

func (l *zeroLogger) Init(opts ...logger.Option) error {
	// Default options
	options := &Options{
		Options: logger.Options{
			Level:   100,
			Fields:  make(map[string]interface{}),
			Out:     os.Stderr,
			Context: context.Background(),
		},
		ReportCaller: false,
		UseAsDefault: false,
		Mode:         Production,
		ExitFunc:     os.Exit,
	}

	for _, o := range opts {
		o(&options.Options)
	}

	if hs, ok := options.Context.Value(hooksKey{}).([]zerolog.Hook); ok {
		options.Hooks = hs
	}
	if tf, ok := options.Context.Value(timeFormatKey{}).(string); ok {
		options.TimeFormat = tf
	}
	if exitFunction, ok := options.Context.Value(exitKey{}).(func(int)); ok {
		options.ExitFunc = exitFunction
	}
	if caller, ok := options.Context.Value(reportCallerKey{}).(bool); ok && caller {
		options.ReportCaller = caller
	}
	if useDefault, ok := options.Context.Value(useAsDefaultKey{}).(bool); ok && useDefault {
		options.UseAsDefault = useDefault
	}
	if devMode, ok := options.Context.Value(developmentModeKey{}).(bool); ok && devMode {
		options.Mode = Development
	}
	if prodMode, ok := options.Context.Value(productionModeKey{}).(bool); ok && prodMode {
		options.Mode = Production
	}

	// RESET
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.ErrorStackMarshaler = nil
	l.exitFunc = options.ExitFunc

	switch options.Mode {
	case Development:
		zerolog.ErrorStackMarshaler = func(err error) interface{} {
			fmt.Println(string(debug.Stack()))
			return nil
		}
		consOut := zerolog.NewConsoleWriter(
			func(w *zerolog.ConsoleWriter) {
				if len(options.TimeFormat) > 0 {
					w.TimeFormat = options.TimeFormat
				}
				w.Out = options.Out
				w.NoColor = false
			},
		)
		//level = logger.DebugLevel
		l.zLog = zerolog.New(consOut).
			Level(zerolog.DebugLevel).
			With().Timestamp().Stack().Logger()
	default: // Production
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		l.zLog = zerolog.New(options.Out).
			Level(zerolog.InfoLevel).
			With().Timestamp().Stack().Logger()
	}

	// Set log Level if not default
	if options.Level != 100 {
		zerolog.SetGlobalLevel(loggerToZerologLevel(options.Level))
		l.zLog = l.zLog.Level(loggerToZerologLevel(options.Level))
	}

	// Adding hooks if exist
	if options.ReportCaller {
		l.zLog = l.zLog.With().Caller().Logger()
	}
	for _, hook := range options.Hooks {
		l.zLog = l.zLog.Hook(hook)
	}

	// Setting timeFormat
	if len(options.TimeFormat) > 0 {
		zerolog.TimeFieldFormat = options.TimeFormat
	}

	// Adding seed fields if exist
	if options.Fields != nil {
		l.zLog = l.zLog.With().Fields(options.Fields).Logger()
	}

	// Also set it as zerolog's Default logger
	if options.UseAsDefault {
		zlog.Logger = l.zLog
	}

	return nil
}

func (l *zeroLogger) Log(level logger.Level, template string, fmtArgs []interface{}, fields logger.Fields) {

	// Format with Sprint, Sprintf, or neither.
	msg := template
	if msg == "" && len(fmtArgs) > 0 {
		msg = fmt.Sprint(fmtArgs...)
	} else if msg != "" && len(fmtArgs) > 0 {
		msg = fmt.Sprintf(template, fmtArgs...)
	}

	l.zLog.WithLevel(loggerToZerologLevel(level)).Fields(fields).Msg(msg)
	// Invoke os.Exit because unlike zerolog.Logger.Fatal zerolog.Logger.WithLevel won't stop the execution.
	if level == logger.FatalLevel {
		l.exitFunc(1)
	}
}

func (l *zeroLogger) Error(level logger.Level, template string, fmtArgs []interface{}, err error) {

	// Format with Sprint, Sprintf, or neither.
	msg := template
	if msg == "" && len(fmtArgs) > 0 {
		msg = fmt.Sprint(fmtArgs...)
	} else if msg != "" && len(fmtArgs) > 0 {
		msg = fmt.Sprintf(template, fmtArgs...)
	}

	l.zLog.WithLevel(loggerToZerologLevel(level)).Stack().Err(err).Msg(msg)
	// Invoke os.Exit because unlike zerolog.Logger.Fatal zerolog.Logger.WithLevel won't stop the execution.
	if level == logger.FatalLevel {
		l.exitFunc(1)
	}
}

func (l *zeroLogger) String() string {
	return "zerolog"
}

// NewLogger builds a new logger based on options
func NewLogger(opts ...logger.Option) logger.Logger {
	l := &zeroLogger{}
	_ = l.Init(opts...)
	return l
}

func loggerToZerologLevel(level logger.Level) zerolog.Level {
	switch level {
	case logger.TraceLevel:
		return zerolog.TraceLevel
	case logger.DebugLevel:
		return zerolog.DebugLevel
	case logger.InfoLevel:
		return zerolog.InfoLevel
	case logger.WarnLevel:
		return zerolog.WarnLevel
	case logger.ErrorLevel:
		return zerolog.ErrorLevel
	case logger.PanicLevel:
		return zerolog.PanicLevel
	case logger.FatalLevel:
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}

func ZerologToLoggerLevel(level zerolog.Level) logger.Level {
	switch level {
	case zerolog.TraceLevel:
		return logger.TraceLevel
	case zerolog.DebugLevel:
		return logger.DebugLevel
	case zerolog.InfoLevel:
		return logger.InfoLevel
	case zerolog.WarnLevel:
		return logger.WarnLevel
	case zerolog.ErrorLevel:
		return logger.ErrorLevel
	case zerolog.PanicLevel:
		return logger.PanicLevel
	case zerolog.FatalLevel:
		return logger.FatalLevel
	default:
		return logger.InfoLevel
	}
}
