package logger

import (
	"testing"
)

func TestZap(t *testing.T) {
	var logger = Zap()
	logger.Debug("abc")
}
func TestZeroLog(t *testing.T) {
	var f ZeroLogFormatter
	var logger = ZeroLog(f.String)

	logger.Info().Msg("hello world")
}
