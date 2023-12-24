package crypto

import (
	"fmt"
	"github.com/davidkhala/goutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var certFilePath = goutils.Absolute("../testdata/tlsca.hyperledger-cert.pem")

func TestCertificate(t *testing.T) {
	var pemBytes = goutils.ReadFileOrPanic(certFilePath)
	var cert = ParseCertPemOrPanic(pemBytes)
	t.Run("GetDN", func(t *testing.T) {

		fmt.Println("issuer: ", GetDN(cert.Issuer))
		fmt.Println("subject: ", GetDN(cert.Subject))
	})
	t.Run("write Bytes", func(t *testing.T) {
		var outBytes = ToCertPem(cert)
		assert.Equal(t, pemBytes, outBytes)
	})
}
