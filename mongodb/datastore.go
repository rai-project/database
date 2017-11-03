package mongodb

import (
	db "upper.io/db.v3"
)

type DataStore struct {
	db         *mongoDatabase
	collection string
}

func NewDataStore(db *mongoDatabase, collection string) *DataStore {
	return &DataStore{
		db:         db,
		collection: collection,
	}
}

func (ds *DataStore) Insert(model interface{}) (err error) {
	query := func(c db.Collection) error {
		_, err := c.Insert(model)
		return err
	}

	create := func() error {
		return ds.db.WithCollection(ds.collection, query)
	}
	err = create()
	return
}

func (ds *DataStore) Find(q interface{}, skip int, limit int, result interface{}) (err error) {
	query := func(c db.Collection) error {
		fn := c.Find(q).Offset(skip).Limit(limit).All(result)
		if limit < 0 {
			fn = c.Find(q).Offset(skip).All(result)
		}
		return fn
	}
	find := func() error {
		return ds.db.WithCollection(ds.collection, query)
	}
	err = find()

	return
}

func (ds *DataStore) FindAll(q interface{}, result interface{}) error {
	return ds.Find(q, 0, -1, result)
}
