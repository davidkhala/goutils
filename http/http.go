package http

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	. "github.com/davidkhala/goutils"
	"io/ioutil"
	"net/http"
)

func BeginTLSConfig() (config *tls.Config) {
	if http.DefaultTransport.(*http.Transport).TLSClientConfig == nil {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{}
	}
	config = http.DefaultTransport.(*http.Transport).TLSClientConfig
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
	var config = BeginTLSConfig()
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
func InsecuredClient() http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return http.Client{Transport: tr}

}
