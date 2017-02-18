package ql

import (
	"context"
	"path/filepath"

	"github.com/rai-project/database"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/ql"
)

type qlDatabase struct {
	session      sqlbuilder.Database
	databaseName string
	opts         database.Options
}

func NewDatabase(databaseName string, opts ...database.Option) (database.Database, error) {
	options := database.Options{
		Endpoints:      Config.Endpoints,
		TLSConfig:      nil,
		MaxConnections: Config.MaxConnections,
		Context:        context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	connectionURL := ql.ConnectionURL{
		Database: filepath.Join(options.Endpoints[0], databaseName),
	}

	sess, err := ql.Open(connectionURL)
	if err != nil {
		return nil, err
	}
	return &qlDatabase{
		session:      sess,
		databaseName: databaseName,
		opts:         options,
	}, nil
}

func (conn *qlDatabase) Options() database.Options {
	return conn.opts
}

func (conn *qlDatabase) Close() error {
	return conn.session.Close()
}

func (conn *qlDatabase) String() string {
	return "ql"
}
