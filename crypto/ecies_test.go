package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestECPub_Encrypt(t *testing.T) {
	var dsaPriv = NewECPriv(nil)
	var rawData = []byte("david secret")
	var ecPub = ECPub{&dsaPriv.PublicKey}
	var cipher2 = ecPub.Encrypt(rand.Reader, rawData)
	fmt.Println("loc cipher", hex.EncodeToString(cipher2))
	fmt.Println("loc cipher,byte len", len(cipher2))
	var recovered = dsaPriv.Decrypt(cipher2)
	fmt.Println(string(recovered))
}
