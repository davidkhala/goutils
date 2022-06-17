package logger

import (
	"github.com/davidkhala/goutils"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"
)

func Zap() *zap.Logger {
	return zap.NewExample()
}

type ZeroLogFormatter struct {
}

func (_ ZeroLogFormatter) UnixSecTimestamp(i interface{}) string {
	var t, err = time.Parse(time.RFC3339, i.(string))
	goutils.PanicError(err)
	return strconv.FormatInt(t.Unix()*1000, 10)
}
func (_ ZeroLogFormatter) String(i interface{}) string {
	return i.(string)
}
func ZeroLog(formatter zerolog.Formatter) zerolog.Logger {
	// UNIX Time is faster and smaller than most timestamps
	if formatter == nil {
		var f ZeroLogFormatter
		formatter = f.UnixSecTimestamp
	}

	return zerolog.New(zerolog.ConsoleWriter{
		Out:             os.Stdout,
		FormatTimestamp: formatter,
	}).With().Timestamp().Logger()
}
