package crypto

import (
	"crypto/sha256"
	"fmt"
	. "github.com/davidkhala/goutils"
	"testing"
)

func TestSha256(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("hello world"))
	fmt.Printf("%x \n", h.Sum(nil))
	fmt.Printf("%x \n","prefix")
	fmt.Printf("%x \n", h.Sum([]byte("prefix:")))
	fmt.Println()
	//simple hash
	{
		sum := sha256.Sum256([]byte("hello world"))
		fmt.Printf("%x \n", sum)
		fmt.Println(HexEncode(sum[:]))
	}

}
