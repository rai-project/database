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

func NewTable(conn database.Connection, db database.Database, tableName string) (database.Table, error) {
	rconn, ok := conn.(*rethinkConnection)
	if !ok {
		return nil, errors.New("Invalid database connection input. Expecting a rethinkdb connection instance.")
	}
	rdb, ok := db.(*rethinkDatabase)
	if !ok {
		return nil, errors.New("Invalid database input. Expecting a rethinkdb database instance.")
	}
	return &rethinkTable{
		session:   rconn.session,
		dbName:    rdb.name,
		tableName: tableName,
	}, nil
}

func (tbl *rethinkTable) Name() string {
	return tbl.tableName
}

func (tbl *rethinkTable) Create() error {
	return r.DB(tbl.dbName).TableCreate(tbl.tableName).Exec(tbl.session)
}

func (tbl *rethinkTable) Delete() error {
	return r.DB(tbl.dbName).TableDrop(tbl.tableName).Exec(tbl.session)
}

func (tbl *rethinkTable) Insert(elem interface{}) error {
	_, err := r.DB(tbl.dbName).Table(tbl.tableName).Insert(elem).RunWrite(tbl.session)
	return err
}
