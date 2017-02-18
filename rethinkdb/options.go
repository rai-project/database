package rethinkdb

import (
	"context"

	"github.com/rai-project/database"
)

const (
	authKeyKey         = "github.com/rai-project/database/rethinkdb/authKey"
	initialCapacityKey = "github.com/rai-project/database/rethinkdb/initialCapacity"
)

var (
	DefaultInitialCapacity = 10
)

func AuthKey(s string) database.Option {
	return func(o *database.Options) {
		o.Context = context.WithValue(o.Context, authKeyKey, s)
	}
}

func InitialCapacity(n int) database.Option {
	return func(o *database.Options) {
		o.Context = context.WithValue(o.Context, initialCapacityKey, n)
	}
}
