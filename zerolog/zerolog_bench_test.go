package zerolog_test

import (
	"os"
	"testing"

	zlog "github.com/rs/zerolog/log"
	"github.com/xmlking/logger"
	"github.com/xmlking/logger/log"
	"github.com/xmlking/logger/zerolog"
)

func BenchmarkZeroLogger_Info(b *testing.B) {
	logger.DefaultLogger = zerolog.NewLogger(logger.WithOutput(os.Stdout))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		log.Info("Benchmarking: Info")
	}
}

func BenchmarkZeroNative_Info(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		zlog.Info().Msg("Benchmarking: Info")
	}
}
