package gormlog_test

import (
	"os"
	"testing"
	"time"

	"github.com/xmlking/logger"
	"github.com/xmlking/logger/gormlog"
)

func BenchmarkLogger_Print(b *testing.B) {
	mLogger := logger.NewLogger(logger.WithOutput(os.Stdout), logger.WithLevel(logger.DebugLevel))
	l := gormlog.NewGormLogger(mLogger)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Print(
			"sql",
			"/some/file.go:34",
			time.Millisecond*5,
			"SELECT * FROM test WHERE id = $1",
			[]interface{}{42},
			int64(1),
		)
	}
}
