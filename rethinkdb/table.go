package rethinkdb

import (
	"errors"

	"github.com/rai-project/database"
	r "gopkg.in/dancannon/gorethink.v3"
)

type rethinkTable struct {
	session   *r.Session
	dbName    string
	tableName string
}

// NewTable ...
func NewTable(db database.Database, tableName string) (database.Table, error) {
	rdb, ok := db.(*rethinkDatabase)
	if !ok {
		return nil, errors.New("Invalid database input. Expecting a rethinkdb database instance.")
	}
	return &rethinkTable{
		session:   rdb.session,
		dbName:    rdb.databaseName,
		tableName: tableName,
	}, nil
}

// Name ...
func (tbl *rethinkTable) Name() string {
	return tbl.tableName
}

// Create ...
func (tbl *rethinkTable) Create(interface{}) error {
	return r.DB(tbl.dbName).TableCreate(tbl.tableName).Exec(tbl.session)
}

// Delete ...
func (tbl *rethinkTable) Delete() error {
	return r.DB(tbl.dbName).TableDrop(tbl.tableName).Exec(tbl.session)
}

// Insert ...
func (tbl *rethinkTable) Insert(elem interface{}) error {
	_, err := r.DB(tbl.dbName).Table(tbl.tableName).Insert(elem).RunWrite(tbl.session)
	return err
}
