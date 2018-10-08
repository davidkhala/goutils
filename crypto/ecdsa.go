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

type ECDSAPriv struct {
	*ecdsa.PrivateKey
}

//generate an EC private key (default to use P256 curve)
func (ECDSAPriv) New(curve elliptic.Curve) (ECDSAPriv) {
	if curve == nil {
		curve = elliptic.P256()
	}
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	PanicError(err)

	return ECDSAPriv{priv}
}
func (t ECDSAPriv) Sign(digest []byte) []byte {
	var r, s, err = ecdsa.Sign(rand.Reader, t.PrivateKey, digest)

	PanicError(err)
	return ECDSASignature{r, s}.Marshal()
}

type ECDSASignature struct {
	R, S *big.Int
}

func (t ECDSASignature) Marshal() []byte {
	var result, err = asn1.Marshal(t)
	PanicError(err)
	return result
}
func (ECDSASignature) Unmarshal(signature []byte) (ecdsaSignature ECDSASignature) {
	var rest, err = asn1.Unmarshal(signature, &ecdsaSignature)
	AssertEmpty(rest, "asn1 unmarshal failed:"+string(rest))
	PanicError(err)
	return
}
func (ECDSAPriv) LoadPem(pemBytes []byte) (ECDSAPriv) {
	block, rest := pem.Decode(pemBytes)
	AssertEmpty(rest, "pem decode failed:"+string(rest))
	privKey, err := x509.ParseECPrivateKey(block.Bytes)
	PanicError(err)
	return ECDSAPriv{privKey}
}

func (t ECDSAPriv) ToPem() []byte {
	writer := bytes.NewBufferString("")
	keyBytes, err := x509.MarshalECPrivateKey(t.PrivateKey)
	PanicError(err)
	pem.Encode(writer, &pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes})
	return writer.Bytes()
}

// PublicKey is a representation of an elliptic curve public key.
type ECDSAPub struct {
	*ecdsa.PublicKey
}

func (t ECDSAPub) Verify(digest []byte, signature []byte) bool {
	var ecdsaSignature = ECDSASignature{}.Unmarshal(signature)
	return ecdsa.Verify(t.PublicKey, digest, ecdsaSignature.R, ecdsaSignature.S)
}
func (ECDSAPub) LoadCert(pemBytes []byte) ECDSAPub {
	block, rest := pem.Decode(pemBytes)
	AssertEmpty(rest, "pem decode failed:"+string(rest))
	var cert, err = x509.ParseCertificate(block.Bytes)
	PanicError(err)

	var pubkey = cert.PublicKey.(*ecdsa.PublicKey)
	return ECDSAPub{pubkey}
}
