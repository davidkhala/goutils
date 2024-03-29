package goutils

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Atoi(str string) int {
	i, err := strconv.Atoi(str)
	PanicError(err)
	return i
}
func ParseFloat(str string) float64 {
	var result, err = strconv.ParseFloat(str, 64)
	PanicError(err)
	return result
}
func ToInt(bytes []byte) int {
	if bytes == nil {
		return 0
	}
	return Atoi(string(bytes))
}
func FormatFloat(f float64, precision int) string {
	return strconv.FormatFloat(f, 'f', precision, 64)
}

func RoundFloat(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	var round = func(num float64) int {
		return int(num + math.Copysign(0.5, num))
	}
	return float64(round(num*output)) / output
}

type DeferHandler func(err error, params ...interface{}) (success bool)

func Deferred(handler DeferHandler, params ...interface{}) {
	recovered := recover()

	var err error
	switch recovered.(type) {
	case error:
		err = recovered.(error)
	case string:
		err = errors.New(recovered.(string))
	default:
		if recovered == nil {
			return
		} else {
			panic(recovered)
		}
	}

	var success = handler(err, params...)
	if !success {
		panic(recovered)
	}
}

// TimeLong unix nano
type TimeLong int64

func (TimeLong) FromTime(t time.Time) TimeLong {
	return TimeLong(t.UnixNano())
}
func (TimeLong) FromString(s string) TimeLong {
	i, err := strconv.ParseInt(s, 10, 64)
	PanicError(err)
	return TimeLong(i)
}
func (TimeLong) FromUnixMilliSecond(t int64) TimeLong {
	return TimeLong(t * int64(time.Millisecond))
}
func (t TimeLong) UnixMilliSecond() int64 {
	return int64(t) / int64(time.Millisecond)
}
func (t TimeLong) String() string {
	return strconv.FormatInt(int64(t), 10)
}

// FromJson a wrapper to panic Unmarshal(non-pointer v)
func FromJson(jsonString []byte, v interface{}) {
	err := json.Unmarshal(jsonString, v)
	PanicError(err)
}

func ToJson(v interface{}) []byte {
	result, err := json.Marshal(v)
	PanicError(err)
	return result
}

// RandString not thread safe
func RandString(length int, letterBytes string) string {
	var src = rand.NewSource(time.Now().UnixNano())
	if letterBytes == "" {
		PanicString("RandString: empty letter array")
	}

	var letterIdxBits = uint(6)                     // 6 bits to represent a letter index
	var letterIdxMask = int64(1<<letterIdxBits - 1) // All 1-bits, as many as letterIdxBits
	var letterIdxMax = 63 / letterIdxBits           // # of letter indices fitting in 63 bits
	b := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

var Base64Encode = base64.StdEncoding.EncodeToString

func Base64DecodeOrPanic(s string) []byte {
	result, err := base64.StdEncoding.DecodeString(s)
	PanicError(err)
	return result
}

var HexEncode = hex.EncodeToString

func HexDecodeOrPanic(s string) []byte {
	result, err := hex.DecodeString(s)
	PanicError(err)
	return result
}

func ItoRunes(i int, runes []rune) string {
	var d = len(runes)
	var forward = func(p int) (int, int) {
		rest := p / d
		module := p % d
		return rest, module
	}

	r, m := forward(i)
	var output = []rune{runes[m]}
	for r > 0 {
		r, m = forward(r)
		output = append([]rune{runes[m]}, output...)
	}

	return string(output)
}

func PaddingLeft(str string, length int, pad rune) string {
	return strings.Repeat(string(pad), length-len(str)) + str
}

func PaddingRight(str string, length int, pad rune) string {
	return str + strings.Repeat(string(pad), length-len(str))
}

func IsEmpty[T any](arr []T) bool {
	return arr == nil || len(arr) == 0
}
