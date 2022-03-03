package goutils

import "context"

// GetGoContext used by the initialization
func GetGoContext() context.Context {
	return context.Background()
}
