package zap

import (
	"github.com/xmlking/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	logger.Options
}

type configKey struct{}

// WithConfig pass zap.Config to logger
func WithConfig(c zap.Config) logger.Option {
	return setOption(configKey{}, c)
}

type encoderConfigKey struct{}

// WithEncoderConfig pass zapcore.EncoderConfig to logger
func WithEncoderConfig(c zapcore.EncoderConfig) logger.Option {
	return setOption(encoderConfigKey{}, c)
}

type levelKey struct{}

// WithLevel pass log level
func WithLevel(l logger.Level) logger.Option {
	return setOption(levelKey{}, l)
}

type fieldsKey struct{}

func WithFields(fields logger.Fields) logger.Option {
	return setOption(fieldsKey{}, fields)
}

type namespaceKey struct{}

func WithNamespace(namespace string) logger.Option {
	return setOption(namespaceKey{}, namespace)
}
