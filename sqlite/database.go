package sqlite

import (
	"context"
	"path/filepath"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rai-project/database"
	"github.com/rai-project/database/relational"
	"upper.io/db.v3/sqlite"
)

const (
	gormDialect = "sqlite"
)

type sqliteDatabase struct {
	database.Database
}

// NewDatabase ...
func NewDatabase(databaseName string, opts ...database.Option) (database.Database, error) {

	options := database.Options{
		Endpoints:      Config.Endpoints,
		TLSConfig:      nil,
		MaxConnections: Config.MaxConnections,
		Context:        context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	connectionURL := sqlite.ConnectionURL{
		Database: filepath.Join(options.Endpoints[0], databaseName),
	}

	d, err := relational.NewDatabase(gormDialect, databaseName, connectionURL, options)
	if err != nil {
		return nil, err
	}
	return &sqliteDatabase{d}, nil
}

// String ...
func (conn *sqliteDatabase) String() string {
	return "SQLite"
}
