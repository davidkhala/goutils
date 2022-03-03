package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

const (
	// TODO make it to context
	// GRPC max message size (same as Fabric)
	maxCallRecvMsgSize = 100 * 1024 * 1024
	maxCallSendMsgSize = 100 * 1024 * 1024
)

type Params struct {
	SslTargetNameOverride string            // "ssl-target-name-override"
	Certificate           *x509.Certificate // tlscacert
	ClientCertificate     *tls.Certificate  // for mutual tls
	KeepAliveParams       *keepalive.ClientParameters
	WaitForReady          bool
}

func DialOptionsFrom(params Params) ([]grpc.DialOption, error) {
	var dialOpts []grpc.DialOption

	if params.KeepAliveParams != nil {
		dialOpts = append(dialOpts, grpc.WithKeepaliveParams(*params.KeepAliveParams))
	}

	dialOpts = append(dialOpts, grpc.WithDefaultCallOptions(grpc.WaitForReady(params.WaitForReady)))

	if params.Certificate != nil {
		tlsConfig, err := TLSConfigFrom(params.ClientCertificate, params.SslTargetNameOverride, params.Certificate)
		if err != nil {
			return nil, err
		}

		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	} else {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}

	dialOpts = append(dialOpts, grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(maxCallRecvMsgSize),
		grpc.MaxCallSendMsgSize(maxCallSendMsgSize),
	))

	return dialOpts, nil
}
