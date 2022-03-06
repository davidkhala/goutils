package goutils

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

// Filename is the __filename equivalent
func Filename() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

// Dirname is the __dirname equivalent
func Dirname() (string, error) {
	filename, err := Filename()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filename), nil
}

func ReadFile(file string) ([]byte, error) {
	if file == "" {
		return nil, nil
	}

	in, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return in, nil
}
func ReadFileOrPanic(file string) []byte {
	byteSlice, err := ReadFile(file)
	PanicError(err)
	return byteSlice
}
