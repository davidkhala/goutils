package goutils

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"math/big"
	"reflect"
)

type ECDSAPriv struct {
	*ecdsa.PrivateKey
}

//generate an EC private key (P256 curve)
func (ECDSAPriv) New() (ECDSAPriv) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
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
	if rest != nil && len(rest) > 0 {
		PanicString("asn1 unmarshal failed:" + string(rest))
	}
	PanicError(err)
	return
}
func (ECDSAPriv) LoadPem(pemBytes []byte) (ECDSAPriv) {
	block, rest := pem.Decode(pemBytes)
	if rest != nil && len(rest) > 0 {
		PanicString("pem decode failed:" + string(rest))
	}

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

type ECDSAPub struct {
	*ecdsa.PublicKey
}

func (t ECDSAPub) Verify(digest []byte, signature []byte) bool {
	var ecdsaSignature = ECDSASignature{}.Unmarshal(signature)
	return ecdsa.Verify(t.PublicKey, digest, ecdsaSignature.R, ecdsaSignature.S)
}
func (ECDSAPub) LoadPem(pemBytes []byte) ECDSAPub {
	block, rest := pem.Decode(pemBytes)
	if rest != nil && len(rest) > 0 {
		PanicString("pem decode failed:" + string(rest))
	}
	var cert, err = x509.ParseCertificate(block.Bytes)
	PanicError(err)

	var pubkey = cert.PublicKey.(*ecdsa.PublicKey)
	return ECDSAPub{pubkey}
}

//default in nodejs sdk
type PKCS8 struct {
	pem.Block
	Key interface{}
	reflect.Type
}

func (PKCS8) LoadPem(pemBytes []byte) (PKCS8) {
	block, rest := pem.Decode(pemBytes)
	if rest != nil && len(rest) > 0 {
		PanicString("pem decode failed:" + string(rest))
	}

	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	PanicError(err)
	return PKCS8{*block, privKey, reflect.TypeOf(privKey)}
}
func (t PKCS8) FormatECDSA() *ecdsa.PrivateKey {
	var result = t.Key.(*ecdsa.PrivateKey)
	return result
}

func (t PKCS8) ToPem() []byte {
	writer := bytes.NewBufferString("")
	pem.Encode(writer, &t.Block)
	return writer.Bytes()
}
