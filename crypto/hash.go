package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
)

func HashSha256(data []byte) []byte {
	return sha256.Sum256(data)[:]
}
func HashSha512(data []byte) []byte {
	return sha512.Sum512(data)[:]
}
