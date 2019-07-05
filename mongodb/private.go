package mongodb

import (
	"sync"
	"time"

	"gopkg.in/mgo.v2"
	"upper.io/db.v3"
	"upper.io/db.v3/mongo"
)

// this needs to mirror the source in db.v3
type source struct {
	db.Settings

	name          string
	connURL       db.ConnectionURL
	session       *mgo.Session
	database      *mgo.Database
	version       []int
	collections   map[string]*mongo.Collection
	collectionsMu sync.Mutex
}

// Open attempts to connect to the database.
func (s *source) Open(connURL db.ConnectionURL, connTimeout time.Duration) error {
	s.connURL = connURL
	return s.open(connTimeout)
}

func (s *source) open(connTimeout time.Duration) error {
	var err error
	s.session, err = mgo.DialWithTimeout(s.connURL.String(), connTimeout)
	if err != nil {
		return err
	}

	s.collections = map[string]*mongo.Collection{}
	s.database = s.session.DB("")

	return nil
}
