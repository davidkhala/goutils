package goutils

func AssertEmpty[T any](rest []T, message string) {
	AssertOK(IsEmpty[T](rest), message)
}
func AssertOK(condition bool, message string) {
	if !condition {
		panic(message)
	}
}
