package mysql

import (
	"context"

	_ "github.com/jinzhu/gorm/dialects/mysql" // include mysql dialect
	"github.com/rai-project/database"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

const (
	gormDialect = "myqsl"
)

type mysqlConnection struct {
	conn     sqlbuilder.Database
	opts     database.ConnectionOptions
	settings *mysql.ConnectionURL
}

func NewConnection(opts ...database.ConnectionOption) (database.Connection, error) {

	options := database.ConnectionOptions{
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

	settings := &mysql.ConnectionURL{
		User:     options.Username,
		Password: options.Password,
		Host:     options.Endpoints[0],
		Database: "",
	}

	return &mysqlConnection{
		conn:     nil,
		opts:     options,
		settings: settings,
	}, nil
}

func (conn *mysqlConnection) Options() database.ConnectionOptions {
	return conn.opts
}

func (conn *mysqlConnection) Close() error {
	if conn.conn == nil {
		return nil
	}
	return conn.conn.Close()
}

func (conn *mysqlConnection) String() string {
	return "RethinkDB"
}
