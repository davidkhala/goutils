package goutils

import (
	"fmt"
	"github.com/kortschak/utter"
	"github.com/stretchr/testify/assert"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
	"math"
	"testing"
	"time"
)

var Map = map[string]string{}

func TestRange(t *testing.T) {
	Map["abc"] = "1"
	for key, value := range Map {
		fmt.Println(key, value)
	}
}
func TestMap2JSON(t *testing.T) {
	Map["abc"] = "2"
	assert.Equal(t, `{"abc":"2"}`, string(ToJson(Map)))
}
func TestInf(t *testing.T) {
	var positiveInf = math.Inf(1)
	assert.Equal(t, "+Inf", fmt.Sprint(positiveInf))
	assert.True(t, 1.00-positiveInf < 0)

}
func TestSlice(t *testing.T) {
	var array = []int{1, 2, 3}
	fmt.Println(append([]int{}, array[1:]...))
}
func TestTimeStamp(t *testing.T) {
	var now = time.Now()
	var unixTime = now.Unix()
	var unixNano = now.UnixNano()

	fmt.Println(now.String(), unixTime, unixNano)
	var stamp timestamp.Timestamp
	stamp.Seconds = now.Unix()
	stamp.Nanos = int32(now.Nanosecond())
	fmt.Println(stamp.Seconds, stamp.Nanos)
}
func TestUtter(t *testing.T) {
	fmt.Println(Map)
	utter.Dump(Map)
}
func TestNil(t *testing.T) {
	var arrays []string
	arrays = nil
	assert.Equal(t, 0, len(arrays))
}
