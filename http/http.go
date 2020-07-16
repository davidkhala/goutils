package http

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	. "github.com/davidkhala/goutils"
	"io/ioutil"
	"net/http"
)

// return the pointer to global config
func GetTLSConfigGlobal() (globalConfig *tls.Config) {
	if http.DefaultTransport.(*http.Transport).TLSClientConfig == nil {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{}
	}
	globalConfig = http.DefaultTransport.(*http.Transport).TLSClientConfig
	return
}
func SetRootCAs(config *tls.Config, rootPEMs []string) {
	roots := x509.NewCertPool()
	for _, rootPEM := range rootPEMs {
		ok := roots.AppendCertsFromPEM([]byte(rootPEM))
		if !ok {
			panic("failed to parse root certificate \n [" + rootPEM + "]")
		}
	}
	config.RootCAs = roots
}
func SetInsuredGlobal() {
	var config = GetTLSConfigGlobal()
	config.InsecureSkipVerify = true
}

type Response http.Response
type ResponseJSON struct {
	StatusCode int // e.g. 200
	Body       string
}

func (t Response) BodyBytes() []byte {
	content, err := ioutil.ReadAll(t.Body)
	PanicError(err)
	return content
}
func (t Response) Trim() ResponseJSON {
	return ResponseJSON{t.StatusCode, string(t.BodyBytes())}
}
func Get(url string, client *http.Client) Response {
	var err error
	var resp *http.Response
	if client == nil {
		resp, err = http.Get(url)
	} else {
		resp, err = client.Get(url)
	}
	PanicError(err)
	return Response(*resp)
}

func PostJson(url string, body interface{}, client *http.Client) Response {
	return Post(url, "application/json", ToJson(body), client)
}
func Post(url string, contentType string, body []byte, client *http.Client) Response {
	var err error
	var resp *http.Response
	if client == nil {
		resp, err = http.Post(url, contentType, bytes.NewBuffer(body))
	} else {
		resp, err = client.Post(url, contentType, bytes.NewBuffer(body))
	}
	PanicError(err)
	return Response(*resp)
}
func GetHttpClient(config *tls.Config) *http.Client {
	var tr = &http.Transport{
		TLSClientConfig: config,
	}
	return &http.Client{Transport: tr}

}
