package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	. "github.com/davidkhala/goutils"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"testing"
)

func TestEthECIES(t *testing.T) {
	var dsaPriv = ECPriv{}.New(nil)
	var rawData = []byte("david secret")
	var ethECPriv = ecies.ImportECDSA(dsaPriv.PrivateKey)
	var ethECPub = ecies.ImportECDSAPublic(&dsaPriv.PrivateKey.PublicKey)
	var cipher, err = ecies.Encrypt(rand.Reader, ethECPub, rawData, nil, nil)
	PanicError(err)
	recovered, err := ethECPriv.Decrypt(cipher, nil, nil)
	PanicError(err)
	fmt.Println(string(recovered))
}
func TestECPub_Encrypt(t *testing.T) {
	var dsaPriv = ECPriv{}.New(nil)
	var rawData = []byte("david secret")
	var ecPub = ECPub{&dsaPriv.PublicKey}
	var cipher2 = ecPub.Encrypt(rand.Reader, rawData)
	fmt.Println("loc cipher", hex.EncodeToString(cipher2))
	fmt.Println("loc cipher,byte len", len(cipher2))
	var recovered = dsaPriv.Decrypt(cipher2)
	fmt.Println(string(recovered))
}
