package crypto

import (
	"crypto"
	"fmt"
	"github.com/davidkhala/goutils"
	"testing"
)

func TestHMAC(t *testing.T) {
	var message = []byte("abc")
	var key = []byte("fabric")
	var hmacMessage = NewHMAC(crypto.SHA512.New, message, key)
	var hexHMACMessage =goutils.HexEncode(hmacMessage)
	fmt.Println(len(hexHMACMessage), hexHMACMessage)

	var isValid = CheckHMAC(crypto.SHA512.New, message, hmacMessage, key)
	fmt.Println("isValid", isValid)

}
