package database

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
)

// DefaultMaxConnections ...
var (
	DefaultMaxConnections = 30
)

// Options ...
type Options struct {
	Endpoints      []string
	Username       string
	Password       string
	TLSConfig      *tls.Config
	MaxConnections int
	Context        context.Context
}

// Option ...
type Option func(*Options)

// Username ...
func Username(s string) Option {
	return func(o *Options) {
		o.Username = s
	}
}

// Password ...
func Password(s string) Option {
	return func(o *Options) {
		o.Password = s
	}
}

// UsernamePassword ...
func UsernamePassword(u string, p string) Option {
	return func(o *Options) {
		o.Username = u
		o.Password = p
	}
}

// Endpoints ...
func Endpoints(addrs []string) Option {
	return func(o *Options) {
		o.Endpoints = addrs
	}
}

// TLSCertificate ...
func TLSCertificate(s string) Option {
	return func(o *Options) {
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

// TLSConfig ...
func TLSConfig(t *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = t
	}
}

// MaxConnections ...
func MaxConnections(n int) Option {
	return func(o *Options) {
		o.MaxConnections = n
	}
}
