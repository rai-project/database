package rethinkdb

import (
	"context"

	"github.com/rai-project/database"
)

const (
	authKeyKey         = "github.com/rai-project/database/rethinkdb/authKey"
	initialCapacityKey = "github.com/rai-project/database/rethinkdb/initialCapacity"
	maxConnectionsKey  = "github.com/rai-project/broker/sqs/maxConnections"
)

var (
	DefaultInitialCapacity = 10
	DefaultMaxConnections  = 30
)

func AuthKey(s string) database.ConnectionOption {
	return func(o *database.ConnectionOptions) {
		o.Context = context.WithValue(o.Context, authKeyKey, s)
	}
}

func InitialCapacity(n int) database.ConnectionOption {
	return func(o *database.ConnectionOptions) {
		o.Context = context.WithValue(o.Context, initialCapacityKey, n)
	}
}

func MaxConnections(n int) database.ConnectionOption {
	return func(o *database.ConnectionOptions) {
		o.Context = context.WithValue(o.Context, maxConnectionsKey, n)
	}
}
