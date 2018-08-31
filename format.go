package goutils

import (
	"strconv"
	"encoding/json"
	"errors"
	"time"
)

func Atoi(str string) int {
	i, err := strconv.Atoi(str)
	PanicError(err)
	return i
}
func ToInt(bytes []byte) int {
	if bytes == nil {
		return 0
	}
	return Atoi(string(bytes))
}
func ToString(integer int64) string {
	return strconv.FormatInt(integer, 10)
}

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

type DeferHandler func(errString string) (toThrow bool)

func Deferred(handler DeferHandler) {
	err := recover()
	if err == nil {
		return
	}
	var errString = err.(error).Error()
	var toThrow = handler(errString)
	if toThrow {
		panic(new)
	}
}
func PanicString(err string) {
	if err != "" {
		panic(errors.New(err))
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
