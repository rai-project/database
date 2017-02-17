package database

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
)

type ConnectionOptions struct {
	Endpoints      []string
	Username       string
	Password       string
	TLSConfig      *tls.Config
	MaxConnections int
	Context        context.Context
}

type ConnectionOption func(*ConnectionOptions)

func Username(s string) ConnectionOption {
	return func(o *ConnectionOptions) {
		o.Username = s
	}
}

func Password(s string) ConnectionOption {
	return func(o *ConnectionOptions) {
		o.Password = s
	}
}

func UsernamePassword(u string, p string) ConnectionOption {
	return func(o *ConnectionOptions) {
		o.Username = u
		o.Password = p
	}
}

func Endpoints(addrs []string) ConnectionOption {
	return func(o *ConnectionOptions) {
		o.Endpoints = addrs
	}
}

func TLSCertificate(s string) ConnectionOption {
	return func(o *ConnectionOptions) {
		var roots *x509.CertPool
		if o.TLSConfig != nil && o.TLSConfig.RootCAs != nil {
			roots = o.TLSConfig.RootCAs
		} else {
			roots = x509.NewCertPool()
		}
		cert, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			cert = []byte(s)
		}
		roots.AppendCertsFromPEM(cert)

		o.TLSConfig = &tls.Config{
			RootCAs: roots,
		}
	}
}

func TLSConfig(t *tls.Config) ConnectionOption {
	return func(o *ConnectionOptions) {
		o.TLSConfig = t
	}
}

func MaxConnections(n int) ConnectionOption {
	return func(o *ConnectionOptions) {
		o.MaxConnections = n
	}
}
