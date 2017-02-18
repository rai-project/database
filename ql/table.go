package ql

import (
	"errors"

	"upper.io/db.v3/lib/sqlbuilder"

	"github.com/cznic/ql"
	"github.com/rai-project/database"
)

type qlTable struct {
	session   sqlbuilder.Database
	dbName    string
	tableName string
}

func NewTable(db database.Database, tableName string) (database.Table, error) {
	rdb, ok := db.(*qlDatabase)
	if !ok {
		return nil, errors.New("invalid database input. Expecting a ql database instance")
	}
	return &qlTable{
		session:   rdb.session,
		dbName:    rdb.databaseName,
		tableName: tableName,
	}, nil
}

func (tbl *qlTable) Name() string {
	return tbl.tableName
}

func (tbl *qlTable) Create(e interface{}) error {
	if tbl.session.Collection(tbl.tableName).Exists() {
		tbl.session.Exec("DROP TABLE IF EXISTS " + tbl.tableName)
	}
	schema, err := ql.Schema(e, tbl.tableName, nil)
	if err != nil {
		return err
	}
	_, err = tbl.session.Exec(schema)
	return err
}

func (tbl *qlTable) Delete() error {
	return tbl.session.Collection(tbl.tableName).Truncate()
}

func (tbl *qlTable) Insert(elem interface{}) error {
	_, err := tbl.session.InsertInto(tbl.tableName).Values(elem).Exec()
	return err
}
