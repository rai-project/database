package mysql

import (
	"context"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rai-project/database"
	"github.com/rai-project/database/relational"
	"upper.io/db.v3/mysql"
)

const (
	gormDialect = "mysql"
)

type mysqlDatabase struct {
	database.Database
}

func NewDatabase(databaseName string, opts ...database.Option) (database.Database, error) {
	options := database.Options{
		Endpoints:      Config.Endpoints,
		Username:       Config.Username,
		Password:       Config.Password,
		TLSConfig:      nil,
		MaxConnections: Config.MaxConnections,
		Context:        context.Background(),
	}

	if Config.Certificate != "" {
		database.TLSCertificate(Config.Certificate)(&options)
	}

	for _, o := range opts {
		o(&options)
	}

	connectionURL := mysql.ConnectionURL{
		User:     options.Username,
		Password: options.Password,
		Host:     options.Endpoints[0],
		Database: databaseName,
	}

	d, err := relational.NewDatabase(gormDialect, databaseName, connectionURL, options)
	if err != nil {
		return nil, err
	}
	return &mysqlDatabase{d}, nil
}

func (conn *mysqlDatabase) String() string {
	return "MySQL"
}
