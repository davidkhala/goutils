package grpc

import (
	"crypto/tls"
	"crypto/x509"
)

// TLSConfigFrom returns the appropriate config for TLS including the root CAs, certs for mutual TLS, and server host override.
func TLSConfigFrom(clientTLSCert *tls.Certificate, serverName string, rootCAs ...*x509.Certificate) (*tls.Config, error) {
	var err error
	var certPool *x509.CertPool
	var clientCertificates []tls.Certificate
	certPool, err = x509.SystemCertPool()
	if err != nil {
		return nil, err
	}
	for _, rootCA := range rootCAs {
		certPool.AddCert(rootCA)
	}

	if clientTLSCert != nil {
		clientCertificates = append(clientCertificates, *clientTLSCert)
	}
	return &tls.Config{RootCAs: certPool, Certificates: clientCertificates, ServerName: serverName}, nil
}
