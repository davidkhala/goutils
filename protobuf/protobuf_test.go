package protobuf

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestTimeStamp(t *testing.T) {
	var now = time.Now()
	var unixTime = now.Unix()
	var unixNano = now.UnixNano()

	fmt.Println("now.String() =", now.String())
	fmt.Println("now.Unix() =", unixTime)
	fmt.Println("now.UnixNano() =", unixNano)
	var stamp timestamp.Timestamp
	stamp.Seconds = now.Unix()
	stamp.Nanos = int32(now.Nanosecond())
	var timeL = FromTimeStamp(&stamp)
	assert.EqualValues(t, unixNano, timeL)

}
