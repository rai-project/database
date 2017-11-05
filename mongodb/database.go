package mongodb

import (
	"context"
	"strings"
	"time"

	"github.com/rai-project/config"
	"github.com/rai-project/database"
	"github.com/rai-project/utils"
	"upper.io/db.v3"
	"upper.io/db.v3/mongo"
)

type mongoDatabase struct {
	session      db.Database
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

	decrypt := func(s string) string {
		if strings.HasPrefix(s, utils.CryptoHeader) && config.App.Secret != "" {
			if val, err := utils.DecryptStringBase64(config.App.Secret, s); err == nil {
				return val
			}
		}
		return s
	}

	connectionURL := mongo.ConnectionURL{
		User:     options.Username,
		Password: decrypt(options.Password),
		Host:     options.Endpoints[0],
		Database: databaseName,
	}

	sess := &mongo.Source{
		Settings: db.NewSettings(),
	}
	sess.Settings.SetConnMaxLifetime(5 * time.Hour)
	sess.Settings.SetMaxIdleConns(options.MaxConnections)
	sess.Settings.SetMaxOpenConns(options.MaxConnections)

	err := sess.Open(connectionURL)
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
	return "mongodb"
}

func (s *mongoDatabase) WithCollection(collection string, f func(db.Collection) error) error {
	c := s.session.Collection(collection)
	return f(c)
}
