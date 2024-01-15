package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	. "github.com/davidkhala/goutils"
	"math/big"
)

// ECPriv Used for both ECDSA and ECDH
type ECPriv struct {
	*ecdsa.PrivateKey
}

// NewECPriv generate an EC private key (default to use P256 curve)
func NewECPriv(curve elliptic.Curve) ECPriv {
	if curve == nil {
		curve = elliptic.P256()
	}
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	PanicError(err)

	return ECPriv{privateKey}
}
func (t ECPriv) Sign(digest []byte) []byte {
	var r, s, err = ecdsa.Sign(rand.Reader, t.PrivateKey, digest)

	PanicError(err)
	return ECDSASignature{r, s}.MarshalOrPanic()
}

type ECDSASignature struct {
	R, S *big.Int
}

func (t ECDSASignature) MarshalOrPanic() []byte {
	var result, err = asn1.Marshal(t)
	PanicError(err)
	return result
}
func (ECDSASignature) UnmarshalOrPanic(signature []byte) (ecdsaSignature ECDSASignature) {
	var rest, err = asn1.Unmarshal(signature, &ecdsaSignature)
	PanicError(err)
	AssertEmptyOrPanic[byte](rest, "asn1 unmarshal failed")
	return
}
func (ECPriv) LoadPem(pemBytes []byte) ECPriv {
	block, rest := pem.Decode(pemBytes)
	AssertEmptyOrPanic[byte](rest, "pem decode failed")
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	PanicError(err)
	return ECPriv{privateKey}
}

func (t ECPriv) ToPem() []byte {
	writer := bytes.NewBufferString("")
	keyBytes, err := x509.MarshalECPrivateKey(t.PrivateKey)
	PanicError(err)
	err = pem.Encode(writer, &pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes})
	PanicError(err)
	return writer.Bytes()
}

// ECPub is a representation of an elliptic curve public key.
type ECPub struct {
	*ecdsa.PublicKey
}

func (t ECPub) Verify(digest []byte, signature []byte) bool {
	var ecdsaSignature = ECDSASignature{}.UnmarshalOrPanic(signature)
	return ecdsa.Verify(t.PublicKey, digest, ecdsaSignature.R, ecdsaSignature.S)
}
func (t *ECPub) LoadCert(pemBytes []byte) {
	var cert = ParseCertPemOrPanic(pemBytes)
	t.PublicKey = cert.PublicKey.(*ecdsa.PublicKey)
}
