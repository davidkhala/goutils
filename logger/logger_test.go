package logger

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogrus(t *testing.T) {
	SetupLogrus()
	logrus.Debug("babc")
}

func TestZap(t *testing.T) {
	var logger = Zap()
	logger.Debug("abc")
}
