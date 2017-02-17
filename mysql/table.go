package mysql

import (
	"errors"

	db "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"

	"github.com/rai-project/database"
)

type mysqlTable struct {
	session   sqlbuilder.Database
	dbName    string
	tableName string
}

func NewTable(conn database.Connection, db database.Database, tableName string) (database.Table, error) {
	_, ok := conn.(*mysqlConnection)
	if !ok {
		return nil, errors.New("Invalid database connection input. Expecting a mysql connection instance.")
	}
	rdb, ok := db.(*mysqlDatabase)
	if !ok {
		return nil, errors.New("Invalid database input. Expecting a mysql database instance.")
	}
	return &mysqlTable{
		session:   rdb.session,
		dbName:    rdb.name,
		tableName: tableName,
	}, nil
}

func (tbl *mysqlTable) Name() string {
	return tbl.tableName
}

func (tbl *mysqlTable) Create() error {
	collection := tbl.collection()
	if collection.Exists() {
		return nil
	}
	return collection.Truncate()
}

func (tbl *mysqlTable) Delete() error {
	collection := tbl.collection()
	if !collection.Exists() {
		return nil
	}
	return collection.Truncate()
}

func (tbl *mysqlTable) Insert(elem interface{}) error {
	collection := tbl.collection()
	_, err := collection.Insert(elem)
	return err
}

func (tbl *mysqlTable) collection() db.Collection {
	return tbl.session.Collection(tbl.tableName)
}
