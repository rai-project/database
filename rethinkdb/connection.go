package rethinkdb

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rai-project/database"
	r "gopkg.in/dancannon/gorethink.v3"
)

type rethinkConnection struct {
	session *r.Session
	opts    database.ConnectionOptions
}

func NewConnection(opts ...database.ConnectionOption) (database.Connection, error) {
	pass := Config.Password
	if pass == "" {
		pass = Config.AuthKey
	}

	options := database.ConnectionOptions{
		Endpoints: Config.Endpoints,
		Username:  Config.Username,
		Password:  pass,
		TLSConfig: nil,
		Context:   context.Background(),
	}

	if Config.Cert != "" {
		opts = append(
			[]database.ConnectionOption{
				database.TLSCertificate(Config.Cert),
			},
			opts...,
		)
	}

	for _, o := range opts {
		o(&options)
	}

	authKey, ok := options.Context.Value(authKeyKey).(string)
	if !ok {
		authKey = options.Password
	}

	initialCapacity, ok := options.Context.Value(initialCapacityKey).(int)
	if !ok {
		initialCapacity = DefaultInitialCapacity
	}

	maxConnections, ok := options.Context.Value(maxConnectionsKey).(int)
	if !ok {
		maxConnections = DefaultMaxConnections
	}

	sess, err := r.Connect(r.ConnectOpts{
		Address:    options.Endpoints[0],
		Username:   options.Username,
		Password:   options.Password,
		AuthKey:    authKey,
		TLSConfig:  options.TLSConfig,
		InitialCap: initialCapacity,
		MaxOpen:    maxConnections,
	})
	if err != nil {
		log.WithError(err).Error("Failed to connect to rethinkdb database")
		return nil, errors.Wrap(err, "Failed to connect to rethinkdb database")
	}

	return &rethinkConnection{
		session: sess,
		opts:    options,
	}, nil
}

func (conn *rethinkConnection) Options() database.ConnectionOptions {
	return conn.opts
}

func (conn *rethinkConnection) Close() error {
	return conn.session.Close()
}
