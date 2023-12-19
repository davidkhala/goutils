package protobuf

import (
	"github.com/davidkhala/goutils"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func FromTimeStamp(t *timestamppb.Timestamp) goutils.TimeLong {
	return goutils.TimeLong(t.GetSeconds()*int64(time.Second) + int64(t.GetNanos()))
}
