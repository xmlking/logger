package zap

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/xmlking/logger"
)

type zaplog struct {
	*zap.Logger
	opts logger.Options
}

func (l *zaplog) Init(opts ...logger.Option) error {
	var err error

	for _, o := range opts {
		o(&l.opts)
	}

	zapConfig := zap.NewProductionConfig()
	if zconfig, ok := l.opts.Context.Value(configKey{}).(zap.Config); ok {
		zapConfig = zconfig
	}

	if zcconfig, ok := l.opts.Context.Value(encoderConfigKey{}).(zapcore.EncoderConfig); ok {
		zapConfig.EncoderConfig = zcconfig

	}

	// Set log Level if not default
	zapConfig.Level = zap.NewAtomicLevel()
	if l.opts.Level != logger.InfoLevel {
		zapConfig.Level.SetLevel(loggerToZapLevel(l.opts.Level))
	}

	// Adding seed fields if exist
	if l.opts.Fields != nil {
		zapConfig.InitialFields = l.opts.Fields
	}

	log, err := zapConfig.Build()
	if err != nil {
		return err
	}

	// Adding namespace
	if namespace, ok := l.opts.Context.Value(namespaceKey{}).(string); ok {
		log = log.With(zap.Namespace(namespace))
	}

	// defer log.Sync() ??
	l.Logger = log

	return nil
}

func (l *zaplog) Options() logger.Options {
	return l.opts
}

func (l *zaplog) WithFields(fields map[string]interface{}) logger.Record {
	return &zapRecord{l.Logger, fields}
}

func (l *zaplog) WithError(err error) logger.Record {
	return &zapRecord{l.Logger, map[string]interface{}{"error": err}}
}

func (l *zaplog) Log(level logger.Level, args ...interface{}) {
	lvl := loggerToZapLevel(level)
	msg := fmt.Sprint(args...)
	if ce := l.Logger.Check(lvl, msg); ce != nil {
		ce.Write()
	}
}

func (l *zaplog) Logf(level logger.Level, format string, args ...interface{}) {
	lvl := loggerToZapLevel(level)
	msg := fmt.Sprintf(format, args...)
	if ce := l.Logger.Check(lvl, msg); ce != nil {
		ce.Write()
	}
}

func (l *zaplog) String() string {
	return "zap"
}

// New builds a new logger based on options
func NewLogger(opts ...logger.Option) (logger.Logger, error) {
	// Default options
	options := logger.Options{
		Level:   logger.InfoLevel,
		Fields:  make(map[string]interface{}),
		Out:     os.Stderr,
		Context: context.Background(),
	}

	l := &zaplog{opts: options}
	if err := l.Init(opts...); err != nil {
		return nil, err
	}

	return l, nil
}

func loggerToZapLevel(level logger.Level) zapcore.Level {
	switch level {
	case logger.TraceLevel, logger.DebugLevel:
		return zap.DebugLevel
	case logger.InfoLevel:
		return zap.InfoLevel
	case logger.WarnLevel:
		return zap.WarnLevel
	case logger.ErrorLevel:
		return zap.ErrorLevel
	case logger.PanicLevel:
		return zap.PanicLevel
	case logger.FatalLevel:
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}
