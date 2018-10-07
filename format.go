package goutils

import (
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
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
func FormatInt(integer int64) string {
	return strconv.FormatInt(integer, 10)
}

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

type DeferHandler func(errString string, params ...interface{}) (success bool)

func Deferred(handler DeferHandler, params ...interface{}) {
	err := recover()
	if err == nil {
		return
	}
	var errString = err.(error).Error()
	var success = handler(errString, params...)
	if ! success {
		panic(err)
	}
}
func PanicString(err string) {
	if err != "" {
		panic(errors.New(err))
	}
}
func AssertEmpty(rest []byte, message string) {
	if rest != nil && len(rest) > 0 {
		PanicString(message)
	}
}
func UnixMilliSecond(t time.Time) TimeLong {
	return TimeLong(t.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)))
}

type TimeLong int64

func (t TimeLong) ToString() string {
	return strconv.FormatInt(int64(t), 10)
}

/**
	a wrapper to panic Unmarshal(non-pointer v)
 */
func FromJson(jsonString []byte, v interface{}) {
	err := json.Unmarshal(jsonString, v)
	PanicError(err)
}

func ToJson(v interface{}) []byte {
	result, err := json.Marshal(v)
	PanicError(err)
	return result
}

//not thread safe
func RandString(length int) string {
	var src = rand.NewSource(time.Now().UnixNano())

	var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
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
