package crypto

import (
	"fmt"
	"testing"
)

var testPrivPem = `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQghaUlXR2fnIt5Ogvi
M3C3tCECJxW2XSEpKXjRPXVFX9GhRANCAASfIOZya/nqC0tXVEK/mE6UdwsEsjdj
U+kTVfJPv/wsLayovBrUQvsT+XeDZXpMhCX2z35eSbSEkGD+5DSAOcdc
-----END PRIVATE KEY-----
`
var testPubPem = `-----BEGIN CERTIFICATE-----
MIICYjCCAgigAwIBAgIUTTWiUslu3NcIAj5hGokwMywqxoswCgYIKoZIzj0EAwIw
XjELMAkGA1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMRQwEgYDVQQK
EwtIeXBlcmxlZGdlcjEPMA0GA1UECxMGRmFicmljMQ8wDQYDVQQDEwZjYS5NQ0Mw
HhcNMTgxMDA3MDg1MTAwWhcNMTkxMDA3MDg1NjAwWjA8MSYwCwYDVQQLEwRwZWVy
MAoGA1UECxMDTUNDMAsGA1UECxMEcGVlcjESMBAGA1UEAxMJcGVlcjAuTUNDMFkw
EwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEnyDmcmv56gtLV1RCv5hOlHcLBLI3Y1Pp
E1XyT7/8LC2sqLwa1EL7E/l3g2V6TIQl9s9+Xkm0hJBg/uQ0gDnHXKOBxTCBwjAO
BgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUKftt/BTkMsrx
5URCspmEbtAsbeIwHwYDVR0jBBgwFoAUR26/dzqlPE03XPcL7hdB3d2osjUwYgYI
KgMEBQYHCAEEVnsiYXR0cnMiOnsiaGYuQWZmaWxpYXRpb24iOiJNQ0MucGVlciIs
ImhmLkVucm9sbG1lbnRJRCI6InBlZXIwLk1DQyIsImhmLlR5cGUiOiJwZWVyIn19
MAoGCCqGSM49BAMCA0gAMEUCIQDmsdZZ8wY7l7eiUVUB/ybXvzc6ry9UxSAk8n0H
pJpGegIgRHIxIvDzKUpVFK7hSO8dMRiaS+6iwZl25Pi7sb5Zn54=
-----END CERTIFICATE-----
`

func TestECDSAPriv_ToPem(t *testing.T) {
	var dsaObj = ECPriv{}.New(nil)
	var pemBytes = dsaObj.ToPem()
	fmt.Println(string(pemBytes))
	dsaObj = ECPriv{}.LoadPem(pemBytes)
	fmt.Println(string(dsaObj.ToPem()))
}
func TestECDSAPriv_Sign(t *testing.T) {
	var dsaObj = ECPriv{}.New(nil)
	var rawData = []byte("david secret")
	var signature = dsaObj.Sign(rawData)
	var dsaPubObj = ECPub{&dsaObj.PrivateKey.PublicKey}
	var result = dsaPubObj.Verify(rawData, signature)
	fmt.Println("is Valid", result)
}

func TestECDSAPriv_LoadPKCS8Pem(t *testing.T) {
	var pkcs8Obj = PKCS8{}.LoadPem([]byte(testPrivPem))
	var rawObj = pkcs8Obj.FormatECDSA()
	var dsaObj = ECPriv{rawObj}

	var dsapubObj = ECPub{}.LoadCert([]byte(testPubPem))
	var digest = []byte("anytest")
	var signature = dsaObj.Sign(digest)
	var isRight = dsapubObj.Verify(digest, signature)
	fmt.Println("isRight", isRight)
}
