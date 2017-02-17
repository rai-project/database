package rethinkdb

import (
	"errors"

	"github.com/rai-project/database"
	r "gopkg.in/dancannon/gorethink.v3"
)

type rethinkDatabase struct {
	session *r.Session
	name    string
}

func NewDatabase(conn database.Connection, name string) (database.Database, error) {
	sess, ok := conn.(*rethinkConnection)
	if !ok {
		return nil, errors.New("Invalid database connection input. Expecting a rethinkdb connection instance.")
	}
	return &rethinkDatabase{
		session: sess.session,
		name:    name,
	}, nil
}

func (db *rethinkDatabase) Name() string {
	return db.name
}

func (db *rethinkDatabase) Create() error {
	return r.DBCreate(db.name).Exec(db.session)
}

func (db *rethinkDatabase) Delete() error {
	return r.DBDrop(db.name).Exec(db.session)
}
