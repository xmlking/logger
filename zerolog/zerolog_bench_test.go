package zerolog_test

import (
	"testing"

	zlog "github.com/rs/zerolog/log"

	"github.com/xmlking/logger"
	"github.com/xmlking/logger/log"
	"github.com/xmlking/logger/zerolog"
)

func BenchmarkInfoLog(b *testing.B) {
	b.Run("zerolog", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				zlog.Info().Msg("Benchmarking: InfoZ")
			}
		})
	})
	b.Run("wrapper", func(b *testing.B) {
		logger.DefaultLogger = zerolog.NewLogger()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				log.Info("Benchmarking: InfoW")
			}
		})
	})
}
