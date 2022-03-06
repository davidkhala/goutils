package goutils

import (
	"fmt"
	"testing"
	"time"
)

func TestDirname(t *testing.T) {
	dir, err := Dirname()
	PanicError(err)
	println(dir)
	fileName, err := Filename()
	PanicError(err)
	println(fileName)
}

func TestFloat(t *testing.T) {
	var amountStr = "123456789.12"
	var amountFloat = ParseFloat(amountStr)
	fmt.Println(amountFloat, FormatFloat(amountFloat, 2))

}

var randomBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func TestRandString(t *testing.T) {
	var result = RandString(12, randomBytes)
	fmt.Println(result, len(result))
}

func TestInt2Float(t *testing.T) {
	var amount = "1"
	var amountFloat float64
	amountFloat = float64(Atoi(amount))
	fmt.Println(amountFloat)
}
func TestIntToByte(t *testing.T) {
	var i int64 = 1
	fmt.Println(byte(i))
	i = 256
	fmt.Println(byte(i))
	i = -1
	fmt.Println(byte(i))
}
func TestItoRunes(t *testing.T) {
	const charSpace = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var runes = []rune(charSpace)
	var result = ItoRunes(50, runes)
	fmt.Println(result)
	result = ItoRunes(51, runes)
	fmt.Println(result)
}

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("20060102")) //correct
}

func TestPaddingLeft(t *testing.T) {
	var str = "123"
	var result = PaddingLeft(str, 7, '0')
	fmt.Println(result)
}
func TestHexEncoding(t *testing.T) {
	var cert = `-----BEGIN CERTIFICATE-----
MIICFzCCAb2gAwIBAgIUS8lAQ16ZG6iTQD7H4E/UQ9d2YrwwCgYIKoZIzj0EAwIw
aDELMAkGA1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMRQwEgYDVQQK
EwtIeXBlcmxlZGdlcjEPMA0GA1UECxMGRmFicmljMRkwFwYDVQQDExBmYWJyaWMt
Y2Etc2VydmVyMB4XDTIyMDMwMzA4MjMwMFoXDTM3MDIyNzA4MjMwMFowaDELMAkG
A1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMRQwEgYDVQQKEwtIeXBl
cmxlZGdlcjEPMA0GA1UECxMGRmFicmljMRkwFwYDVQQDExBmYWJyaWMtY2Etc2Vy
dmVyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE0qy6fs9TWREZ/vspZMSgjK2X
lHMDcTAikBLmjpp63zxzbNkYYIWZrVlrmpdtV5XWlRIMbDkY+1c/lCStuVT7KaNF
MEMwDgYDVR0PAQH/BAQDAgEGMBIGA1UdEwEB/wQIMAYBAf8CAQEwHQYDVR0OBBYE
FB6QJa7MCqssjbnQ7Ral4vUGpsVvMAoGCCqGSM49BAMCA0gAMEUCIQDWb+GO0rZ8
vLbOgtIOBwbIcK13Gi2yMb0AIM5ropJTygIgBoOyOOrXcboyjfAiidNvfNTClpSD
4DWMi2X5N0Z5S8k=
-----END CERTIFICATE-----
`
	var encoded = HexEncode([]byte(cert))
	println(encoded)
	var decoded = HexDecodeOrPanic(encoded)
	if cert != string(decoded) {
		panic("encode failure")
	}
}
