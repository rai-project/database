package relational

import (
	"context"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"    // include mysql dialect
	_ "github.com/jinzhu/gorm/dialects/postgres" // include postgres dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // include postgres sqlite
	"github.com/pkg/errors"
	"github.com/rai-project/database"
)

type relationalDatabase struct {
	dialect string
	*gorm.DB
	url    string
	dbName string
	opts   database.ConnectionOptions
}

func NewDatabase(dialect string, url string, dbName string, opts ...database.ConnectionOption) (database.Connection, error) {
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

	return &relationalDatabase{
		dialect: dialect,
		DB:      db,
		url:     url,
		opts:    options,
	}, nil
}

func (conn *relationalDatabase) Options() database.ConnectionOptions {
	return conn.opts
}

func (conn *relationalDatabase) Close() error {
	return conn.Close()
}

func (conn *relationalDatabase) String() string {
	return strings.Title(conn.dialect)
}
