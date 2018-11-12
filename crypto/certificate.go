package crypto

import (
	"crypto/x509"
	"encoding/pem"
	. "github.com/davidkhala/goutils"
)

func ParseCertPem(pemBytes []byte) *x509.Certificate {
	block, rest := pem.Decode(pemBytes)
	assertEmpty(rest, "pem decode failed")
	cert, err := x509.ParseCertificate(block.Bytes)
	PanicError(err)
	return cert
}
