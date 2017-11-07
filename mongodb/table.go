package mongodb

import (
	"errors"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/rai-project/database"
	"upper.io/db.v3"
)

type MongoTable struct {
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
	return &MongoTable{
		session:   rdb.session,
		dbName:    rdb.databaseName,
		tableName: tableName,
	}, nil
}

// Name ...
func (tbl *MongoTable) Name() string {
	return tbl.tableName
}

func (tbl *MongoTable) Exists() bool {
	return tbl.session.Collection(tbl.tableName).Exists()
}

// Create ...
func (tbl *MongoTable) Create(e interface{}) error {
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
func (tbl *MongoTable) Delete() error {
	return tbl.session.Collection(tbl.tableName).Truncate()
}

// insert ...
func (tbl *MongoTable) insert(elem interface{}) error {
	_, err := tbl.session.Collection(tbl.tableName).Insert(elem)
	return err
}

// Insert ...
func (tbl *MongoTable) Insert(elem interface{}) error {
	insert := func() error {
		return tbl.insert(elem)
	}
	alg := backoff.NewExponentialBackOff()
	alg.InitialInterval = 10 * time.Millisecond
	alg.Multiplier = 1.2
	alg.MaxElapsedTime = 5 * time.Minute
	return backoff.Retry(insert, alg)
}

func (tbl *MongoTable) Find(q interface{}, skip int, limit int, result interface{}) (err error) {

	collection := tbl.session.Collection(tbl.tableName)

	if limit < 0 {
		return collection.Find(q).Offset(skip).All(result)
	}
	return collection.Find(q).Offset(skip).Limit(limit).All(result)
}

func (ds *MongoTable) FindAll(q interface{}, result interface{}) error {
	return ds.Find(q, 0, -1, result)
}
