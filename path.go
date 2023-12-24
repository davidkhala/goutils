package goutils

import (
	"os/user"
	"path/filepath"
)

func HomeResolve(tokens ...string) string {
	_user, err := user.Current()
	PanicError(err)
	paths := append([]string{_user.HomeDir}, tokens...)
	return filepath.Join(paths...)
}
func Absolute(_path string) string {
	_absPath, err := filepath.Abs(_path)
	PanicError(err)
	return _absPath
}
