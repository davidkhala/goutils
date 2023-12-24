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
