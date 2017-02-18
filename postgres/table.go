package postgres

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/rai-project/database"
)

type postgresTable struct {
	conn      *gorm.DB
	dbName    string
	tableName string
}

func NewTable(db database.Database, tableName string) (database.Table, error) {
	rdb, ok := db.(*postgresDatabase)
	if !ok {
		return nil, errors.New("Invalid database input. Expecting a postgres database instance.")
	}
	return &postgresTable{
		conn:      rdb.conn,
		dbName:    rdb.databaseName,
		tableName: tableName,
	}, nil
}

func (tbl *postgresTable) Name() string {
	return tbl.tableName
}

func (tbl *postgresTable) Create(e interface{}) error {
	err := tbl.conn.DropTableIfExists(tbl.Name).Error
	if err != nil {
		return err
	}
	return tbl.conn.AutoMigrate(e).Error
}

func (tbl *postgresTable) Delete() error {
	return tbl.conn.DropTableIfExists(tbl.Name).Error
}

func (tbl *postgresTable) Insert(elem interface{}) error {
	return tbl.conn.Create(elem).Error
}
