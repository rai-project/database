package mongodb

import (
	"errors"

	"github.com/cenkalti/backoff"
	"github.com/rai-project/database"
	"upper.io/db.v3"
)

type mongoTable struct {
	session   db.Database
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

func (tbl *mongoTable) Exists() bool {
	return tbl.session.Collection(tbl.tableName).Exists()
}

// Create ...
func (tbl *mongoTable) Create(e interface{}) error {
	if tbl.Exists() {
		return nil
	}
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

// insert ...
func (tbl *mongoTable) insert(elem interface{}) error {
	_, err := tbl.session.Collection(tbl.tableName).Insert(elem)
	return err
}

// Insert ...
func (tbl *mongoTable) Insert(elem interface{}) error {
	insert := func() error {
		return tbl.insert(elem)
	}
	return backoff.Retry(insert, backoff.NewExponentialBackOff())
}
