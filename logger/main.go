package logger

import (
	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"os"
)

func Zap() *zap.Logger {
	return zap.NewExample()
}
func ZeroLog() zerolog.Logger {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	return zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})
}
