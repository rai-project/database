package rethinkdb

import (
	"context"

	"github.com/rai-project/database"
)

const (
	authKeyKey         = "github.com/rai-project/database/rethinkdb/authKey"
	initialCapacityKey = "github.com/rai-project/database/rethinkdb/initialCapacity"
)

// DefaultInitialCapacity ...
var (
	DefaultInitialCapacity = 10
)

// AuthKey ...
func AuthKey(s string) database.Option {
	return func(o *database.Options) {
		o.Context = context.WithValue(o.Context, authKeyKey, s)
	}
}

// InitialCapacity ...
func InitialCapacity(n int) database.Option {
	return func(o *database.Options) {
		o.Context = context.WithValue(o.Context, initialCapacityKey, n)
	}
}
