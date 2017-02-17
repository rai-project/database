package mysql

import (
	"errors"

	"github.com/rai-project/database"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

type mysqlDatabase struct {
	session sqlbuilder.Database
	name    string
}

func NewDatabase(conn database.Connection, name string) (database.Database, error) {
	sess, ok := conn.(*mysqlConnection)
	if !ok {
		return nil, errors.New("Invalid database connection input. Expecting a mysql connection instance.")
	}

	sess.settings.Database = name
	c, err := mysql.Open(*sess.settings)
	if err != nil {
		return nil, err
	}

	maxConnections := sess.opts.MaxConnections

	c.SetMaxOpenConns(maxConnections)

	sess.conn = c

	return &mysqlDatabase{
		session: c,
		name:    name,
	}, nil
}

func (db *mysqlDatabase) Name() string {
	return db.name
}

func (db *mysqlDatabase) Create() error {
	return nil
}

func (db *mysqlDatabase) Delete() error {
	return nil
}
