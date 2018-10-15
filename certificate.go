package goutils

import (
	"crypto/x509"
	"encoding/pem"
)

func ParseCertPem(pemBytes []byte) *x509.Certificate {
	block, rest := pem.Decode(pemBytes)
	AssertEmpty(rest, "pem decode failed:"+string(rest))
	cert, err := x509.ParseCertificate(block.Bytes)
	PanicError(err)
	return cert
}
