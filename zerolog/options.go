package zerolog

import (
	"github.com/rs/zerolog"

	"github.com/xmlking/logger"
)

type Options struct {
	logger.Options

	// Flag for whether to log caller info (off by default)
	ReportCaller bool
	// Use this logger as system wide default logger  (off by default)
	UseAsDefault bool
	// zerolog hooks
	Hooks []zerolog.Hook
	// Runtime mode. (Production by default)
	Mode Mode
	// Exit Function to call when FatalLevel log
	ExitFunc func(int)
}

type reportCallerKey struct{}

func ReportCaller() logger.Option {
	return logger.SetOption(reportCallerKey{}, true)
}

type useAsDefaultKey struct{}

func UseAsDefault() logger.Option {
	return logger.SetOption(useAsDefaultKey{}, true)
}

type developmentModeKey struct{}

func WithDevelopmentMode() logger.Option {
	return logger.SetOption(developmentModeKey{}, true)
}

type productionModeKey struct{}

func WithProductionMode() logger.Option {
	return logger.SetOption(productionModeKey{}, true)
}

type gcpModeKey struct{}

func WithGCPMode() logger.Option {
	return logger.SetOption(gcpModeKey{}, true)
}

type hooksKey struct{}

func WithHooks(hooks []zerolog.Hook) logger.Option {
	return logger.SetOption(hooksKey{}, hooks)
}

type exitKey struct{}

func WithExitFunc(exit func(int)) logger.Option {
	return logger.SetOption(exitKey{}, exit)
}
