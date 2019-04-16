package mongodb

import (
	"time"

	"github.com/pkg/errors"

	"github.com/cenkalti/backoff"
	"github.com/rai-project/database"
	"upper.io/db.v3"
)

type MongoTable struct {
	Session   db.Database
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
		Session:   rdb.session,
		dbName:    rdb.databaseName,
		tableName: tableName,
	}, nil
}

// Name ...
func (tbl *MongoTable) Name() string {
	return tbl.tableName
}

func (tbl *MongoTable) Exists() bool {
	return tbl.Session.Collection(tbl.tableName).Exists()
}

// Create ...
func (tbl *MongoTable) Create(e interface{}) error {
	//if tbl.Exists() {
  if true {
		return nil
	}
	err := tbl.Session.Collection(tbl.tableName).Truncate()
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func (tbl *MongoTable) Delete() error {
	return tbl.Session.Collection(tbl.tableName).Truncate()
}

// insert ...
func (tbl *MongoTable) insert(elem interface{}) error {
	_, err := tbl.Session.Collection(tbl.tableName).Insert(elem)
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
	alg.MaxElapsedTime = time.Minute
	return backoff.Retry(insert, alg)
}

func (tbl *MongoTable) Find(q interface{}, skip int, limit int, result interface{}) (err error) {

	collection := tbl.Session.Collection(tbl.tableName)

	var res db.Result
	if q == nil {
		res = collection.Find()
	} else {
		res = collection.Find(q)
	}

	if skip != 0 {
		res = res.Offset(skip)
	}

	if limit > 0 {
		res = res.Limit(limit)
	}
	return res.All(result)
}

func (tbl *MongoTable) FindOne(q interface{}, result interface{}) error {
	collection := tbl.Session.Collection(tbl.tableName)
	return collection.Find(q).One(result)
}

func (tbl *MongoTable) FindAll(q interface{}, result interface{}) error {
	return tbl.Find(q, 0, -1, result)
}
