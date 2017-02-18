package relational

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/rai-project/database"
)

type relationalTable struct {
	conn      *gorm.DB
	dbName    string
	tableName string
}

func NewTable(db database.Database, tableName string) (database.Table, error) {
	rdb, ok := db.(*relationalDatabase)
	if !ok {
		return nil, errors.New("invalid database input. Expecting a relational database instance")
	}
	return &relationalTable{
		conn:      rdb.conn,
		dbName:    rdb.databaseName,
		tableName: tableName,
	}, nil
}

func (tbl *relationalTable) Name() string {
	return tbl.tableName
}

func (tbl *relationalTable) Create(e interface{}) error {
	err := tbl.conn.DropTableIfExists(tbl.Name).Error
	if err != nil {
		return err
	}
	return tbl.conn.AutoMigrate(e).Error
}

func (tbl *relationalTable) Delete() error {
	return tbl.conn.DropTableIfExists(tbl.Name).Error
}

func (tbl *relationalTable) Insert(elem interface{}) error {
	return tbl.conn.Create(elem).Error
}
