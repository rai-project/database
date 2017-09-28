package relational

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/rai-project/database"
	db "upper.io/db.v3"
)

type relationalDatabase struct {
	conn          *gorm.DB
	databaseName  string
	opts          database.Options
	gormDialect   string
	connectionURL db.ConnectionURL
}

// NewDatabase ...
func NewDatabase(gormDialect string, databaseName string, connectionURL db.ConnectionURL, opts database.Options) (database.Database, error) {
	c, err := gorm.Open(gormDialect, connectionURL.String())
	if err != nil {
		return nil, err
	}

	maxConnections := opts.MaxConnections
	c.DB().SetMaxIdleConns(maxConnections)

	return &relationalDatabase{
		conn:          c,
		databaseName:  databaseName,
		opts:          opts,
		gormDialect:   gormDialect,
		connectionURL: connectionURL,
	}, nil
}

// Session ...
func (conn *relationalDatabase) Session() interface{} {
	return conn.conn
}

// Options ...
func (conn *relationalDatabase) Options() database.Options {
	return conn.opts
}

// Close ...
func (conn *relationalDatabase) Close() error {
	return conn.conn.Close()
}

// String ...
func (conn *relationalDatabase) String() string {
	return strings.Title(conn.gormDialect)
}
