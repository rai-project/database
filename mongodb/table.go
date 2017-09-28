package mongodb

import (
	"errors"

	"github.com/rai-project/database"
	"upper.io/db.v3/lib/sqlbuilder"
)

type mongoTable struct {
	session   sqlbuilder.Database
	dbName    string
	tableName string
}

// NewTable ...
func NewTable(db database.Database, tableName string) (database.Table, error) {
	rdb, ok := db.(*mongoDatabase)
	if !ok {
		return nil, errors.New("invalid database input. Expecting a mongodb database instance")
	}
	return &mongoTable{
		session:   rdb.session,
		dbName:    rdb.databaseName,
		tableName: tableName,
	}, nil
}

// Name ...
func (tbl *mongoTable) Name() string {
	return tbl.tableName
}

// Create ...
func (tbl *mongoTable) Create(e interface{}) error {
	err := tbl.session.Collection(tbl.tableName).Truncate()
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func (tbl *mongoTable) Delete() error {
	return tbl.session.Collection(tbl.tableName).Truncate()
}

// Insert ...
func (tbl *mongoTable) Insert(elem interface{}) error {
	_, err := tbl.session.InsertInto(tbl.tableName).Values(elem).Exec()
	return err
}
