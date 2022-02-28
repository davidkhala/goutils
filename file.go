package goutils

import (
	"os"
	"path/filepath"
)

func Which() string {
	ex, err := os.Executable()
	PanicError(err)
	exPath := filepath.Dir(ex)
	return exPath
}
