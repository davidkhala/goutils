package crypto

func assertEmpty(rest []byte, message string) {
	if rest != nil && len(rest) > 0 {
		panic(message + ":" + string(rest))
	}
}
