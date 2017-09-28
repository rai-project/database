package mongodb

import (
	"context"

	"github.com/rai-project/database"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mongo"
	"upper.io/db.v3/ql"
)

type mongoDatabase struct {
	session      sqlbuilder.Database
	databaseName string
	opts         database.Options
}

// NewDatabase ...
func NewDatabase(databaseName string, opts ...database.Option) (database.Database, error) {
	options := database.Options{
		Endpoints:      Config.Endpoints,
		Username:       Config.Username,
		Password:       Config.Password,
		TLSConfig:      nil,
		MaxConnections: Config.MaxConnections,
		Context:        context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	connectionURL := mongo.ConnectionURL{
		User:     options.Username,
		Password: options.Password,
		Host:     options.Endpoints[0],
		Database: databaseName,
	}

	sess, err := ql.Open(connectionURL)
	if err != nil {
		return nil, err
	}
	return &mongoDatabase{
		session:      sess,
		databaseName: databaseName,
		opts:         options,
	}, nil
}

// Session ...
func (conn *mongoDatabase) Session() interface{} {
	return conn.session
}

// Options ...
func (conn *mongoDatabase) Options() database.Options {
	return conn.opts
}

// Close ...
func (conn *mongoDatabase) Close() error {
	return conn.session.Close()
}

// String ...
func (conn *mongoDatabase) String() string {
	return "ql"
}
