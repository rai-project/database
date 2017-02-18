package rethinkdb

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rai-project/database"
	r "gopkg.in/dancannon/gorethink.v3"
)

type rethinkDatabase struct {
	session      *r.Session
	databaseName string
	opts         database.Options
}

func NewDatabase(databaseName string, opts ...database.Option) (database.Database, error) {
	pass := Config.Password
	if pass == "" {
		pass = Config.AuthKey
	}

	options := database.Options{
		Endpoints:      Config.Endpoints,
		Username:       Config.Username,
		Password:       pass,
		TLSConfig:      nil,
		MaxConnections: Config.MaxConnections,
		Context:        context.Background(),
	}

	if Config.Cert != "" {
		opts = append(
			[]database.Option{
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

	maxConnections := options.MaxConnections

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

	r.DBCreate(databaseName).Exec(sess)

	return &rethinkDatabase{
		session:      sess,
		databaseName: databaseName,
		opts:         options,
	}, nil
}

func (conn *rethinkDatabase) Options() database.Options {
	return conn.opts
}

func (conn *rethinkDatabase) Close() error {
	return conn.session.Close()
}

func (conn *rethinkDatabase) String() string {
	return "RethinkDB"
}
