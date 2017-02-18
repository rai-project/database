package postgres

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rai-project/database"
	"upper.io/db.v3/postgresql"
)

const (
	gormDialect = "postgres"
)

type postgresDatabase struct {
	conn         *gorm.DB
	databaseName string
	opts         database.Options
	settings     postgresql.ConnectionURL
}

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

	settings := postgresql.ConnectionURL{
		User:     options.Username,
		Password: options.Password,
		Host:     options.Endpoints[0],
		Database: databaseName,
	}

	c, err := gorm.Open(gormDialect, settings.String())
	if err != nil {
		return nil, err
	}

	maxConnections := options.MaxConnections

	c.DB().SetMaxIdleConns(maxConnections)

	return &postgresDatabase{
		conn:         c,
		databaseName: databaseName,
		opts:         options,
		settings:     settings,
	}, nil
}

func (conn *postgresDatabase) Options() database.Options {
	return conn.opts
}

func (conn *postgresDatabase) Close() error {
	return conn.conn.Close()
}

func (conn *postgresDatabase) String() string {
	return "PostgreSQL"
}
