package crypto

import (
	"crypto/rand"
	"fmt"
	. "github.com/davidkhala/goutils"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"testing"
)

func TestEthECIES(t *testing.T) {
	var dsaPriv = ECDSAPriv{}.New(nil)
	var rawData = []byte("david secret")
	var ethECPriv = ecies.ImportECDSA(dsaPriv.PrivateKey)
	var ethECPub = ecies.ImportECDSAPublic(&dsaPriv.PrivateKey.PublicKey)
	var cipher, err = ecies.Encrypt(rand.Reader, ethECPub, rawData, nil, nil)
	PanicError(err)
	recovered, err := ethECPriv.Decrypt(cipher, nil, nil)
	PanicError(err)
	fmt.Println(string(recovered))
}
