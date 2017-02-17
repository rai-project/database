package relational

import (
	"context"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/rai-project/database"
)

type relationalDBConnection struct {
	dialect string
	*gorm.DB
	url    string
	dbName string
	opts   database.ConnectionOptions
}

func NewConnection(dialect string, url string, dbName string, opts ...database.ConnectionOption) (database.Connection, error) {
	db, err := gorm.Open(dialect, url)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to connect to %v database.", dialect)
	}

	options := database.ConnectionOptions{
		Context: context.Background(),
	}
	for _, o := range opts {
		o(&options)
	}

	db.DB().SetMaxIdleConns(options.MaxConnections)

	return &relationalDBConnection{
		dialect: dialect,
		DB:      db,
		url:     url,
		opts:    options,
	}, nil
}

func (conn *relationalDBConnection) Options() database.ConnectionOptions {
	return conn.opts
}

func (conn *relationalDBConnection) Close() error {
	return conn.Close()
}

func (conn *relationalDBConnection) Name() string {
	return conn.dbName
}

func (conn *relationalDBConnection) Create() error {
	panic("Relational database create has not been implemented....")
	return nil
}

func (conn *relationalDBConnection) Delete() error {
	panic("Relational database delete has not been implemented....")
	return nil
}

func (conn *relationalDBConnection) String() string {
	return strings.Title(conn.dialect)
}
