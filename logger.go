// Package log provides a log interface
package logger

var (
	// Default logger
	DefaultLogger Logger = NewLogger()
)

type Fields map[string]interface{}

// Logger is a generic logging interface
type Logger interface {
	// Init initializes options
	Init(options ...Option) error
	// The Logger options
	Options() Options
	// log at given level with message, fmtArgs and context fields
	Log(level Level, template string, fmtArgs []interface{}, fields Fields)
	// log error at given level with message, fmtArgs and stack if enabled.
	Error(level Level, template string, fmtArgs []interface{}, err error)
	// String returns the name of logger
	String() string
}

// func Init(opts ...Option) error {
// 	return DefaultLogger.Init(opts...)
// }

// Set DefaultLogger Level
func SetLevel(lvl Level) {
	if err := DefaultLogger.Init(WithLevel(lvl)); err != nil {
		print(err)
	}
}

// Log to DefaultLogger
func Log(level Level, template string, fmtArgs []interface{}, fields Fields) {
	DefaultLogger.Log(level, template, fmtArgs, fields)
}

// Error log to DefaultLogger
func Error(level Level, template string, fmtArgs []interface{}, err error) {
	DefaultLogger.Error(level, template, fmtArgs, err)
}

// Get DefaultLogger name
func String() string {
	return DefaultLogger.String()
}
